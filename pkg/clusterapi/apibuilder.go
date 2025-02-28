package clusterapi

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/constants"
)

const (
	clusterKind               = "Cluster"
	kubeadmControlPlaneKind   = "KubeadmControlPlane"
	etcdadmClusterKind        = "EtcdadmCluster"
	kubeadmConfigTemplateKind = "KubeadmConfigTemplate"
	machineDeploymentKind     = "MachineDeployment"
)

var (
	clusterAPIVersion             = clusterv1.GroupVersion.String()
	kubeadmControlPlaneAPIVersion = controlplanev1.GroupVersion.String()
	bootstrapAPIVersion           = bootstrapv1.GroupVersion.String()
	etcdClusterAPIVersion         = fmt.Sprintf("etcdcluster.%s/%s", clusterv1.GroupVersion.Group, clusterv1.GroupVersion.Version)
)

type APIObject interface {
	runtime.Object
	GetName() string
}

func InfrastructureAPIVersion() string {
	return fmt.Sprintf("infrastructure.%s/%s", clusterv1.GroupVersion.Group, clusterv1.GroupVersion.Version)
}

func clusterLabels(clusterName string) map[string]string {
	return map[string]string{clusterv1.ClusterLabelName: clusterName}
}

func Cluster(clusterSpec *cluster.Spec, infrastructureObject, controlPlaneObject APIObject) *clusterv1.Cluster {
	clusterName := clusterSpec.Cluster.GetName()
	cluster := &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterAPIVersion,
			Kind:       clusterKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: constants.EksaSystemNamespace,
			Labels:    clusterLabels(clusterName),
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				Pods: &clusterv1.NetworkRanges{
					CIDRBlocks: clusterSpec.Cluster.Spec.ClusterNetwork.Pods.CidrBlocks,
				},
				Services: &clusterv1.NetworkRanges{
					CIDRBlocks: clusterSpec.Cluster.Spec.ClusterNetwork.Services.CidrBlocks,
				},
			},
			ControlPlaneRef: &v1.ObjectReference{
				APIVersion: controlPlaneObject.GetObjectKind().GroupVersionKind().GroupVersion().String(),
				Name:       controlPlaneObject.GetName(),
				Kind:       controlPlaneObject.GetObjectKind().GroupVersionKind().Kind,
			},
			InfrastructureRef: &v1.ObjectReference{
				APIVersion: infrastructureObject.GetObjectKind().GroupVersionKind().GroupVersion().String(),
				Name:       infrastructureObject.GetName(),
				Kind:       infrastructureObject.GetObjectKind().GroupVersionKind().Kind,
			},
		},
	}

	if clusterSpec.Cluster.Spec.ExternalEtcdConfiguration != nil {
		cluster.Spec.ManagedExternalEtcdRef = &v1.ObjectReference{
			APIVersion: etcdClusterAPIVersion,
			Kind:       etcdadmClusterKind,
			Name:       clusterName,
		}
	}

	return cluster
}

func KubeadmControlPlane(clusterSpec *cluster.Spec, infrastructureObject APIObject) *controlplanev1.KubeadmControlPlane {
	replicas := int32(clusterSpec.Cluster.Spec.ControlPlaneConfiguration.Count)

	etcd := bootstrapv1.Etcd{}
	if clusterSpec.Cluster.Spec.ExternalEtcdConfiguration != nil {
		etcd.External = &bootstrapv1.ExternalEtcd{
			Endpoints: []string{},
		}
	} else {
		etcd.Local = &bootstrapv1.LocalEtcd{
			ImageMeta: bootstrapv1.ImageMeta{
				ImageRepository: clusterSpec.VersionsBundle.KubeDistro.Etcd.Repository,
				ImageTag:        clusterSpec.VersionsBundle.KubeDistro.Etcd.Tag,
			},
			ExtraArgs: map[string]string{},
		}
	}

	return &controlplanev1.KubeadmControlPlane{
		TypeMeta: metav1.TypeMeta{
			APIVersion: kubeadmControlPlaneAPIVersion,
			Kind:       kubeadmControlPlaneKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterSpec.Cluster.GetName(),
			Namespace: constants.EksaSystemNamespace,
		},
		Spec: controlplanev1.KubeadmControlPlaneSpec{
			MachineTemplate: controlplanev1.KubeadmControlPlaneMachineTemplate{
				InfrastructureRef: v1.ObjectReference{
					APIVersion: infrastructureObject.GetObjectKind().GroupVersionKind().GroupVersion().String(),
					Kind:       infrastructureObject.GetObjectKind().GroupVersionKind().Kind,
					Name:       infrastructureObject.GetName(),
				},
			},
			KubeadmConfigSpec: bootstrapv1.KubeadmConfigSpec{
				ClusterConfiguration: &bootstrapv1.ClusterConfiguration{
					ImageRepository: clusterSpec.VersionsBundle.KubeDistro.Kubernetes.Repository,
					DNS: bootstrapv1.DNS{
						ImageMeta: bootstrapv1.ImageMeta{
							ImageRepository: clusterSpec.VersionsBundle.KubeDistro.CoreDNS.Repository,
							ImageTag:        clusterSpec.VersionsBundle.KubeDistro.CoreDNS.Tag,
						},
					},
					Etcd: etcd,
					APIServer: bootstrapv1.APIServer{
						ControlPlaneComponent: bootstrapv1.ControlPlaneComponent{
							ExtraArgs: map[string]string{},
						},
					},
					ControllerManager: bootstrapv1.ControlPlaneComponent{
						ExtraArgs: map[string]string{},
					},
				},
				InitConfiguration: &bootstrapv1.InitConfiguration{
					NodeRegistration: bootstrapv1.NodeRegistrationOptions{
						KubeletExtraArgs: map[string]string{},
					},
				},
				JoinConfiguration: &bootstrapv1.JoinConfiguration{
					NodeRegistration: bootstrapv1.NodeRegistrationOptions{
						KubeletExtraArgs: map[string]string{},
					},
				},
				PreKubeadmCommands:  []string{},
				PostKubeadmCommands: []string{},
			},
			Replicas: &replicas,
			Version:  clusterSpec.VersionsBundle.KubeDistro.Kubernetes.Tag,
		},
	}
}

