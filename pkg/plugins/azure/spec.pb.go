// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec.proto

/*
Package azure is a generated protocol buffer package.

It is generated from these files:
	spec.proto

It has these top-level messages:
	UpstreamSpec
	FunctionSpec
*/
package azure

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for Azure Functions Upstreams
// Azure Upstreams represent a collection of Azure Functions for a particular Azure Account (IAM Role or User account)
// within a particular Function App
type UpstreamSpec struct {
	// The Name of the Azure Function App where the functions are grouped
	FunctionAppName string `protobuf:"bytes,1,opt,name=function_app_name,json=functionAppName,proto3" json:"function_app_name,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an [Azure Publish Profile JSON file](https://azure.microsoft.com/en-us/downloads/publishing-profile-overview/).
	// {{ hide_not_implemented "Azure Secrets can be created with `glooctl secret create azure ...`" }}
	// Note that this secret is not required unless Function Discovery is enabled
	SecretRef string `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
}

func (m *UpstreamSpec) Reset()                    { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string            { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()               {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{0} }

func (m *UpstreamSpec) GetFunctionAppName() string {
	if m != nil {
		return m.FunctionAppName
	}
	return ""
}

func (m *UpstreamSpec) GetSecretRef() string {
	if m != nil {
		return m.SecretRef
	}
	return ""
}

// Function Spec for Functions on Azure Functions Upstreams
// The Function Spec contains data necessary for Gloo to invoke Azure functions
type FunctionSpec struct {
	// The Name of the Azure Function as it appears in the Azure Functions Portal
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
	// Auth Level can bve either "anonymous" "function" or "admin"
	// See https://vincentlauzon.com/2017/12/04/azure-functions-http-authorization-levels/ for more details
	AuthLevel string `protobuf:"bytes,2,opt,name=auth_level,json=authLevel,proto3" json:"auth_level,omitempty"`
}

func (m *FunctionSpec) Reset()                    { *m = FunctionSpec{} }
func (m *FunctionSpec) String() string            { return proto.CompactTextString(m) }
func (*FunctionSpec) ProtoMessage()               {}
func (*FunctionSpec) Descriptor() ([]byte, []int) { return fileDescriptorSpec, []int{1} }

func (m *FunctionSpec) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *FunctionSpec) GetAuthLevel() string {
	if m != nil {
		return m.AuthLevel
	}
	return ""
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.api.azure.v1.UpstreamSpec")
	proto.RegisterType((*FunctionSpec)(nil), "gloo.api.azure.v1.FunctionSpec")
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionAppName != that1.FunctionAppName {
		return false
	}
	if this.SecretRef != that1.SecretRef {
		return false
	}
	return true
}
func (this *FunctionSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FunctionSpec)
	if !ok {
		that2, ok := that.(FunctionSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FunctionName != that1.FunctionName {
		return false
	}
	if this.AuthLevel != that1.AuthLevel {
		return false
	}
	return true
}

func init() { proto.RegisterFile("spec.proto", fileDescriptorSpec) }

var fileDescriptorSpec = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0x89, 0x85, 0x70, 0x43, 0x44, 0x2e, 0x58, 0x1c, 0x82, 0x22, 0x67, 0xa3, 0x82, 0xbb,
	0x88, 0x4f, 0xa0, 0x85, 0x95, 0x58, 0x44, 0x2c, 0xb4, 0x09, 0x7b, 0xcb, 0x64, 0x6f, 0x71, 0x93,
	0x19, 0xf6, 0xcf, 0x15, 0x3e, 0x91, 0xcf, 0xe5, 0x93, 0x48, 0x36, 0xae, 0x5c, 0x37, 0xfc, 0xbe,
	0x6f, 0x7e, 0xf0, 0x01, 0x04, 0x46, 0x2d, 0xd8, 0x53, 0xa4, 0x66, 0x69, 0x1c, 0x91, 0x50, 0x6c,
	0x85, 0xfa, 0x4a, 0x1e, 0xc5, 0xee, 0xee, 0xf4, 0xc4, 0x90, 0xa1, 0x9c, 0xca, 0xe9, 0x9a, 0x8b,
	0xeb, 0x77, 0xa8, 0xdf, 0x38, 0x44, 0x8f, 0x6a, 0x78, 0x65, 0xd4, 0xcd, 0x0d, 0x2c, 0xfb, 0x34,
	0xea, 0x68, 0x69, 0xec, 0x14, 0x73, 0x37, 0xaa, 0x01, 0x57, 0xd5, 0x45, 0x75, 0xb5, 0x68, 0x8f,
	0x4b, 0xf0, 0xc0, 0xfc, 0xa2, 0x06, 0x6c, 0xce, 0x00, 0x02, 0x6a, 0x8f, 0xb1, 0xf3, 0xd8, 0xaf,
	0x0e, 0x72, 0x69, 0x31, 0x93, 0x16, 0xfb, 0x75, 0x0b, 0xf5, 0xd3, 0xdf, 0x47, 0x56, 0x5f, 0xc2,
	0xd1, 0xbf, 0x7a, 0x4f, 0x5b, 0x17, 0x58, 0x9c, 0x2a, 0xc5, 0x6d, 0xe7, 0x70, 0x87, 0xae, 0x38,
	0x27, 0xf2, 0x3c, 0x81, 0x47, 0xf9, 0xfd, 0x73, 0x5e, 0x7d, 0x5c, 0x1b, 0x1b, 0xb7, 0x69, 0x23,
	0x34, 0x0d, 0x32, 0x90, 0xa3, 0x5b, 0x4b, 0x72, 0x1a, 0x2c, 0xf9, 0xd3, 0x48, 0x76, 0xc9, 0xd8,
	0x31, 0xc8, 0x3c, 0x7c, 0x73, 0x98, 0x67, 0xde, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x10,
	0x00, 0xc3, 0x1d, 0x01, 0x00, 0x00,
}
