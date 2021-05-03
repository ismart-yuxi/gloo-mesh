// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/virtual_host.proto

package v1beta1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *VirtualHostSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHostSpec)
	if !ok {
		that2, ok := that.(VirtualHostSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetDomains()) != len(target.GetDomains()) {
		return false
	}
	for idx, v := range m.GetDomains() {

		if strings.Compare(v, target.GetDomains()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetRoutes()) != len(target.GetRoutes()) {
		return false
	}
	for idx, v := range m.GetRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRoutes()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetVirtualHostOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVirtualHostOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVirtualHostOptions(), target.GetVirtualHostOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualHostStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHostStatus)
	if !ok {
		that2, ok := that.(VirtualHostStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetWarnings()) != len(target.GetWarnings()) {
		return false
	}
	for idx, v := range m.GetWarnings() {

		if strings.Compare(v, target.GetWarnings()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAttachedVirtualGateways()) != len(target.GetAttachedVirtualGateways()) {
		return false
	}
	for idx, v := range m.GetAttachedVirtualGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAttachedVirtualGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAttachedVirtualGateways()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSelectedVirtualHosts()) != len(target.GetSelectedVirtualHosts()) {
		return false
	}
	for idx, v := range m.GetSelectedVirtualHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedVirtualHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedVirtualHosts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualHostSpec_VirtualHostOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHostSpec_VirtualHostOptions)
	if !ok {
		that2, ok := that.(VirtualHostSpec_VirtualHostOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetTodoAddOptions(), target.GetTodoAddOptions()) != 0 {
		return false
	}

	return true
}