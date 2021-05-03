// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/settings/v1/settings.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	v11 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configure system-wide settings and defaults. Settings specified in networking policies take precedence over those specified here.
type SettingsSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configure default mTLS settings for Destinations.
	Mtls *v1.TrafficPolicySpec_Policy_MTLS `protobuf:"bytes,1,opt,name=mtls,proto3" json:"mtls,omitempty"`
	// Configure Gloo Mesh networking to communicate with one or more external gRPC NetworkingExtensions servers.
	// Updates will be applied by the servers in the order they are listed (servers towards the end of the list take precedence).
	// Note: Extension Servers have full write access to the output objects written by Gloo Mesh.
	NetworkingExtensionServers []*GrpcServer `protobuf:"bytes,2,rep,name=networking_extension_servers,json=networkingExtensionServers,proto3" json:"networking_extension_servers,omitempty"`
	// Settings for Gloo Mesh discovery.
	Discovery *DiscoverySettings `protobuf:"bytes,3,opt,name=discovery,proto3" json:"discovery,omitempty"`
	// Enable and configure use of Relay mode to communicate with remote clusters. This is an enterprise-only feature.
	Relay *RelaySettings `protobuf:"bytes,4,opt,name=relay,proto3" json:"relay,omitempty"`
}

func (x *SettingsSpec) Reset() {
	*x = SettingsSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsSpec) ProtoMessage() {}

func (x *SettingsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsSpec.ProtoReflect.Descriptor instead.
func (*SettingsSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsSpec) GetMtls() *v1.TrafficPolicySpec_Policy_MTLS {
	if x != nil {
		return x.Mtls
	}
	return nil
}

func (x *SettingsSpec) GetNetworkingExtensionServers() []*GrpcServer {
	if x != nil {
		return x.NetworkingExtensionServers
	}
	return nil
}

func (x *SettingsSpec) GetDiscovery() *DiscoverySettings {
	if x != nil {
		return x.Discovery
	}
	return nil
}

func (x *SettingsSpec) GetRelay() *RelaySettings {
	if x != nil {
		return x.Relay
	}
	return nil
}

// RelaySettings contains options for configuring Gloo Mesh to use Relay for cluster management.
// Relay provides a way for connecting Gloo Mesh to remote Kubernetes Clusters
// without the need to share credentials and access to remote Kube API Servers
// from the management cluster (the Gloo Mesh controllers).
//
// Relay instead uses a streaming gRPC API to pass discovery data
// from remote clusters to the management cluster, and push
// configuration from the management cluster to the remote clusters.
//
// Architecturally, it includes a Relay-agent which is installed to remote Kube clusters at
// registration time, which then connects directly to the Relay Server in the management cluster.
// to push its discovery data and pull its mesh configuration.
//
//
// To configure Gloo Mesh to use Relay, make sure to read the
// [relay installation guide]({{< versioned_link_path fromRoot="/guides/setup/install_gloo_mesh" >}}) and
// [relay cluster registration guide]({{< versioned_link_path fromRoot="/guides/setup/register_cluster" >}}).
type RelaySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enable the use of Relay for cluster management.
	// If relay is enabled, make sure to follow the [relay cluster registration guide]({{< versioned_link_path fromRoot="/guides/setup/register_cluster#relay" >}})
	// for registering your clusters.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Connection info for the Relay Server. Gloo Mesh will fetch discovery resources from this server
	// and push translated outputs to this server.
	// Note: currently this field has no effect as the relay server runs in-process of the networking Pod.
	Server *GrpcServer `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *RelaySettings) Reset() {
	*x = RelaySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelaySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelaySettings) ProtoMessage() {}

func (x *RelaySettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelaySettings.ProtoReflect.Descriptor instead.
func (*RelaySettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{1}
}

func (x *RelaySettings) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *RelaySettings) GetServer() *GrpcServer {
	if x != nil {
		return x.Server
	}
	return nil
}

// Settings for Gloo Mesh discovery.
type DiscoverySettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Istio-specific discovery settings
	Istio *DiscoverySettings_Istio `protobuf:"bytes,1,opt,name=istio,proto3" json:"istio,omitempty"`
}

func (x *DiscoverySettings) Reset() {
	*x = DiscoverySettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverySettings) ProtoMessage() {}

func (x *DiscoverySettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverySettings.ProtoReflect.Descriptor instead.
func (*DiscoverySettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverySettings) GetIstio() *DiscoverySettings_Istio {
	if x != nil {
		return x.Istio
	}
	return nil
}

// Options for connecting to an external gRPC server.
type GrpcServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TCP address of the gRPC Server (including port).
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// If true communicate over HTTP rather than HTTPS.
	Insecure bool `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	// If true Gloo Mesh will automatically attempt to reconnect to the server after encountering network failures.
	ReconnectOnNetworkFailures bool `protobuf:"varint,3,opt,name=reconnect_on_network_failures,json=reconnectOnNetworkFailures,proto3" json:"reconnect_on_network_failures,omitempty"`
}

