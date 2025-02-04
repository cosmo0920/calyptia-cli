// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package aws

import (
	"context"
	"sync"
)

// Ensure, that ClientMock does implement Client.
// If this is not the case, regenerate this file with moq.
var _ Client = &ClientMock{}

// ClientMock is a mock implementation of Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked Client
// 		mockedClient := &ClientMock{
// 			CreateInstanceFunc: func(ctx context.Context, in *CreateInstanceParams) (CreatedInstance, error) {
// 				panic("mock out the CreateInstance method")
// 			},
// 			CreateUserdataFunc: func(in *CreateUserDataParams) (string, error) {
// 				panic("mock out the CreateUserdata method")
// 			},
// 			DeleteInstanceFunc: func(ctx context.Context, instanceID string) error {
// 				panic("mock out the DeleteInstance method")
// 			},
// 			DeleteKeyPairFunc: func(ctx context.Context, keyPairID string) error {
// 				panic("mock out the DeleteKeyPair method")
// 			},
// 			DeleteResourcesFunc: func(ctx context.Context, resources []Resource) error {
// 				panic("mock out the DeleteResources method")
// 			},
// 			DeleteSecurityGroupFunc: func(ctx context.Context, securityGroupName string) error {
// 				panic("mock out the DeleteSecurityGroup method")
// 			},
// 			EnsureAndAssociateElasticIPv4AddressFunc: func(ctx context.Context, instanceID string, environment string, elasticIPv4AddressPool string, elasticIPv4Address string) (string, error) {
// 				panic("mock out the EnsureAndAssociateElasticIPv4Address method")
// 			},
// 			EnsureInstanceTypeFunc: func(ctx context.Context, instanceTypeName string) (string, error) {
// 				panic("mock out the EnsureInstanceType method")
// 			},
// 			EnsureKeyPairFunc: func(ctx context.Context, keyPairName string, environment string) (string, error) {
// 				panic("mock out the EnsureKeyPair method")
// 			},
// 			EnsureSecurityGroupFunc: func(ctx context.Context, securityGroupName string, environment string, vpcID string) (string, error) {
// 				panic("mock out the EnsureSecurityGroup method")
// 			},
// 			EnsureSecurityGroupIngressRulesFunc: func(ctx context.Context, securityGroupID string) error {
// 				panic("mock out the EnsureSecurityGroupIngressRules method")
// 			},
// 			EnsureSubnetFunc: func(ctx context.Context, subNetID string) (string, error) {
// 				panic("mock out the EnsureSubnet method")
// 			},
// 			FindMatchingAMIFunc: func(ctx context.Context, region string, version string) (string, error) {
// 				panic("mock out the FindMatchingAMI method")
// 			},
// 			GetResourcesByTagsFunc: func(ctx context.Context, tags TagSpec) ([]Resource, error) {
// 				panic("mock out the GetResourcesByTags method")
// 			},
// 			InstanceStateFunc: func(ctx context.Context, instanceID string) (string, error) {
// 				panic("mock out the InstanceState method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// CreateInstanceFunc mocks the CreateInstance method.
	CreateInstanceFunc func(ctx context.Context, in *CreateInstanceParams) (CreatedInstance, error)

	// CreateUserdataFunc mocks the CreateUserdata method.
	CreateUserdataFunc func(in *CreateUserDataParams) (string, error)

	// DeleteInstanceFunc mocks the DeleteInstance method.
	DeleteInstanceFunc func(ctx context.Context, instanceID string) error

	// DeleteKeyPairFunc mocks the DeleteKeyPair method.
	DeleteKeyPairFunc func(ctx context.Context, keyPairID string) error

	// DeleteResourcesFunc mocks the DeleteResources method.
	DeleteResourcesFunc func(ctx context.Context, resources []Resource) error

	// DeleteSecurityGroupFunc mocks the DeleteSecurityGroup method.
	DeleteSecurityGroupFunc func(ctx context.Context, securityGroupName string) error

	// EnsureAndAssociateElasticIPv4AddressFunc mocks the EnsureAndAssociateElasticIPv4Address method.
	EnsureAndAssociateElasticIPv4AddressFunc func(ctx context.Context, instanceID string, environment string, elasticIPv4AddressPool string, elasticIPv4Address string) (string, error)

	// EnsureInstanceTypeFunc mocks the EnsureInstanceType method.
	EnsureInstanceTypeFunc func(ctx context.Context, instanceTypeName string) (string, error)

	// EnsureKeyPairFunc mocks the EnsureKeyPair method.
	EnsureKeyPairFunc func(ctx context.Context, keyPairName string, environment string) (string, error)

	// EnsureSecurityGroupFunc mocks the EnsureSecurityGroup method.
	EnsureSecurityGroupFunc func(ctx context.Context, securityGroupName string, environment string, vpcID string) (string, error)

	// EnsureSecurityGroupIngressRulesFunc mocks the EnsureSecurityGroupIngressRules method.
	EnsureSecurityGroupIngressRulesFunc func(ctx context.Context, securityGroupID string) error

	// EnsureSubnetFunc mocks the EnsureSubnet method.
	EnsureSubnetFunc func(ctx context.Context, subNetID string) (string, error)

	// FindMatchingAMIFunc mocks the FindMatchingAMI method.
	FindMatchingAMIFunc func(ctx context.Context, region string, version string) (string, error)

	// GetResourcesByTagsFunc mocks the GetResourcesByTags method.
	GetResourcesByTagsFunc func(ctx context.Context, tags TagSpec) ([]Resource, error)

	// InstanceStateFunc mocks the InstanceState method.
	InstanceStateFunc func(ctx context.Context, instanceID string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateInstance holds details about calls to the CreateInstance method.
		CreateInstance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *CreateInstanceParams
		}
		// CreateUserdata holds details about calls to the CreateUserdata method.
		CreateUserdata []struct {
			// In is the in argument value.
			In *CreateUserDataParams
		}
		// DeleteInstance holds details about calls to the DeleteInstance method.
		DeleteInstance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// InstanceID is the instanceID argument value.
			InstanceID string
		}
		// DeleteKeyPair holds details about calls to the DeleteKeyPair method.
		DeleteKeyPair []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// KeyPairID is the keyPairID argument value.
			KeyPairID string
		}
		// DeleteResources holds details about calls to the DeleteResources method.
		DeleteResources []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Resources is the resources argument value.
			Resources []Resource
		}
		// DeleteSecurityGroup holds details about calls to the DeleteSecurityGroup method.
		DeleteSecurityGroup []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SecurityGroupName is the securityGroupName argument value.
			SecurityGroupName string
		}
		// EnsureAndAssociateElasticIPv4Address holds details about calls to the EnsureAndAssociateElasticIPv4Address method.
		EnsureAndAssociateElasticIPv4Address []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// InstanceID is the instanceID argument value.
			InstanceID string
			// Environment is the environment argument value.
			Environment string
			// ElasticIPv4AddressPool is the elasticIPv4AddressPool argument value.
			ElasticIPv4AddressPool string
			// ElasticIPv4Address is the elasticIPv4Address argument value.
			ElasticIPv4Address string
		}
		// EnsureInstanceType holds details about calls to the EnsureInstanceType method.
		EnsureInstanceType []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// InstanceTypeName is the instanceTypeName argument value.
			InstanceTypeName string
		}
		// EnsureKeyPair holds details about calls to the EnsureKeyPair method.
		EnsureKeyPair []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// KeyPairName is the keyPairName argument value.
			KeyPairName string
			// Environment is the environment argument value.
			Environment string
		}
		// EnsureSecurityGroup holds details about calls to the EnsureSecurityGroup method.
		EnsureSecurityGroup []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SecurityGroupName is the securityGroupName argument value.
			SecurityGroupName string
			// Environment is the environment argument value.
			Environment string
			// VpcID is the vpcID argument value.
			VpcID string
		}
		// EnsureSecurityGroupIngressRules holds details about calls to the EnsureSecurityGroupIngressRules method.
		EnsureSecurityGroupIngressRules []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SecurityGroupID is the securityGroupID argument value.
			SecurityGroupID string
		}
		// EnsureSubnet holds details about calls to the EnsureSubnet method.
		EnsureSubnet []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SubNetID is the subNetID argument value.
			SubNetID string
		}
		// FindMatchingAMI holds details about calls to the FindMatchingAMI method.
		FindMatchingAMI []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Region is the region argument value.
			Region string
			// Version is the version argument value.
			Version string
		}
		// GetResourcesByTags holds details about calls to the GetResourcesByTags method.
		GetResourcesByTags []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tags is the tags argument value.
			Tags TagSpec
		}
		// InstanceState holds details about calls to the InstanceState method.
		InstanceState []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// InstanceID is the instanceID argument value.
			InstanceID string
		}
	}
	lockCreateInstance                       sync.RWMutex
	lockCreateUserdata                       sync.RWMutex
	lockDeleteInstance                       sync.RWMutex
	lockDeleteKeyPair                        sync.RWMutex
	lockDeleteResources                      sync.RWMutex
	lockDeleteSecurityGroup                  sync.RWMutex
	lockEnsureAndAssociateElasticIPv4Address sync.RWMutex
	lockEnsureInstanceType                   sync.RWMutex
	lockEnsureKeyPair                        sync.RWMutex
	lockEnsureSecurityGroup                  sync.RWMutex
	lockEnsureSecurityGroupIngressRules      sync.RWMutex
	lockEnsureSubnet                         sync.RWMutex
	lockFindMatchingAMI                      sync.RWMutex
	lockGetResourcesByTags                   sync.RWMutex
	lockInstanceState                        sync.RWMutex
}

