// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/enterprise/networking/v1beta1/virtual_gateway.proto

package v1beta1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"

	v1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
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

	_ = v1.ApprovalState(0)
)

// Equal function
func (m *VirtualGatewaySpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec)
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

	if len(m.GetIngressGatewaySelectors()) != len(target.GetIngressGatewaySelectors()) {
		return false
	}
	for idx, v := range m.GetIngressGatewaySelectors() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetIngressGatewaySelectors()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetIngressGatewaySelectors()[idx]) {
				return false
			}
		}

	}

	if len(m.GetConnectionHandlers()) != len(target.GetConnectionHandlers()) {
		return false
	}
	for idx, v := range m.GetConnectionHandlers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConnectionHandlers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConnectionHandlers()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewayStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewayStatus)
	if !ok {
		that2, ok := that.(VirtualGatewayStatus)
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

	if m.GetState() != target.GetState() {
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

	if len(m.GetAppliedIngressGateways()) != len(target.GetAppliedIngressGateways()) {
		return false
	}
	for idx, v := range m.GetAppliedIngressGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAppliedIngressGateways()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAppliedIngressGateways()[idx]) {
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

	if len(m.GetSelectedRouteTables()) != len(target.GetSelectedRouteTables()) {
		return false
	}
	for idx, v := range m.GetSelectedRouteTables() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelectedRouteTables()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSelectedRouteTables()[idx]) {
				return false
			}
		}

	}

	if len(m.GetCreatedIstioGateways()) != len(target.GetCreatedIstioGateways()) {
		return false
	}
	for k, v := range m.GetCreatedIstioGateways() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetCreatedIstioGateways()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetCreatedIstioGateways()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *SslConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SslConfig)
	if !ok {
		that2, ok := that.(SslConfig)
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

	if len(m.GetVerifySubjectAltName()) != len(target.GetVerifySubjectAltName()) {
		return false
	}
	for idx, v := range m.GetVerifySubjectAltName() {

		if strings.Compare(v, target.GetVerifySubjectAltName()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetParameters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetParameters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetParameters(), target.GetParameters()) {
			return false
		}
	}

	if m.GetTlsMode() != target.GetTlsMode() {
		return false
	}

	switch m.Certificates.(type) {

	case *SslConfig_SecretName:
		if _, ok := target.Certificates.(*SslConfig_SecretName); !ok {
			return false
		}

		if strings.Compare(m.GetSecretName(), target.GetSecretName()) != 0 {
			return false
		}

	case *SslConfig_SslFiles:
		if _, ok := target.Certificates.(*SslConfig_SslFiles); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSslFiles()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslFiles()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSslFiles(), target.GetSslFiles()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Certificates != target.Certificates {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler)
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

	if h, ok := interface{}(m.GetConnectionMatch()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionMatch()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionMatch(), target.GetConnectionMatch()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConnectionOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionOptions(), target.GetConnectionOptions()) {
			return false
		}
	}

	switch m.HandlerType.(type) {

	case *VirtualGatewaySpec_ConnectionHandler_Http:
		if _, ok := target.HandlerType.(*VirtualGatewaySpec_ConnectionHandler_Http); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttp()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttp()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttp(), target.GetHttp()) {
				return false
			}
		}

	case *VirtualGatewaySpec_ConnectionHandler_Tcp:
		if _, ok := target.HandlerType.(*VirtualGatewaySpec_ConnectionHandler_Tcp); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcp()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcp()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcp(), target.GetTcp()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.HandlerType != target.HandlerType {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_GatewayOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_GatewayOptions)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_GatewayOptions)
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

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerConnectionBufferLimitBytes(), target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	}

	if strings.Compare(m.GetBindAddress(), target.GetBindAddress()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_ConnectionMatch) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_ConnectionMatch)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_ConnectionMatch)
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

	if len(m.GetServerNames()) != len(target.GetServerNames()) {
		return false
	}
	for idx, v := range m.GetServerNames() {

		if strings.Compare(v, target.GetServerNames()[idx]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetTransportProtocol(), target.GetTransportProtocol()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_ConnectionOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_ConnectionOptions)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_ConnectionOptions)
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

	if m.GetStrictFilterManagement() != target.GetStrictFilterManagement() {
		return false
	}

	if m.GetEnableProxyProtocol() != target.GetEnableProxyProtocol() {
		return false
	}

	switch m.SslSettings.(type) {

	case *VirtualGatewaySpec_ConnectionHandler_ConnectionOptions_SslConfig:
		if _, ok := target.SslSettings.(*VirtualGatewaySpec_ConnectionHandler_ConnectionOptions_SslConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
				return false
			}
		}

	case *VirtualGatewaySpec_ConnectionHandler_ConnectionOptions_HttpsRedirect:
		if _, ok := target.SslSettings.(*VirtualGatewaySpec_ConnectionHandler_ConnectionOptions_HttpsRedirect); !ok {
			return false
		}

		if m.GetHttpsRedirect() != target.GetHttpsRedirect() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.SslSettings != target.SslSettings {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_HttpRoutes) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_HttpRoutes)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_HttpRoutes)
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

	if len(m.GetRouteConfig()) != len(target.GetRouteConfig()) {
		return false
	}
	for idx, v := range m.GetRouteConfig() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteConfig()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRouteConfig()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRouteOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteOptions(), target.GetRouteOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_TcpRoutes) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_TcpRoutes)
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

	if len(m.GetTcpHosts()) != len(target.GetTcpHosts()) {
		return false
	}
	for idx, v := range m.GetTcpHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTcpHosts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier)
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

	switch m.RouteType.(type) {

	case *VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_VirtualHostSelector:
		if _, ok := target.RouteType.(*VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_VirtualHostSelector); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVirtualHostSelector()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualHostSelector()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVirtualHostSelector(), target.GetVirtualHostSelector()) {
				return false
			}
		}

	case *VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_VirtualHost:
		if _, ok := target.RouteType.(*VirtualGatewaySpec_ConnectionHandler_HttpRoutes_RouteSpecifier_VirtualHost); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVirtualHost()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualHost()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVirtualHost(), target.GetVirtualHost()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RouteType != target.RouteType {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_HttpRoutes_HttpOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_HttpRoutes_HttpOptions)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_HttpRoutes_HttpOptions)
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

	if h, ok := interface{}(m.GetCsrf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCsrf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCsrf(), target.GetCsrf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRateLimit()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRateLimit()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRateLimit(), target.GetRateLimit()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetDestination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestination(), target.GetDestination()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions)
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

	if h, ok := interface{}(m.GetTcpProxySettings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcpProxySettings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcpProxySettings(), target.GetTcpProxySettings()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction)
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

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	switch m.Destination.(type) {

	case *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Static:
		if _, ok := target.Destination.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Static); !ok {
			return false
		}

		if h, ok := interface{}(m.GetStatic()).(equality.Equalizer); ok {
			if !h.Equal(target.GetStatic()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetStatic(), target.GetStatic()) {
				return false
			}
		}

	case *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Virtual:
		if _, ok := target.Destination.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Virtual); !ok {
			return false
		}

		if h, ok := interface{}(m.GetVirtual()).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtual()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetVirtual(), target.GetVirtual()) {
				return false
			}
		}

	case *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Kube:
		if _, ok := target.Destination.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_Kube); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKube()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKube()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKube(), target.GetKube()) {
				return false
			}
		}

	case *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_ForwardSniClusterName:
		if _, ok := target.Destination.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpHost_TcpAction_ForwardSniClusterName); !ok {
			return false
		}

		if h, ok := interface{}(m.GetForwardSniClusterName()).(equality.Equalizer); ok {
			if !h.Equal(target.GetForwardSniClusterName()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetForwardSniClusterName(), target.GetForwardSniClusterName()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Destination != target.Destination {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings)
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

	if h, ok := interface{}(m.GetMaxConnectAttempts()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxConnectAttempts()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxConnectAttempts(), target.GetMaxConnectAttempts()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTunnelingConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTunnelingConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTunnelingConfig(), target.GetTunnelingConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings_TunnelingConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings_TunnelingConfig)
	if !ok {
		that2, ok := that.(VirtualGatewaySpec_ConnectionHandler_TcpRoutes_TcpOptions_TcpProxySettings_TunnelingConfig)
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

	if strings.Compare(m.GetHostname(), target.GetHostname()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *SslConfig_SSLFiles) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SslConfig_SSLFiles)
	if !ok {
		that2, ok := that.(SslConfig_SSLFiles)
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

	if strings.Compare(m.GetTlsCert(), target.GetTlsCert()) != 0 {
		return false
	}

	if strings.Compare(m.GetTlsKey(), target.GetTlsKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetRootCa(), target.GetRootCa()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *SslConfig_SslParameters) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SslConfig_SslParameters)
	if !ok {
		that2, ok := that.(SslConfig_SslParameters)
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

	if m.GetMinimumProtocolVersion() != target.GetMinimumProtocolVersion() {
		return false
	}

	if m.GetMaximumProtocolVersion() != target.GetMaximumProtocolVersion() {
		return false
	}

	if len(m.GetCipherSuites()) != len(target.GetCipherSuites()) {
		return false
	}
	for idx, v := range m.GetCipherSuites() {

		if strings.Compare(v, target.GetCipherSuites()[idx]) != 0 {
			return false
		}

	}

	return true
}
