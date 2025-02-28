// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/clustermanager (interfaces: ClusterClient,Networking,AwsIamAuth)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	cluster "github.com/aws/eks-anywhere/pkg/cluster"
	executables "github.com/aws/eks-anywhere/pkg/executables"
	filewriter "github.com/aws/eks-anywhere/pkg/filewriter"
	providers "github.com/aws/eks-anywhere/pkg/providers"
	types "github.com/aws/eks-anywhere/pkg/types"
	v1alpha10 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	v1alpha11 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockClusterClient is a mock of ClusterClient interface.
type MockClusterClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterClientMockRecorder
}

// MockClusterClientMockRecorder is the mock recorder for MockClusterClient.
type MockClusterClientMockRecorder struct {
	mock *MockClusterClient
}

// NewMockClusterClient creates a new mock instance.
func NewMockClusterClient(ctrl *gomock.Controller) *MockClusterClient {
	mock := &MockClusterClient{ctrl: ctrl}
	mock.recorder = &MockClusterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterClient) EXPECT() *MockClusterClientMockRecorder {
	return m.recorder
}

// ApplyKubeSpecFromBytes mocks base method.
func (m *MockClusterClient) ApplyKubeSpecFromBytes(arg0 context.Context, arg1 *types.Cluster, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyKubeSpecFromBytes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyKubeSpecFromBytes indicates an expected call of ApplyKubeSpecFromBytes.
func (mr *MockClusterClientMockRecorder) ApplyKubeSpecFromBytes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyKubeSpecFromBytes", reflect.TypeOf((*MockClusterClient)(nil).ApplyKubeSpecFromBytes), arg0, arg1, arg2)
}

// ApplyKubeSpecFromBytesForce mocks base method.
func (m *MockClusterClient) ApplyKubeSpecFromBytesForce(arg0 context.Context, arg1 *types.Cluster, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyKubeSpecFromBytesForce", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyKubeSpecFromBytesForce indicates an expected call of ApplyKubeSpecFromBytesForce.
func (mr *MockClusterClientMockRecorder) ApplyKubeSpecFromBytesForce(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyKubeSpecFromBytesForce", reflect.TypeOf((*MockClusterClient)(nil).ApplyKubeSpecFromBytesForce), arg0, arg1, arg2)
}

// ApplyKubeSpecFromBytesWithNamespace mocks base method.
func (m *MockClusterClient) ApplyKubeSpecFromBytesWithNamespace(arg0 context.Context, arg1 *types.Cluster, arg2 []byte, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyKubeSpecFromBytesWithNamespace", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyKubeSpecFromBytesWithNamespace indicates an expected call of ApplyKubeSpecFromBytesWithNamespace.
func (mr *MockClusterClientMockRecorder) ApplyKubeSpecFromBytesWithNamespace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyKubeSpecFromBytesWithNamespace", reflect.TypeOf((*MockClusterClient)(nil).ApplyKubeSpecFromBytesWithNamespace), arg0, arg1, arg2, arg3)
}

// CreateNamespace mocks base method.
func (m *MockClusterClient) CreateNamespace(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockClusterClientMockRecorder) CreateNamespace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockClusterClient)(nil).CreateNamespace), arg0, arg1, arg2)
}

// DeleteAWSIamConfig mocks base method.
func (m *MockClusterClient) DeleteAWSIamConfig(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAWSIamConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAWSIamConfig indicates an expected call of DeleteAWSIamConfig.
func (mr *MockClusterClientMockRecorder) DeleteAWSIamConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAWSIamConfig", reflect.TypeOf((*MockClusterClient)(nil).DeleteAWSIamConfig), arg0, arg1, arg2, arg3)
}

// DeleteCluster mocks base method.
func (m *MockClusterClient) DeleteCluster(arg0 context.Context, arg1, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockClusterClientMockRecorder) DeleteCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterClient)(nil).DeleteCluster), arg0, arg1, arg2)
}

// DeleteEKSACluster mocks base method.
func (m *MockClusterClient) DeleteEKSACluster(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEKSACluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEKSACluster indicates an expected call of DeleteEKSACluster.
func (mr *MockClusterClientMockRecorder) DeleteEKSACluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEKSACluster", reflect.TypeOf((*MockClusterClient)(nil).DeleteEKSACluster), arg0, arg1, arg2, arg3)
}

