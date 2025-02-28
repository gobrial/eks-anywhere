package cloudstack

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"

	etcdv1beta1 "github.com/mrajashree/etcdadm-controller/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	kubeadmv1beta1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/bootstrapper"
	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/clusterapi"
	"github.com/aws/eks-anywhere/pkg/constants"
	"github.com/aws/eks-anywhere/pkg/crypto"
	"github.com/aws/eks-anywhere/pkg/executables"
	"github.com/aws/eks-anywhere/pkg/filewriter"
	"github.com/aws/eks-anywhere/pkg/logger"
	"github.com/aws/eks-anywhere/pkg/providers"
	"github.com/aws/eks-anywhere/pkg/providers/cloudstack/decoder"
	"github.com/aws/eks-anywhere/pkg/providers/common"
	"github.com/aws/eks-anywhere/pkg/templater"
	"github.com/aws/eks-anywhere/pkg/types"
	releasev1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

const (
	eksaLicense                = "EKSA_LICENSE"
	controlEndpointDefaultPort = "6443"
)

//go:embed config/template-cp.yaml
var defaultCAPIConfigCP string

//go:embed config/template-md.yaml
var defaultClusterConfigMD string

//go:embed config/machine-health-check-template.yaml
var mhcTemplate []byte

var requiredEnvs = []string{decoder.CloudStackCloudConfigB64SecretKey}

var (
	eksaCloudStackDeploymentResourceType = fmt.Sprintf("cloudstackdatacenterconfigs.%s", v1alpha1.GroupVersion.Group)
	eksaCloudStackMachineResourceType    = fmt.Sprintf("cloudstackmachineconfigs.%s", v1alpha1.GroupVersion.Group)
)

type cloudstackProvider struct {
	datacenterConfig       *v1alpha1.CloudStackDatacenterConfig
	machineConfigs         map[string]*v1alpha1.CloudStackMachineConfig
	clusterConfig          *v1alpha1.Cluster
	providerKubectlClient  ProviderKubectlClient
	writer                 filewriter.FileWriter
	selfSigned             bool
	controlPlaneSshAuthKey string
	workerSshAuthKey       string
	etcdSshAuthKey         string
	templateBuilder        *CloudStackTemplateBuilder
	skipIpCheck            bool
	validator              *Validator
}

func (p *cloudstackProvider) PreBootstrapSetup(ctx context.Context, cluster *types.Cluster) error {
	return nil
}

func (p *cloudstackProvider) PostBootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	return nil
}

func (p *cloudstackProvider) UpdateSecrets(ctx context.Context, cluster *types.Cluster) error {
	return nil
}

func (p *cloudstackProvider) ValidateNewSpec(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	return fmt.Errorf("cloudstack provider does not support this functionality currently")
}

func (p *cloudstackProvider) ChangeDiff(currentSpec, newSpec *cluster.Spec) *types.ComponentChangeDiff {
	panic("implement me")
}

func (p *cloudstackProvider) MachineDeploymentsToDelete(workloadCluster *types.Cluster, currentSpec, newSpec *cluster.Spec) []string {
	panic("implement me")
}

func (p *cloudstackProvider) RunPostControlPlaneUpgrade(ctx context.Context, oldClusterSpec *cluster.Spec, clusterSpec *cluster.Spec, workloadCluster *types.Cluster, managementCluster *types.Cluster) error {
	// Nothing to do
	return nil
}

func (p *cloudstackProvider) RunPostControlPlaneCreation(ctx context.Context, clusterSpec *cluster.Spec, cluster *types.Cluster) error {
	// Nothing to do
	return nil
}

type ProviderKubectlClient interface {
	ApplyKubeSpecFromBytes(ctx context.Context, cluster *types.Cluster, data []byte) error
	CreateNamespace(ctx context.Context, kubeconfig string, namespace string) error
	LoadSecret(ctx context.Context, secretObject string, secretObjType string, secretObjectName string, kubeConfFile string) error
	GetEksaCluster(ctx context.Context, cluster *types.Cluster, clusterName string) (*v1alpha1.Cluster, error)
	GetEksaCloudStackDatacenterConfig(ctx context.Context, cloudstackDatacenterConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.CloudStackDatacenterConfig, error)
	GetEksaCloudStackMachineConfig(ctx context.Context, cloudstackMachineConfigName string, kubeconfigFile string, namespace string) (*v1alpha1.CloudStackMachineConfig, error)
	GetKubeadmControlPlane(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*kubeadmv1beta1.KubeadmControlPlane, error)
	GetMachineDeployment(ctx context.Context, workerNodeGroupName string, opts ...executables.KubectlOpt) (*clusterv1.MachineDeployment, error)
	GetEtcdadmCluster(ctx context.Context, cluster *types.Cluster, clusterName string, opts ...executables.KubectlOpt) (*etcdv1beta1.EtcdadmCluster, error)
	GetSecret(ctx context.Context, secretObjectName string, opts ...executables.KubectlOpt) (*corev1.Secret, error)
	UpdateAnnotation(ctx context.Context, resourceType, objectName string, annotations map[string]string, opts ...executables.KubectlOpt) error
	SearchCloudStackMachineConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.CloudStackMachineConfig, error)
	SearchCloudStackDatacenterConfig(ctx context.Context, name string, kubeconfigFile string, namespace string) ([]*v1alpha1.CloudStackDatacenterConfig, error)
	DeleteEksaCloudStackDatacenterConfig(ctx context.Context, cloudstackDatacenterConfigName string, kubeconfigFile string, namespace string) error
	DeleteEksaCloudStackMachineConfig(ctx context.Context, cloudstackMachineConfigName string, kubeconfigFile string, namespace string) error
}

