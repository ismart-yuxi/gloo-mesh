// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/discovery/v1alpha2/traffic_target.proto

package v1alpha2

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
func (m *TrafficTargetSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec)
	if !ok {
		that2, ok := that.(TrafficTargetSpec)
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

	if h, ok := interface{}(m.GetMesh()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMesh()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMesh(), target.GetMesh()) {
			return false
		}
	}

	switch m.Type.(type) {

	case *TrafficTargetSpec_KubeService_:
		if _, ok := target.Type.(*TrafficTargetSpec_KubeService_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKubeService()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKubeService()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKubeService(), target.GetKubeService()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Type != target.Type {
			return false
		}
	}

	return true
}

// Equal function
func (m *TrafficTargetStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetStatus)
	if !ok {
		that2, ok := that.(TrafficTargetStatus)
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

	if len(m.GetAppliedTrafficPolicies()) != len(target.GetAppliedTrafficPolicies()) {
		return false
	}
	for idx, v := range m.GetAppliedTrafficPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedTrafficPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedTrafficPolicies()[idx]) {
				return false
			}
		}

	}

	if len(m.GetAppliedAccessPolicies()) != len(target.GetAppliedAccessPolicies()) {
		return false
	}
	for idx, v := range m.GetAppliedAccessPolicies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedAccessPolicies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedAccessPolicies()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetLocalFqdn(), target.GetLocalFqdn()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetAppliedFederation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAppliedFederation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAppliedFederation(), target.GetAppliedFederation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TrafficTargetSpec_KubeService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec_KubeService)
	if !ok {
		that2, ok := that.(TrafficTargetSpec_KubeService)
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

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if len(m.GetWorkloadSelectorLabels()) != len(target.GetWorkloadSelectorLabels()) {
		return false
	}
	for k, v := range m.GetWorkloadSelectorLabels() {

		if strings.Compare(v, target.GetWorkloadSelectorLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	if len(m.GetSubsets()) != len(target.GetSubsets()) {
		return false
	}
	for k, v := range m.GetSubsets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSubsets()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSubsets()[k]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if len(m.GetEndpointSubsets()) != len(target.GetEndpointSubsets()) {
		return false
	}
	for idx, v := range m.GetEndpointSubsets() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEndpointSubsets()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEndpointSubsets()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *TrafficTargetSpec_KubeService_KubeServicePort) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec_KubeService_KubeServicePort)
	if !ok {
		that2, ok := that.(TrafficTargetSpec_KubeService_KubeServicePort)
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

	if m.GetPort() != target.GetPort() {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetProtocol(), target.GetProtocol()) != 0 {
		return false
	}

	if strings.Compare(m.GetAppProtocol(), target.GetAppProtocol()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TrafficTargetSpec_KubeService_Subset) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec_KubeService_Subset)
	if !ok {
		that2, ok := that.(TrafficTargetSpec_KubeService_Subset)
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

	if len(m.GetValues()) != len(target.GetValues()) {
		return false
	}
	for idx, v := range m.GetValues() {

		if strings.Compare(v, target.GetValues()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *TrafficTargetSpec_KubeService_EndpointsSubset) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec_KubeService_EndpointsSubset)
	if !ok {
		that2, ok := that.(TrafficTargetSpec_KubeService_EndpointsSubset)
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

	if len(m.GetEndpoints()) != len(target.GetEndpoints()) {
		return false
	}
	for idx, v := range m.GetEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEndpoints()[idx]) {
				return false
			}
		}

	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *TrafficTargetSpec_KubeService_EndpointsSubset_Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec_KubeService_EndpointsSubset_Endpoint)
	if !ok {
		that2, ok := that.(TrafficTargetSpec_KubeService_EndpointsSubset_Endpoint)
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

	if strings.Compare(m.GetIpAddress(), target.GetIpAddress()) != 0 {
		return false
	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetSubLocality()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSubLocality()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSubLocality(), target.GetSubLocality()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TrafficTargetSpec_KubeService_EndpointsSubset_Endpoint_SubLocality) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetSpec_KubeService_EndpointsSubset_Endpoint_SubLocality)
	if !ok {
		that2, ok := that.(TrafficTargetSpec_KubeService_EndpointsSubset_Endpoint_SubLocality)
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

	if strings.Compare(m.GetZone(), target.GetZone()) != 0 {
		return false
	}

	if strings.Compare(m.GetSubzone(), target.GetSubzone()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TrafficTargetStatus_AppliedTrafficPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetStatus_AppliedTrafficPolicy)
	if !ok {
		that2, ok := that.(TrafficTargetStatus_AppliedTrafficPolicy)
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

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TrafficTargetStatus_AppliedAccessPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetStatus_AppliedAccessPolicy)
	if !ok {
		that2, ok := that.(TrafficTargetStatus_AppliedAccessPolicy)
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

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TrafficTargetStatus_AppliedFederation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TrafficTargetStatus_AppliedFederation)
	if !ok {
		that2, ok := that.(TrafficTargetStatus_AppliedFederation)
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

	if strings.Compare(m.GetFederatedHostname(), target.GetFederatedHostname()) != 0 {
		return false
	}

	if len(m.GetFederatedToMeshes()) != len(target.GetFederatedToMeshes()) {
		return false
	}
	for idx, v := range m.GetFederatedToMeshes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetFederatedToMeshes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetFederatedToMeshes()[idx]) {
				return false
			}
		}

	}

	if m.GetFlatNetwork() != target.GetFlatNetwork() {
		return false
	}

	return true
}
