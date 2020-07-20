// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	networking_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type TrafficPolicySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_smh_solo_io_v1alpha1.TrafficPolicy) bool) []*networking_smh_solo_io_v1alpha1.TrafficPolicy
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_smh_solo_io_v1alpha1.TrafficPolicy
	// Insert a resource into the set.
	Insert(trafficPolicy ...*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(trafficPolicySet TrafficPolicySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(trafficPolicy ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(trafficPolicy ezkube.ResourceId)
	// Return the union with the provided set
	Union(set TrafficPolicySet) TrafficPolicySet
	// Return the difference with the provided set
	Difference(set TrafficPolicySet) TrafficPolicySet
	// Return the intersection with the provided set
	Intersection(set TrafficPolicySet) TrafficPolicySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_smh_solo_io_v1alpha1.TrafficPolicy, error)
	// Get the length of the set
	Length() int
}

func makeGenericTrafficPolicySet(trafficPolicyList []*networking_smh_solo_io_v1alpha1.TrafficPolicy) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range trafficPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type trafficPolicySet struct {
	set sksets.ResourceSet
}

func NewTrafficPolicySet(trafficPolicyList ...*networking_smh_solo_io_v1alpha1.TrafficPolicy) TrafficPolicySet {
	return &trafficPolicySet{set: makeGenericTrafficPolicySet(trafficPolicyList)}
}

func NewTrafficPolicySetFromList(trafficPolicyList *networking_smh_solo_io_v1alpha1.TrafficPolicyList) TrafficPolicySet {
	list := make([]*networking_smh_solo_io_v1alpha1.TrafficPolicy, 0, len(trafficPolicyList.Items))
	for idx := range trafficPolicyList.Items {
		list = append(list, &trafficPolicyList.Items[idx])
	}
	return &trafficPolicySet{set: makeGenericTrafficPolicySet(list)}
}

func (s *trafficPolicySet) Keys() sets.String {
	return s.set.Keys()
}

func (s *trafficPolicySet) List(filterResource ...func(*networking_smh_solo_io_v1alpha1.TrafficPolicy) bool) []*networking_smh_solo_io_v1alpha1.TrafficPolicy {

	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_smh_solo_io_v1alpha1.TrafficPolicy))
		})
	}

	var trafficPolicyList []*networking_smh_solo_io_v1alpha1.TrafficPolicy
	for _, obj := range s.set.List(genericFilters...) {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_smh_solo_io_v1alpha1.TrafficPolicy))
	}
	return trafficPolicyList
}

func (s *trafficPolicySet) Map() map[string]*networking_smh_solo_io_v1alpha1.TrafficPolicy {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.TrafficPolicy{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	}
	return newMap
}

func (s *trafficPolicySet) Insert(
	trafficPolicyList ...*networking_smh_solo_io_v1alpha1.TrafficPolicy,
) {
	for _, obj := range trafficPolicyList {
		s.set.Insert(obj)
	}
}

func (s *trafficPolicySet) Has(trafficPolicy ezkube.ResourceId) bool {
	return s.set.Has(trafficPolicy)
}

func (s *trafficPolicySet) Equal(
	trafficPolicySet TrafficPolicySet,
) bool {
	return s.set.Equal(makeGenericTrafficPolicySet(trafficPolicySet.List()))
}

func (s *trafficPolicySet) Delete(TrafficPolicy ezkube.ResourceId) {
	s.set.Delete(TrafficPolicy)
}

func (s *trafficPolicySet) Union(set TrafficPolicySet) TrafficPolicySet {
	return NewTrafficPolicySet(append(s.List(), set.List()...)...)
}

func (s *trafficPolicySet) Difference(set TrafficPolicySet) TrafficPolicySet {
	newSet := s.set.Difference(makeGenericTrafficPolicySet(set.List()))
	return &trafficPolicySet{set: newSet}
}

func (s *trafficPolicySet) Intersection(set TrafficPolicySet) TrafficPolicySet {
	newSet := s.set.Intersection(makeGenericTrafficPolicySet(set.List()))
	var trafficPolicyList []*networking_smh_solo_io_v1alpha1.TrafficPolicy
	for _, obj := range newSet.List() {
		trafficPolicyList = append(trafficPolicyList, obj.(*networking_smh_solo_io_v1alpha1.TrafficPolicy))
	}
	return NewTrafficPolicySet(trafficPolicyList...)
}

func (s *trafficPolicySet) Find(id ezkube.ResourceId) (*networking_smh_solo_io_v1alpha1.TrafficPolicy, error) {
	obj, err := s.set.Find(&networking_smh_solo_io_v1alpha1.TrafficPolicy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_smh_solo_io_v1alpha1.TrafficPolicy), nil
}