func NewProvider(datacenterConfig *v1alpha1.CloudStackDatacenterConfig, machineConfigs map[string]*v1alpha1.CloudStackMachineConfig, clusterConfig *v1alpha1.Cluster, providerKubectlClient ProviderKubectlClient, providerCmkClient ProviderCmkClient, writer filewriter.FileWriter, now types.NowFunc, skipIpCheck bool) *cloudstackProvider {
	return NewProviderCustomNet(
		datacenterConfig,
		machineConfigs,
		clusterConfig,
		providerKubectlClient,
		providerCmkClient,
		writer,
		now,
		skipIpCheck,
	)
}

func NewProviderCustomNet(datacenterConfig *v1alpha1.CloudStackDatacenterConfig, machineConfigs map[string]*v1alpha1.CloudStackMachineConfig, clusterConfig *v1alpha1.Cluster, providerKubectlClient ProviderKubectlClient, providerCmkClient ProviderCmkClient, writer filewriter.FileWriter, now types.NowFunc, skipIpCheck bool) *cloudstackProvider {
	var controlPlaneMachineSpec, workerNodeGroupMachineSpec, etcdMachineSpec *v1alpha1.CloudStackMachineConfigSpec
	if clusterConfig.Spec.ControlPlaneConfiguration.MachineGroupRef != nil && machineConfigs[clusterConfig.Spec.ControlPlaneConfiguration.MachineGroupRef.Name] != nil {
		controlPlaneMachineSpec = &machineConfigs[clusterConfig.Spec.ControlPlaneConfiguration.MachineGroupRef.Name].Spec
	}
	if len(clusterConfig.Spec.WorkerNodeGroupConfigurations) > 0 && clusterConfig.Spec.WorkerNodeGroupConfigurations[0].MachineGroupRef != nil && machineConfigs[clusterConfig.Spec.WorkerNodeGroupConfigurations[0].MachineGroupRef.Name] != nil {
		workerNodeGroupMachineSpec = &machineConfigs[clusterConfig.Spec.WorkerNodeGroupConfigurations[0].MachineGroupRef.Name].Spec
	}
	if clusterConfig.Spec.ExternalEtcdConfiguration != nil {
		if clusterConfig.Spec.ExternalEtcdConfiguration.MachineGroupRef != nil && machineConfigs[clusterConfig.Spec.ExternalEtcdConfiguration.MachineGroupRef.Name] != nil {
			etcdMachineSpec = &machineConfigs[clusterConfig.Spec.ExternalEtcdConfiguration.MachineGroupRef.Name].Spec
		}
	}
	return &cloudstackProvider{
		datacenterConfig:      datacenterConfig,
		machineConfigs:        machineConfigs,
		clusterConfig:         clusterConfig,
		providerKubectlClient: providerKubectlClient,
		writer:                writer,
		selfSigned:            false,
		templateBuilder: &CloudStackTemplateBuilder{
			datacenterConfigSpec:       &datacenterConfig.Spec,
			controlPlaneMachineSpec:    controlPlaneMachineSpec,
			workerNodeGroupMachineSpec: workerNodeGroupMachineSpec,
			etcdMachineSpec:            etcdMachineSpec,
			now:                        now,
		},
		skipIpCheck: skipIpCheck,
		validator:   NewValidator(providerCmkClient),
	}
}

func (p *cloudstackProvider) UpdateKubeConfig(_ *[]byte, _ string) error {
	// customize generated kube config
	return nil
}

func (p *cloudstackProvider) BootstrapClusterOpts() ([]bootstrapper.BootstrapClusterOption, error) {
	execConfig, err := decoder.ParseCloudStackSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to parse environment variable exec config: %v", err)
	}
	return common.BootstrapClusterOpts(execConfig.ManagementUrl, p.clusterConfig)
}

func (p *cloudstackProvider) Name() string {
	return constants.CloudStackProviderName
}

