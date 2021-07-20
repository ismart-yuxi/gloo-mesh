// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/networking/v1/transformation/transformation.proto

package transformation

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	v1 "github.com/solo-io/gloo-mesh/pkg/api/common.mesh.gloo.solo.io/v1"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// GatewayTransformation enables use of the Transformation feature on Gateway.
type GatewayTransformations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Explicitly enable Transformations on a gateway. Only required if strict filter management is set on the gateway.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *GatewayTransformations) Reset() {
	*x = GatewayTransformations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayTransformations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayTransformations) ProtoMessage() {}

func (x *GatewayTransformations) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayTransformations.ProtoReflect.Descriptor instead.
func (*GatewayTransformations) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayTransformations) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// A RouteTransformation defines a text transformations for the content of an HTTP Request and/or Response on a matched route.
// Transformation takes the existing HTTP Headers (and, optionally, Body) and transforms them into a new set of headers (and body).
// Transformations can be used on outbound Request data as well as inbound Response data.
// Various types of transformations can be performed depending on the format of the input and output data types.
// Currently, Inja (for JSON-to-JSON transformations) and XSLT (for JSON-to-XML and XML-to-XML transformations) are currently supported.
// Transformations can optionally define a set of request/response HTTP match criteria. The first matched transformation in a list will
// be applied to the HTTP request/response.
type RouteTransformations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transformations to apply on the outbound HTTP request before it arrives at the routing Destination.
	// Only the first matched transformation will be applied to the request.
	Request []*RequestTransformation `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty"`
	// Transformations to apply on the inbound HTTP request before it returns to the HTTP client (traffic source).
	// Only the first matched transformation will be applied to the response.
	Response []*ResponseTransformation `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *RouteTransformations) Reset() {
	*x = RouteTransformations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTransformations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTransformations) ProtoMessage() {}

func (x *RouteTransformations) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTransformations.ProtoReflect.Descriptor instead.
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{1}
}

func (x *RouteTransformations) GetRequest() []*RequestTransformation {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RouteTransformations) GetResponse() []*ResponseTransformation {
	if x != nil {
		return x.Response
	}
	return nil
}

// match and transform the contents of an HTTP Request
type RequestTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify criteria that HTTP requests must satisfy for the RequestTransformation to apply
	// Omit to apply to any HTTP request.
	Matchers []*v1.HttpMatcher `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
	// the text transformation to apply to to the matched request
	Transformation *TextTransformation `protobuf:"bytes,2,opt,name=transformation,proto3" json:"transformation,omitempty"`
	// If the request was transformed such that it would match a different route within the same Gateway,
	// recalculate the routing destination (select a new route) based on the transformed content of the request.
	RecalculateRoutingDestination bool `protobuf:"varint,3,opt,name=recalculate_routing_destination,json=recalculateRoutingDestination,proto3" json:"recalculate_routing_destination,omitempty"`
	// Apply this transformation before Auth and Rate Limit checks are performed on the request.
	// This can be used to modify the request headers before they are captured by the ExtAuth & Rate Limiter services.
	ApplyBeforeAuth bool `protobuf:"varint,4,opt,name=apply_before_auth,json=applyBeforeAuth,proto3" json:"apply_before_auth,omitempty"`
}

func (x *RequestTransformation) Reset() {
	*x = RequestTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTransformation) ProtoMessage() {}

