// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	admin_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/admin.enterprise.mesh.gloo.solo.io/v1alpha1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type IstioInstallationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation) bool) []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation) bool) []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation
	// Return the Set as a map of key to resource.
	Map() map[string]*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation
	// Insert a resource into the set.
	Insert(istioInstallation ...*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(istioInstallationSet IstioInstallationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(istioInstallation ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(istioInstallation ezkube.ResourceId)
	// Return the union with the provided set
	Union(set IstioInstallationSet) IstioInstallationSet
	// Return the difference with the provided set
	Difference(set IstioInstallationSet) IstioInstallationSet
	// Return the intersection with the provided set
	Intersection(set IstioInstallationSet) IstioInstallationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another IstioInstallationSet
	Delta(newSet IstioInstallationSet) sksets.ResourceDelta
	// Create a deep copy of the current IstioInstallationSet
	Clone() IstioInstallationSet
}

func makeGenericIstioInstallationSet(istioInstallationList []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range istioInstallationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type istioInstallationSet struct {
	set sksets.ResourceSet
}

func NewIstioInstallationSet(istioInstallationList ...*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation) IstioInstallationSet {
	return &istioInstallationSet{set: makeGenericIstioInstallationSet(istioInstallationList)}
}

func NewIstioInstallationSetFromList(istioInstallationList *admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallationList) IstioInstallationSet {
	list := make([]*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation, 0, len(istioInstallationList.Items))
	for idx := range istioInstallationList.Items {
		list = append(list, &istioInstallationList.Items[idx])
	}
	return &istioInstallationSet{set: makeGenericIstioInstallationSet(list)}
}

func (s *istioInstallationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *istioInstallationSet) List(filterResource ...func(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation) bool) []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation))
		})
	}

	objs := s.Generic().List(genericFilters...)
	istioInstallationList := make([]*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation, 0, len(objs))
	for _, obj := range objs {
		istioInstallationList = append(istioInstallationList, obj.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation))
	}
	return istioInstallationList
}

func (s *istioInstallationSet) UnsortedList(filterResource ...func(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation) bool) []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation))
		})
	}

	var istioInstallationList []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		istioInstallationList = append(istioInstallationList, obj.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation))
	}
	return istioInstallationList
}

func (s *istioInstallationSet) Map() map[string]*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation {
	if s == nil {
		return nil
	}

	newMap := map[string]*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation)
	}
	return newMap
}

func (s *istioInstallationSet) Insert(
	istioInstallationList ...*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range istioInstallationList {
		s.Generic().Insert(obj)
	}
}

func (s *istioInstallationSet) Has(istioInstallation ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(istioInstallation)
}

func (s *istioInstallationSet) Equal(
	istioInstallationSet IstioInstallationSet,
) bool {
	if s == nil {
		return istioInstallationSet == nil
	}
	return s.Generic().Equal(istioInstallationSet.Generic())
}

func (s *istioInstallationSet) Delete(IstioInstallation ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(IstioInstallation)
}

func (s *istioInstallationSet) Union(set IstioInstallationSet) IstioInstallationSet {
	if s == nil {
		return set
	}
	return NewIstioInstallationSet(append(s.List(), set.List()...)...)
}

func (s *istioInstallationSet) Difference(set IstioInstallationSet) IstioInstallationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &istioInstallationSet{set: newSet}
}

func (s *istioInstallationSet) Intersection(set IstioInstallationSet) IstioInstallationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var istioInstallationList []*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation
	for _, obj := range newSet.List() {
		istioInstallationList = append(istioInstallationList, obj.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation))
	}
	return NewIstioInstallationSet(istioInstallationList...)
}

func (s *istioInstallationSet) Find(id ezkube.ResourceId) (*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find IstioInstallation %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*admin_enterprise_mesh_gloo_solo_io_v1alpha1.IstioInstallation), nil
}

func (s *istioInstallationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *istioInstallationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *istioInstallationSet) Delta(newSet IstioInstallationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

func (s *istioInstallationSet) Clone() IstioInstallationSet {
	if s == nil {
		return nil
	}
	return &istioInstallationSet{set: sksets.NewResourceSet(s.Generic().Clone().List()...)}
}