func (s *trafficPolicySet) Length() int {
	return s.set.Length()
}

type AccessPolicySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_smh_solo_io_v1alpha1.AccessPolicy) bool) []*networking_smh_solo_io_v1alpha1.AccessPolicy
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_smh_solo_io_v1alpha1.AccessPolicy
	// Insert a resource into the set.
	Insert(accessPolicy ...*networking_smh_solo_io_v1alpha1.AccessPolicy)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(accessPolicySet AccessPolicySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(accessPolicy ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(accessPolicy ezkube.ResourceId)
	// Return the union with the provided set
	Union(set AccessPolicySet) AccessPolicySet
	// Return the difference with the provided set
	Difference(set AccessPolicySet) AccessPolicySet
	// Return the intersection with the provided set
	Intersection(set AccessPolicySet) AccessPolicySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_smh_solo_io_v1alpha1.AccessPolicy, error)
	// Get the length of the set
	Length() int
}

func makeGenericAccessPolicySet(accessPolicyList []*networking_smh_solo_io_v1alpha1.AccessPolicy) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range accessPolicyList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type accessPolicySet struct {
	set sksets.ResourceSet
}

func NewAccessPolicySet(accessPolicyList ...*networking_smh_solo_io_v1alpha1.AccessPolicy) AccessPolicySet {
	return &accessPolicySet{set: makeGenericAccessPolicySet(accessPolicyList)}
}

func NewAccessPolicySetFromList(accessPolicyList *networking_smh_solo_io_v1alpha1.AccessPolicyList) AccessPolicySet {
	list := make([]*networking_smh_solo_io_v1alpha1.AccessPolicy, 0, len(accessPolicyList.Items))
	for idx := range accessPolicyList.Items {
		list = append(list, &accessPolicyList.Items[idx])
	}
	return &accessPolicySet{set: makeGenericAccessPolicySet(list)}
}

func (s *accessPolicySet) Keys() sets.String {
	return s.set.Keys()
}

func (s *accessPolicySet) List(filterResource ...func(*networking_smh_solo_io_v1alpha1.AccessPolicy) bool) []*networking_smh_solo_io_v1alpha1.AccessPolicy {

	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_smh_solo_io_v1alpha1.AccessPolicy))
		})
	}

	var accessPolicyList []*networking_smh_solo_io_v1alpha1.AccessPolicy
	for _, obj := range s.set.List(genericFilters...) {
		accessPolicyList = append(accessPolicyList, obj.(*networking_smh_solo_io_v1alpha1.AccessPolicy))
	}
	return accessPolicyList
}

func (s *accessPolicySet) Map() map[string]*networking_smh_solo_io_v1alpha1.AccessPolicy {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.AccessPolicy{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.AccessPolicy)
	}
	return newMap
}

func (s *accessPolicySet) Insert(
	accessPolicyList ...*networking_smh_solo_io_v1alpha1.AccessPolicy,
) {
	for _, obj := range accessPolicyList {
		s.set.Insert(obj)
	}
}

func (s *accessPolicySet) Has(accessPolicy ezkube.ResourceId) bool {
	return s.set.Has(accessPolicy)
}

func (s *accessPolicySet) Equal(
	accessPolicySet AccessPolicySet,
) bool {
	return s.set.Equal(makeGenericAccessPolicySet(accessPolicySet.List()))
}

func (s *accessPolicySet) Delete(AccessPolicy ezkube.ResourceId) {
	s.set.Delete(AccessPolicy)
}

func (s *accessPolicySet) Union(set AccessPolicySet) AccessPolicySet {
	return NewAccessPolicySet(append(s.List(), set.List()...)...)
}

func (s *accessPolicySet) Difference(set AccessPolicySet) AccessPolicySet {
	newSet := s.set.Difference(makeGenericAccessPolicySet(set.List()))
	return &accessPolicySet{set: newSet}
}

func (s *accessPolicySet) Intersection(set AccessPolicySet) AccessPolicySet {
	newSet := s.set.Intersection(makeGenericAccessPolicySet(set.List()))
	var accessPolicyList []*networking_smh_solo_io_v1alpha1.AccessPolicy
	for _, obj := range newSet.List() {
		accessPolicyList = append(accessPolicyList, obj.(*networking_smh_solo_io_v1alpha1.AccessPolicy))
	}
	return NewAccessPolicySet(accessPolicyList...)
}

