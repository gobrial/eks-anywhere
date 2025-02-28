// +build e2e

package e2e

import (
	"testing"

	"github.com/aws/eks-anywhere/internal/pkg/api"
	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/test/framework"
)

func runWorkloadClusterFlow(test *framework.MulticlusterE2ETest) {
	test.CreateManagementCluster()
	test.RunInWorkloadClusters(func(w *framework.WorkloadCluster) {
		w.GenerateClusterConfig()
		w.CreateCluster()
		w.DeleteCluster()
	})
	test.DeleteManagementCluster()
}

func runWorkloadClusterFlowWithGitOps(test *framework.MulticlusterE2ETest, clusterOpts ...framework.ClusterE2ETestOpt) {
	test.CreateManagementCluster()
	test.RunInWorkloadClusters(func(w *framework.WorkloadCluster) {
		w.GenerateClusterConfig()
		w.CreateCluster()
		w.UpgradeWithGitOps(clusterOpts...)
		w.DeleteCluster()
	})
	test.DeleteManagementCluster()
}

func TestVSphereKubernetes121MulticlusterWorkloadClusterDemo(t *testing.T) {
	provider := framework.NewVSphere(t, framework.WithUbuntu121())
	test := framework.NewMulticlusterE2ETest(
		t,
		framework.NewClusterE2ETest(
			t,
			provider,
			framework.WithClusterFiller(
				api.WithKubernetesVersion(v1alpha1.Kube121),
				api.WithControlPlaneCount(1),
				api.WithWorkerNodeCount(1),
				api.WithStackedEtcdTopology(),
			),
		),
		framework.NewClusterE2ETest(
			t,
			provider,
			framework.WithClusterFiller(
				api.WithKubernetesVersion(v1alpha1.Kube121),
				api.WithControlPlaneCount(1),
				api.WithWorkerNodeCount(1),
				api.WithStackedEtcdTopology(),
			),
		),
	)
	runWorkloadClusterFlow(test)
}

func TestVSphereUpgradeMulticlusterWorkloadClusterWithFlux(t *testing.T) {
	provider := framework.NewVSphere(t, framework.WithUbuntu120())
	test := framework.NewMulticlusterE2ETest(
		t,
		framework.NewClusterE2ETest(
			t,
			provider,
			framework.WithFlux(),
			framework.WithClusterFiller(
				api.WithKubernetesVersion(v1alpha1.Kube120),
				api.WithControlPlaneCount(1),
				api.WithWorkerNodeCount(1),
				api.WithStackedEtcdTopology(),
			),
		),
		framework.NewClusterE2ETest(
			t,
			provider,
			framework.WithFlux(),
			framework.WithClusterFiller(
				api.WithKubernetesVersion(v1alpha1.Kube120),
				api.WithControlPlaneCount(1),
				api.WithWorkerNodeCount(1),
				api.WithStackedEtcdTopology(),
			),
		),
	)
	runWorkloadClusterFlowWithGitOps(
		test,
		framework.WithClusterUpgradeGit(
			api.WithKubernetesVersion(v1alpha1.Kube121),
			api.WithControlPlaneCount(3),
			api.WithWorkerNodeCount(3),
		),
		provider.WithProviderUpgradeGit(
			framework.UpdateUbuntuTemplate121Var(),
		),
	)
}

func TestDockerUpgradeWorkloadClusterWithFlux(t *testing.T) {
	provider := framework.NewDocker(t)
	test := framework.NewMulticlusterE2ETest(
		t,
		framework.NewClusterE2ETest(
			t,
			provider,
			framework.WithFlux(),
			framework.WithClusterFiller(
				api.WithKubernetesVersion(v1alpha1.Kube120),
				api.WithControlPlaneCount(1),
				api.WithWorkerNodeCount(1),
			),
		),
		framework.NewClusterE2ETest(
			t,
			provider,
			framework.WithFlux(),
			framework.WithClusterFiller(
				api.WithKubernetesVersion(v1alpha1.Kube120),
				api.WithControlPlaneCount(1),
				api.WithWorkerNodeCount(1),
			),
		),
	)
	runWorkloadClusterFlowWithGitOps(
		test,
		framework.WithClusterUpgradeGit(
			api.WithKubernetesVersion(v1alpha1.Kube121),
			api.WithControlPlaneCount(2),
			api.WithWorkerNodeCount(2),
		),
	)
}
