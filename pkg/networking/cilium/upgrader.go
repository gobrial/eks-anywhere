package cilium

import (
	"context"
	"fmt"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/logger"
	"github.com/aws/eks-anywhere/pkg/types"
)

type upgraderClient interface {
	Apply(ctx context.Context, cluster *types.Cluster, data []byte) error
	Delete(ctx context.Context, cluster *types.Cluster, data []byte) error
	WaitForPreflightDaemonSet(ctx context.Context, cluster *types.Cluster) error
	WaitForPreflightDeployment(ctx context.Context, cluster *types.Cluster) error
	WaitForCiliumDaemonSet(ctx context.Context, cluster *types.Cluster) error
	WaitForCiliumDeployment(ctx context.Context, cluster *types.Cluster) error
}

type Upgrader struct {
	templater *Templater
	client    upgraderClient
}

func NewUpgrader(client Client, helm Helm) *Upgrader {
	return &Upgrader{
		templater: NewTemplater(helm),
		client:    newRetrier(client),
	}
}

func (u *Upgrader) Upgrade(ctx context.Context, cluster *types.Cluster, currentSpec, newSpec *cluster.Spec) (*types.ChangeDiff, error) {
	diff := ciliumChangeDiff(currentSpec, newSpec)
	if diff == nil {
		logger.V(1).Info("Nothing to upgrade for Cilium, skipping")
		return nil, nil
	}

	logger.V(1).Info("Upgrading Cilium", "oldVersion", diff.ComponentReports[0].OldVersion, "newVersion", diff.ComponentReports[0].NewVersion)
	logger.V(4).Info("Generating Cilium upgrade preflight manifest")
	preflight, err := u.templater.GenerateUpgradePreflightManifest(ctx, newSpec)
	if err != nil {
		return nil, err
	}

	logger.V(2).Info("Installing Cilium upgrade preflight manifest")
	if err := u.client.Apply(ctx, cluster, preflight); err != nil {
		return nil, fmt.Errorf("failed applying cilium preflight check: %v", err)
	}

	logger.V(3).Info("Waiting for Cilium upgrade preflight checks to be up")
	if err := u.waitForPreflight(ctx, cluster); err != nil {
		return nil, err
	}

	logger.V(3).Info("Deleting Cilium upgrade preflight")
	if err := u.client.Delete(ctx, cluster, preflight); err != nil {
		return nil, fmt.Errorf("failed deleting cilium preflight check: %v", err)
	}

	logger.V(3).Info("Generating Cilium upgrade manifest")
	upgradeManifest, err := u.templater.GenerateUpgradeManifest(ctx, currentSpec, newSpec)
	if err != nil {
		return nil, err
	}

	logger.V(2).Info("Installing new Cilium version")
	if err := u.client.Apply(ctx, cluster, upgradeManifest); err != nil {
		return nil, fmt.Errorf("failed applying cilium upgrade: %v", err)
	}

	logger.V(3).Info("Waiting for upgraded Cilium to be ready")
	if err := u.waitForCilium(ctx, cluster); err != nil {
		return nil, err
	}

	return diff, nil
}

func (u *Upgrader) waitForPreflight(ctx context.Context, cluster *types.Cluster) error {
	if err := u.client.WaitForPreflightDaemonSet(ctx, cluster); err != nil {
		return err
	}

	if err := u.client.WaitForPreflightDeployment(ctx, cluster); err != nil {
		return err
	}

	return nil
}

func (u *Upgrader) waitForCilium(ctx context.Context, cluster *types.Cluster) error {
	if err := u.client.WaitForCiliumDaemonSet(ctx, cluster); err != nil {
		return err
	}

	if err := u.client.WaitForCiliumDeployment(ctx, cluster); err != nil {
		return err
	}

	return nil
}

func ciliumChangeDiff(currentSpec, newSpec *cluster.Spec) *types.ChangeDiff {
	if currentSpec.VersionsBundle.Cilium.Version == newSpec.VersionsBundle.Cilium.Version {
		return nil
	}

	return &types.ChangeDiff{
		ComponentReports: []types.ComponentChangeDiff{
			{
				ComponentName: "cilium",
				OldVersion:    currentSpec.VersionsBundle.Cilium.Version,
				NewVersion:    newSpec.VersionsBundle.Cilium.Version,
			},
		},
	}
}

func ChangeDiff(currentSpec, newSpec *cluster.Spec) *types.ChangeDiff {
	return ciliumChangeDiff(currentSpec, newSpec)
}
