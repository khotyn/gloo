/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	solo_io_v1 "github.com/solo-io/gloo/pkg/storage/crd/solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAttributes implements AttributeInterface
type FakeAttributes struct {
	Fake *FakeGlooV1
	ns   string
}

var attributesResource = schema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "attributes"}

var attributesKind = schema.GroupVersionKind{Group: "gloo.solo.io", Version: "v1", Kind: "Attribute"}

// Get takes name of the attribute, and returns the corresponding attribute object, and an error if there is any.
func (c *FakeAttributes) Get(name string, options v1.GetOptions) (result *solo_io_v1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(attributesResource, c.ns, name), &solo_io_v1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Attribute), err
}

// List takes label and field selectors, and returns the list of Attributes that match those selectors.
func (c *FakeAttributes) List(opts v1.ListOptions) (result *solo_io_v1.AttributeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(attributesResource, attributesKind, c.ns, opts), &solo_io_v1.AttributeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &solo_io_v1.AttributeList{}
	for _, item := range obj.(*solo_io_v1.AttributeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested attributes.
func (c *FakeAttributes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(attributesResource, c.ns, opts))

}

// Create takes the representation of a attribute and creates it.  Returns the server's representation of the attribute, and an error, if there is any.
func (c *FakeAttributes) Create(attribute *solo_io_v1.Attribute) (result *solo_io_v1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(attributesResource, c.ns, attribute), &solo_io_v1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Attribute), err
}

// Update takes the representation of a attribute and updates it. Returns the server's representation of the attribute, and an error, if there is any.
func (c *FakeAttributes) Update(attribute *solo_io_v1.Attribute) (result *solo_io_v1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(attributesResource, c.ns, attribute), &solo_io_v1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Attribute), err
}

// Delete takes name of the attribute and deletes it. Returns an error if one occurs.
func (c *FakeAttributes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(attributesResource, c.ns, name), &solo_io_v1.Attribute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAttributes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(attributesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &solo_io_v1.AttributeList{})
	return err
}

// Patch applies the patch and returns the patched attribute.
func (c *FakeAttributes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *solo_io_v1.Attribute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(attributesResource, c.ns, name, data, subresources...), &solo_io_v1.Attribute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Attribute), err
}
