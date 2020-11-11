package initialize

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/gobuffalo/packr"
	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo-mesh/codegen/helm"
	"github.com/solo-io/gloo-mesh/pkg/common/defaults"
	"github.com/solo-io/gloo-mesh/pkg/common/version"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/install/gloomesh"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/registration"
	"github.com/solo-io/skv2/pkg/multicluster/register"
)

const (
	// The default version of k8s under Linux is 1.18 https://github.com/solo-io/gloo-mesh/issues/700
	kindImage      = "kindest/node:v1.17.5"
	managementPort = "32001"
	remotePort     = "32000"
)

func createKindCluster(cluster string, port string, box packr.Box) error {
	fmt.Printf("Creating cluster %s with ingress port %s\n", cluster, port)

	script, err := box.FindString("create_kind_cluster.sh")
	if err != nil {
		return eris.Wrap(err, "Error loading script")
	}

	cmd := exec.Command("bash", "-c", script, cluster, port, kindImage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return eris.Wrapf(err, "Error creating cluster %s", cluster)
	}

	fmt.Printf("Successfully created cluster %s\n", cluster)
	return nil
}

func installGlooMesh(ctx context.Context, cluster string, box packr.Box) error {
	fmt.Printf("Deploying Gloo Mesh to %s from images\n", cluster)

	apiServerAddress, err := getApiAddress(cluster, box)
	if err != nil {
		return err
	}

	kubeConfig := ""
	kubeContext := fmt.Sprintf("kind-%s", cluster)

	namespace := defaults.DefaultPodNamespace
	verbose := true
	gloomeshChartUri := fmt.Sprintf(gloomesh.GlooMeshChartUriTemplate, version.Version)
	certAgentChartUri := fmt.Sprintf(gloomesh.CertAgentChartUriTemplate, version.Version)

	err = gloomesh.Installer{
		HelmChartPath:  gloomeshChartUri,
		HelmValuesPath: "",
		KubeConfig:     "",
		KubeContext:    kubeContext,
		Namespace:      namespace,
		ReleaseName:    helm.Chart.Data.Name,
		Verbose:        verbose,
		DryRun:         false,
	}.InstallGlooMesh(
		ctx,
	)
	if err != nil {
		return eris.Wrap(err, "Error installing Gloo Mesh")
	}

	registrantOpts := &registration.RegistrantOptions{
		KubeConfigPath: kubeConfig,
		MgmtContext:    kubeContext,
		RemoteContext:  kubeContext,
		Registration: register.RegistrationOptions{
			ClusterName:      cluster,
			RemoteCtx:        kubeContext,
			Namespace:        namespace,
			RemoteNamespace:  namespace,
			APIServerAddress: apiServerAddress,
			ClusterDomain:    "",
		},
		CertAgent: registration.AgentInstallOptions{
			ChartPath:   certAgentChartUri,
			ChartValues: "",
		},
		Verbose: verbose,
	}

	registrant, err := registration.NewRegistrant(registrantOpts)
	if err != nil {
		return eris.Wrapf(err, "initializing registrant for cluster %s", cluster)
	}
	err = registrant.RegisterCluster(ctx)
	if err != nil {
		return eris.Wrapf(err, "registering cluster %s", cluster)
	}

	script, err := box.FindString("post_install_gloomesh.sh")
	if err != nil {
		return eris.Wrap(err, "Error loading script")
	}

	cmd := exec.Command("bash", "-c", script, cluster)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return eris.Wrap(err, "Error running post-install script")
	}

	fmt.Printf("Successfully set up Gloo Mesh on cluster %s\n", cluster)
	return nil
}

func registerCluster(ctx context.Context, mgmtCluster string, cluster string, box packr.Box) error {
	fmt.Printf("Registering cluster %s with cert-agent image\n", cluster)

	apiServerAddress, err := getApiAddress(cluster, box)
	if err != nil {
		return err
	}

	kubeConfig := ""
	mgmtKubeContext := fmt.Sprintf("kind-%s", mgmtCluster)
	remoteKubeContext := fmt.Sprintf("kind-%s", cluster)

	namespace := defaults.DefaultPodNamespace
	certAgentChartUri := fmt.Sprintf(gloomesh.CertAgentChartUriTemplate, version.Version)

	registrantOpts := &registration.RegistrantOptions{
		KubeConfigPath: kubeConfig,
		MgmtContext:    mgmtKubeContext,
		RemoteContext:  remoteKubeContext,
		Registration: register.RegistrationOptions{
			ClusterName:      cluster,
			RemoteCtx:        remoteKubeContext,
			Namespace:        namespace,
			RemoteNamespace:  namespace,
			APIServerAddress: apiServerAddress,
			ClusterDomain:    "",
		},
		CertAgent: registration.AgentInstallOptions{
			ChartPath:   certAgentChartUri,
			ChartValues: "",
		},
		Verbose: true,
	}

	registrant, err := registration.NewRegistrant(registrantOpts)
	if err != nil {
		return eris.Wrapf(err, "initializing registrant for cluster %s", cluster)
	}
	err = registrant.RegisterCluster(ctx)
	if err != nil {
		return eris.Wrapf(err, "registering cluster %s", cluster)
	}

	fmt.Printf("Successfully registered cluster %s\n", cluster)
	return nil
}

func getApiAddress(cluster string, box packr.Box) (string, error) {
	script, err := box.FindString("get_api_address.sh")
	if err != nil {
		return "", eris.Wrap(err, "Error loading script")
	}
	cmd := exec.Command("bash", "-c", script, cluster)
	bytes, err := cmd.Output()
	if err != nil {
		return "", eris.Wrap(err, "Error getting API server address")
	}

	return string(bytes), nil
}

func switchContext(cluster string, box packr.Box) error {
	script, err := box.FindString("switch_context.sh")
	if err != nil {
		return err
	}
	cmd := exec.Command("bash", "-c", script, cluster)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return eris.Wrapf(err, "Could not switch context to %s", fmt.Sprintf("kind-%s", cluster))
	}
	return nil
}