func (p *cloudstackProvider) DatacenterResourceType() string {
	return eksaCloudStackDeploymentResourceType
}

func (p *cloudstackProvider) MachineResourceType() string {
	return eksaCloudStackMachineResourceType
}

func (p *cloudstackProvider) setupSSHAuthKeysForCreate() error {
	var useKeyGeneratedForControlplane, useKeyGeneratedForWorker bool
	var err error
	controlPlaneUser := p.machineConfigs[p.clusterConfig.Spec.ControlPlaneConfiguration.MachineGroupRef.Name].Spec.Users[0]
	p.controlPlaneSshAuthKey = controlPlaneUser.SshAuthorizedKeys[0]
	if len(p.controlPlaneSshAuthKey) > 0 {
		p.controlPlaneSshAuthKey, err = common.StripSshAuthorizedKeyComment(p.controlPlaneSshAuthKey)
		if err != nil {
			return err
		}
	} else {
		logger.Info("Provided control plane sshAuthorizedKey is not set or is empty, auto-generating new key pair...")
		generatedKey, err := common.GenerateSSHAuthKey(p.writer)
		if err != nil {
			return err
		}
		p.controlPlaneSshAuthKey = generatedKey
		useKeyGeneratedForControlplane = true
	}
	workerUser := p.machineConfigs[p.clusterConfig.Spec.WorkerNodeGroupConfigurations[0].MachineGroupRef.Name].Spec.Users[0]
	p.workerSshAuthKey = workerUser.SshAuthorizedKeys[0]
	if len(p.workerSshAuthKey) > 0 {
		p.workerSshAuthKey, err = common.StripSshAuthorizedKeyComment(p.workerSshAuthKey)
		if err != nil {
			return err
		}
	} else {
		if useKeyGeneratedForControlplane { // use the same key
			p.workerSshAuthKey = p.controlPlaneSshAuthKey
		} else {
			logger.Info("Provided worker sshAuthorizedKey is not set or is empty, auto-generating new key pair...")
			generatedKey, err := common.GenerateSSHAuthKey(p.writer)
			if err != nil {
				return err
			}
			p.workerSshAuthKey = generatedKey
			useKeyGeneratedForWorker = true
		}
	}
	if p.clusterConfig.Spec.ExternalEtcdConfiguration != nil {
		etcdUser := p.machineConfigs[p.clusterConfig.Spec.ExternalEtcdConfiguration.MachineGroupRef.Name].Spec.Users[0]
		p.etcdSshAuthKey = etcdUser.SshAuthorizedKeys[0]
		if len(p.etcdSshAuthKey) > 0 {
			p.etcdSshAuthKey, err = common.StripSshAuthorizedKeyComment(p.etcdSshAuthKey)
			if err != nil {
				return err
			}
		} else {
			if useKeyGeneratedForControlplane { // use the same key as for controlplane
				p.etcdSshAuthKey = p.controlPlaneSshAuthKey
			} else if useKeyGeneratedForWorker {
				p.etcdSshAuthKey = p.workerSshAuthKey // if cp key was provided by user, check if worker key was generated by cli and use that
			} else {
				logger.Info("Provided etcd sshAuthorizedKey is not set or is empty, auto-generating new key pair...")
				generatedKey, err := common.GenerateSSHAuthKey(p.writer)
				if err != nil {
					return err
				}
				p.etcdSshAuthKey = generatedKey
			}
		}
		etcdUser.SshAuthorizedKeys[0] = p.etcdSshAuthKey
	}
	controlPlaneUser.SshAuthorizedKeys[0] = p.controlPlaneSshAuthKey
	workerUser.SshAuthorizedKeys[0] = p.workerSshAuthKey
	return nil
}

func (p *cloudstackProvider) validateManagementApiEndpoint(rawurl string) error {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return fmt.Errorf("CloudStack managementApiEndpoint is invalid: #{err}")
	}
	return nil
}

func getHostnameFromUrl(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", fmt.Errorf("%s is not a valid url", rawurl)
	}

	return url.Hostname(), nil
}

func (p *cloudstackProvider) validateEnv(ctx context.Context) error {
	var cloudStackB64EncodedSecret string
	var ok bool

	if cloudStackB64EncodedSecret, ok = os.LookupEnv(decoder.EksacloudStackCloudConfigB64SecretKey); ok && len(cloudStackB64EncodedSecret) > 0 {
		if err := os.Setenv(decoder.CloudStackCloudConfigB64SecretKey, cloudStackB64EncodedSecret); err != nil {
			return fmt.Errorf("unable to set %s: %v", decoder.CloudStackCloudConfigB64SecretKey, err)
		}
	} else {
		return fmt.Errorf("%s is not set or is empty", decoder.EksacloudStackCloudConfigB64SecretKey)
	}
	execConfig, err := decoder.ParseCloudStackSecret()
	if err != nil {
		return fmt.Errorf("failed to parse environment variable exec config: %v", err)
	}
	if len(execConfig.ManagementUrl) <= 0 {
		return errors.New("cloudstack management api url is not set or is empty")
	}
	if err := p.validateManagementApiEndpoint(execConfig.ManagementUrl); err != nil {
		return errors.New("CloudStackDatacenterConfig managementApiEndpoint is invalid")
	}
	if _, ok := os.LookupEnv(eksaLicense); !ok {
		if err := os.Setenv(eksaLicense, ""); err != nil {
			return fmt.Errorf("unable to set %s: %v", eksaLicense, err)
		}
	}
	return nil
}