// CreateInstance calls CreateInstanceFunc.
func (mock *ClientMock) CreateInstance(ctx context.Context, in *CreateInstanceParams) (CreatedInstance, error) {
	if mock.CreateInstanceFunc == nil {
		panic("ClientMock.CreateInstanceFunc: method is nil but Client.CreateInstance was just called")
	}
	callInfo := struct {
		Ctx context.Context
		In  *CreateInstanceParams
	}{
		Ctx: ctx,
		In:  in,
	}
	mock.lockCreateInstance.Lock()
	mock.calls.CreateInstance = append(mock.calls.CreateInstance, callInfo)
	mock.lockCreateInstance.Unlock()
	return mock.CreateInstanceFunc(ctx, in)
}

// CreateInstanceCalls gets all the calls that were made to CreateInstance.
// Check the length with:
//     len(mockedClient.CreateInstanceCalls())
func (mock *ClientMock) CreateInstanceCalls() []struct {
	Ctx context.Context
	In  *CreateInstanceParams
} {
	var calls []struct {
		Ctx context.Context
		In  *CreateInstanceParams
	}
	mock.lockCreateInstance.RLock()
	calls = mock.calls.CreateInstance
	mock.lockCreateInstance.RUnlock()
	return calls
}

// CreateUserdata calls CreateUserdataFunc.
func (mock *ClientMock) CreateUserdata(in *CreateUserDataParams) (string, error) {
	if mock.CreateUserdataFunc == nil {
		panic("ClientMock.CreateUserdataFunc: method is nil but Client.CreateUserdata was just called")
	}
	callInfo := struct {
		In *CreateUserDataParams
	}{
		In: in,
	}
	mock.lockCreateUserdata.Lock()
	mock.calls.CreateUserdata = append(mock.calls.CreateUserdata, callInfo)
	mock.lockCreateUserdata.Unlock()
	return mock.CreateUserdataFunc(in)
}

