// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./appmesh_snapshot.go -destination mocks/appmesh_snapshot.go

// source: codegen/custom/appmesh/appmesh_snapshot.gotmpl

package appmesh

import (
	"context"
	"encoding/json"

	appmesh_aws_solo_io_v1alpha1 "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.aws.solo.io/v1alpha1"
	appmesh_aws_solo_io_v1alpha1_sets "github.com/solo-io/external-apis/pkg/api/appmesh/appmesh.aws.solo.io/v1alpha1/sets"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/contrib/pkg/sets"
)

// the snapshot of output resources produced by a translation
type Snapshot interface {
	// return the set of VirtualNodes with a given set of labels
	VirtualNodes() appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet
	// return the set of VirtualRouters with a given set of labels
	VirtualRouters() appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet
	// return the set of Routes with a given set of labels
	Routes() appmesh_aws_solo_io_v1alpha1_sets.RouteSet
	// return the set of VirtualServices with a given set of labels
	VirtualServices() appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet

	// return the set of Meshes the snapshot is built for
	Meshes() []string

	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

type snapshot struct {
	name            string
	virtualNodes    appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet
	virtualRouters  appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet
	routes          appmesh_aws_solo_io_v1alpha1_sets.RouteSet
	virtualServices appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet
	meshes          []string
}

func NewSnapshot(
	name string,
	virtualNodes appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet,
	virtualRouters appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet,
	routes appmesh_aws_solo_io_v1alpha1_sets.RouteSet,
	virtualServices appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet,
	meshes ...string, // the set of meshes to which to apply the snapshot
) Snapshot {
	return &snapshot{
		name:            name,
		virtualNodes:    virtualNodes,
		virtualRouters:  virtualRouters,
		routes:          routes,
		virtualServices: virtualServices,
		meshes:          meshes,
	}
}

func (s snapshot) VirtualNodes() appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet {
	return s.virtualNodes
}

func (s snapshot) VirtualRouters() appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet {
	return s.virtualRouters
}

func (s snapshot) Routes() appmesh_aws_solo_io_v1alpha1_sets.RouteSet {
	return s.routes
}

func (s snapshot) VirtualServices() appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet {
	return s.virtualServices
}

func (s snapshot) Meshes() []string {
	return s.meshes
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}
	snapshotMap["virtualNodes"] = s.virtualNodes.List()
	snapshotMap["virtualRouters"] = s.virtualRouters.List()
	snapshotMap["routes"] = s.routes.List()
	snapshotMap["virtualServices"] = s.virtualServices.List()

	snapshotMap["meshes"] = s.meshes

	return json.Marshal(snapshotMap)
}

type builder struct {
	ctx             context.Context
	name            string
	meshes          []string
	virtualNodes    appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet
	virtualRouters  appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet
	routes          appmesh_aws_solo_io_v1alpha1_sets.RouteSet
	virtualServices appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet
}

func NewBuilder(ctx context.Context, name string) *builder {
	return &builder{
		ctx:             ctx,
		name:            name,
		virtualNodes:    appmesh_aws_solo_io_v1alpha1_sets.NewVirtualNodeSet(),
		virtualRouters:  appmesh_aws_solo_io_v1alpha1_sets.NewVirtualRouterSet(),
		routes:          appmesh_aws_solo_io_v1alpha1_sets.NewRouteSet(),
		virtualServices: appmesh_aws_solo_io_v1alpha1_sets.NewVirtualServiceSet(),
	}
}

// the output Builder uses a builder pattern to allow
// iteratively collecting outputs before producing a final snapshot
type Builder interface {

	// add VirtualNodes to the collected outputs
	AddVirtualNodes(virtualNodes ...*appmesh_aws_solo_io_v1alpha1.VirtualNode)

	// get the collected VirtualNodes
	GetVirtualNodes() appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet

	// add VirtualRouters to the collected outputs
	AddVirtualRouters(virtualRouters ...*appmesh_aws_solo_io_v1alpha1.VirtualRouter)

	// get the collected VirtualRouters
	GetVirtualRouters() appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet

	// add Routes to the collected outputs
	AddRoutes(routes ...*appmesh_aws_solo_io_v1alpha1.Route)

	// get the collected Routes
	GetRoutes() appmesh_aws_solo_io_v1alpha1_sets.RouteSet

	// add VirtualServices to the collected outputs
	AddVirtualServices(virtualServices ...*appmesh_aws_solo_io_v1alpha1.VirtualService)

	// get the collected VirtualServices
	GetVirtualServices() appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet

	// build the collected outputs into a snapshot
	BuildSnapshot() Snapshot

	// add a mesh to the collected meshes.
	// this can be used to collect meshes for use with MultiMesh snapshots.
	AddMesh(mesh string)
}

func (b *builder) AddVirtualNodes(virtualNodes ...*appmesh_aws_solo_io_v1alpha1.VirtualNode) {
	for _, obj := range virtualNodes {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output VirtualNode %v", sets.Key(obj))
		b.virtualNodes.Insert(obj)
	}
}

func (b *builder) GetVirtualNodes() appmesh_aws_solo_io_v1alpha1_sets.VirtualNodeSet {
	return b.virtualNodes
}

func (b *builder) AddVirtualRouters(virtualRouters ...*appmesh_aws_solo_io_v1alpha1.VirtualRouter) {
	for _, obj := range virtualRouters {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output VirtualRouter %v", sets.Key(obj))
		b.virtualRouters.Insert(obj)
	}
}

func (b *builder) GetVirtualRouters() appmesh_aws_solo_io_v1alpha1_sets.VirtualRouterSet {
	return b.virtualRouters
}

func (b *builder) AddRoutes(routes ...*appmesh_aws_solo_io_v1alpha1.Route) {
	for _, obj := range routes {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output Route %v", sets.Key(obj))
		b.routes.Insert(obj)
	}
}

func (b *builder) GetRoutes() appmesh_aws_solo_io_v1alpha1_sets.RouteSet {
	return b.routes
}

func (b *builder) AddVirtualServices(virtualServices ...*appmesh_aws_solo_io_v1alpha1.VirtualService) {
	for _, obj := range virtualServices {
		if obj == nil {
			continue
		}
		contextutils.LoggerFrom(b.ctx).Debugf("added output VirtualService %v", sets.Key(obj))
		b.virtualServices.Insert(obj)
	}
}

func (b *builder) GetVirtualServices() appmesh_aws_solo_io_v1alpha1_sets.VirtualServiceSet {
	return b.virtualServices
}

func (b *builder) BuildSnapshot() Snapshot {
	return NewSnapshot(
		b.name,
		b.virtualNodes,
		b.virtualRouters,
		b.routes,
		b.virtualServices,
		b.meshes...,
	)
}

func (b *builder) AddMesh(mesh string) {
	b.meshes = append(b.meshes, mesh)
}
