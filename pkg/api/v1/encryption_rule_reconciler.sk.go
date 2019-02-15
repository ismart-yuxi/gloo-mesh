// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionEncryptionRuleFunc func(original, desired *EncryptionRule) (bool, error)

type EncryptionRuleReconciler interface {
	Reconcile(namespace string, desiredResources EncryptionRuleList, transition TransitionEncryptionRuleFunc, opts clients.ListOpts) error
}

func encryptionRulesToResources(list EncryptionRuleList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, encryptionRule := range list {
		resourceList = append(resourceList, encryptionRule)
	}
	return resourceList
}

func NewEncryptionRuleReconciler(client EncryptionRuleClient) EncryptionRuleReconciler {
	return &encryptionRuleReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type encryptionRuleReconciler struct {
	base reconcile.Reconciler
}

func (r *encryptionRuleReconciler) Reconcile(namespace string, desiredResources EncryptionRuleList, transition TransitionEncryptionRuleFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "encryptionRule_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*EncryptionRule), desired.(*EncryptionRule))
		}
	}
	return r.base.Reconcile(namespace, encryptionRulesToResources(desiredResources), transitionResources, opts)
}