func (x *GrpcServer) Reset() {
	*x = GrpcServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcServer) ProtoMessage() {}

func (x *GrpcServer) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcServer.ProtoReflect.Descriptor instead.
func (*GrpcServer) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{3}
}

func (x *GrpcServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GrpcServer) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *GrpcServer) GetReconnectOnNetworkFailures() bool {
	if x != nil {
		return x.ReconnectOnNetworkFailures
	}
	return false
}

type SettingsStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the Settings metadata.
	// If the `observedGeneration` does not match `metadata.generation`, Gloo Mesh has not processed the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	// It will only show accepted if no processing errors encountered.
	State v11.ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=common.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// Any errors encountered while processing Settings object.
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *SettingsStatus) Reset() {
	*x = SettingsStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsStatus) ProtoMessage() {}

func (x *SettingsStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsStatus.ProtoReflect.Descriptor instead.
func (*SettingsStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{4}
}

func (x *SettingsStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *SettingsStatus) GetState() v11.ApprovalState {
	if x != nil {
		return x.State
	}
	return v11.ApprovalState_PENDING
}

func (x *SettingsStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

// Istio-specific discovery settings
type DiscoverySettings_Istio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configure discovery of ingress gateways per cluster. The key to the map is either a Gloo Mesh cluster name or
	// `*` denoting all clusters. If an entry is found for a given cluster, it will be used. Otherwise, the
	// wildcard entry will be used if it exists. Lastly, we will fall back to a set of default values.
	IngressGatewayDetectors map[string]*DiscoverySettings_Istio_IngressGatewayDetector `protobuf:"bytes,1,rep,name=ingress_gateway_detectors,json=ingressGatewayDetectors,proto3" json:"ingress_gateway_detectors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DiscoverySettings_Istio) Reset() {
	*x = DiscoverySettings_Istio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverySettings_Istio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverySettings_Istio) ProtoMessage() {}

func (x *DiscoverySettings_Istio) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverySettings_Istio.ProtoReflect.Descriptor instead.
func (*DiscoverySettings_Istio) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DiscoverySettings_Istio) GetIngressGatewayDetectors() map[string]*DiscoverySettings_Istio_IngressGatewayDetector {
	if x != nil {
		return x.IngressGatewayDetectors
	}
	return nil
}

// Configure discovery of ingress gateways.
type DiscoverySettings_Istio_IngressGatewayDetector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workload labels used to detect ingress gateways for an Istio deployment.
	// If not specified, will default to `{"istio": "ingressgateway"}`.
	GatewayWorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=gateway_workload_labels,json=gatewayWorkloadLabels,proto3" json:"gateway_workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the TLS port used to detect ingress gateways. Kubernetes services must have a port with this name
	// in order to be recognized as an ingress gateway. If not specified, will default to `tls`.
	GatewayTlsPortName string `protobuf:"bytes,2,opt,name=gateway_tls_port_name,json=gatewayTlsPortName,proto3" json:"gateway_tls_port_name,omitempty"`
}

func (x *DiscoverySettings_Istio_IngressGatewayDetector) Reset() {
	*x = DiscoverySettings_Istio_IngressGatewayDetector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverySettings_Istio_IngressGatewayDetector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverySettings_Istio_IngressGatewayDetector) ProtoMessage() {}

func (x *DiscoverySettings_Istio_IngressGatewayDetector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverySettings_Istio_IngressGatewayDetector.ProtoReflect.Descriptor instead.
func (*DiscoverySettings_Istio_IngressGatewayDetector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP(), []int{2, 0, 1}
}

func (x *DiscoverySettings_Istio_IngressGatewayDetector) GetGatewayWorkloadLabels() map[string]string {
	if x != nil {
		return x.GatewayWorkloadLabels
	}
	return nil
}

func (x *DiscoverySettings_Istio_IngressGatewayDetector) GetGatewayTlsPortName() string {
	if x != nil {
		return x.GatewayTlsPortName
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4f, 0x0a, 0x04, 0x6d, 0x74, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x54, 0x4c,
	0x53, 0x52, 0x04, 0x6d, 0x74, 0x6c, 0x73, 0x12, 0x68, 0x0a, 0x1c, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x1a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x4b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x3f,
	0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x22,
	0x69, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xc8, 0x05, 0x0a, 0x11, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x49, 0x0a, 0x05, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x49,
	0x73, 0x74, 0x69, 0x6f, 0x52, 0x05, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x1a, 0xe7, 0x04, 0x0a, 0x05,
	0x49, 0x73, 0x74, 0x69, 0x6f, 0x12, 0x8c, 0x01, 0x0a, 0x19, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x1a, 0x96, 0x01, 0x0a, 0x1c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x60, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xb5, 0x02,
	0x0a, 0x16, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x9d, 0x01, 0x0a, 0x17, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x65, 0x2e, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x54, 0x6c, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x48, 0x0a, 0x1a, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x85, 0x01, 0x0a, 0x0a, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1a, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4f, 0x6e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x22, 0x98, 0x01,
	0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x48, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0xc0, 0xf5,
	0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_goTypes = []interface{}{
	(*SettingsSpec)(nil),            // 0: settings.mesh.gloo.solo.io.SettingsSpec
	(*RelaySettings)(nil),           // 1: settings.mesh.gloo.solo.io.RelaySettings
	(*DiscoverySettings)(nil),       // 2: settings.mesh.gloo.solo.io.DiscoverySettings
	(*GrpcServer)(nil),              // 3: settings.mesh.gloo.solo.io.GrpcServer
	(*SettingsStatus)(nil),          // 4: settings.mesh.gloo.solo.io.SettingsStatus
	(*DiscoverySettings_Istio)(nil), // 5: settings.mesh.gloo.solo.io.DiscoverySettings.Istio
	nil,                             // 6: settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetectorsEntry
	(*DiscoverySettings_Istio_IngressGatewayDetector)(nil), // 7: settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetector
	nil,                                      // 8: settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetector.GatewayWorkloadLabelsEntry
	(*v1.TrafficPolicySpec_Policy_MTLS)(nil), // 9: networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS
	(v11.ApprovalState)(0),                   // 10: common.mesh.gloo.solo.io.ApprovalState
}
var file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_depIdxs = []int32{
	9,  // 0: settings.mesh.gloo.solo.io.SettingsSpec.mtls:type_name -> networking.mesh.gloo.solo.io.TrafficPolicySpec.Policy.MTLS
	3,  // 1: settings.mesh.gloo.solo.io.SettingsSpec.networking_extension_servers:type_name -> settings.mesh.gloo.solo.io.GrpcServer
	2,  // 2: settings.mesh.gloo.solo.io.SettingsSpec.discovery:type_name -> settings.mesh.gloo.solo.io.DiscoverySettings
	1,  // 3: settings.mesh.gloo.solo.io.SettingsSpec.relay:type_name -> settings.mesh.gloo.solo.io.RelaySettings
	3,  // 4: settings.mesh.gloo.solo.io.RelaySettings.server:type_name -> settings.mesh.gloo.solo.io.GrpcServer
	5,  // 5: settings.mesh.gloo.solo.io.DiscoverySettings.istio:type_name -> settings.mesh.gloo.solo.io.DiscoverySettings.Istio
	10, // 6: settings.mesh.gloo.solo.io.SettingsStatus.state:type_name -> common.mesh.gloo.solo.io.ApprovalState
	6,  // 7: settings.mesh.gloo.solo.io.DiscoverySettings.Istio.ingress_gateway_detectors:type_name -> settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetectorsEntry
	7,  // 8: settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetectorsEntry.value:type_name -> settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetector
	8,  // 9: settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetector.gateway_workload_labels:type_name -> settings.mesh.gloo.solo.io.DiscoverySettings.Istio.IngressGatewayDetector.GatewayWorkloadLabelsEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelaySettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverySettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcServer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverySettings_Istio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverySettings_Istio_IngressGatewayDetector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_settings_v1_settings_proto_depIdxs = nil
}