func (p *cloudstackProvider) SetupAndValidateCreateCluster(ctx context.Context, clusterSpec *cluster.Spec) error {
	err := p.validateEnv(ctx)
	if err != nil {
		return fmt.Errorf("failed setup and validations: %v", err)
	}

	cloudStackClusterSpec := NewSpec(clusterSpec, p.machineConfigs, p.datacenterConfig)

	if err := p.validator.validateCloudStackAccess(ctx); err != nil {
		return err
	}
	if err := p.validator.ValidateCloudStackDatacenterConfig(ctx, p.datacenterConfig); err != nil {
		return err
	}
	if err := p.validator.ValidateClusterMachineConfigs(ctx, cloudStackClusterSpec); err != nil {
		return err
	}

	if err := p.setupSSHAuthKeysForCreate(); err != nil {
		return fmt.Errorf("failed setup and validations: %v", err)
	}

	if clusterSpec.Cluster.IsManaged() {
		for _, mc := range p.MachineConfigs(clusterSpec) {
			em, err := p.providerKubectlClient.SearchCloudStackMachineConfig(ctx, mc.GetName(), clusterSpec.ManagementCluster.KubeconfigFile, mc.GetNamespace())
			if err != nil {
				return err
			}
			if len(em) > 0 {
				return fmt.Errorf("CloudStackMachineConfig %s already exists", mc.GetName())
			}
		}
		existingDatacenter, err := p.providerKubectlClient.SearchCloudStackDatacenterConfig(ctx, p.datacenterConfig.Name, clusterSpec.ManagementCluster.KubeconfigFile, clusterSpec.Cluster.Namespace)
		if err != nil {
			return err
		}
		if len(existingDatacenter) > 0 {
			return fmt.Errorf("CloudStackDeployment %s already exists", p.datacenterConfig.Name)
		}
	}
	if p.skipIpCheck {
		logger.Info("Skipping check for whether control plane ip is in use")
		return nil
	}

	return nil
}

func (p *cloudstackProvider) SetupAndValidateUpgradeCluster(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) error {
	return fmt.Errorf("upgrade is not yet supported for CloudStack cluster")
}

func (p *cloudstackProvider) SetupAndValidateDeleteCluster(ctx context.Context) error {
	err := p.validateEnv(ctx)
	if err != nil {
		return fmt.Errorf("failed setup and validations: %v", err)
	}
	return nil
}

type CloudStackTemplateBuilder struct {
	datacenterConfigSpec       *v1alpha1.CloudStackDatacenterConfigSpec
	controlPlaneMachineSpec    *v1alpha1.CloudStackMachineConfigSpec
	workerNodeGroupMachineSpec *v1alpha1.CloudStackMachineConfigSpec
	etcdMachineSpec            *v1alpha1.CloudStackMachineConfigSpec
	now                        types.NowFunc
}

