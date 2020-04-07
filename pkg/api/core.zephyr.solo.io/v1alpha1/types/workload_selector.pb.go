// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/core/v1alpha1/workload_selector.proto

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
//Select Kubernetes workloads directly using label and/or namespace criteria. See comments on the fields for
//detailed semantics.
type WorkloadSelector struct {
	// If specified, all labels must exist on workloads, else match on any labels.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If specified, match workloads if they exist in one of the specified namespaces. If not specified, match on any namespace.
	Namespaces           []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadSelector) Reset()         { *m = WorkloadSelector{} }
func (m *WorkloadSelector) String() string { return proto.CompactTextString(m) }
func (*WorkloadSelector) ProtoMessage()    {}
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8aa2157c287aedb, []int{0}
}
func (m *WorkloadSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSelector.Unmarshal(m, b)
}
func (m *WorkloadSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSelector.Marshal(b, m, deterministic)
}
func (m *WorkloadSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSelector.Merge(m, src)
}
func (m *WorkloadSelector) XXX_Size() int {
	return xxx_messageInfo_WorkloadSelector.Size(m)
}
func (m *WorkloadSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSelector.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSelector proto.InternalMessageInfo

func (m *WorkloadSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WorkloadSelector) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkloadSelector)(nil), "core.zephyr.solo.io.WorkloadSelector")
	proto.RegisterMapType((map[string]string)(nil), "core.zephyr.solo.io.WorkloadSelector.LabelsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/core/v1alpha1/workload_selector.proto", fileDescriptor_b8aa2157c287aedb)
}

var fileDescriptor_b8aa2157c287aedb = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcf, 0x4a, 0xf3, 0x40,
	0x1c, 0x64, 0x5b, 0xbe, 0x42, 0xb7, 0x97, 0x92, 0xaf, 0x87, 0xd0, 0x43, 0x08, 0x9e, 0x72, 0x30,
	0xbb, 0x44, 0x2f, 0xea, 0x51, 0xf0, 0x20, 0xa8, 0x87, 0x28, 0x08, 0x5e, 0x64, 0x13, 0x7f, 0xdd,
	0x84, 0x6c, 0xfa, 0x5b, 0x76, 0x93, 0x94, 0xf8, 0x44, 0x3e, 0x82, 0xcf, 0xe3, 0x93, 0x48, 0xfe,
	0xa0, 0xa5, 0x78, 0xe8, 0x6d, 0x76, 0x98, 0x99, 0x9d, 0x9d, 0xa5, 0xf7, 0x32, 0xaf, 0xb2, 0x3a,
	0x61, 0x29, 0x96, 0xdc, 0xa2, 0xc2, 0x30, 0x47, 0x6e, 0xc1, 0x34, 0x79, 0x0a, 0x61, 0x09, 0x36,
	0x0b, 0xb3, 0x3a, 0xe1, 0x42, 0xe7, 0x3c, 0x45, 0x03, 0xbc, 0x89, 0x84, 0xd2, 0x99, 0x88, 0xf8,
	0x0e, 0x4d, 0xa1, 0x50, 0xbc, 0xbd, 0x5a, 0x50, 0x90, 0x56, 0x68, 0x98, 0x36, 0x58, 0xa1, 0xf3,
	0xbf, 0x93, 0xb1, 0x77, 0xd0, 0x59, 0x6b, 0x58, 0x97, 0xc7, 0x72, 0x5c, 0x7b, 0x12, 0x51, 0x2a,
	0xe0, 0xbd, 0x24, 0xa9, 0x37, 0x7c, 0x67, 0x84, 0xd6, 0x60, 0xec, 0x60, 0x5a, 0x9f, 0x1e, 0x71,
	0xa1, 0x81, 0xcd, 0xa8, 0x5e, 0x49, 0x94, 0xd8, 0x43, 0xde, 0xa1, 0x81, 0x3d, 0xf9, 0x24, 0x74,
	0xf9, 0x3c, 0x96, 0x7a, 0x1c, 0x3b, 0x39, 0xb7, 0x74, 0xa6, 0x44, 0x02, 0xca, 0xba, 0xc4, 0x9f,
	0x06, 0x8b, 0xb3, 0x88, 0xfd, 0x51, 0x8f, 0x1d, 0xda, 0xd8, 0x5d, 0xef, 0xb9, 0xd9, 0x56, 0xa6,
	0x8d, 0xc7, 0x00, 0xc7, 0xa3, 0x74, 0x2b, 0x4a, 0xb0, 0x5a, 0xa4, 0x60, 0xdd, 0x89, 0x3f, 0x0d,
	0xe6, 0xf1, 0x1e, 0xb3, 0xbe, 0xa4, 0x8b, 0x3d, 0x9b, 0xb3, 0xa4, 0xd3, 0x02, 0x5a, 0x97, 0xf8,
	0x24, 0x98, 0xc7, 0x1d, 0x74, 0x56, 0xf4, 0x5f, 0x23, 0x54, 0x0d, 0xee, 0xa4, 0xe7, 0x86, 0xc3,
	0xd5, 0xe4, 0x82, 0x5c, 0x3f, 0x7d, 0x7c, 0x79, 0xe4, 0xe5, 0xe1, 0x98, 0x8f, 0xd0, 0x85, 0xfc,
	0xd9, 0xe6, 0xe0, 0x19, 0xbf, 0x53, 0x55, 0xad, 0x06, 0x9b, 0xcc, 0xfa, 0x5d, 0xce, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0xf1, 0xeb, 0x6d, 0xe1, 0x01, 0x00, 0x00,
}

func (this *WorkloadSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkloadSelector)
	if !ok {
		that2, ok := that.(WorkloadSelector)
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
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
