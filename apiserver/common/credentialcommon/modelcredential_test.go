// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package credentialcommon_test

import (
	stdcontext "context"

	"github.com/juju/errors"
	"github.com/juju/names/v4"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/common/credentialcommon"
	"github.com/juju/juju/apiserver/common/credentialcommon/mocks"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/caas"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/core/instance"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/environs/context"
	"github.com/juju/juju/environs/instances"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
	jujuesting "github.com/juju/juju/testing"
)

var _ = gc.Suite(&CheckMachinesSuite{})
var _ = gc.Suite(&ModelCredentialSuite{})

type CheckMachinesSuite struct {
	testing.IsolationSuite

	provider    *mockProvider
	instance    *mockInstance
	callContext context.ProviderCallContext

	backend *mockPersistedBackend
	machine *mockMachine
}

func (s *CheckMachinesSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
	s.backend = createModelBackend()

	// This is what the test gets from the state.
	s.machine = createTestMachine("1", "wind-up")
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine}, nil
	}

	// This is what the test gets from the cloud.
	s.instance = &mockInstance{id: "wind-up"}
	s.provider = &mockProvider{
		Stub: &testing.Stub{},
		allInstancesFunc: func(ctx context.ProviderCallContext) ([]instances.Instance, error) {
			return []instances.Instance{s.instance}, nil
		},
	}
	s.callContext = context.NewEmptyCloudCallContext()
}

func (s *CheckMachinesSuite) TestCheckMachinesSuccess(c *gc.C) {
	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *CheckMachinesSuite) TestCheckMachinesInstancesMissing(c *gc.C) {
	machine1 := createTestMachine("2", "birds")
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine1}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(results.Results, gc.HasLen, 1)
	c.Assert(results.Results[0].Error, gc.ErrorMatches, `couldn't find instance "birds" for machine 2`)
}

func (s *CheckMachinesSuite) TestCheckMachinesExtraInstances(c *gc.C) {
	instance2 := &mockInstance{id: "analyse"}
	s.provider.allInstancesFunc = func(ctx context.ProviderCallContext) ([]instances.Instance, error) {
		return []instances.Instance{s.instance, instance2}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results.Results, gc.IsNil)
}

func (s *CheckMachinesSuite) TestCheckMachinesExtraInstancesWhenMigrating(c *gc.C) {
	instance2 := &mockInstance{id: "analyse"}
	s.provider.allInstancesFunc = func(ctx context.ProviderCallContext) ([]instances.Instance, error) {
		return []instances.Instance{s.instance, instance2}, nil
	}
	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, true)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(results.Results, gc.HasLen, 1)
	c.Assert(results.Results[0].Error, gc.ErrorMatches, `no machine with instance "analyse"`)
}

func (s *CheckMachinesSuite) TestCheckMachinesErrorGettingMachines(c *gc.C) {
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return nil, errors.New("boom")
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, gc.ErrorMatches, "boom")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *CheckMachinesSuite) TestCheckMachinesErrorGettingInstances(c *gc.C) {
	s.provider.allInstancesFunc = func(ctx context.ProviderCallContext) ([]instances.Instance, error) {
		return nil, errors.New("kaboom")
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, gc.ErrorMatches, "kaboom")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *CheckMachinesSuite) TestCheckMachinesHandlesContainers(c *gc.C) {
	machine1 := createTestMachine("1", "")
	machine1.container = true
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine1}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *CheckMachinesSuite) TestCheckMachinesHandlesManual(c *gc.C) {
	machine1 := createTestMachine("2", "")
	machine1.manualFunc = func() (bool, error) { return false, errors.New("manual retrieval failure") }
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine1}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, gc.ErrorMatches, "manual retrieval failure")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})

	machine1.manualFunc = func() (bool, error) { return true, nil }
	results, err = credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *CheckMachinesSuite) TestCheckMachinesErrorGettingMachineInstanceId(c *gc.C) {
	machine1 := createTestMachine("2", "")
	machine1.instanceIdFunc = func() (instance.Id, error) { return "", errors.New("retrieval failure") }
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine1}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{
			{Error: apiservererrors.ServerError(errors.Errorf("getting instance id for machine 2: retrieval failure"))},
		},
	})
}