func (cs *CloudStackTemplateBuilder) GenerateCAPISpecControlPlane(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	var etcdMachineSpec v1alpha1.CloudStackMachineConfigSpec
	if clusterSpec.Cluster.Spec.ExternalEtcdConfiguration != nil {
		etcdMachineSpec = *cs.etcdMachineSpec
	}
	execConfig, err := decoder.ParseCloudStackSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to parse environment variable exec config: %v", err)
	}
	values := buildTemplateMapCP(clusterSpec, *cs.datacenterConfigSpec, *cs.controlPlaneMachineSpec, etcdMachineSpec, execConfig.ManagementUrl, execConfig.VerifySsl)

	for _, buildOption := range buildOptions {
		buildOption(values)
	}

	bytes, err := templater.Execute(defaultCAPIConfigCP, values)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (cs *CloudStackTemplateBuilder) GenerateCAPISpecWorkers(clusterSpec *cluster.Spec, buildOptions ...providers.BuildMapOption) (content []byte, err error) {
	execConfig, err := decoder.ParseCloudStackSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to parse environment variable exec config: %v", err)
	}
	values := buildTemplateMapMD(clusterSpec, *cs.datacenterConfigSpec, *cs.workerNodeGroupMachineSpec, execConfig.ManagementUrl)

	for _, buildOption := range buildOptions {
		buildOption(values)
	}

	bytes, err := templater.Execute(defaultClusterConfigMD, values)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func buildTemplateMapCP(clusterSpec *cluster.Spec, datacenterConfigSpec v1alpha1.CloudStackDatacenterConfigSpec, controlPlaneMachineSpec, etcdMachineSpec v1alpha1.CloudStackMachineConfigSpec, managementApiEndpoint string, verifySsl string) map[string]interface{} {
	bundle := clusterSpec.VersionsBundle
	format := "cloud-config"
	host, port, _ := net.SplitHostPort(clusterSpec.Cluster.Spec.ControlPlaneConfiguration.Endpoint.Host)
	etcdExtraArgs := clusterapi.SecureEtcdTlsCipherSuitesExtraArgs()
	sharedExtraArgs := clusterapi.SecureTlsCipherSuitesExtraArgs()
	kubeletExtraArgs := clusterapi.SecureTlsCipherSuitesExtraArgs().
		Append(clusterapi.ResolvConfExtraArgs(clusterSpec.Cluster.Spec.ClusterNetwork.DNS.ResolvConf)).
		Append(clusterapi.ControlPlaneNodeLabelsExtraArgs(clusterSpec.Cluster.Spec.ControlPlaneConfiguration))
	apiServerExtraArgs := clusterapi.OIDCToExtraArgs(clusterSpec.OIDCConfig).
		Append(clusterapi.AwsIamAuthExtraArgs(clusterSpec.AWSIamConfig)).
		Append(clusterapi.PodIAMAuthExtraArgs(clusterSpec.Cluster.Spec.PodIAMConfig)).
		Append(sharedExtraArgs)

	values := map[string]interface{}{
		"clusterName":                                clusterSpec.Cluster.Name,
		"controlPlaneEndpointHost":                   host,
		"controlPlaneEndpointPort":                   port,
		"controlPlaneReplicas":                       clusterSpec.Cluster.Spec.ControlPlaneConfiguration.Count,
		"kubernetesRepository":                       bundle.KubeDistro.Kubernetes.Repository,
		"kubernetesVersion":                          bundle.KubeDistro.Kubernetes.Tag,
		"etcdRepository":                             bundle.KubeDistro.Etcd.Repository,
		"etcdImageTag":                               bundle.KubeDistro.Etcd.Tag,
		"corednsRepository":                          bundle.KubeDistro.CoreDNS.Repository,
		"corednsVersion":                             bundle.KubeDistro.CoreDNS.Tag,
		"nodeDriverRegistrarImage":                   bundle.KubeDistro.NodeDriverRegistrar.VersionedImage(),
		"livenessProbeImage":                         bundle.KubeDistro.LivenessProbe.VersionedImage(),
		"externalAttacherImage":                      bundle.KubeDistro.ExternalAttacher.VersionedImage(),
		"externalProvisionerImage":                   bundle.KubeDistro.ExternalProvisioner.VersionedImage(),
		"cloudstackManagementApiEndpoint":            managementApiEndpoint,
		"managerImage":                               bundle.CloudStack.ClusterAPIController.VersionedImage(),
		"verifySsl":                                  verifySsl,
		"cloudstackDomain":                           datacenterConfigSpec.Domain,
		"cloudstackZones":                            datacenterConfigSpec.Zones,
		"cloudstackAccount":                          datacenterConfigSpec.Account,
		"cloudstackControlPlaneComputeOfferingId":    controlPlaneMachineSpec.ComputeOffering.Id,
		"cloudstackControlPlaneComputeOfferingName":  controlPlaneMachineSpec.ComputeOffering.Name,
		"cloudstackControlPlaneTemplateOfferingId":   controlPlaneMachineSpec.Template.Id,
		"cloudstackControlPlaneTemplateOfferingName": controlPlaneMachineSpec.Template.Name,
		"cloudstackControlPlaneCustomDetails":        controlPlaneMachineSpec.UserCustomDetails,
		"affinityGroupIds":                           controlPlaneMachineSpec.AffinityGroupIds,
		"cloudstackEtcdComputeOfferingId":            etcdMachineSpec.ComputeOffering.Id,
		"cloudstackEtcdComputeOfferingName":          etcdMachineSpec.ComputeOffering.Name,
		"cloudstackEtcdTemplateOfferingId":           etcdMachineSpec.Template.Id,
		"cloudstackEtcdTemplateOfferingName":         etcdMachineSpec.Template.Name,
		"cloudstackEtcdCustomDetails":                etcdMachineSpec.UserCustomDetails,
		"cloudstackEtcdAffinityGroupIds":             etcdMachineSpec.AffinityGroupIds,
		"controlPlaneSshUsername":                    controlPlaneMachineSpec.Users[0].Name,
		"podCidrs":                                   clusterSpec.Cluster.Spec.ClusterNetwork.Pods.CidrBlocks,
		"serviceCidrs":                               clusterSpec.Cluster.Spec.ClusterNetwork.Services.CidrBlocks,
		"apiserverExtraArgs":                         apiServerExtraArgs.ToPartialYaml(),
		"kubeletExtraArgs":                           kubeletExtraArgs.ToPartialYaml(),
		"etcdExtraArgs":                              etcdExtraArgs.ToPartialYaml(),
		"etcdCipherSuites":                           crypto.SecureCipherSuitesString(),
		"controllermanagerExtraArgs":                 sharedExtraArgs.ToPartialYaml(),
		"schedulerExtraArgs":                         sharedExtraArgs.ToPartialYaml(),
		"format":                                     format,
		"externalEtcdVersion":                        bundle.KubeDistro.EtcdVersion,
		"etcdImage":                                  bundle.KubeDistro.EtcdImage.VersionedImage(),
		"eksaSystemNamespace":                        constants.EksaSystemNamespace,
		"auditPolicy":                                common.GetAuditPolicy(),
	}

	if clusterSpec.Cluster.Spec.RegistryMirrorConfiguration != nil {
		values["registryMirrorConfiguration"] = clusterSpec.Cluster.Spec.RegistryMirrorConfiguration.Endpoint
		if len(clusterSpec.Cluster.Spec.RegistryMirrorConfiguration.CACertContent) > 0 {
			values["registryCACert"] = clusterSpec.Cluster.Spec.RegistryMirrorConfiguration.CACertContent
		}
	}

	if clusterSpec.Cluster.Spec.ProxyConfiguration != nil {
		values["proxyConfig"] = true
		capacity := len(clusterSpec.Cluster.Spec.ClusterNetwork.Pods.CidrBlocks) +
			len(clusterSpec.Cluster.Spec.ClusterNetwork.Services.CidrBlocks) +
			len(clusterSpec.Cluster.Spec.ProxyConfiguration.NoProxy) + 4
		noProxyList := make([]string, 0, capacity)
		noProxyList = append(noProxyList, clusterSpec.Cluster.Spec.ClusterNetwork.Pods.CidrBlocks...)
		noProxyList = append(noProxyList, clusterSpec.Cluster.Spec.ClusterNetwork.Services.CidrBlocks...)
		noProxyList = append(noProxyList, clusterSpec.Cluster.Spec.ProxyConfiguration.NoProxy...)

		// Add no-proxy defaults
		noProxyList = append(noProxyList, common.NoProxyDefaults...)
		cloudStackManagementApiEndpointHostname, err := getHostnameFromUrl(managementApiEndpoint)
		if err == nil {
			noProxyList = append(noProxyList, cloudStackManagementApiEndpointHostname)
		}
		noProxyList = append(noProxyList,
			clusterSpec.Cluster.Spec.ControlPlaneConfiguration.Endpoint.Host,
		)

		values["httpProxy"] = clusterSpec.Cluster.Spec.ProxyConfiguration.HttpProxy
		values["httpsProxy"] = clusterSpec.Cluster.Spec.ProxyConfiguration.HttpsProxy
		values["noProxy"] = noProxyList
	}

	if clusterSpec.Cluster.Spec.ExternalEtcdConfiguration != nil {
		values["externalEtcd"] = true
		values["externalEtcdReplicas"] = clusterSpec.Cluster.Spec.ExternalEtcdConfiguration.Count
		values["etcdSshUsername"] = etcdMachineSpec.Users[0].Name
	}

	if clusterSpec.AWSIamConfig != nil {
		values["awsIamAuth"] = true
	}

	return values
}

