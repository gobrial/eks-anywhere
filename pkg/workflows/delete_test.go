package workflows_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/aws/eks-anywhere/internal/test"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	providermocks "github.com/aws/eks-anywhere/pkg/providers/mocks"
	"github.com/aws/eks-anywhere/pkg/types"
	"github.com/aws/eks-anywhere/pkg/workflows"
	"github.com/aws/eks-anywhere/pkg/workflows/interfaces/mocks"
)

type deleteTestSetup struct {
	t                *testing.T
	bootstrapper     *mocks.MockBootstrapper
	clusterManager   *mocks.MockClusterManager
	addonManager     *mocks.MockAddonManager
	provider         *providermocks.MockProvider
	workflow         *workflows.Delete
	ctx              context.Context
	clusterSpec      *cluster.Spec
	forceCleanup     bool
	bootstrapCluster *types.Cluster
	workloadCluster  *types.Cluster
}

func newDeleteTest(t *testing.T) *deleteTestSetup {
	mockCtrl := gomock.NewController(t)
	mockBootstrapper := mocks.NewMockBootstrapper(mockCtrl)
	clusterManager := mocks.NewMockClusterManager(mockCtrl)
	addonManager := mocks.NewMockAddonManager(mockCtrl)
	provider := providermocks.NewMockProvider(mockCtrl)
	workflow := workflows.NewDelete(mockBootstrapper, provider, clusterManager, addonManager)

	return &deleteTestSetup{
		t:                t,
		bootstrapper:     mockBootstrapper,
		clusterManager:   clusterManager,
		addonManager:     addonManager,
		provider:         provider,
		workflow:         workflow,
		ctx:              context.Background(),
		clusterSpec:      test.NewClusterSpec(func(s *cluster.Spec) { s.Cluster.Name = "cluster-name" }),
		bootstrapCluster: &types.Cluster{Name: "bootstrap"},
		workloadCluster:  &types.Cluster{Name: "workload"},
	}
}

func (c *deleteTestSetup) expectSetup() {
	c.provider.EXPECT().SetupAndValidateDeleteCluster(c.ctx)
}

func (c *deleteTestSetup) expectCreateBootstrap() {
	opts := []bootstrapper.BootstrapClusterOption{
		bootstrapper.WithDefaultCNIDisabled(), bootstrapper.WithExtraDockerMounts(),
	}

	gomock.InOrder(
		c.provider.EXPECT().BootstrapClusterOpts().Return(opts, nil),
		c.bootstrapper.EXPECT().CreateBootstrapCluster(
			c.ctx, gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
		).Return(c.bootstrapCluster, nil),

		c.clusterManager.EXPECT().InstallCAPI(c.ctx, gomock.Not(gomock.Nil()), c.bootstrapCluster, c.provider),
	)
}

func (c *deleteTestSetup) expectNotToCreateBootstrap() {
	opts := []bootstrapper.BootstrapClusterOption{
		bootstrapper.WithDefaultCNIDisabled(), bootstrapper.WithExtraDockerMounts(),
	}

	c.provider.EXPECT().BootstrapClusterOpts().Return(opts, nil).Times(0)
	c.bootstrapper.EXPECT().CreateBootstrapCluster(
		c.ctx, gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil()),
	).Return(c.bootstrapCluster, nil).Times(0)

	c.clusterManager.EXPECT().InstallCAPI(c.ctx, gomock.Not(gomock.Nil()), c.bootstrapCluster, c.provider).Times(0)
}

func (c *deleteTestSetup) expectDeleteBootstrap() {
	gomock.InOrder(
		c.bootstrapper.EXPECT().DeleteBootstrapCluster(
			c.ctx, c.bootstrapCluster,
			gomock.Any()).Return(nil),
	)
}

func (c *deleteTestSetup) expectNotToDeleteBootstrap() {
	c.bootstrapper.EXPECT().DeleteBootstrapCluster(c.ctx, c.bootstrapCluster, gomock.Any()).Return(nil).Times(0)
}

func (c *deleteTestSetup) expectDeleteWorkload(cluster *types.Cluster) {
	gomock.InOrder(
		c.clusterManager.EXPECT().DeleteCluster(
			c.ctx, cluster, c.workloadCluster, c.provider, c.clusterSpec,
		).Return(nil),
	)
}

func (c *deleteTestSetup) expectCleanupGitRepo() {
	gomock.InOrder(
		c.addonManager.EXPECT().CleanupGitRepo(
			c.ctx, c.clusterSpec,
		).Return(nil),
	)
}

func (c *deleteTestSetup) expectMoveManagement() {
	gomock.InOrder(
		c.clusterManager.EXPECT().MoveCAPI(
			c.ctx, c.workloadCluster, c.bootstrapCluster, c.workloadCluster.Name, c.clusterSpec, gomock.Any(),
		),
	)
}

func (c *deleteTestSetup) expectNotToMoveManagement() {
	gomock.InOrder(
		c.clusterManager.EXPECT().MoveCAPI(
			c.ctx, c.workloadCluster, c.bootstrapCluster, c.workloadCluster.Name, gomock.Any(),
		).Times(0),
	)
}

func (c *deleteTestSetup) run() error {
	// ctx context.Context, workloadCluster *types.Cluster, forceCleanup bool
	return c.workflow.Run(c.ctx, c.workloadCluster, c.clusterSpec, c.forceCleanup, "")
}

func TestDeleteRunSuccess(t *testing.T) {
	test := newDeleteTest(t)
	test.expectSetup()
	test.expectCreateBootstrap()
	test.expectDeleteWorkload(test.bootstrapCluster)
	test.expectCleanupGitRepo()
	test.expectMoveManagement()
	test.expectDeleteBootstrap()

	err := test.run()
	if err != nil {
		t.Fatalf("Delete.Run() err = %v, want err = nil", err)
	}
}

func TestDeleteWorkloadRunSuccess(t *testing.T) {
	test := newDeleteTest(t)
	test.expectSetup()
	test.expectNotToCreateBootstrap()
	test.clusterSpec.ManagementCluster = &types.Cluster{
		Name:               "management-cluster",
		KubeconfigFile:     "kc.kubeconfig",
		ExistingManagement: true,
	}
	test.clusterSpec.Cluster.SetManagedBy(test.clusterSpec.ManagementCluster.Name)
	test.expectDeleteWorkload(test.clusterSpec.ManagementCluster)
	test.expectCleanupGitRepo()
	test.expectNotToMoveManagement()
	test.expectNotToDeleteBootstrap()

	err := test.run()
	if err != nil {
		t.Fatalf("Delete.Run() err = %v, want err = nil", err)
	}
}