func (s *CheckMachinesSuite) TestCheckMachinesErrorGettingMachineInstanceIdNonFatal(c *gc.C) {
	machine1 := createTestMachine("2", "")
	machine1.instanceIdFunc = func() (instance.Id, error) { return "", errors.New("retrieval failure") }
	s.machine.instanceIdFunc = machine1.instanceIdFunc
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine1}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{
			{Error: apiservererrors.ServerError(errors.New("getting instance id for machine 1: retrieval failure"))},
			{Error: apiservererrors.ServerError(errors.New("getting instance id for machine 2: retrieval failure"))},
		},
	})
}

func (s *CheckMachinesSuite) TestCheckMachinesErrorGettingMachineInstanceIdNonFatalWhenMigrating(c *gc.C) {
	machine1 := createTestMachine("2", "")
	machine1.instanceIdFunc = func() (instance.Id, error) { return "", errors.New("retrieval failure") }
	s.machine.instanceIdFunc = machine1.instanceIdFunc
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine1}, nil
	}

	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, true)
	c.Assert(err, jc.ErrorIsNil)
	// There should be 3 errors here:
	// * 2 of them because failing to get an instance id from one machine should not stop the processing the rest of the machines;
	// * 1 because we no longer can link test instance (s.instance) to a test machine (s.machine).
	c.Assert(results, gc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{
			{Error: apiservererrors.ServerError(errors.New("getting instance id for machine 1: retrieval failure"))},
			{Error: apiservererrors.ServerError(errors.New("getting instance id for machine 2: retrieval failure"))},
			{Error: apiservererrors.ServerError(errors.New(`no machine with instance "wind-up"`))},
		},
	})
}

func (s *CheckMachinesSuite) TestCheckMachinesNotProvisionedError(c *gc.C) {
	machine2 := createTestMachine("2", "")
	machine2.instanceIdFunc = func() (instance.Id, error) { return "", errors.NotProvisionedf("machine 2") }
	s.backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{s.machine, machine2}, nil
	}

	// We should ignore the unprovisioned machine - we wouldn't expect
	// the cloud to know about it.
	results, err := credentialcommon.CheckMachineInstances(s.callContext, s.backend, s.provider, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

type ModelCredentialSuite struct {
	testing.IsolationSuite

	backend                 *mockPersistedBackend
	credentialService       *mocks.MockCredentialService
	controllerConfigService *mocks.MockControllerConfigService
	callContext             context.ProviderCallContext
}

func (s *ModelCredentialSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
	s.backend = createModelBackend()
	s.callContext = context.NewEmptyCloudCallContext()
}

func (s *ModelCredentialSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)
	s.credentialService = mocks.NewMockCredentialService(ctrl)
	s.controllerConfigService = mocks.NewMockControllerConfigService(ctrl)
	return ctrl
}