func buildTemplateMapMD(clusterSpec *cluster.Spec, datacenterConfigSpec v1alpha1.CloudStackDatacenterConfigSpec, workerNodeGroupMachineSpec v1alpha1.CloudStackMachineConfigSpec, managementApiEndpoint string) map[string]interface{} {
	bundle := clusterSpec.VersionsBundle
	format := "cloud-config"

	values := map[string]interface{}{
		"clusterName":                clusterSpec.Cluster.Name,
		"kubernetesVersion":          bundle.KubeDistro.Kubernetes.Tag,
		"cloudstackTemplateId":       workerNodeGroupMachineSpec.Template.Id,
		"cloudstackTemplateName":     workerNodeGroupMachineSpec.Template.Name,
		"cloudstackOfferingId":       workerNodeGroupMachineSpec.ComputeOffering.Id,
		"cloudstackOfferingName":     workerNodeGroupMachineSpec.ComputeOffering.Name,
		"cloudstackCustomDetails":    workerNodeGroupMachineSpec.UserCustomDetails,
		"cloudstackAffinityGroupIds": workerNodeGroupMachineSpec.AffinityGroupIds,
		"workerReplicas":             clusterSpec.Cluster.Spec.WorkerNodeGroupConfigurations[0].Count,
		"workerSshUsername":          workerNodeGroupMachineSpec.Users[0].Name,
		"format":                     format,
		"eksaSystemNamespace":        constants.EksaSystemNamespace,
	}

	if clusterSpec.Cluster.Spec.RegistryMirrorConfiguration != nil {
		values["registryMirrorConfiguration"] = clusterSpec.Cluster.Spec.RegistryMirrorConfiguration.Endpoint
		if len(clusterSpec.Cluster.Spec.RegistryMirrorConfiguration.CACertContent) > 0 {
			values["registryCACert"] = clusterSpec.Cluster.Spec.RegistryMirrorConfiguration.CACertContent
		}
	}

	if clusterSpec.Cluster.Spec.ProxyConfiguration != nil {
		values["proxyConfig"] = true
		capacity := len(clusterSpec.Cluster.Spec.ClusterNetwork.Pods.CidrBlocks) +
			len(clusterSpec.Cluster.Spec.ClusterNetwork.Services.CidrBlocks) +
			len(clusterSpec.Cluster.Spec.ProxyConfiguration.NoProxy) + 4
		noProxyList := make([]string, 0, capacity)
		noProxyList = append(noProxyList, clusterSpec.Cluster.Spec.ClusterNetwork.Pods.CidrBlocks...)
		noProxyList = append(noProxyList, clusterSpec.Cluster.Spec.ClusterNetwork.Services.CidrBlocks...)
		noProxyList = append(noProxyList, clusterSpec.Cluster.Spec.ProxyConfiguration.NoProxy...)

		// Add no-proxy defaults
		noProxyList = append(noProxyList, common.NoProxyDefaults...)
		cloudStackManagementApiEndpointHostname, err := getHostnameFromUrl(managementApiEndpoint)
		if err == nil {
			noProxyList = append(noProxyList, cloudStackManagementApiEndpointHostname)
		}
		noProxyList = append(noProxyList,
			clusterSpec.Cluster.Spec.ControlPlaneConfiguration.Endpoint.Host,
		)

		values["httpProxy"] = clusterSpec.Cluster.Spec.ProxyConfiguration.HttpProxy
		values["httpsProxy"] = clusterSpec.Cluster.Spec.ProxyConfiguration.HttpsProxy
		values["noProxy"] = noProxyList
	}

	return values
}

