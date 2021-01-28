// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller



import (
	"context"

	networking_enterprise_mesh_gloo_solo_io_v1alpha1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the WasmDeployment Resource across clusters.
// implemented by the user
type MulticlusterWasmDeploymentReconciler interface {
	ReconcileWasmDeployment(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) (reconcile.Result, error)
}

// Reconcile deletion events for the WasmDeployment Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterWasmDeploymentDeletionReconciler interface {
	ReconcileWasmDeploymentDeletion(clusterName string, req reconcile.Request) error
}

type MulticlusterWasmDeploymentReconcilerFuncs struct {
	OnReconcileWasmDeployment         func(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) (reconcile.Result, error)
	OnReconcileWasmDeploymentDeletion func(clusterName string, req reconcile.Request) error
}

func (f *MulticlusterWasmDeploymentReconcilerFuncs) ReconcileWasmDeployment(clusterName string, obj *networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment) (reconcile.Result, error) {
	if f.OnReconcileWasmDeployment == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileWasmDeployment(clusterName, obj)
}

func (f *MulticlusterWasmDeploymentReconcilerFuncs) ReconcileWasmDeploymentDeletion(clusterName string, req reconcile.Request) error {
	if f.OnReconcileWasmDeploymentDeletion == nil {
		return nil
	}
	return f.OnReconcileWasmDeploymentDeletion(clusterName, req)
}

type MulticlusterWasmDeploymentReconcileLoop interface {
	// AddMulticlusterWasmDeploymentReconciler adds a MulticlusterWasmDeploymentReconciler to the MulticlusterWasmDeploymentReconcileLoop.
	AddMulticlusterWasmDeploymentReconciler(ctx context.Context, rec MulticlusterWasmDeploymentReconciler, predicates ...predicate.Predicate)
}

type multiclusterWasmDeploymentReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterWasmDeploymentReconcileLoop) AddMulticlusterWasmDeploymentReconciler(ctx context.Context, rec MulticlusterWasmDeploymentReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericWasmDeploymentMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterWasmDeploymentReconcileLoop(name string, cw multicluster.ClusterWatcher, options reconcile.Options) MulticlusterWasmDeploymentReconcileLoop {
	return &multiclusterWasmDeploymentReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment{}, options)}
}

type genericWasmDeploymentMulticlusterReconciler struct {
	reconciler MulticlusterWasmDeploymentReconciler
}

func (g genericWasmDeploymentMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) error {
	if deletionReconciler, ok := g.reconciler.(MulticlusterWasmDeploymentDeletionReconciler); ok {
		return deletionReconciler.ReconcileWasmDeploymentDeletion(cluster, req)
	}
	return nil
}

func (g genericWasmDeploymentMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_enterprise_mesh_gloo_solo_io_v1alpha1.WasmDeployment)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: WasmDeployment handler received event for %T", object)
	}
	return g.reconciler.ReconcileWasmDeployment(cluster, obj)
}
