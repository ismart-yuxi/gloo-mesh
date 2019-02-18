// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	istio_authentication_v1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
	istio_networking_v1alpha3 "github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"

	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"go.uber.org/zap"
)

type ConfigSnapshot struct {
	Meshes           MeshesByNamespace
	Meshgroups       MeshgroupsByNamespace
	Upstreams        gloo_solo_io.UpstreamsByNamespace
	Routingrules     RoutingrulesByNamespace
	Tlssecrets       TlssecretsByNamespace
	Destinationrules istio_networking_v1alpha3.DestinationrulesByNamespace
	Virtualservices  istio_networking_v1alpha3.VirtualservicesByNamespace
	Meshpolicies     istio_authentication_v1alpha1.MeshpoliciesByNamespace
}

func (s ConfigSnapshot) Clone() ConfigSnapshot {
	return ConfigSnapshot{
		Meshes:           s.Meshes.Clone(),
		Meshgroups:       s.Meshgroups.Clone(),
		Upstreams:        s.Upstreams.Clone(),
		Routingrules:     s.Routingrules.Clone(),
		Tlssecrets:       s.Tlssecrets.Clone(),
		Destinationrules: s.Destinationrules.Clone(),
		Virtualservices:  s.Virtualservices.Clone(),
		Meshpolicies:     s.Meshpolicies.Clone(),
	}
}

func (s ConfigSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashMeshes(),
		s.hashMeshgroups(),
		s.hashUpstreams(),
		s.hashRoutingrules(),
		s.hashTlssecrets(),
		s.hashDestinationrules(),
		s.hashVirtualservices(),
		s.hashMeshpolicies(),
	)
}

func (s ConfigSnapshot) hashMeshes() uint64 {
	return hashutils.HashAll(s.Meshes.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashMeshgroups() uint64 {
	return hashutils.HashAll(s.Meshgroups.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashRoutingrules() uint64 {
	return hashutils.HashAll(s.Routingrules.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashTlssecrets() uint64 {
	return hashutils.HashAll(s.Tlssecrets.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashDestinationrules() uint64 {
	return hashutils.HashAll(s.Destinationrules.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashVirtualservices() uint64 {
	return hashutils.HashAll(s.Virtualservices.List().AsInterfaces()...)
}

func (s ConfigSnapshot) hashMeshpolicies() uint64 {
	return hashutils.HashAll(s.Meshpolicies.List().AsInterfaces()...)
}

func (s ConfigSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("meshes", s.hashMeshes()))
	fields = append(fields, zap.Uint64("meshgroups", s.hashMeshgroups()))
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))
	fields = append(fields, zap.Uint64("routingrules", s.hashRoutingrules()))
	fields = append(fields, zap.Uint64("tlssecrets", s.hashTlssecrets()))
	fields = append(fields, zap.Uint64("destinationrules", s.hashDestinationrules()))
	fields = append(fields, zap.Uint64("virtualservices", s.hashVirtualservices()))
	fields = append(fields, zap.Uint64("meshpolicies", s.hashMeshpolicies()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}