func (x *RequestTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTransformation.ProtoReflect.Descriptor instead.
func (*RequestTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{2}
}

func (x *RequestTransformation) GetMatchers() []*v1.HttpMatcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

func (x *RequestTransformation) GetTransformation() *TextTransformation {
	if x != nil {
		return x.Transformation
	}
	return nil
}

func (x *RequestTransformation) GetRecalculateRoutingDestination() bool {
	if x != nil {
		return x.RecalculateRoutingDestination
	}
	return false
}

func (x *RequestTransformation) GetApplyBeforeAuth() bool {
	if x != nil {
		return x.ApplyBeforeAuth
	}
	return false
}

// match and transform the contents of an HTTP Response
type ResponseTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Match elements of the Response in order to apply a response transformation.
	// If no response matchers are specified, the transformation will always be applied.
	Matchers []*ResponseTransformation_ResponseMatcher `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
	// the text transformation to apply to to the matched response
	Transformation *TextTransformation `protobuf:"bytes,2,opt,name=transformation,proto3" json:"transformation,omitempty"`
}

func (x *ResponseTransformation) Reset() {
	*x = ResponseTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseTransformation) ProtoMessage() {}

func (x *ResponseTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseTransformation.ProtoReflect.Descriptor instead.
func (*ResponseTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseTransformation) GetMatchers() []*ResponseTransformation_ResponseMatcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

func (x *ResponseTransformation) GetTransformation() *TextTransformation {
	if x != nil {
		return x.Transformation
	}
	return nil
}

// Transform the HTTP Headers/Body of a request / response using one of the supported mechanisms
type TextTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// each transformation can leverage one several available mechanisms for transforming HTTP data
	//
	// Types that are assignable to TransformationType:
	//	*TextTransformation_XsltTransformation
	TransformationType isTextTransformation_TransformationType `protobuf_oneof:"transformation_type"`
}

func (x *TextTransformation) Reset() {
	*x = TextTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextTransformation) ProtoMessage() {}

func (x *TextTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextTransformation.ProtoReflect.Descriptor instead.
func (*TextTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{4}
}

func (m *TextTransformation) GetTransformationType() isTextTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (x *TextTransformation) GetXsltTransformation() *XsltTransformation {
	if x, ok := x.GetTransformationType().(*TextTransformation_XsltTransformation); ok {
		return x.XsltTransformation
	}
	return nil
}

type isTextTransformation_TransformationType interface {
	isTextTransformation_TransformationType()
}

type TextTransformation_XsltTransformation struct {
	// transform HTTP body using XSLT styling language.
	XsltTransformation *XsltTransformation `protobuf:"bytes,2,opt,name=xslt_transformation,json=xsltTransformation,proto3,oneof"`
}

func (*TextTransformation_XsltTransformation) isTextTransformation_TransformationType() {}

// transform HTTP body and headers using Inja templates.
type InjaTemplateTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InjaTemplateTransformation) Reset() {
	*x = InjaTemplateTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InjaTemplateTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InjaTemplateTransformation) ProtoMessage() {}

func (x *InjaTemplateTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InjaTemplateTransformation.ProtoReflect.Descriptor instead.
func (*InjaTemplateTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{5}
}

// transform HTTP body using XSLT styling language.
type XsltTransformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// XSLT transformation template which you want to transform requests/responses with.
	// Invalid XSLT transformation templates will result will result in an invalid route.
	Xslt string `protobuf:"bytes,1,opt,name=xslt,proto3" json:"xslt,omitempty"`
	// Changes the content-type header of the HTTP request/response to what is set here.
	// This is useful in situations where an XSLT transformation is used to transform XML to JSON and the content-type
	// should be changed from `application/xml` to `application/json`.
	// If left empty, the content-type header remains unmodified by default.
	SetContentType string `protobuf:"bytes,2,opt,name=set_content_type,json=setContentType,proto3" json:"set_content_type,omitempty"`
	// This should be set to true if the content being transformed is not XML.
	// For example, if the content being transformed is from JSON to XML, this should be set to true.
	// XSLT transformations can only take valid XML as input to be transformed. If the body is not a valid XML
	// (e.g. using JSON as input in a JSON-to-XML transformation), setting `non_xml_transform` to true will allow the
	// XSLT to accept the non-XML input without throwing an error by passing the input as XML CDATA.
	// defaults to false.
	NonXmlTransform bool `protobuf:"varint,3,opt,name=non_xml_transform,json=nonXmlTransform,proto3" json:"non_xml_transform,omitempty"`
}

func (x *XsltTransformation) Reset() {
	*x = XsltTransformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XsltTransformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XsltTransformation) ProtoMessage() {}

func (x *XsltTransformation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XsltTransformation.ProtoReflect.Descriptor instead.
func (*XsltTransformation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{6}
}

func (x *XsltTransformation) GetXslt() string {
	if x != nil {
		return x.Xslt
	}
	return ""
}

func (x *XsltTransformation) GetSetContentType() string {
	if x != nil {
		return x.SetContentType
	}
	return ""
}

func (x *XsltTransformation) GetNonXmlTransform() bool {
	if x != nil {
		return x.NonXmlTransform
	}
	return false
}

// specifies a set of criteria for matching an HTTP response
type ResponseTransformation_ResponseMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specify a set of response headers which must match in entirety (all headers must match).
	Headers []*v1.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// Response code detail to match on. To see the response code details for your usecase,
	// you can use the envoy access log %RESPONSE_CODE_DETAILS% formatter to log it.
	ResponseCodeDetails string `protobuf:"bytes,2,opt,name=response_code_details,json=responseCodeDetails,proto3" json:"response_code_details,omitempty"`
}

func (x *ResponseTransformation_ResponseMatcher) Reset() {
	*x = ResponseTransformation_ResponseMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseTransformation_ResponseMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseTransformation_ResponseMatcher) ProtoMessage() {}

func (x *ResponseTransformation_ResponseMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseTransformation_ResponseMatcher.ProtoReflect.Descriptor instead.
func (*ResponseTransformation_ResponseMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ResponseTransformation_ResponseMatcher) GetHeaders() []*v1.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ResponseTransformation_ResponseMatcher) GetResponseCodeDetails() string {
	if x != nil {
		return x.ResponseCodeDetails
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x16, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xd5, 0x01, 0x0a,
	0x14, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41,
	0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x12, 0x67, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x1f, 0x72, 0x65,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x1d, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x22, 0xfd,
	0x02, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6f, 0x0a, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x67, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x88, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x9f,
	0x01, 0x0a, 0x12, 0x54, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a, 0x13, 0x78, 0x73, 0x6c, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x58, 0x73, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12, 0x78, 0x73, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x15, 0x0a, 0x13, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x1c, 0x0a, 0x1a, 0x49, 0x6e, 0x6a, 0x61, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7e,
	0x0a, 0x12, 0x58, 0x73, 0x6c, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x78, 0x73, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x78, 0x73, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x6f, 0x6e, 0x5f, 0x78, 0x6d, 0x6c, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6e,
	0x6f, 0x6e, 0x58, 0x6d, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x59,
	0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_goTypes = []interface{}{
	(*GatewayTransformations)(nil),                 // 0: transformation.networking.mesh.gloo.solo.io.GatewayTransformations
	(*RouteTransformations)(nil),                   // 1: transformation.networking.mesh.gloo.solo.io.RouteTransformations
	(*RequestTransformation)(nil),                  // 2: transformation.networking.mesh.gloo.solo.io.RequestTransformation
	(*ResponseTransformation)(nil),                 // 3: transformation.networking.mesh.gloo.solo.io.ResponseTransformation
	(*TextTransformation)(nil),                     // 4: transformation.networking.mesh.gloo.solo.io.TextTransformation
	(*InjaTemplateTransformation)(nil),             // 5: transformation.networking.mesh.gloo.solo.io.InjaTemplateTransformation
	(*XsltTransformation)(nil),                     // 6: transformation.networking.mesh.gloo.solo.io.XsltTransformation
	(*ResponseTransformation_ResponseMatcher)(nil), // 7: transformation.networking.mesh.gloo.solo.io.ResponseTransformation.ResponseMatcher
	(*v1.HttpMatcher)(nil),                         // 8: common.mesh.gloo.solo.io.HttpMatcher
	(*v1.HeaderMatcher)(nil),                       // 9: common.mesh.gloo.solo.io.HeaderMatcher
}
var file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_depIdxs = []int32{
	2, // 0: transformation.networking.mesh.gloo.solo.io.RouteTransformations.request:type_name -> transformation.networking.mesh.gloo.solo.io.RequestTransformation
	3, // 1: transformation.networking.mesh.gloo.solo.io.RouteTransformations.response:type_name -> transformation.networking.mesh.gloo.solo.io.ResponseTransformation
	8, // 2: transformation.networking.mesh.gloo.solo.io.RequestTransformation.matchers:type_name -> common.mesh.gloo.solo.io.HttpMatcher
	4, // 3: transformation.networking.mesh.gloo.solo.io.RequestTransformation.transformation:type_name -> transformation.networking.mesh.gloo.solo.io.TextTransformation
	7, // 4: transformation.networking.mesh.gloo.solo.io.ResponseTransformation.matchers:type_name -> transformation.networking.mesh.gloo.solo.io.ResponseTransformation.ResponseMatcher
	4, // 5: transformation.networking.mesh.gloo.solo.io.ResponseTransformation.transformation:type_name -> transformation.networking.mesh.gloo.solo.io.TextTransformation
	6, // 6: transformation.networking.mesh.gloo.solo.io.TextTransformation.xslt_transformation:type_name -> transformation.networking.mesh.gloo.solo.io.XsltTransformation
	9, // 7: transformation.networking.mesh.gloo.solo.io.ResponseTransformation.ResponseMatcher.headers:type_name -> common.mesh.gloo.solo.io.HeaderMatcher
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_init()
}
func file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayTransformations); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTransformations); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTransformation); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseTransformation); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextTransformation); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InjaTemplateTransformation); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XsltTransformation); i {
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
		file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseTransformation_ResponseMatcher); i {
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
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*TextTransformation_XsltTransformation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_networking_v1_transformation_transformation_proto_depIdxs = nil
}