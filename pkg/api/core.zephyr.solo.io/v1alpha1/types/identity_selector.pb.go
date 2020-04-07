// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/core/v1alpha1/identity_selector.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Selector capable of selecting specific service identities. Useful for binding policy rules.
//Either (namespaces, cluster, service_account_names) or service_accounts can be specified.
//If all fields are omitted, any source identity is permitted.
type IdentitySelector struct {
	// Types that are valid to be assigned to IdentitySelectorType:
	//	*IdentitySelector_Matcher_
	//	*IdentitySelector_ServiceAccountRefs_
	IdentitySelectorType isIdentitySelector_IdentitySelectorType `protobuf_oneof:"identity_selector_type"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *IdentitySelector) Reset()         { *m = IdentitySelector{} }
func (m *IdentitySelector) String() string { return proto.CompactTextString(m) }
func (*IdentitySelector) ProtoMessage()    {}
func (*IdentitySelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae7559b65fb129e0, []int{0}
}
func (m *IdentitySelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentitySelector.Unmarshal(m, b)
}
func (m *IdentitySelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentitySelector.Marshal(b, m, deterministic)
}
func (m *IdentitySelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentitySelector.Merge(m, src)
}
func (m *IdentitySelector) XXX_Size() int {
	return xxx_messageInfo_IdentitySelector.Size(m)
}
func (m *IdentitySelector) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentitySelector.DiscardUnknown(m)
}

var xxx_messageInfo_IdentitySelector proto.InternalMessageInfo

type isIdentitySelector_IdentitySelectorType interface {
	isIdentitySelector_IdentitySelectorType()
	Equal(interface{}) bool
}

type IdentitySelector_Matcher_ struct {
	Matcher *IdentitySelector_Matcher `protobuf:"bytes,1,opt,name=matcher,proto3,oneof" json:"matcher,omitempty"`
}
type IdentitySelector_ServiceAccountRefs_ struct {
	ServiceAccountRefs *IdentitySelector_ServiceAccountRefs `protobuf:"bytes,2,opt,name=service_account_refs,json=serviceAccountRefs,proto3,oneof" json:"service_account_refs,omitempty"`
}

func (*IdentitySelector_Matcher_) isIdentitySelector_IdentitySelectorType()            {}
func (*IdentitySelector_ServiceAccountRefs_) isIdentitySelector_IdentitySelectorType() {}

func (m *IdentitySelector) GetIdentitySelectorType() isIdentitySelector_IdentitySelectorType {
	if m != nil {
		return m.IdentitySelectorType
	}
	return nil
}

func (m *IdentitySelector) GetMatcher() *IdentitySelector_Matcher {
	if x, ok := m.GetIdentitySelectorType().(*IdentitySelector_Matcher_); ok {
		return x.Matcher
	}
	return nil
}

func (m *IdentitySelector) GetServiceAccountRefs() *IdentitySelector_ServiceAccountRefs {
	if x, ok := m.GetIdentitySelectorType().(*IdentitySelector_ServiceAccountRefs_); ok {
		return x.ServiceAccountRefs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IdentitySelector) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IdentitySelector_Matcher_)(nil),
		(*IdentitySelector_ServiceAccountRefs_)(nil),
	}
}

type IdentitySelector_Matcher struct {
	// Namespaces to allow. If not set, any namespace is allowed.
	Namespaces []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// Cluster to allow. If not set, any cluster is allowed.
	Clusters             []string `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentitySelector_Matcher) Reset()         { *m = IdentitySelector_Matcher{} }
func (m *IdentitySelector_Matcher) String() string { return proto.CompactTextString(m) }
func (*IdentitySelector_Matcher) ProtoMessage()    {}
func (*IdentitySelector_Matcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae7559b65fb129e0, []int{0, 0}
}
func (m *IdentitySelector_Matcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentitySelector_Matcher.Unmarshal(m, b)
}
func (m *IdentitySelector_Matcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentitySelector_Matcher.Marshal(b, m, deterministic)
}
func (m *IdentitySelector_Matcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentitySelector_Matcher.Merge(m, src)
}
func (m *IdentitySelector_Matcher) XXX_Size() int {
	return xxx_messageInfo_IdentitySelector_Matcher.Size(m)
}
func (m *IdentitySelector_Matcher) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentitySelector_Matcher.DiscardUnknown(m)
}