// DeleteGitOpsConfig mocks base method.
func (m *MockClusterClient) DeleteGitOpsConfig(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGitOpsConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGitOpsConfig indicates an expected call of DeleteGitOpsConfig.
func (mr *MockClusterClientMockRecorder) DeleteGitOpsConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGitOpsConfig", reflect.TypeOf((*MockClusterClient)(nil).DeleteGitOpsConfig), arg0, arg1, arg2, arg3)
}

// DeleteOIDCConfig mocks base method.
func (m *MockClusterClient) DeleteOIDCConfig(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOIDCConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOIDCConfig indicates an expected call of DeleteOIDCConfig.
func (mr *MockClusterClientMockRecorder) DeleteOIDCConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOIDCConfig", reflect.TypeOf((*MockClusterClient)(nil).DeleteOIDCConfig), arg0, arg1, arg2, arg3)
}

// DeleteOldWorkerNodeGroup mocks base method.
func (m *MockClusterClient) DeleteOldWorkerNodeGroup(arg0 context.Context, arg1 *v1beta1.MachineDeployment, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOldWorkerNodeGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOldWorkerNodeGroup indicates an expected call of DeleteOldWorkerNodeGroup.
func (mr *MockClusterClientMockRecorder) DeleteOldWorkerNodeGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOldWorkerNodeGroup", reflect.TypeOf((*MockClusterClient)(nil).DeleteOldWorkerNodeGroup), arg0, arg1, arg2)
}

// GetApiServerUrl mocks base method.
func (m *MockClusterClient) GetApiServerUrl(arg0 context.Context, arg1 *types.Cluster) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiServerUrl", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiServerUrl indicates an expected call of GetApiServerUrl.
func (mr *MockClusterClientMockRecorder) GetApiServerUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiServerUrl", reflect.TypeOf((*MockClusterClient)(nil).GetApiServerUrl), arg0, arg1)
}

