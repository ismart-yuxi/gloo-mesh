// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/v1/transformation/transformation.proto

package transformation

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
func (m *GatewayTransformations) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayTransformations)
	if !ok {
		that2, ok := that.(GatewayTransformations)
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

	if m.GetEnabled() != target.GetEnabled() {
		return false
	}

	return true
}

// Equal function
func (m *RouteTransformations) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteTransformations)
	if !ok {
		that2, ok := that.(RouteTransformations)
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

	if len(m.GetRequest()) != len(target.GetRequest()) {
		return false
	}
	for idx, v := range m.GetRequest() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequest()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRequest()[idx]) {
				return false
			}
		}

	}

	if len(m.GetResponse()) != len(target.GetResponse()) {
		return false
	}
	for idx, v := range m.GetResponse() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponse()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResponse()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *RequestTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RequestTransformation)
	if !ok {
		that2, ok := that.(RequestTransformation)
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

	if len(m.GetMatchers()) != len(target.GetMatchers()) {
		return false
	}
	for idx, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchers()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransformation(), target.GetTransformation()) {
			return false
		}
	}

	if m.GetRecalculateRoutingDestination() != target.GetRecalculateRoutingDestination() {
		return false
	}

	if m.GetApplyBeforeAuth() != target.GetApplyBeforeAuth() {
		return false
	}

	return true
}

// Equal function
func (m *ResponseTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResponseTransformation)
	if !ok {
		that2, ok := that.(ResponseTransformation)
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

	if len(m.GetMatchers()) != len(target.GetMatchers()) {
		return false
	}
	for idx, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchers()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransformation(), target.GetTransformation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TextTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TextTransformation)
	if !ok {
		that2, ok := that.(TextTransformation)
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

	switch m.TransformationType.(type) {

	case *TextTransformation_XsltTransformation:
		if _, ok := target.TransformationType.(*TextTransformation_XsltTransformation); !ok {
			return false
		}

		if h, ok := interface{}(m.GetXsltTransformation()).(equality.Equalizer); ok {
			if !h.Equal(target.GetXsltTransformation()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetXsltTransformation(), target.GetXsltTransformation()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.TransformationType != target.TransformationType {
			return false
		}
	}

	return true
}

// Equal function
func (m *InjaTemplateTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*InjaTemplateTransformation)
	if !ok {
		that2, ok := that.(InjaTemplateTransformation)
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

	return true
}

// Equal function
func (m *XsltTransformation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*XsltTransformation)
	if !ok {
		that2, ok := that.(XsltTransformation)
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

	if strings.Compare(m.GetXslt(), target.GetXslt()) != 0 {
		return false
	}

	if strings.Compare(m.GetSetContentType(), target.GetSetContentType()) != 0 {
		return false
	}

	if m.GetNonXmlTransform() != target.GetNonXmlTransform() {
		return false
	}

	return true
}

// Equal function
func (m *ResponseTransformation_ResponseMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResponseTransformation_ResponseMatcher)
	if !ok {
		that2, ok := that.(ResponseTransformation_ResponseMatcher)
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

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for idx, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeaders()[idx]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetResponseCodeDetails(), target.GetResponseCodeDetails()) != 0 {
		return false
	}

	return true
}