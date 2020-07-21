package main

import (
	"log"

	externalapis "github.com/solo-io/external-apis/codegen"
	"github.com/solo-io/service-mesh-hub/codegen/groups"
	"github.com/solo-io/service-mesh-hub/codegen/helm"
	"github.com/solo-io/service-mesh-hub/codegen/io"
	skv1alpha1 "github.com/solo-io/skv2/api/multicluster/v1alpha1"
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/contrib"
	"github.com/solo-io/solo-kit/pkg/code-generator/sk_anyvendor"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

var (
	appName = "service-mesh-hub"

	discoveryInputSnapshotCodePath  = "pkg/api/discovery.smh.solo.io/snapshot/input/snapshot.go"
	discoveryReconcilerCodePath     = "pkg/api/discovery.smh.solo.io/snapshot/input/reconciler.go"
	discoveryOutputSnapshotCodePath = "pkg/api/discovery.smh.solo.io/snapshot/output/snapshot.go"

	networkingInputSnapshotCodePath       = "pkg/api/networking.smh.solo.io/snapshot/input/snapshot.go"
	networkingReconcilerSnapshotCodePath  = "pkg/api/networking.smh.solo.io/snapshot/input/reconciler.go"
	networkingOutputIstioSnapshotCodePath = "pkg/api/networking.smh.solo.io/snapshot/output/istio/snapshot.go"

	smhManifestRoot = "install/helm/charts/service-mesh-hub"
	csrManifestRoot = "install/helm/charts/csr-agent/"

	allApiGroups = map[string][]model.Group{
		"":                                 groups.SMHGroups,
		"github.com/solo-io/external-apis": externalapis.Groups,
		"github.com/solo-io/skv2":          []model.Group{skv1alpha1.Group},
	}

	// define custom templates
	discoveryInputSnapshot = makeTopLevelTemplate(
		contrib.InputSnapshot,
		discoveryInputSnapshotCodePath,
		io.DiscoveryInputTypes,
	)

	discoveryReconciler = makeTopLevelTemplate(
		contrib.InputReconciler,
		discoveryReconcilerCodePath,
		io.DiscoveryInputTypes,
	)

	discoveryOutputSnapshot = makeTopLevelTemplate(
		contrib.OutputSnapshot,
		discoveryOutputSnapshotCodePath,
		io.DiscoveryOutputTypes,
	)

	networkingInputSnapshot = makeTopLevelTemplate(
		contrib.InputSnapshot,
		networkingInputSnapshotCodePath,
		io.NetworkingInputTypes,
	)

	networkingReconciler = makeTopLevelTemplate(
		contrib.InputReconciler,
		networkingReconcilerSnapshotCodePath,
		io.NetworkingInputTypes,
	)

	networkingOutputIstioSnapshot = makeTopLevelTemplate(
		contrib.OutputSnapshot,
		networkingOutputIstioSnapshotCodePath,
		io.NetworkingOutputIstioTypes,
	)

	topLevelTemplates = []model.CustomTemplates{
		discoveryInputSnapshot,
		discoveryReconciler,
		discoveryOutputSnapshot,
		networkingInputSnapshot,
		networkingReconciler,
		networkingOutputIstioSnapshot,
	}

	protoImports = sk_anyvendor.CreateDefaultMatchOptions([]string{
		"api/**/*.proto",
	})
)

func run() error {
	log.Printf("generating smh")
	if err := makeSmhCommand().Execute(); err != nil {
		return err
	}
	return nil
	// TODO(ilackarms): revive this when reviving csr agent
	log.Printf("generating csr-agent")
	if err := makeCsrCommand().Execute(); err != nil {
		return err
	}
	return nil
}

func makeSmhCommand() codegen.Command {

	protoImports.External["github.com/solo-io/skv2"] = []string{
		"api/**/*.proto",
	}

	return codegen.Command{
		AppName:           appName,
		AnyVendorConfig:   protoImports,
		ManifestRoot:      smhManifestRoot,
		TopLevelTemplates: topLevelTemplates,
		Groups:            groups.SMHGroups,
		RenderProtos:      true,
		Chart:             helm.Chart,
	}
}

func makeCsrCommand() codegen.Command {
	return codegen.Command{
		AppName:         appName,
		AnyVendorConfig: protoImports,
		ManifestRoot:    csrManifestRoot,
		Groups:          groups.CSRGroups,
		RenderProtos:    true,
	}
}

func makeTopLevelTemplate(templateFunc func(params contrib.CrossGroupTemplateParameters) model.CustomTemplates, outPath string, resourceSnapshot io.Snapshot) model.CustomTemplates {
	return templateFunc(contrib.CrossGroupTemplateParameters{
		OutputFilename:    outPath,
		SelectFromGroups:  allApiGroups,
		ResourcesToSelect: resourceSnapshot,
	})
}