// GetBundles mocks base method.
func (m *MockClusterClient) GetBundles(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha10.Bundles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundles", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha10.Bundles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundles indicates an expected call of GetBundles.
func (mr *MockClusterClientMockRecorder) GetBundles(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundles", reflect.TypeOf((*MockClusterClient)(nil).GetBundles), arg0, arg1, arg2, arg3)
}

// GetClusterCATlsCert mocks base method.
func (m *MockClusterClient) GetClusterCATlsCert(arg0 context.Context, arg1 string, arg2 *types.Cluster, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterCATlsCert", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterCATlsCert indicates an expected call of GetClusterCATlsCert.
func (mr *MockClusterClientMockRecorder) GetClusterCATlsCert(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterCATlsCert", reflect.TypeOf((*MockClusterClient)(nil).GetClusterCATlsCert), arg0, arg1, arg2, arg3)
}

// GetClusters mocks base method.
func (m *MockClusterClient) GetClusters(arg0 context.Context, arg1 *types.Cluster) ([]types.CAPICluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusters", arg0, arg1)
	ret0, _ := ret[0].([]types.CAPICluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusters indicates an expected call of GetClusters.
func (mr *MockClusterClientMockRecorder) GetClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusters", reflect.TypeOf((*MockClusterClient)(nil).GetClusters), arg0, arg1)
}

// GetEksaCluster mocks base method.
func (m *MockClusterClient) GetEksaCluster(arg0 context.Context, arg1 *types.Cluster, arg2 string) (*v1alpha1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaCluster", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaCluster indicates an expected call of GetEksaCluster.
func (mr *MockClusterClientMockRecorder) GetEksaCluster(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaCluster", reflect.TypeOf((*MockClusterClient)(nil).GetEksaCluster), arg0, arg1, arg2)
}

// GetEksaGitOpsConfig mocks base method.
func (m *MockClusterClient) GetEksaGitOpsConfig(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha1.GitOpsConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaGitOpsConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.GitOpsConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaGitOpsConfig indicates an expected call of GetEksaGitOpsConfig.
func (mr *MockClusterClientMockRecorder) GetEksaGitOpsConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaGitOpsConfig", reflect.TypeOf((*MockClusterClient)(nil).GetEksaGitOpsConfig), arg0, arg1, arg2, arg3)
}

// GetEksaOIDCConfig mocks base method.
func (m *MockClusterClient) GetEksaOIDCConfig(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha1.OIDCConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaOIDCConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.OIDCConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaOIDCConfig indicates an expected call of GetEksaOIDCConfig.
func (mr *MockClusterClientMockRecorder) GetEksaOIDCConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaOIDCConfig", reflect.TypeOf((*MockClusterClient)(nil).GetEksaOIDCConfig), arg0, arg1, arg2, arg3)
}

// GetEksaVSphereDatacenterConfig mocks base method.
func (m *MockClusterClient) GetEksaVSphereDatacenterConfig(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha1.VSphereDatacenterConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaVSphereDatacenterConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.VSphereDatacenterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaVSphereDatacenterConfig indicates an expected call of GetEksaVSphereDatacenterConfig.
func (mr *MockClusterClientMockRecorder) GetEksaVSphereDatacenterConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaVSphereDatacenterConfig", reflect.TypeOf((*MockClusterClient)(nil).GetEksaVSphereDatacenterConfig), arg0, arg1, arg2, arg3)
}

// GetEksaVSphereMachineConfig mocks base method.
func (m *MockClusterClient) GetEksaVSphereMachineConfig(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha1.VSphereMachineConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksaVSphereMachineConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha1.VSphereMachineConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksaVSphereMachineConfig indicates an expected call of GetEksaVSphereMachineConfig.
func (mr *MockClusterClientMockRecorder) GetEksaVSphereMachineConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksaVSphereMachineConfig", reflect.TypeOf((*MockClusterClient)(nil).GetEksaVSphereMachineConfig), arg0, arg1, arg2, arg3)
}

// GetEksdRelease mocks base method.
func (m *MockClusterClient) GetEksdRelease(arg0 context.Context, arg1, arg2, arg3 string) (*v1alpha11.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEksdRelease", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha11.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEksdRelease indicates an expected call of GetEksdRelease.
func (mr *MockClusterClientMockRecorder) GetEksdRelease(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEksdRelease", reflect.TypeOf((*MockClusterClient)(nil).GetEksdRelease), arg0, arg1, arg2, arg3)
}

// GetMachineDeployment mocks base method.
func (m *MockClusterClient) GetMachineDeployment(arg0 context.Context, arg1 string, arg2 ...executables.KubectlOpt) (*v1beta1.MachineDeployment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMachineDeployment", varargs...)
	ret0, _ := ret[0].(*v1beta1.MachineDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineDeployment indicates an expected call of GetMachineDeployment.
func (mr *MockClusterClientMockRecorder) GetMachineDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineDeployment", reflect.TypeOf((*MockClusterClient)(nil).GetMachineDeployment), varargs...)
}

// GetMachines mocks base method.
func (m *MockClusterClient) GetMachines(arg0 context.Context, arg1 *types.Cluster, arg2 string) ([]types.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachines", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachines indicates an expected call of GetMachines.
func (mr *MockClusterClientMockRecorder) GetMachines(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachines", reflect.TypeOf((*MockClusterClient)(nil).GetMachines), arg0, arg1, arg2)
}

// GetNamespace mocks base method.
func (m *MockClusterClient) GetNamespace(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockClusterClientMockRecorder) GetNamespace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockClusterClient)(nil).GetNamespace), arg0, arg1, arg2)
}

// GetWorkloadKubeconfig mocks base method.
func (m *MockClusterClient) GetWorkloadKubeconfig(arg0 context.Context, arg1 string, arg2 *types.Cluster) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkloadKubeconfig", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkloadKubeconfig indicates an expected call of GetWorkloadKubeconfig.
func (mr *MockClusterClientMockRecorder) GetWorkloadKubeconfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkloadKubeconfig", reflect.TypeOf((*MockClusterClient)(nil).GetWorkloadKubeconfig), arg0, arg1, arg2)
}

// InitInfrastructure mocks base method.
func (m *MockClusterClient) InitInfrastructure(arg0 context.Context, arg1 *cluster.Spec, arg2 *types.Cluster, arg3 providers.Provider) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitInfrastructure", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitInfrastructure indicates an expected call of InitInfrastructure.
func (mr *MockClusterClientMockRecorder) InitInfrastructure(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitInfrastructure", reflect.TypeOf((*MockClusterClient)(nil).InitInfrastructure), arg0, arg1, arg2, arg3)
}

// KubeconfigSecretAvailable mocks base method.
func (m *MockClusterClient) KubeconfigSecretAvailable(arg0 context.Context, arg1, arg2, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KubeconfigSecretAvailable", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KubeconfigSecretAvailable indicates an expected call of KubeconfigSecretAvailable.
func (mr *MockClusterClientMockRecorder) KubeconfigSecretAvailable(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KubeconfigSecretAvailable", reflect.TypeOf((*MockClusterClient)(nil).KubeconfigSecretAvailable), arg0, arg1, arg2, arg3)
}

// MoveManagement mocks base method.
func (m *MockClusterClient) MoveManagement(arg0 context.Context, arg1, arg2 *types.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveManagement", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveManagement indicates an expected call of MoveManagement.
func (mr *MockClusterClientMockRecorder) MoveManagement(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveManagement", reflect.TypeOf((*MockClusterClient)(nil).MoveManagement), arg0, arg1, arg2)
}

// RemoveAnnotationInNamespace mocks base method.
func (m *MockClusterClient) RemoveAnnotationInNamespace(arg0 context.Context, arg1, arg2, arg3 string, arg4 *types.Cluster, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAnnotationInNamespace", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAnnotationInNamespace indicates an expected call of RemoveAnnotationInNamespace.
func (mr *MockClusterClientMockRecorder) RemoveAnnotationInNamespace(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAnnotationInNamespace", reflect.TypeOf((*MockClusterClient)(nil).RemoveAnnotationInNamespace), arg0, arg1, arg2, arg3, arg4, arg5)
}

// SaveLog mocks base method.
func (m *MockClusterClient) SaveLog(arg0 context.Context, arg1 *types.Cluster, arg2 *types.Deployment, arg3 string, arg4 filewriter.FileWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLog", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLog indicates an expected call of SaveLog.
func (mr *MockClusterClientMockRecorder) SaveLog(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLog", reflect.TypeOf((*MockClusterClient)(nil).SaveLog), arg0, arg1, arg2, arg3, arg4)
}

// UpdateAnnotationInNamespace mocks base method.
func (m *MockClusterClient) UpdateAnnotationInNamespace(arg0 context.Context, arg1, arg2 string, arg3 map[string]string, arg4 *types.Cluster, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnnotationInNamespace", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAnnotationInNamespace indicates an expected call of UpdateAnnotationInNamespace.
func (mr *MockClusterClientMockRecorder) UpdateAnnotationInNamespace(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnnotationInNamespace", reflect.TypeOf((*MockClusterClient)(nil).UpdateAnnotationInNamespace), arg0, arg1, arg2, arg3, arg4, arg5)
}

// UpdateEnvironmentVariablesInNamespace mocks base method.
func (m *MockClusterClient) UpdateEnvironmentVariablesInNamespace(arg0 context.Context, arg1, arg2 string, arg3 map[string]string, arg4 *types.Cluster, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEnvironmentVariablesInNamespace", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnvironmentVariablesInNamespace indicates an expected call of UpdateEnvironmentVariablesInNamespace.
func (mr *MockClusterClientMockRecorder) UpdateEnvironmentVariablesInNamespace(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvironmentVariablesInNamespace", reflect.TypeOf((*MockClusterClient)(nil).UpdateEnvironmentVariablesInNamespace), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ValidateControlPlaneNodes mocks base method.
func (m *MockClusterClient) ValidateControlPlaneNodes(arg0 context.Context, arg1 *types.Cluster, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateControlPlaneNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateControlPlaneNodes indicates an expected call of ValidateControlPlaneNodes.
func (mr *MockClusterClientMockRecorder) ValidateControlPlaneNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateControlPlaneNodes", reflect.TypeOf((*MockClusterClient)(nil).ValidateControlPlaneNodes), arg0, arg1, arg2)
}

// ValidateWorkerNodes mocks base method.
func (m *MockClusterClient) ValidateWorkerNodes(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateWorkerNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateWorkerNodes indicates an expected call of ValidateWorkerNodes.
func (mr *MockClusterClientMockRecorder) ValidateWorkerNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateWorkerNodes", reflect.TypeOf((*MockClusterClient)(nil).ValidateWorkerNodes), arg0, arg1, arg2)
}

// WaitForControlPlaneReady mocks base method.
func (m *MockClusterClient) WaitForControlPlaneReady(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForControlPlaneReady", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForControlPlaneReady indicates an expected call of WaitForControlPlaneReady.
func (mr *MockClusterClientMockRecorder) WaitForControlPlaneReady(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForControlPlaneReady", reflect.TypeOf((*MockClusterClient)(nil).WaitForControlPlaneReady), arg0, arg1, arg2, arg3)
}

// WaitForDeployment mocks base method.
func (m *MockClusterClient) WaitForDeployment(arg0 context.Context, arg1 *types.Cluster, arg2, arg3, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForDeployment", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForDeployment indicates an expected call of WaitForDeployment.
func (mr *MockClusterClientMockRecorder) WaitForDeployment(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForDeployment", reflect.TypeOf((*MockClusterClient)(nil).WaitForDeployment), arg0, arg1, arg2, arg3, arg4, arg5)
}

// WaitForManagedExternalEtcdReady mocks base method.
func (m *MockClusterClient) WaitForManagedExternalEtcdReady(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForManagedExternalEtcdReady", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForManagedExternalEtcdReady indicates an expected call of WaitForManagedExternalEtcdReady.
func (mr *MockClusterClientMockRecorder) WaitForManagedExternalEtcdReady(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForManagedExternalEtcdReady", reflect.TypeOf((*MockClusterClient)(nil).WaitForManagedExternalEtcdReady), arg0, arg1, arg2, arg3)
}

// MockNetworking is a mock of Networking interface.
type MockNetworking struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingMockRecorder
}

// MockNetworkingMockRecorder is the mock recorder for MockNetworking.
type MockNetworkingMockRecorder struct {
	mock *MockNetworking
}

// NewMockNetworking creates a new mock instance.
func NewMockNetworking(ctrl *gomock.Controller) *MockNetworking {
	mock := &MockNetworking{ctrl: ctrl}
	mock.recorder = &MockNetworkingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworking) EXPECT() *MockNetworkingMockRecorder {
	return m.recorder
}

// GenerateManifest mocks base method.
func (m *MockNetworking) GenerateManifest(arg0 context.Context, arg1 *cluster.Spec, arg2 []string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateManifest", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateManifest indicates an expected call of GenerateManifest.
func (mr *MockNetworkingMockRecorder) GenerateManifest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateManifest", reflect.TypeOf((*MockNetworking)(nil).GenerateManifest), arg0, arg1, arg2)
}

// Upgrade mocks base method.
func (m *MockNetworking) Upgrade(arg0 context.Context, arg1 *types.Cluster, arg2, arg3 *cluster.Spec) (*types.ChangeDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.ChangeDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockNetworkingMockRecorder) Upgrade(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockNetworking)(nil).Upgrade), arg0, arg1, arg2, arg3)
}

// MockAwsIamAuth is a mock of AwsIamAuth interface.
type MockAwsIamAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAwsIamAuthMockRecorder
}

// MockAwsIamAuthMockRecorder is the mock recorder for MockAwsIamAuth.
type MockAwsIamAuthMockRecorder struct {
	mock *MockAwsIamAuth
}

// NewMockAwsIamAuth creates a new mock instance.
func NewMockAwsIamAuth(ctrl *gomock.Controller) *MockAwsIamAuth {
	mock := &MockAwsIamAuth{ctrl: ctrl}
	mock.recorder = &MockAwsIamAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAwsIamAuth) EXPECT() *MockAwsIamAuthMockRecorder {
	return m.recorder
}

// GenerateAwsIamAuthKubeconfig mocks base method.
func (m *MockAwsIamAuth) GenerateAwsIamAuthKubeconfig(arg0 *cluster.Spec, arg1, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAwsIamAuthKubeconfig", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAwsIamAuthKubeconfig indicates an expected call of GenerateAwsIamAuthKubeconfig.
func (mr *MockAwsIamAuthMockRecorder) GenerateAwsIamAuthKubeconfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAwsIamAuthKubeconfig", reflect.TypeOf((*MockAwsIamAuth)(nil).GenerateAwsIamAuthKubeconfig), arg0, arg1, arg2)
}

// GenerateCertKeyPairSecret mocks base method.
func (m *MockAwsIamAuth) GenerateCertKeyPairSecret() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCertKeyPairSecret")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCertKeyPairSecret indicates an expected call of GenerateCertKeyPairSecret.
func (mr *MockAwsIamAuthMockRecorder) GenerateCertKeyPairSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCertKeyPairSecret", reflect.TypeOf((*MockAwsIamAuth)(nil).GenerateCertKeyPairSecret))
}

// GenerateManifest mocks base method.
func (m *MockAwsIamAuth) GenerateManifest(arg0 *cluster.Spec) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateManifest", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateManifest indicates an expected call of GenerateManifest.
func (mr *MockAwsIamAuthMockRecorder) GenerateManifest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateManifest", reflect.TypeOf((*MockAwsIamAuth)(nil).GenerateManifest), arg0)
}
