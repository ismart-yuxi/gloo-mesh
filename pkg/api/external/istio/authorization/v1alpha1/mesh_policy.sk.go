// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"log"
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewMeshPolicy(namespace, name string) *MeshPolicy {
	meshpolicy := &MeshPolicy{}
	meshpolicy.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return meshpolicy
}

func (r *MeshPolicy) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *MeshPolicy) SetStatus(status core.Status) {
	r.Status = status
}

func (r *MeshPolicy) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	metaCopy.Generation = 0
	// investigate zeroing out owner refs as well
	return hashutils.HashAll(
		metaCopy,
		r.Targets,
		r.Peers,
		r.PeerIsOptional,
		r.Origins,
		r.OriginIsOptional,
		r.PrincipalBinding,
	)
}

func (r *MeshPolicy) GroupVersionKind() schema.GroupVersionKind {
	return MeshPolicyGVK
}

type MeshPolicyList []*MeshPolicy

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list MeshPolicyList) Find(namespace, name string) (*MeshPolicy, error) {
	for _, meshPolicy := range list {
		if meshPolicy.GetMetadata().Name == name {
			if namespace == "" || meshPolicy.GetMetadata().Namespace == namespace {
				return meshPolicy, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find meshPolicy %v.%v", namespace, name)
}

func (list MeshPolicyList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, meshPolicy := range list {
		ress = append(ress, meshPolicy)
	}
	return ress
}

func (list MeshPolicyList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, meshPolicy := range list {
		ress = append(ress, meshPolicy)
	}
	return ress
}

func (list MeshPolicyList) Names() []string {
	var names []string
	for _, meshPolicy := range list {
		names = append(names, meshPolicy.GetMetadata().Name)
	}
	return names
}

func (list MeshPolicyList) NamespacesDotNames() []string {
	var names []string
	for _, meshPolicy := range list {
		names = append(names, meshPolicy.GetMetadata().Namespace+"."+meshPolicy.GetMetadata().Name)
	}
	return names
}

func (list MeshPolicyList) Sort() MeshPolicyList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list MeshPolicyList) Clone() MeshPolicyList {
	var meshPolicyList MeshPolicyList
	for _, meshPolicy := range list {
		meshPolicyList = append(meshPolicyList, resources.Clone(meshPolicy).(*MeshPolicy))
	}
	return meshPolicyList
}

func (list MeshPolicyList) Each(f func(element *MeshPolicy)) {
	for _, meshPolicy := range list {
		f(meshPolicy)
	}
}

func (list MeshPolicyList) EachResource(f func(element resources.Resource)) {
	for _, meshPolicy := range list {
		f(meshPolicy)
	}
}

func (list MeshPolicyList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *MeshPolicy) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for MeshPolicy

func (o *MeshPolicy) GetObjectKind() schema.ObjectKind {
	t := MeshPolicyCrd.TypeMeta()
	return &t
}

func (o *MeshPolicy) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*MeshPolicy)
}

func (o *MeshPolicy) DeepCopyInto(out *MeshPolicy) {
	clone := resources.Clone(o).(*MeshPolicy)
	*out = *clone
}

var (
	MeshPolicyCrd = crd.NewCrd(
		"meshpolicies",
		MeshPolicyGVK.Group,
		MeshPolicyGVK.Version,
		MeshPolicyGVK.Kind,
		"meshpolicy",
		true,
		&MeshPolicy{})
)

func init() {
	if err := crd.AddCrd(MeshPolicyCrd); err != nil {
		log.Fatalf("could not add crd to global registry")
	}
}

var (
	MeshPolicyGVK = schema.GroupVersionKind{
		Version: "v1alpha1",
		Group:   "authentication.istio.io",
		Kind:    "MeshPolicy",
	}
)