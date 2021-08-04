// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/discovery/v1/mesh.proto

package v1

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
func (m *MeshSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec)
	if !ok {
		that2, ok := that.(MeshSpec)
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

	if h, ok := interface{}(m.GetAgentInfo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAgentInfo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAgentInfo(), target.GetAgentInfo()) {
			return false
		}
	}

	switch m.Type.(type) {

	case *MeshSpec_Istio_:
		if _, ok := target.Type.(*MeshSpec_Istio_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetIstio()).(equality.Equalizer); ok {
			if !h.Equal(target.GetIstio()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetIstio(), target.GetIstio()) {
				return false
			}
		}

	case *MeshSpec_AwsAppMesh_:
		if _, ok := target.Type.(*MeshSpec_AwsAppMesh_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAwsAppMesh()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAwsAppMesh()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAwsAppMesh(), target.GetAwsAppMesh()) {
				return false
			}
		}

	case *MeshSpec_Linkerd:
		if _, ok := target.Type.(*MeshSpec_Linkerd); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLinkerd()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLinkerd()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLinkerd(), target.GetLinkerd()) {
				return false
			}
		}

	case *MeshSpec_ConsulConnect:
		if _, ok := target.Type.(*MeshSpec_ConsulConnect); !ok {
			return false
		}

		if h, ok := interface{}(m.GetConsulConnect()).(equality.Equalizer); ok {
			if !h.Equal(target.GetConsulConnect()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetConsulConnect(), target.GetConsulConnect()) {
				return false
			}
		}

	case *MeshSpec_Osm:
		if _, ok := target.Type.(*MeshSpec_Osm); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOsm()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOsm()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOsm(), target.GetOsm()) {
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
func (m *MeshInstallation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshInstallation)
	if !ok {
		that2, ok := that.(MeshInstallation)
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

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	if len(m.GetPodLabels()) != len(target.GetPodLabels()) {
		return false
	}
	for k, v := range m.GetPodLabels() {

		if strings.Compare(v, target.GetPodLabels()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetVersion(), target.GetVersion()) != 0 {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *MeshStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshStatus)
	if !ok {
		that2, ok := that.(MeshStatus)
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

	if h, ok := interface{}(m.GetAppliedVirtualMesh()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAppliedVirtualMesh()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAppliedVirtualMesh(), target.GetAppliedVirtualMesh()) {
			return false
		}
	}

	if len(m.GetAppliedVirtualDestinations()) != len(target.GetAppliedVirtualDestinations()) {
		return false
	}
	for idx, v := range m.GetAppliedVirtualDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedVirtualDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedVirtualDestinations()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *MeshSpec_Istio) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_Istio)
	if !ok {
		that2, ok := that.(MeshSpec_Istio)
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

	if h, ok := interface{}(m.GetInstallation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInstallation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInstallation(), target.GetInstallation()) {
			return false
		}
	}

	if strings.Compare(m.GetTrustDomain(), target.GetTrustDomain()) != 0 {
		return false
	}

	if strings.Compare(m.GetIstiodServiceAccount(), target.GetIstiodServiceAccount()) != 0 {
		return false
	}

	if len(m.GetIngressGateways()) != len(target.GetIngressGateways()) {
		return false
	}
	for idx, v := range m.GetIngressGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetIngressGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetIngressGateways()[idx]) {
				return false
			}
		}

	}

	if m.GetSmartDnsProxyingEnabled() != target.GetSmartDnsProxyingEnabled() {
		return false
	}

	return true
}

// Equal function
func (m *MeshSpec_AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_AwsAppMesh)
	if !ok {
		that2, ok := that.(MeshSpec_AwsAppMesh)
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

	if strings.Compare(m.GetAwsName(), target.GetAwsName()) != 0 {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if strings.Compare(m.GetAwsAccountId(), target.GetAwsAccountId()) != 0 {
		return false
	}

	if strings.Compare(m.GetArn(), target.GetArn()) != 0 {
		return false
	}

	if len(m.GetClusters()) != len(target.GetClusters()) {
		return false
	}
	for idx, v := range m.GetClusters() {

		if strings.Compare(v, target.GetClusters()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *MeshSpec_LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_LinkerdMesh)
	if !ok {
		that2, ok := that.(MeshSpec_LinkerdMesh)
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

	if h, ok := interface{}(m.GetInstallation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInstallation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInstallation(), target.GetInstallation()) {
			return false
		}
	}

	if strings.Compare(m.GetClusterDomain(), target.GetClusterDomain()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *MeshSpec_ConsulConnectMesh) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_ConsulConnectMesh)
	if !ok {
		that2, ok := that.(MeshSpec_ConsulConnectMesh)
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

	if h, ok := interface{}(m.GetInstallation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInstallation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInstallation(), target.GetInstallation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *MeshSpec_OSM) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_OSM)
	if !ok {
		that2, ok := that.(MeshSpec_OSM)
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

	if h, ok := interface{}(m.GetInstallation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInstallation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInstallation(), target.GetInstallation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *MeshSpec_AgentInfo) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_AgentInfo)
	if !ok {
		that2, ok := that.(MeshSpec_AgentInfo)
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

	if strings.Compare(m.GetAgentNamespace(), target.GetAgentNamespace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *MeshSpec_Istio_IngressGatewayInfo) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshSpec_Istio_IngressGatewayInfo)
	if !ok {
		that2, ok := that.(MeshSpec_Istio_IngressGatewayInfo)
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

	if len(m.GetWorkloadLabels()) != len(target.GetWorkloadLabels()) {
		return false
	}
	for k, v := range m.GetWorkloadLabels() {

		if strings.Compare(v, target.GetWorkloadLabels()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetExternalAddress(), target.GetExternalAddress()) != 0 {
		return false
	}

	if m.GetExternalTlsPort() != target.GetExternalTlsPort() {
		return false
	}

	if m.GetTlsContainerPort() != target.GetTlsContainerPort() {
		return false
	}

	return true
}

// Equal function
func (m *MeshStatus_AppliedVirtualMesh) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshStatus_AppliedVirtualMesh)
	if !ok {
		that2, ok := that.(MeshStatus_AppliedVirtualMesh)
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
func (m *MeshStatus_AppliedVirtualDestination) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MeshStatus_AppliedVirtualDestination)
	if !ok {
		that2, ok := that.(MeshStatus_AppliedVirtualDestination)
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

	if len(m.GetErrors()) != len(target.GetErrors()) {
		return false
	}
	for idx, v := range m.GetErrors() {

		if strings.Compare(v, target.GetErrors()[idx]) != 0 {
			return false
		}

	}

	return true
}