func (s *ModelCredentialSuite) TestValidateNewModelCredentialUnknownModelType(c *gc.C) {
	unknownModel := createTestModel()
	unknownModel.modelType = "unknown"
	s.backend.modelFunc = func() (credentialcommon.Model, error) {
		return unknownModel, nil
	}

	results, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.NewCloudCredentialTag("dummy/bob/default"),
		&testCredential,
		testCloud, false,
	)
	c.Assert(err, gc.ErrorMatches, `model type "unknown" not supported`)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestBuildingOpenParamsErrorGettingModel(c *gc.C) {
	s.backend.SetErrors(errors.New("get model error"))
	results, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.CloudCredentialTag{},
		nil, testCloud, false)
	c.Assert(err, gc.ErrorMatches, "get model error")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestBuildingOpenParamsErrorGettingModelConfig(c *gc.C) {
	model := createTestModel()
	model.configFunc = func() (*config.Config, error) {
		return nil, errors.New("get model config error")
	}

	s.backend.modelFunc = func() (credentialcommon.Model, error) {
		return model, nil
	}

	results, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.NewCloudCredentialTag("dummy/bob/default"),
		&testCredential, testCloud, false)
	c.Assert(err, gc.ErrorMatches, "get model config error")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestBuildingOpenParamsErrorValidateCredentialForModelCloud(c *gc.C) {
	c.Skip("TODO(wallyworld) - need to move to dqlite for this test")
	model := createTestModel()
	model.validateCredentialFunc = func(tag names.CloudCredentialTag, credential cloud.Credential) error {
		return errors.New("credential not for model cloud error")
	}

	s.backend.modelFunc = func() (credentialcommon.Model, error) {
		return model, nil
	}

	results, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.CloudCredentialTag{},
		&testCredential, testCloud, false)
	c.Assert(err, gc.ErrorMatches, "credential not for model cloud error")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestValidateExistingModelCredentialErrorGettingModel(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.backend.SetErrors(errors.New("get model error"))
	results, err := credentialcommon.ValidateExistingModelCredential(s.callContext, s.backend, s.controllerConfigService, s.credentialService, testCloud, false)
	c.Assert(err, gc.ErrorMatches, "get model error")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestValidateExistingModelCredentialUnsetCloudCredential(c *gc.C) {
	defer s.setupMocks(c).Finish()

	model := createTestModel()
	model.cloudCredentialTagFunc = func() (names.CloudCredentialTag, bool) {
		return names.CloudCredentialTag{}, false
	}

	s.backend.modelFunc = func() (credentialcommon.Model, error) {
		return model, nil
	}

	results, err := credentialcommon.ValidateExistingModelCredential(s.callContext, s.backend, s.controllerConfigService, s.credentialService, testCloud, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestValidateExistingModelCredentialErrorGettingCredential(c *gc.C) {
	defer s.setupMocks(c).Finish()

	tag := names.NewCloudCredentialTag("cirrus/fred/default")
	s.credentialService.EXPECT().CloudCredential(gomock.Any(), tag).Return(cloud.Credential{}, errors.New("no nope niet"))

	results, err := credentialcommon.ValidateExistingModelCredential(s.callContext, s.backend, s.controllerConfigService, s.credentialService, testCloud, false)
	c.Assert(err, gc.ErrorMatches, "no nope niet")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestValidateExistingModelCredentialInvalidCredential(c *gc.C) {
	defer s.setupMocks(c).Finish()

	tag := names.NewCloudCredentialTag("cirrus/fred/default")
	cred := cloud.NewEmptyCredential()
	cred.Label = "cred"
	cred.Invalid = true
	s.credentialService.EXPECT().CloudCredential(gomock.Any(), tag).Return(cred, nil)

	results, err := credentialcommon.ValidateExistingModelCredential(s.callContext, s.backend, s.controllerConfigService, s.credentialService, testCloud, false)
	c.Assert(err, gc.ErrorMatches, `credential "cred" not valid`)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
	s.backend.CheckCallNames(c, "Model")
}

func (s *ModelCredentialSuite) TestOpeningProviderFails(c *gc.C) {
	s.PatchValue(credentialcommon.NewEnv, func(stdcontext.Context, environs.OpenParams) (environs.Environ, error) {
		return nil, errors.New("explosive")
	})
	results, err := credentialcommon.CheckIAASModelCredential(s.callContext, environs.OpenParams{}, s.backend, false)
	c.Assert(err, gc.ErrorMatches, "explosive")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestValidateNewModelCredentialForIAASModel(c *gc.C) {
	s.ensureEnvForIAASModel(c)
	results, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.NewCloudCredentialTag("dummy/bob/default"),
		&testCredential, testCloud, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestValidateModelCredentialCloudMismatch(c *gc.C) {
	s.ensureEnvForIAASModel(c)
	_, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.NewCloudCredentialTag("other/bob/default"),
		&testCredential, testCloud, false)
	c.Assert(err, gc.ErrorMatches, `credential "other/bob/default" not valid`)
}

func (s *ModelCredentialSuite) TestValidateExistingModelCredentialForIAASModel(c *gc.C) {
	s.ensureEnvForIAASModel(c)
	results, err := credentialcommon.ValidateExistingModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		s.credentialService, testCloud, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestOpeningCAASBrokerFails(c *gc.C) {
	s.PatchValue(credentialcommon.NewCAASBroker, func(stdcontext.Context, environs.OpenParams) (caas.Broker, error) {
		return nil, errors.New("explosive")
	})
	results, err := credentialcommon.CheckCAASModelCredential(stdcontext.Background(), environs.OpenParams{})
	c.Assert(err, gc.ErrorMatches, "explosive")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestCAASCredentialCheckFailed(c *gc.C) {
	s.PatchValue(credentialcommon.NewCAASBroker, func(stdcontext.Context, environs.OpenParams) (caas.Broker, error) {
		return &mockCaasBroker{
			namespacesFunc: func() ([]string, error) { return nil, errors.New("fail auth") },
		}, nil
	})
	results, err := credentialcommon.CheckCAASModelCredential(stdcontext.Background(), environs.OpenParams{})
	c.Assert(err, gc.ErrorMatches, "fail auth")
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestCAASCredentialCheckSucceeds(c *gc.C) {
	s.PatchValue(credentialcommon.NewCAASBroker, func(stdcontext.Context, environs.OpenParams) (caas.Broker, error) {
		return &mockCaasBroker{
			namespacesFunc: func() ([]string, error) { return []string{}, nil },
		}, nil
	})
	results, err := credentialcommon.CheckCAASModelCredential(stdcontext.Background(), environs.OpenParams{})
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestValidateNewModelCredentialForCAASModel(c *gc.C) {
	s.ensureEnvForCAASModel(c)
	results, err := credentialcommon.ValidateNewModelCredential(
		s.callContext,
		s.backend,
		s.controllerConfigService,
		names.NewCloudCredentialTag("dummy/bob/default"),
		&testCredential, testCloud, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) TestValidateExistingModelCredentialForCAASSuccess(c *gc.C) {
	s.ensureEnvForCAASModel(c)
	results, err := credentialcommon.ValidateExistingModelCredential(s.callContext, s.backend, s.controllerConfigService, s.credentialService, testCloud, false)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, gc.DeepEquals, params.ErrorResults{})
}

func (s *ModelCredentialSuite) ensureEnvForCAASModel(c *gc.C) {
	caasModel := createTestModel()
	caasModel.modelType = state.ModelTypeCAAS
	s.backend.modelFunc = func() (credentialcommon.Model, error) {
		return caasModel, nil
	}
	s.PatchValue(credentialcommon.NewCAASBroker, func(stdcontext.Context, environs.OpenParams) (caas.Broker, error) {
		return &mockCaasBroker{
			namespacesFunc: func() ([]string, error) { return []string{}, nil },
		}, nil
	})
}

func (s *ModelCredentialSuite) ensureEnvForIAASModel(c *gc.C) {
	s.PatchValue(credentialcommon.NewEnv, func(stdcontext.Context, environs.OpenParams) (environs.Environ, error) {
		return &mockEnviron{
			mockProvider: &mockProvider{
				Stub: &testing.Stub{},
				allInstancesFunc: func(ctx context.ProviderCallContext) ([]instances.Instance, error) {
					return []instances.Instance{}, nil
				},
			},
		}, nil
	})
}

func createModelBackend() *mockPersistedBackend {
	backend := mockPersistedBackend{Stub: &testing.Stub{}}
	backend.allMachinesFunc = func() ([]credentialcommon.Machine, error) {
		return []credentialcommon.Machine{}, backend.NextErr()
	}
	backend.modelFunc = func() (credentialcommon.Model, error) {
		return createTestModel(), backend.NextErr()
	}
	backend.controllerConfigFunc = func() (credentialcommon.ControllerConfig, error) {
		return testControllerConfig, backend.NextErr()
	}

	backend.cloudFunc = func(name string) (cloud.Cloud, error) {
		return cloud.Cloud{
			Name:      "nuage",
			Type:      "dummy",
			AuthTypes: []cloud.AuthType{cloud.EmptyAuthType, cloud.UserPassAuthType},
			Regions:   []cloud.Region{{Name: "nine", Endpoint: "endpoint"}},
		}, backend.NextErr()
	}
	return &backend
}

type mockProvider struct {
	*testing.Stub
	allInstancesFunc func(ctx context.ProviderCallContext) ([]instances.Instance, error)
}

func (m *mockProvider) AllInstances(ctx context.ProviderCallContext) ([]instances.Instance, error) {
	m.MethodCall(m, "AllInstances", ctx)
	return m.allInstancesFunc(ctx)
}

type mockInstance struct {
	instances.Instance
	id string
}

func (i *mockInstance) Id() instance.Id {
	return instance.Id(i.id)
}

type mockMachine struct {
	id             string
	container      bool
	manualFunc     func() (bool, error)
	instanceIdFunc func() (instance.Id, error)
}

func (m *mockMachine) IsManual() (bool, error) {
	return m.manualFunc()
}

func (m *mockMachine) IsContainer() bool {
	return m.container
}

func (m *mockMachine) InstanceId() (instance.Id, error) {
	return m.instanceIdFunc()
}

func (m *mockMachine) InstanceNames() (instance.Id, string, error) {
	instId, err := m.instanceIdFunc()
	return instId, "", err
}

func (m *mockMachine) Id() string {
	return m.id
}

func createTestMachine(id, instanceId string) *mockMachine {
	return &mockMachine{
		id:             id,
		manualFunc:     func() (bool, error) { return false, nil },
		instanceIdFunc: func() (instance.Id, error) { return instance.Id(instanceId), nil },
	}
}

type mockPersistedBackend struct {
	*testing.Stub
	allMachinesFunc func() ([]credentialcommon.Machine, error)

	modelFunc            func() (credentialcommon.Model, error)
	controllerConfigFunc func() (credentialcommon.ControllerConfig, error)
	cloudFunc            func(name string) (cloud.Cloud, error)
}

func (m *mockPersistedBackend) AllMachines() ([]credentialcommon.Machine, error) {
	m.MethodCall(m, "AllMachines")
	return m.allMachinesFunc()
}

func (m *mockPersistedBackend) Model() (credentialcommon.Model, error) {
	m.MethodCall(m, "Model")
	return m.modelFunc()
}

func (m *mockPersistedBackend) ControllerConfig() (credentialcommon.ControllerConfig, error) {
	m.MethodCall(m, "ControllerConfig")
	return m.controllerConfigFunc()
}

func (m *mockPersistedBackend) Cloud(name string) (cloud.Cloud, error) {
	m.MethodCall(m, "Cloud", name)
	return m.cloudFunc(name)
}

type mockModel struct {
	modelType              state.ModelType
	cloudNameFunc          func() string
	cloudRegionFunc        func() string
	configFunc             func() (*config.Config, error)
	validateCredentialFunc func(tag names.CloudCredentialTag, credential cloud.Credential) error
	cloudCredentialTagFunc func() (names.CloudCredentialTag, bool)
}

func (m *mockModel) CloudName() string {
	return m.cloudNameFunc()
}

func (m *mockModel) CloudRegion() string {
	return m.cloudRegionFunc()
}

func (m *mockModel) Config() (*config.Config, error) {
	return m.configFunc()
}

func (m *mockModel) Type() state.ModelType {
	return m.modelType
}

func (m *mockModel) CloudCredentialTag() (names.CloudCredentialTag, bool) {
	return m.cloudCredentialTagFunc()
}

func (m *mockModel) ValidateCloudCredential(tag names.CloudCredentialTag, credential cloud.Credential) error {
	return m.validateCredentialFunc(tag, credential)
}

func createTestModel() *mockModel {
	return &mockModel{
		modelType:       state.ModelTypeIAAS,
		cloudNameFunc:   func() string { return "nuage" },
		cloudRegionFunc: func() string { return "nine" },
		configFunc: func() (*config.Config, error) {
			return nil, nil
		},
		validateCredentialFunc: func(tag names.CloudCredentialTag, credential cloud.Credential) error {
			return nil
		},
		cloudCredentialTagFunc: func() (names.CloudCredentialTag, bool) {
			// return true here since, most of the time, we want to test when the cloud credential is set.
			return names.NewCloudCredentialTag("cirrus/fred/default"), true
		},
	}
}

var (
	testCloud = cloud.Cloud{
		Name:    "dummy",
		Regions: []cloud.Region{{Name: "nine"}},
	}
	testCredential = cloud.NewCredential(
		cloud.UserPassAuthType,
		map[string]string{
			"username": "user",
			"password": "password",
		},
	)
	testControllerConfig = jujuesting.FakeControllerConfig()
)

type mockEnviron struct {
	environs.Environ
	*mockProvider
}

func (m *mockEnviron) AllInstances(ctx context.ProviderCallContext) ([]instances.Instance, error) {
	return m.mockProvider.AllInstances(ctx)
}

type mockCaasBroker struct {
	caas.Broker

	namespacesFunc func() ([]string, error)
}

func (m *mockCaasBroker) Namespaces() ([]string, error) {
	return m.namespacesFunc()
}

func (m *mockCaasBroker) CheckCloudCredentials() error {
	// The k8s provider implements this via a list namespaces call to the cluster
	_, err := m.namespacesFunc()
	return err
}