func (p *cloudstackProvider) generateCAPISpecForCreate(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) (controlPlaneSpec, workersSpec []byte, err error) {
	clusterName := clusterSpec.Cluster.Name

	cpOpt := func(values map[string]interface{}) {
		values["controlPlaneTemplateName"] = common.CPMachineTemplateName(clusterName, p.templateBuilder.now)
		values["cloudstackControlPlaneSshAuthorizedKey"] = p.controlPlaneSshAuthKey
		values["cloudstackEtcdSshAuthorizedKey"] = p.etcdSshAuthKey
		values["etcdTemplateName"] = common.EtcdMachineTemplateName(clusterName, p.templateBuilder.now)
	}
	controlPlaneSpec, err = p.templateBuilder.GenerateCAPISpecControlPlane(clusterSpec, cpOpt)
	if err != nil {
		return nil, nil, err
	}
	workersOpt := func(values map[string]interface{}) {
		values["workloadTemplateName"] = common.WorkerMachineTemplateName(clusterName, clusterSpec.Cluster.Spec.WorkerNodeGroupConfigurations[0].Name, p.templateBuilder.now)
		values["cloudstackWorkerSshAuthorizedKey"] = p.workerSshAuthKey
	}
	workersSpec, err = p.templateBuilder.GenerateCAPISpecWorkers(clusterSpec, workersOpt)
	if err != nil {
		return nil, nil, err
	}
	return controlPlaneSpec, workersSpec, nil
}

func (p *cloudstackProvider) GenerateCAPISpecForUpgrade(ctx context.Context, bootstrapCluster, workloadCluster *types.Cluster, currentSpec, clusterSpec *cluster.Spec) (controlPlaneSpec, workersSpec []byte, err error) {
	return nil, nil, fmt.Errorf("cloudstack provider does not support upgrade yet")
}

func (p *cloudstackProvider) GenerateCAPISpecForCreate(ctx context.Context, cluster *types.Cluster, clusterSpec *cluster.Spec) (controlPlaneSpec, workersSpec []byte, err error) {
	controlPlaneSpec, workersSpec, err = p.generateCAPISpecForCreate(ctx, cluster, clusterSpec)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating cluster api Spec contents: %v", err)
	}
	return controlPlaneSpec, workersSpec, nil
}

func (p *cloudstackProvider) GenerateMHC() ([]byte, error) {
	data := map[string]string{
		"clusterName":         p.clusterConfig.Name,
		"eksaSystemNamespace": constants.EksaSystemNamespace,
	}
	mhc, err := templater.Execute(string(mhcTemplate), data)
	if err != nil {
		return nil, err
	}
	return mhc, nil
}