func KubeadmConfigTemplate(clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration) bootstrapv1.KubeadmConfigTemplate {
	return bootstrapv1.KubeadmConfigTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: bootstrapAPIVersion,
			Kind:       kubeadmConfigTemplateKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      workerNodeGroupConfig.Name, // TODO: diff
			Namespace: constants.EksaSystemNamespace,
		},
		Spec: bootstrapv1.KubeadmConfigTemplateSpec{
			Template: bootstrapv1.KubeadmConfigTemplateResource{
				Spec: bootstrapv1.KubeadmConfigSpec{
					ClusterConfiguration: &bootstrapv1.ClusterConfiguration{
						ControllerManager: bootstrapv1.ControlPlaneComponent{
							ExtraArgs: map[string]string{},
						},
						APIServer: bootstrapv1.APIServer{
							ControlPlaneComponent: bootstrapv1.ControlPlaneComponent{
								ExtraArgs: map[string]string{},
							},
						},
					},
					JoinConfiguration: &bootstrapv1.JoinConfiguration{
						NodeRegistration: bootstrapv1.NodeRegistrationOptions{
							KubeletExtraArgs: map[string]string{},
						},
					},
					PreKubeadmCommands:  []string{},
					PostKubeadmCommands: []string{},
				},
			},
		},
	}
}

func MachineDeployment(clusterSpec *cluster.Spec, workerNodeGroupConfig v1alpha1.WorkerNodeGroupConfiguration, bootstrapObject, infrastructureObject APIObject) clusterv1.MachineDeployment {
	clusterName := clusterSpec.Cluster.GetName()
	replicas := int32(workerNodeGroupConfig.Count)
	version := clusterSpec.VersionsBundle.KubeDistro.Kubernetes.Tag

	return clusterv1.MachineDeployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterAPIVersion,
			Kind:       machineDeploymentKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      workerNodeGroupConfig.Name,
			Namespace: constants.EksaSystemNamespace,
			Labels:    clusterLabels(clusterName),
		},
		Spec: clusterv1.MachineDeploymentSpec{
			ClusterName: clusterName,
			Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{},
			},
			Template: clusterv1.MachineTemplateSpec{
				ObjectMeta: clusterv1.ObjectMeta{
					Labels: clusterLabels(clusterName),
				},
				Spec: clusterv1.MachineSpec{
					Bootstrap: clusterv1.Bootstrap{
						ConfigRef: &v1.ObjectReference{
							APIVersion: bootstrapObject.GetObjectKind().GroupVersionKind().GroupVersion().String(),
							Kind:       bootstrapObject.GetObjectKind().GroupVersionKind().Kind,
							Name:       bootstrapObject.GetName(),
						},
					},
					ClusterName: clusterName,
					InfrastructureRef: v1.ObjectReference{
						APIVersion: infrastructureObject.GetObjectKind().GroupVersionKind().GroupVersion().String(),
						Kind:       infrastructureObject.GetObjectKind().GroupVersionKind().Kind,
						Name:       infrastructureObject.GetName(),
					},
					Version: &version,
				},
			},
			Replicas: &replicas,
		},
	}
}
