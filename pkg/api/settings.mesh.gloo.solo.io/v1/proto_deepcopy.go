// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto for the Settings.Spec
func (in *SettingsSpec) DeepCopyInto(out *SettingsSpec) {
	p := proto.Clone(in).(*SettingsSpec)
	*out = *p
}

// DeepCopyInto for the Settings.Status
func (in *SettingsStatus) DeepCopyInto(out *SettingsStatus) {
	p := proto.Clone(in).(*SettingsStatus)
	*out = *p
}

// DeepCopyInto for the Dashboard.Spec
func (in *DashboardSpec) DeepCopyInto(out *DashboardSpec) {
	p := proto.Clone(in).(*DashboardSpec)
	*out = *p
}

// DeepCopyInto for the Dashboard.Status
func (in *DashboardStatus) DeepCopyInto(out *DashboardStatus) {
	p := proto.Clone(in).(*DashboardStatus)
	*out = *p
}