// CreateUserdataCalls gets all the calls that were made to CreateUserdata.
// Check the length with:
//     len(mockedClient.CreateUserdataCalls())
func (mock *ClientMock) CreateUserdataCalls() []struct {
	In *CreateUserDataParams
} {
	var calls []struct {
		In *CreateUserDataParams
	}
	mock.lockCreateUserdata.RLock()
	calls = mock.calls.CreateUserdata
	mock.lockCreateUserdata.RUnlock()
	return calls
}

// DeleteInstance calls DeleteInstanceFunc.
func (mock *ClientMock) DeleteInstance(ctx context.Context, instanceID string) error {
	if mock.DeleteInstanceFunc == nil {
		panic("ClientMock.DeleteInstanceFunc: method is nil but Client.DeleteInstance was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		InstanceID string
	}{
		Ctx:        ctx,
		InstanceID: instanceID,
	}
	mock.lockDeleteInstance.Lock()
	mock.calls.DeleteInstance = append(mock.calls.DeleteInstance, callInfo)
	mock.lockDeleteInstance.Unlock()
	return mock.DeleteInstanceFunc(ctx, instanceID)
}

// DeleteInstanceCalls gets all the calls that were made to DeleteInstance.
// Check the length with:
//     len(mockedClient.DeleteInstanceCalls())
func (mock *ClientMock) DeleteInstanceCalls() []struct {
	Ctx        context.Context
	InstanceID string
} {
	var calls []struct {
		Ctx        context.Context
		InstanceID string
	}
	mock.lockDeleteInstance.RLock()
	calls = mock.calls.DeleteInstance
	mock.lockDeleteInstance.RUnlock()
	return calls
}

// DeleteKeyPair calls DeleteKeyPairFunc.
func (mock *ClientMock) DeleteKeyPair(ctx context.Context, keyPairID string) error {
	if mock.DeleteKeyPairFunc == nil {
		panic("ClientMock.DeleteKeyPairFunc: method is nil but Client.DeleteKeyPair was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		KeyPairID string
	}{
		Ctx:       ctx,
		KeyPairID: keyPairID,
	}
	mock.lockDeleteKeyPair.Lock()
	mock.calls.DeleteKeyPair = append(mock.calls.DeleteKeyPair, callInfo)
	mock.lockDeleteKeyPair.Unlock()
	return mock.DeleteKeyPairFunc(ctx, keyPairID)
}

// DeleteKeyPairCalls gets all the calls that were made to DeleteKeyPair.
// Check the length with:
//     len(mockedClient.DeleteKeyPairCalls())
func (mock *ClientMock) DeleteKeyPairCalls() []struct {
	Ctx       context.Context
	KeyPairID string
} {
	var calls []struct {
		Ctx       context.Context
		KeyPairID string
	}
	mock.lockDeleteKeyPair.RLock()
	calls = mock.calls.DeleteKeyPair
	mock.lockDeleteKeyPair.RUnlock()
	return calls
}

// DeleteResources calls DeleteResourcesFunc.
func (mock *ClientMock) DeleteResources(ctx context.Context, resources []Resource) error {
	if mock.DeleteResourcesFunc == nil {
		panic("ClientMock.DeleteResourcesFunc: method is nil but Client.DeleteResources was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Resources []Resource
	}{
		Ctx:       ctx,
		Resources: resources,
	}
	mock.lockDeleteResources.Lock()
	mock.calls.DeleteResources = append(mock.calls.DeleteResources, callInfo)
	mock.lockDeleteResources.Unlock()
	return mock.DeleteResourcesFunc(ctx, resources)
}

// DeleteResourcesCalls gets all the calls that were made to DeleteResources.
// Check the length with:
//     len(mockedClient.DeleteResourcesCalls())
func (mock *ClientMock) DeleteResourcesCalls() []struct {
	Ctx       context.Context
	Resources []Resource
} {
	var calls []struct {
		Ctx       context.Context
		Resources []Resource
	}
	mock.lockDeleteResources.RLock()
	calls = mock.calls.DeleteResources
	mock.lockDeleteResources.RUnlock()
	return calls
}

// DeleteSecurityGroup calls DeleteSecurityGroupFunc.
func (mock *ClientMock) DeleteSecurityGroup(ctx context.Context, securityGroupName string) error {
	if mock.DeleteSecurityGroupFunc == nil {
		panic("ClientMock.DeleteSecurityGroupFunc: method is nil but Client.DeleteSecurityGroup was just called")
	}
	callInfo := struct {
		Ctx               context.Context
		SecurityGroupName string
	}{
		Ctx:               ctx,
		SecurityGroupName: securityGroupName,
	}
	mock.lockDeleteSecurityGroup.Lock()
	mock.calls.DeleteSecurityGroup = append(mock.calls.DeleteSecurityGroup, callInfo)
	mock.lockDeleteSecurityGroup.Unlock()
	return mock.DeleteSecurityGroupFunc(ctx, securityGroupName)
}

// DeleteSecurityGroupCalls gets all the calls that were made to DeleteSecurityGroup.
// Check the length with:
//     len(mockedClient.DeleteSecurityGroupCalls())
func (mock *ClientMock) DeleteSecurityGroupCalls() []struct {
	Ctx               context.Context
	SecurityGroupName string
} {
	var calls []struct {
		Ctx               context.Context
		SecurityGroupName string
	}
	mock.lockDeleteSecurityGroup.RLock()
	calls = mock.calls.DeleteSecurityGroup
	mock.lockDeleteSecurityGroup.RUnlock()
	return calls
}

// EnsureAndAssociateElasticIPv4Address calls EnsureAndAssociateElasticIPv4AddressFunc.
func (mock *ClientMock) EnsureAndAssociateElasticIPv4Address(ctx context.Context, instanceID string, environment string, elasticIPv4AddressPool string, elasticIPv4Address string) (string, error) {
	if mock.EnsureAndAssociateElasticIPv4AddressFunc == nil {
		panic("ClientMock.EnsureAndAssociateElasticIPv4AddressFunc: method is nil but Client.EnsureAndAssociateElasticIPv4Address was just called")
	}
	callInfo := struct {
		Ctx                    context.Context
		InstanceID             string
		Environment            string
		ElasticIPv4AddressPool string
		ElasticIPv4Address     string
	}{
		Ctx:                    ctx,
		InstanceID:             instanceID,
		Environment:            environment,
		ElasticIPv4AddressPool: elasticIPv4AddressPool,
		ElasticIPv4Address:     elasticIPv4Address,
	}
	mock.lockEnsureAndAssociateElasticIPv4Address.Lock()
	mock.calls.EnsureAndAssociateElasticIPv4Address = append(mock.calls.EnsureAndAssociateElasticIPv4Address, callInfo)
	mock.lockEnsureAndAssociateElasticIPv4Address.Unlock()
	return mock.EnsureAndAssociateElasticIPv4AddressFunc(ctx, instanceID, environment, elasticIPv4AddressPool, elasticIPv4Address)
}

// EnsureAndAssociateElasticIPv4AddressCalls gets all the calls that were made to EnsureAndAssociateElasticIPv4Address.
// Check the length with:
//     len(mockedClient.EnsureAndAssociateElasticIPv4AddressCalls())
func (mock *ClientMock) EnsureAndAssociateElasticIPv4AddressCalls() []struct {
	Ctx                    context.Context
	InstanceID             string
	Environment            string
	ElasticIPv4AddressPool string
	ElasticIPv4Address     string
} {
	var calls []struct {
		Ctx                    context.Context
		InstanceID             string
		Environment            string
		ElasticIPv4AddressPool string
		ElasticIPv4Address     string
	}
	mock.lockEnsureAndAssociateElasticIPv4Address.RLock()
	calls = mock.calls.EnsureAndAssociateElasticIPv4Address
	mock.lockEnsureAndAssociateElasticIPv4Address.RUnlock()
	return calls
}

// EnsureInstanceType calls EnsureInstanceTypeFunc.
func (mock *ClientMock) EnsureInstanceType(ctx context.Context, instanceTypeName string) (string, error) {
	if mock.EnsureInstanceTypeFunc == nil {
		panic("ClientMock.EnsureInstanceTypeFunc: method is nil but Client.EnsureInstanceType was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		InstanceTypeName string
	}{
		Ctx:              ctx,
		InstanceTypeName: instanceTypeName,
	}
	mock.lockEnsureInstanceType.Lock()
	mock.calls.EnsureInstanceType = append(mock.calls.EnsureInstanceType, callInfo)
	mock.lockEnsureInstanceType.Unlock()
	return mock.EnsureInstanceTypeFunc(ctx, instanceTypeName)
}

// EnsureInstanceTypeCalls gets all the calls that were made to EnsureInstanceType.
// Check the length with:
//     len(mockedClient.EnsureInstanceTypeCalls())
func (mock *ClientMock) EnsureInstanceTypeCalls() []struct {
	Ctx              context.Context
	InstanceTypeName string
} {
	var calls []struct {
		Ctx              context.Context
		InstanceTypeName string
	}
	mock.lockEnsureInstanceType.RLock()
	calls = mock.calls.EnsureInstanceType
	mock.lockEnsureInstanceType.RUnlock()
	return calls
}

// EnsureKeyPair calls EnsureKeyPairFunc.
func (mock *ClientMock) EnsureKeyPair(ctx context.Context, keyPairName string, environment string) (string, error) {
	if mock.EnsureKeyPairFunc == nil {
		panic("ClientMock.EnsureKeyPairFunc: method is nil but Client.EnsureKeyPair was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		KeyPairName string
		Environment string
	}{
		Ctx:         ctx,
		KeyPairName: keyPairName,
		Environment: environment,
	}
	mock.lockEnsureKeyPair.Lock()
	mock.calls.EnsureKeyPair = append(mock.calls.EnsureKeyPair, callInfo)
	mock.lockEnsureKeyPair.Unlock()
	return mock.EnsureKeyPairFunc(ctx, keyPairName, environment)
}

// EnsureKeyPairCalls gets all the calls that were made to EnsureKeyPair.
// Check the length with:
//     len(mockedClient.EnsureKeyPairCalls())
func (mock *ClientMock) EnsureKeyPairCalls() []struct {
	Ctx         context.Context
	KeyPairName string
	Environment string
} {
	var calls []struct {
		Ctx         context.Context
		KeyPairName string
		Environment string
	}
	mock.lockEnsureKeyPair.RLock()
	calls = mock.calls.EnsureKeyPair
	mock.lockEnsureKeyPair.RUnlock()
	return calls
}

// EnsureSecurityGroup calls EnsureSecurityGroupFunc.
func (mock *ClientMock) EnsureSecurityGroup(ctx context.Context, securityGroupName string, environment string, vpcID string) (string, error) {
	if mock.EnsureSecurityGroupFunc == nil {
		panic("ClientMock.EnsureSecurityGroupFunc: method is nil but Client.EnsureSecurityGroup was just called")
	}
	callInfo := struct {
		Ctx               context.Context
		SecurityGroupName string
		Environment       string
		VpcID             string
	}{
		Ctx:               ctx,
		SecurityGroupName: securityGroupName,
		Environment:       environment,
		VpcID:             vpcID,
	}
	mock.lockEnsureSecurityGroup.Lock()
	mock.calls.EnsureSecurityGroup = append(mock.calls.EnsureSecurityGroup, callInfo)
	mock.lockEnsureSecurityGroup.Unlock()
	return mock.EnsureSecurityGroupFunc(ctx, securityGroupName, environment, vpcID)
}

// EnsureSecurityGroupCalls gets all the calls that were made to EnsureSecurityGroup.
// Check the length with:
//     len(mockedClient.EnsureSecurityGroupCalls())
func (mock *ClientMock) EnsureSecurityGroupCalls() []struct {
	Ctx               context.Context
	SecurityGroupName string
	Environment       string
	VpcID             string
} {
	var calls []struct {
		Ctx               context.Context
		SecurityGroupName string
		Environment       string
		VpcID             string
	}
	mock.lockEnsureSecurityGroup.RLock()
	calls = mock.calls.EnsureSecurityGroup
	mock.lockEnsureSecurityGroup.RUnlock()
	return calls
}

// EnsureSecurityGroupIngressRules calls EnsureSecurityGroupIngressRulesFunc.
func (mock *ClientMock) EnsureSecurityGroupIngressRules(ctx context.Context, securityGroupID string) error {
	if mock.EnsureSecurityGroupIngressRulesFunc == nil {
		panic("ClientMock.EnsureSecurityGroupIngressRulesFunc: method is nil but Client.EnsureSecurityGroupIngressRules was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		SecurityGroupID string
	}{
		Ctx:             ctx,
		SecurityGroupID: securityGroupID,
	}
	mock.lockEnsureSecurityGroupIngressRules.Lock()
	mock.calls.EnsureSecurityGroupIngressRules = append(mock.calls.EnsureSecurityGroupIngressRules, callInfo)
	mock.lockEnsureSecurityGroupIngressRules.Unlock()
	return mock.EnsureSecurityGroupIngressRulesFunc(ctx, securityGroupID)
}

// EnsureSecurityGroupIngressRulesCalls gets all the calls that were made to EnsureSecurityGroupIngressRules.
// Check the length with:
//     len(mockedClient.EnsureSecurityGroupIngressRulesCalls())
func (mock *ClientMock) EnsureSecurityGroupIngressRulesCalls() []struct {
	Ctx             context.Context
	SecurityGroupID string
} {
	var calls []struct {
		Ctx             context.Context
		SecurityGroupID string
	}
	mock.lockEnsureSecurityGroupIngressRules.RLock()
	calls = mock.calls.EnsureSecurityGroupIngressRules
	mock.lockEnsureSecurityGroupIngressRules.RUnlock()
	return calls
}

// EnsureSubnet calls EnsureSubnetFunc.
func (mock *ClientMock) EnsureSubnet(ctx context.Context, subNetID string) (string, error) {
	if mock.EnsureSubnetFunc == nil {
		panic("ClientMock.EnsureSubnetFunc: method is nil but Client.EnsureSubnet was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		SubNetID string
	}{
		Ctx:      ctx,
		SubNetID: subNetID,
	}
	mock.lockEnsureSubnet.Lock()
	mock.calls.EnsureSubnet = append(mock.calls.EnsureSubnet, callInfo)
	mock.lockEnsureSubnet.Unlock()
	return mock.EnsureSubnetFunc(ctx, subNetID)
}

// EnsureSubnetCalls gets all the calls that were made to EnsureSubnet.
// Check the length with:
//     len(mockedClient.EnsureSubnetCalls())
func (mock *ClientMock) EnsureSubnetCalls() []struct {
	Ctx      context.Context
	SubNetID string
} {
	var calls []struct {
		Ctx      context.Context
		SubNetID string
	}
	mock.lockEnsureSubnet.RLock()
	calls = mock.calls.EnsureSubnet
	mock.lockEnsureSubnet.RUnlock()
	return calls
}

// FindMatchingAMI calls FindMatchingAMIFunc.
func (mock *ClientMock) FindMatchingAMI(ctx context.Context, region string, version string) (string, error) {
	if mock.FindMatchingAMIFunc == nil {
		panic("ClientMock.FindMatchingAMIFunc: method is nil but Client.FindMatchingAMI was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Region  string
		Version string
	}{
		Ctx:     ctx,
		Region:  region,
		Version: version,
	}
	mock.lockFindMatchingAMI.Lock()
	mock.calls.FindMatchingAMI = append(mock.calls.FindMatchingAMI, callInfo)
	mock.lockFindMatchingAMI.Unlock()
	return mock.FindMatchingAMIFunc(ctx, region, version)
}

// FindMatchingAMICalls gets all the calls that were made to FindMatchingAMI.
// Check the length with:
//     len(mockedClient.FindMatchingAMICalls())
func (mock *ClientMock) FindMatchingAMICalls() []struct {
	Ctx     context.Context
	Region  string
	Version string
} {
	var calls []struct {
		Ctx     context.Context
		Region  string
		Version string
	}
	mock.lockFindMatchingAMI.RLock()
	calls = mock.calls.FindMatchingAMI
	mock.lockFindMatchingAMI.RUnlock()
	return calls
}

// GetResourcesByTags calls GetResourcesByTagsFunc.
func (mock *ClientMock) GetResourcesByTags(ctx context.Context, tags TagSpec) ([]Resource, error) {
	if mock.GetResourcesByTagsFunc == nil {
		panic("ClientMock.GetResourcesByTagsFunc: method is nil but Client.GetResourcesByTags was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Tags TagSpec
	}{
		Ctx:  ctx,
		Tags: tags,
	}
	mock.lockGetResourcesByTags.Lock()
	mock.calls.GetResourcesByTags = append(mock.calls.GetResourcesByTags, callInfo)
	mock.lockGetResourcesByTags.Unlock()
	return mock.GetResourcesByTagsFunc(ctx, tags)
}

// GetResourcesByTagsCalls gets all the calls that were made to GetResourcesByTags.
// Check the length with:
//     len(mockedClient.GetResourcesByTagsCalls())
func (mock *ClientMock) GetResourcesByTagsCalls() []struct {
	Ctx  context.Context
	Tags TagSpec
} {
	var calls []struct {
		Ctx  context.Context
		Tags TagSpec
	}
	mock.lockGetResourcesByTags.RLock()
	calls = mock.calls.GetResourcesByTags
	mock.lockGetResourcesByTags.RUnlock()
	return calls
}

// InstanceState calls InstanceStateFunc.
func (mock *ClientMock) InstanceState(ctx context.Context, instanceID string) (string, error) {
	if mock.InstanceStateFunc == nil {
		panic("ClientMock.InstanceStateFunc: method is nil but Client.InstanceState was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		InstanceID string
	}{
		Ctx:        ctx,
		InstanceID: instanceID,
	}
	mock.lockInstanceState.Lock()
	mock.calls.InstanceState = append(mock.calls.InstanceState, callInfo)
	mock.lockInstanceState.Unlock()
	return mock.InstanceStateFunc(ctx, instanceID)
}

// InstanceStateCalls gets all the calls that were made to InstanceState.
// Check the length with:
//     len(mockedClient.InstanceStateCalls())
func (mock *ClientMock) InstanceStateCalls() []struct {
	Ctx        context.Context
	InstanceID string
} {
	var calls []struct {
		Ctx        context.Context
		InstanceID string
	}
	mock.lockInstanceState.RLock()
	calls = mock.calls.InstanceState
	mock.lockInstanceState.RUnlock()
	return calls
}
