package consul

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"github.com/solo-io/gloo/pkg/api/types/v1"
	"github.com/solo-io/gloo/pkg/backoff"
	"github.com/solo-io/gloo/pkg/config"
	"github.com/solo-io/gloo/pkg/plugins/consul"
	"github.com/solo-io/gloo/pkg/storage"
)

const (
	generatedBy = "consul-upstream-discovery"
)

type UpstreamController struct {
	errors            chan error
	syncer            config.UpstreamSyncer
	consul            *api.Client
	ctx               context.Context
	latestServiceList map[string][]string
}

func NewUpstreamController(cfg *api.Config,
	configStore storage.Interface) (*UpstreamController, error) {
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client: %v", err)
	}

	// attempt to register upstreams if they don't exist
	if err := configStore.V1().Register(); err != nil && !storage.IsAlreadyExists(err) {
		return nil, errors.Wrap(err, "failed to register upstreams")
	}

	c := &UpstreamController{
		errors: make(chan error),
		consul: client,
		syncer: config.UpstreamSyncer{
			Owner:       generatedBy,
			GlooStorage: configStore,
		},
	}

	return c, nil
}

func (c *UpstreamController) Run(stop <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c.ctx = ctx
	discoveredServices := make(chan []consulService)
	go c.watchConsulServices(ctx, discoveredServices)
	for {
		select {
		case <-stop:
			log.Printf("consul upstream discovery stopped")
			return
		case serviceList := <-discoveredServices:
			c.syncGlooUpstreamsWithConsulServices(serviceList)
		}
	}
}

func (c *UpstreamController) Error() <-chan error {
	return c.errors
}

type consulService struct {
	name    string
	tagSets [][]string
	connect bool
}

func (c *UpstreamController) watchConsulServices(ctx context.Context, discoveredServices chan []consulService) {
	var lastIndex uint64
	for {
		select {
		case <-ctx.Done():
			return
		default:
			backoff.UntilSuccess(func() error {
				services, index, err := c.getNextUpdate(ctx, lastIndex)
				if err != nil {
					return errors.Wrapf(err, "getting next endpoints for consul upstream failed")
				}
				lastIndex = index
				if len(services) == 0 {
					return nil
				}
				// get each unique set of tags
				// we will use this to generate an upstream for each unique set
				var consulServices []consulService

				for svcName := range services {
					serviceInstances, _, err := c.consul.Catalog().Service(svcName, "", &api.QueryOptions{RequireConsistent: true})
					if err != nil {
						return errors.Wrapf(err, "failed to get instances of service %s", svcName)
					}
					var allTagSets [][]string
					for _, inst := range serviceInstances {
						allTagSets = append(allTagSets, inst.ServiceTags)
					}
					// add a service with no tags, so the service can be accessed regardless of tags.
					allTagSets = append(allTagSets, []string{})

					consulServices = append(consulServices, consulService{
						name:    svcName,
						tagSets: uniqueTagSets(allTagSets),
						connect: false,
					})

					proxyInstances, _, _ := c.consul.Catalog().Connect(svcName, "", &api.QueryOptions{RequireConsistent: true})

					if len(proxyInstances) > 0 {
						consulServices = append(consulServices, consulService{
							name:    svcName,
							tagSets: uniqueTagSets(allTagSets),
							connect: true,
						})
					}
				}

				discoveredServices <- consulServices
				return nil
			}, ctx)
		}
	}
}

func uniqueTagSets(allTagSets [][]string) [][]string {
	var uniqueTagSets [][]string
	for _, tags := range allTagSets {
		// sort for idempotency
		sort.SliceStable(tags, func(i, j int) bool {
			return tags[i] < tags[j]
		})

		// check if this set already exists in the list
		var duplicate bool
		for _, set := range uniqueTagSets {
			if sliceEqual(tags, set) {
				duplicate = true
				break
			}
		}
		// if this set is accounted for we don't want a new upstream for this instance
		if duplicate {
			continue
		}

		uniqueTagSets = append(uniqueTagSets, tags)
	}
	// sort the set of sets
	sort.SliceStable(uniqueTagSets, func(i, j int) bool {
		tags1 := uniqueTagSets[i]
		tags2 := uniqueTagSets[j]
		if len(tags1) != len(tags2) {
			return len(tags1) < len(tags2)
		}
		for i := range tags1 {
			if tags1[i] != tags2[i] {
				return tags1[i] < tags2[i]
			}
		}
		panic("they're equal!?! THIS SHOULD NOT HAVE HAPPENED")
	})
	return uniqueTagSets
}

func sliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func (c *UpstreamController) getNextUpdate(ctx context.Context, lastIndex uint64) (map[string][]string, uint64, error) {
	opts := &api.QueryOptions{RequireConsistent: true, WaitIndex: lastIndex}
	opts = opts.WithContext(ctx)
	services, meta, err := c.consul.Catalog().Services(opts)
	if err != nil {
		return nil, lastIndex, errors.Wrapf(err, "failed to consul list services")
	}
	return services, meta.LastIndex, nil
}

func (c *UpstreamController) syncGlooUpstreamsWithConsulServices(serviceList []consulService) {
	c.syncer.DesiredUpstreams = c.convertServicesFunc(serviceList)
	if err := c.syncer.SyncDesiredState(); err != nil {
		c.errors <- err
	}
}

func (c *UpstreamController) convertServicesFunc(serviceList []consulService) func() ([]*v1.Upstream, error) {
	return func() ([]*v1.Upstream, error) {
		var upstreams []*v1.Upstream
		for _, svc := range serviceList {
			if svc.connect {
				us := &v1.Upstream{
					Name: consul.UpstreamNameForConnectService(svc.name),
					Type: consul.UpstreamTypeConsul,
					Spec: consul.EncodeUpstreamSpec(consul.UpstreamSpec{
						ServiceName: svc.name,
						Connect: &consul.Connect{
							TlsSecretRef: consul.LeafCertificateSecret,
						},
					}),
				}
				upstreams = append(upstreams, us)
				continue
			}
			for _, tags := range svc.tagSets {
				us := &v1.Upstream{
					Name: upstreamName(svc.name, tags),
					Type: consul.UpstreamTypeConsul,
					Spec: consul.EncodeUpstreamSpec(consul.UpstreamSpec{
						ServiceName: svc.name,
						ServiceTags: tags,
					}),
				}
				upstreams = append(upstreams, us)
			}
		}
		return upstreams, nil
	}
}

func upstreamName(serviceName string, tags []string) string {
	if len(tags) < 1 {
		return serviceName
	}
	return fmt.Sprintf("%s-%s", serviceName, strings.Join(tags, "-"))
}