var xxx_messageInfo_IdentitySelector_Matcher proto.InternalMessageInfo

func (m *IdentitySelector_Matcher) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *IdentitySelector_Matcher) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type IdentitySelector_ServiceAccountRefs struct {
	// List of ServiceAccounts to allow. If not set, any ServiceAccount is allowed.
	ServiceAccounts      []*ResourceRef `protobuf:"bytes,1,rep,name=service_accounts,json=serviceAccounts,proto3" json:"service_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IdentitySelector_ServiceAccountRefs) Reset()         { *m = IdentitySelector_ServiceAccountRefs{} }
func (m *IdentitySelector_ServiceAccountRefs) String() string { return proto.CompactTextString(m) }
func (*IdentitySelector_ServiceAccountRefs) ProtoMessage()    {}
func (*IdentitySelector_ServiceAccountRefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae7559b65fb129e0, []int{0, 1}
}
func (m *IdentitySelector_ServiceAccountRefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentitySelector_ServiceAccountRefs.Unmarshal(m, b)
}
func (m *IdentitySelector_ServiceAccountRefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentitySelector_ServiceAccountRefs.Marshal(b, m, deterministic)
}
func (m *IdentitySelector_ServiceAccountRefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentitySelector_ServiceAccountRefs.Merge(m, src)
}
func (m *IdentitySelector_ServiceAccountRefs) XXX_Size() int {
	return xxx_messageInfo_IdentitySelector_ServiceAccountRefs.Size(m)
}
func (m *IdentitySelector_ServiceAccountRefs) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentitySelector_ServiceAccountRefs.DiscardUnknown(m)
}

var xxx_messageInfo_IdentitySelector_ServiceAccountRefs proto.InternalMessageInfo

func (m *IdentitySelector_ServiceAccountRefs) GetServiceAccounts() []*ResourceRef {
	if m != nil {
		return m.ServiceAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*IdentitySelector)(nil), "core.zephyr.solo.io.IdentitySelector")
	proto.RegisterType((*IdentitySelector_Matcher)(nil), "core.zephyr.solo.io.IdentitySelector.Matcher")
	proto.RegisterType((*IdentitySelector_ServiceAccountRefs)(nil), "core.zephyr.solo.io.IdentitySelector.ServiceAccountRefs")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/core/v1alpha1/identity_selector.proto", fileDescriptor_ae7559b65fb129e0)
}

var fileDescriptor_ae7559b65fb129e0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x10, 0xc7, 0xc9, 0x22, 0x2d, 0xbb, 0xe6, 0xb0, 0xc8, 0x8b, 0x56, 0x51, 0x0e, 0x08, 0xed, 0x89,
	0xc3, 0xc6, 0x16, 0xec, 0xa5, 0xd7, 0x22, 0x55, 0x02, 0x55, 0xf4, 0x10, 0x7a, 0xea, 0x25, 0x72,
	0xdc, 0xc9, 0x87, 0x9a, 0x30, 0x96, 0xed, 0x50, 0xd1, 0x97, 0xe8, 0x6b, 0xf4, 0xb9, 0xfa, 0x24,
	0x55, 0x3e, 0x4a, 0x5b, 0xe0, 0xc0, 0x6d, 0x3c, 0x9e, 0xff, 0x7f, 0x7e, 0xa3, 0x19, 0xb2, 0x4a,
	0x32, 0x9b, 0x96, 0x11, 0x93, 0x58, 0x70, 0x83, 0x39, 0xfa, 0x19, 0x72, 0x03, 0x7a, 0x9b, 0x49,
	0xf0, 0x0b, 0x30, 0xa9, 0x9f, 0x96, 0x11, 0x17, 0x2a, 0xe3, 0x12, 0x35, 0xf0, 0xed, 0x54, 0xe4,
	0x2a, 0x15, 0x53, 0x9e, 0xdd, 0xc3, 0xc6, 0x66, 0x76, 0x17, 0x1a, 0xc8, 0x41, 0x5a, 0xd4, 0x4c,
	0x69, 0xb4, 0x48, 0x7f, 0x57, 0x65, 0xec, 0x09, 0x54, 0xba, 0xd3, 0xac, 0xf2, 0x63, 0x19, 0x7a,
	0xa3, 0x04, 0x31, 0xc9, 0x81, 0xd7, 0x25, 0x51, 0x19, 0xf3, 0x47, 0x2d, 0x94, 0x02, 0x6d, 0x1a,
	0x91, 0xf7, 0xef, 0x8c, 0x86, 0x1a, 0xe2, 0xb6, 0x7a, 0x98, 0x60, 0x82, 0x75, 0xc8, 0xab, 0xa8,
	0xc9, 0xfe, 0x7d, 0xee, 0x92, 0xc1, 0xb2, 0x85, 0x5a, 0xb7, 0x4c, 0x74, 0x49, 0x7a, 0x85, 0xb0,
	0x32, 0x05, 0xed, 0x3a, 0x63, 0x67, 0xd2, 0x9f, 0xf9, 0xec, 0x04, 0x1f, 0x3b, 0xd4, 0xb1, 0x55,
	0x23, 0x5a, 0x74, 0x82, 0x77, 0x3d, 0xcd, 0xc9, 0xb0, 0xa5, 0x0c, 0x85, 0x94, 0x58, 0x6e, 0x6c,
	0xa8, 0x21, 0x36, 0xee, 0xb7, 0xda, 0xf7, 0xe2, 0x3c, 0xdf, 0x75, 0xe3, 0x70, 0xd9, 0x18, 0x04,
	0x10, 0x9b, 0x45, 0x27, 0xa0, 0xe6, 0x28, 0xeb, 0x5d, 0x91, 0x5e, 0xcb, 0x40, 0x47, 0x84, 0x6c,
	0x44, 0x01, 0x46, 0x09, 0x09, 0xc6, 0x75, 0xc6, 0xdd, 0xc9, 0xcf, 0xe0, 0x53, 0x86, 0x7a, 0xe4,
	0x87, 0xcc, 0x4b, 0x63, 0x41, 0x57, 0x30, 0xd5, 0xef, 0xfe, 0xed, 0x09, 0x42, 0x8f, 0x5b, 0xd2,
	0x6b, 0x32, 0x38, 0x18, 0xa5, 0xf1, 0xed, 0xcf, 0xc6, 0x27, 0xc7, 0x08, 0xc0, 0x60, 0xa9, 0x25,
	0x04, 0x10, 0x07, 0xbf, 0xbe, 0xc2, 0x9a, 0xb9, 0x4b, 0xfe, 0x1c, 0xdd, 0x42, 0x68, 0x77, 0x0a,
	0xe6, 0xb7, 0x2f, 0xaf, 0x23, 0xe7, 0xee, 0xe6, 0x9c, 0xfb, 0x52, 0x0f, 0xc9, 0x7e, 0xe5, 0x07,
	0xdd, 0x3f, 0x2e, 0xa0, 0x32, 0x35, 0xd1, 0xf7, 0x7a, 0xdd, 0xff, 0xdf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xe0, 0xfa, 0x97, 0xb8, 0x02, 0x00, 0x00,
}

func (this *IdentitySelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector)
	if !ok {
		that2, ok := that.(IdentitySelector)
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
	if that1.IdentitySelectorType == nil {
		if this.IdentitySelectorType != nil {
			return false
		}
	} else if this.IdentitySelectorType == nil {
		return false
	} else if !this.IdentitySelectorType.Equal(that1.IdentitySelectorType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IdentitySelector_Matcher_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector_Matcher_)
	if !ok {
		that2, ok := that.(IdentitySelector_Matcher_)
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
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	return true
}
func (this *IdentitySelector_ServiceAccountRefs_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector_ServiceAccountRefs_)
	if !ok {
		that2, ok := that.(IdentitySelector_ServiceAccountRefs_)
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
	if !this.ServiceAccountRefs.Equal(that1.ServiceAccountRefs) {
		return false
	}
	return true
}
func (this *IdentitySelector_Matcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector_Matcher)
	if !ok {
		that2, ok := that.(IdentitySelector_Matcher)
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
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IdentitySelector_ServiceAccountRefs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IdentitySelector_ServiceAccountRefs)
	if !ok {
		that2, ok := that.(IdentitySelector_ServiceAccountRefs)
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
	if len(this.ServiceAccounts) != len(that1.ServiceAccounts) {
		return false
	}
	for i := range this.ServiceAccounts {
		if !this.ServiceAccounts[i].Equal(that1.ServiceAccounts[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}