func (p *cloudstackProvider) CleanupProviderInfrastructure(_ context.Context) error {
	return nil
}

func (p *cloudstackProvider) BootstrapSetup(ctx context.Context, clusterConfig *v1alpha1.Cluster, cluster *types.Cluster) error {
	// Nothing to do
	return nil
}

func (p *cloudstackProvider) Version(clusterSpec *cluster.Spec) string {
	return clusterSpec.VersionsBundle.CloudStack.Version
}

func (p *cloudstackProvider) EnvMap(_ *cluster.Spec) (map[string]string, error) {
	envMap := make(map[string]string)
	for _, key := range requiredEnvs {
		if env, ok := os.LookupEnv(key); ok && len(env) > 0 {
			envMap[key] = env
		} else {
			return envMap, fmt.Errorf("warning required env not set %s", key)
		}
	}
	return envMap, nil
}

func (p *cloudstackProvider) GetDeployments() map[string][]string {
	return map[string][]string{
		"capc-system": {"capc-controller-manager"},
	}
}

func (p *cloudstackProvider) GetInfrastructureBundle(clusterSpec *cluster.Spec) *types.InfrastructureBundle {
	bundle := clusterSpec.VersionsBundle
	folderName := fmt.Sprintf("infrastructure-cloudstack/%s/", bundle.CloudStack.Version)

	infraBundle := types.InfrastructureBundle{
		FolderName: folderName,
		Manifests: []releasev1alpha1.Manifest{
			bundle.CloudStack.Components,
			bundle.CloudStack.Metadata,
		},
	}
	return &infraBundle
}

func (p *cloudstackProvider) DatacenterConfig(_ *cluster.Spec) providers.DatacenterConfig {
	return p.datacenterConfig
}

func (p *cloudstackProvider) MachineConfigs(_ *cluster.Spec) []providers.MachineConfig {
	var configs []providers.MachineConfig
	controlPlaneMachineName := p.clusterConfig.Spec.ControlPlaneConfiguration.MachineGroupRef.Name
	workerMachineName := p.clusterConfig.Spec.WorkerNodeGroupConfigurations[0].MachineGroupRef.Name
	p.machineConfigs[controlPlaneMachineName].Annotations = map[string]string{p.clusterConfig.ControlPlaneAnnotation(): "true"}
	if p.clusterConfig.IsManaged() {
		p.machineConfigs[controlPlaneMachineName].SetManagement(p.clusterConfig.ManagedBy())
	}

	configs = append(configs, p.machineConfigs[controlPlaneMachineName])
	if workerMachineName != controlPlaneMachineName {
		configs = append(configs, p.machineConfigs[workerMachineName])
		if p.clusterConfig.IsManaged() {
			p.machineConfigs[workerMachineName].SetManagement(p.clusterConfig.ManagedBy())
		}
	}
	if p.clusterConfig.Spec.ExternalEtcdConfiguration != nil {
		etcdMachineName := p.clusterConfig.Spec.ExternalEtcdConfiguration.MachineGroupRef.Name
		p.machineConfigs[etcdMachineName].Annotations = map[string]string{p.clusterConfig.EtcdAnnotation(): "true"}
		if etcdMachineName != controlPlaneMachineName && etcdMachineName != workerMachineName {
			configs = append(configs, p.machineConfigs[etcdMachineName])
			p.machineConfigs[etcdMachineName].SetManagement(p.clusterConfig.ManagedBy())
		}
	}
	return configs
}

func (p *cloudstackProvider) RunPostUpgrade(ctx context.Context, clusterSpec *cluster.Spec, managementCluster, workloadCluster *types.Cluster) error {
	return fmt.Errorf("upgrade is not supported for CloudStack provider yet")
}

func (p *cloudstackProvider) UpgradeNeeded(ctx context.Context, newSpec, currentSpec *cluster.Spec) (bool, error) {
	return false, fmt.Errorf("upgrade is not supported for CloudStack provider yet")
}

func (p *cloudstackProvider) DeleteResources(ctx context.Context, clusterSpec *cluster.Spec) error {
	for _, mc := range p.machineConfigs {
		if err := p.providerKubectlClient.DeleteEksaCloudStackMachineConfig(ctx, mc.Name, clusterSpec.ManagementCluster.KubeconfigFile, mc.Namespace); err != nil {
			return err
		}
	}
	return p.providerKubectlClient.DeleteEksaCloudStackDatacenterConfig(ctx, p.datacenterConfig.Name, clusterSpec.ManagementCluster.KubeconfigFile, p.datacenterConfig.Namespace)
}

func (p *cloudstackProvider) GenerateStorageClass() []byte {
	return nil
}