func (s *accessPolicySet) Find(id ezkube.ResourceId) (*networking_smh_solo_io_v1alpha1.AccessPolicy, error) {
	obj, err := s.set.Find(&networking_smh_solo_io_v1alpha1.AccessPolicy{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_smh_solo_io_v1alpha1.AccessPolicy), nil
}

func (s *accessPolicySet) Length() int {
	return s.set.Length()
}

type VirtualMeshSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_smh_solo_io_v1alpha1.VirtualMesh) bool) []*networking_smh_solo_io_v1alpha1.VirtualMesh
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_smh_solo_io_v1alpha1.VirtualMesh
	// Insert a resource into the set.
	Insert(virtualMesh ...*networking_smh_solo_io_v1alpha1.VirtualMesh)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualMeshSet VirtualMeshSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualMesh ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualMesh ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualMeshSet) VirtualMeshSet
	// Return the difference with the provided set
	Difference(set VirtualMeshSet) VirtualMeshSet
	// Return the intersection with the provided set
	Intersection(set VirtualMeshSet) VirtualMeshSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_smh_solo_io_v1alpha1.VirtualMesh, error)
	// Get the length of the set
	Length() int
}

func makeGenericVirtualMeshSet(virtualMeshList []*networking_smh_solo_io_v1alpha1.VirtualMesh) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualMeshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualMeshSet struct {
	set sksets.ResourceSet
}

func NewVirtualMeshSet(virtualMeshList ...*networking_smh_solo_io_v1alpha1.VirtualMesh) VirtualMeshSet {
	return &virtualMeshSet{set: makeGenericVirtualMeshSet(virtualMeshList)}
}

func NewVirtualMeshSetFromList(virtualMeshList *networking_smh_solo_io_v1alpha1.VirtualMeshList) VirtualMeshSet {
	list := make([]*networking_smh_solo_io_v1alpha1.VirtualMesh, 0, len(virtualMeshList.Items))
	for idx := range virtualMeshList.Items {
		list = append(list, &virtualMeshList.Items[idx])
	}
	return &virtualMeshSet{set: makeGenericVirtualMeshSet(list)}
}

func (s *virtualMeshSet) Keys() sets.String {
	return s.set.Keys()
}

func (s *virtualMeshSet) List(filterResource ...func(*networking_smh_solo_io_v1alpha1.VirtualMesh) bool) []*networking_smh_solo_io_v1alpha1.VirtualMesh {

	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_smh_solo_io_v1alpha1.VirtualMesh))
		})
	}

	var virtualMeshList []*networking_smh_solo_io_v1alpha1.VirtualMesh
	for _, obj := range s.set.List(genericFilters...) {
		virtualMeshList = append(virtualMeshList, obj.(*networking_smh_solo_io_v1alpha1.VirtualMesh))
	}
	return virtualMeshList
}

func (s *virtualMeshSet) Map() map[string]*networking_smh_solo_io_v1alpha1.VirtualMesh {
	newMap := map[string]*networking_smh_solo_io_v1alpha1.VirtualMesh{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	}
	return newMap
}

func (s *virtualMeshSet) Insert(
	virtualMeshList ...*networking_smh_solo_io_v1alpha1.VirtualMesh,
) {
	for _, obj := range virtualMeshList {
		s.set.Insert(obj)
	}
}

func (s *virtualMeshSet) Has(virtualMesh ezkube.ResourceId) bool {
	return s.set.Has(virtualMesh)
}

func (s *virtualMeshSet) Equal(
	virtualMeshSet VirtualMeshSet,
) bool {
	return s.set.Equal(makeGenericVirtualMeshSet(virtualMeshSet.List()))
}

func (s *virtualMeshSet) Delete(VirtualMesh ezkube.ResourceId) {
	s.set.Delete(VirtualMesh)
}

func (s *virtualMeshSet) Union(set VirtualMeshSet) VirtualMeshSet {
	return NewVirtualMeshSet(append(s.List(), set.List()...)...)
}

func (s *virtualMeshSet) Difference(set VirtualMeshSet) VirtualMeshSet {
	newSet := s.set.Difference(makeGenericVirtualMeshSet(set.List()))
	return &virtualMeshSet{set: newSet}
}

func (s *virtualMeshSet) Intersection(set VirtualMeshSet) VirtualMeshSet {
	newSet := s.set.Intersection(makeGenericVirtualMeshSet(set.List()))
	var virtualMeshList []*networking_smh_solo_io_v1alpha1.VirtualMesh
	for _, obj := range newSet.List() {
		virtualMeshList = append(virtualMeshList, obj.(*networking_smh_solo_io_v1alpha1.VirtualMesh))
	}
	return NewVirtualMeshSet(virtualMeshList...)
}

func (s *virtualMeshSet) Find(id ezkube.ResourceId) (*networking_smh_solo_io_v1alpha1.VirtualMesh, error) {
	obj, err := s.set.Find(&networking_smh_solo_io_v1alpha1.VirtualMesh{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_smh_solo_io_v1alpha1.VirtualMesh), nil
}

func (s *virtualMeshSet) Length() int {
	return s.set.Length()
}
