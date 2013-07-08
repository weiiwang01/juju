// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package azure

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"sync"

	. "launchpad.net/gocheck"
	"launchpad.net/gwacl"
	"launchpad.net/juju-core/environs"
	"launchpad.net/juju-core/environs/config"
	"launchpad.net/juju-core/environs/localstorage"
	"launchpad.net/juju-core/errors"
	"launchpad.net/juju-core/instance"
	"launchpad.net/juju-core/testing"
)

type EnvironSuite struct {
	ProviderSuite
}

var _ = Suite(new(EnvironSuite))

func makeEnviron(c *C) *azureEnviron {
	attrs := makeAzureConfigMap(c)
	cfg, err := config.New(attrs)
	c.Assert(err, IsNil)
	env, err := NewEnviron(cfg)
	c.Assert(err, IsNil)
	return env
}

// setDummyStorage injects the local provider's fake storage implementation
// into the given environment, so that tests can manipulate storage as if it
// were real.
// Returns a cleanup function that must be called when done with the storage.
func setDummyStorage(c *C, env *azureEnviron) func() {
	listener, err := localstorage.Serve("127.0.0.1:0", c.MkDir())
	c.Assert(err, IsNil)
	env.storage = localstorage.Client(listener.Addr().String())
	return func() { listener.Close() }
}

func (EnvironSuite) TestGetSnapshot(c *C) {
	original := azureEnviron{name: "this-env", ecfg: new(azureEnvironConfig)}
	snapshot := original.getSnapshot()

	// The snapshot is identical to the original.
	c.Check(*snapshot, DeepEquals, original)

	// However, they are distinct objects.
	c.Check(snapshot, Not(Equals), &original)

	// It's a shallow copy; they still share pointers.
	c.Check(snapshot.ecfg, Equals, original.ecfg)

	// Neither object is locked at the end of the copy.
	c.Check(original.Mutex, Equals, sync.Mutex{})
	c.Check(snapshot.Mutex, Equals, sync.Mutex{})
}

func (EnvironSuite) TestGetSnapshotLocksEnviron(c *C) {
	original := azureEnviron{}
	testing.TestLockingFunction(&original.Mutex, func() { original.getSnapshot() })
}

func (EnvironSuite) TestName(c *C) {
	env := azureEnviron{name: "foo"}
	c.Check(env.Name(), Equals, env.name)
}

func (EnvironSuite) TestConfigReturnsConfig(c *C) {
	cfg := new(config.Config)
	ecfg := azureEnvironConfig{Config: cfg}
	env := azureEnviron{ecfg: &ecfg}
	c.Check(env.Config(), Equals, cfg)
}

func (EnvironSuite) TestConfigLocksEnviron(c *C) {
	env := azureEnviron{name: "env", ecfg: new(azureEnvironConfig)}
	testing.TestLockingFunction(&env.Mutex, func() { env.Config() })
}

func (EnvironSuite) TestGetManagementAPI(c *C) {
	env := makeEnviron(c)
	context, err := env.getManagementAPI()
	c.Assert(err, IsNil)
	defer env.releaseManagementAPI(context)
	c.Check(context, NotNil)
	c.Check(context.ManagementAPI, NotNil)
	c.Check(context.certFile, NotNil)
}

func (EnvironSuite) TestReleaseManagementAPIAcceptsNil(c *C) {
	env := makeEnviron(c)
	env.releaseManagementAPI(nil)
	// The real test is that this does not panic.
}

func (EnvironSuite) TestReleaseManagementAPIAcceptsIncompleteContext(c *C) {
	env := makeEnviron(c)
	context := azureManagementContext{
		ManagementAPI: nil,
		certFile:      nil,
	}
	env.releaseManagementAPI(&context)
	// The real test is that this does not panic.
}

func patchWithPropertiesResponse(c *C, deployments []gwacl.Deployment) *[]*gwacl.X509Request {
	propertiesS1 := gwacl.HostedService{
		ServiceName: "S1", Deployments: deployments}
	propertiesS1XML, err := propertiesS1.Serialize()
	c.Assert(err, IsNil)
	responses := []gwacl.DispatcherResponse{gwacl.NewDispatcherResponse(
		[]byte(propertiesS1XML),
		http.StatusOK,
		nil,
	)}
	requests := gwacl.PatchManagementAPIResponses(responses)
	return requests
}

func (suite EnvironSuite) TestAllInstances(c *C) {
	deployments := []gwacl.Deployment{{Name: "deployment-1"}, {Name: "deployment-2"}}
	requests := patchWithPropertiesResponse(c, deployments)
	env := makeEnviron(c)
	instances, err := env.AllInstances()
	c.Assert(err, IsNil)
	c.Check(len(instances), Equals, len(deployments))
	c.Check(len(*requests), Equals, 1)
}

func (suite EnvironSuite) TestInstancesReturnsFilteredList(c *C) {
	deployments := []gwacl.Deployment{{Name: "deployment-1"}, {Name: "deployment-2"}}
	requests := patchWithPropertiesResponse(c, deployments)
	env := makeEnviron(c)
	instances, err := env.Instances([]instance.Id{"deployment-1"})
	c.Assert(err, IsNil)
	c.Check(len(instances), Equals, 1)
	c.Check(instances[0].Id(), Equals, instance.Id("deployment-1"))
	c.Check(len(*requests), Equals, 1)
}

func (suite EnvironSuite) TestInstancesReturnsErrNoInstancesIfNoInstancesRequested(c *C) {
	deployments := []gwacl.Deployment{{Name: "deployment-1"}, {Name: "deployment-2"}}
	patchWithPropertiesResponse(c, deployments)
	env := makeEnviron(c)
	instances, err := env.Instances([]instance.Id{})
	c.Check(err, Equals, environs.ErrNoInstances)
	c.Check(instances, IsNil)
}

func (suite EnvironSuite) TestInstancesReturnsErrNoInstancesIfNoInstanceFound(c *C) {
	deployments := []gwacl.Deployment{}
	patchWithPropertiesResponse(c, deployments)
	env := makeEnviron(c)
	instances, err := env.Instances([]instance.Id{"deploy-id"})
	c.Check(err, Equals, environs.ErrNoInstances)
	c.Check(instances, IsNil)
}

func (suite EnvironSuite) TestInstancesReturnsPartialInstancesIfSomeInstancesAreNotFound(c *C) {
	deployments := []gwacl.Deployment{{Name: "deployment-1"}, {Name: "deployment-2"}}
	requests := patchWithPropertiesResponse(c, deployments)
	env := makeEnviron(c)
	instances, err := env.Instances([]instance.Id{"deployment-1", "unknown-deployment"})
	c.Assert(err, Equals, environs.ErrPartialInstances)
	c.Check(len(instances), Equals, 1)
	c.Check(instances[0].Id(), Equals, instance.Id("deployment-1"))
	c.Check(len(*requests), Equals, 1)
}

func (EnvironSuite) TestStorage(c *C) {
	env := makeEnviron(c)
	baseStorage := env.Storage()
	storage, ok := baseStorage.(*azureStorage)
	c.Check(ok, Equals, true)
	c.Assert(storage, NotNil)
	c.Check(storage.storageContext.getContainer(), Equals, env.ecfg.StorageContainerName())
	context, err := storage.getStorageContext()
	c.Assert(err, IsNil)
	c.Check(context.Account, Equals, env.ecfg.StorageAccountName())
	c.Check(context.Key, Equals, env.ecfg.StorageAccountKey())
}

func (EnvironSuite) TestPublicStorage(c *C) {
	env := makeEnviron(c)
	baseStorage := env.PublicStorage()
	storage, ok := baseStorage.(*azureStorage)
	c.Assert(storage, NotNil)
	c.Check(ok, Equals, true)
	c.Check(storage.storageContext.getContainer(), Equals, env.ecfg.PublicStorageContainerName())
	context, err := storage.getStorageContext()
	c.Assert(err, IsNil)
	c.Check(context.Account, Equals, env.ecfg.PublicStorageAccountName())
	c.Check(context.Key, Equals, "")
}

func (EnvironSuite) TestPublicStorageReturnsEmptyStorageIfNoInfo(c *C) {
	attrs := makeAzureConfigMap(c)
	attrs["public-storage-container-name"] = ""
	attrs["public-storage-account-name"] = ""
	cfg, err := config.New(attrs)
	c.Assert(err, IsNil)
	env, err := NewEnviron(cfg)
	c.Assert(err, IsNil)
	c.Check(env.PublicStorage(), Equals, environs.EmptyStorage)
}

func (EnvironSuite) TestGetStorageContext(c *C) {
	env := makeEnviron(c)
	storage, err := env.getStorageContext()
	c.Assert(err, IsNil)
	c.Assert(storage, NotNil)
	c.Check(storage.Account, Equals, env.ecfg.StorageAccountName())
	c.Check(storage.Key, Equals, env.ecfg.StorageAccountKey())
}

func (EnvironSuite) TestGetPublicStorageContext(c *C) {
	env := makeEnviron(c)
	storage, err := env.getPublicStorageContext()
	c.Assert(err, IsNil)
	c.Assert(storage, NotNil)
	c.Check(storage.Account, Equals, env.ecfg.PublicStorageAccountName())
	c.Check(storage.Key, Equals, "")
}

func (EnvironSuite) TestSetConfigValidates(c *C) {
	env := makeEnviron(c)
	originalCfg := env.ecfg
	attrs := makeAzureConfigMap(c)
	// This config is not valid.  It lacks essential information.
	delete(attrs, "management-subscription-id")
	badCfg, err := config.New(attrs)
	c.Assert(err, IsNil)

	err = env.SetConfig(badCfg)

	// Since the config was not valid, SetConfig returns an error.  It
	// does not update the environment's config either.
	c.Check(err, NotNil)
	c.Check(
		err,
		ErrorMatches,
		"management-subscription-id: expected string, got nothing")
	c.Check(env.ecfg, Equals, originalCfg)
}

func (EnvironSuite) TestSetConfigUpdatesConfig(c *C) {
	env := makeEnviron(c)
	// We're going to set a new config.  It can be recognized by its
	// unusual default Ubuntu release series: 7.04 Feisty Fawn.
	attrs := makeAzureConfigMap(c)
	attrs["default-series"] = "feisty"
	cfg, err := config.New(attrs)
	c.Assert(err, IsNil)

	err = env.SetConfig(cfg)
	c.Assert(err, IsNil)

	c.Check(env.ecfg.Config.DefaultSeries(), Equals, "feisty")
}

func (EnvironSuite) TestSetConfigLocksEnviron(c *C) {
	env := makeEnviron(c)
	cfg, err := config.New(makeAzureConfigMap(c))
	c.Assert(err, IsNil)

	testing.TestLockingFunction(&env.Mutex, func() { env.SetConfig(cfg) })
}

func (EnvironSuite) TestSetConfigWillNotUpdateName(c *C) {
	// Once the environment's name has been set, it cannot be updated.
	// Global validation rejects such a change.
	// This matters because the attribute is not protected by a lock.
	env := makeEnviron(c)
	originalName := env.Name()
	attrs := makeAzureConfigMap(c)
	attrs["name"] = "new-name"
	cfg, err := config.New(attrs)
	c.Assert(err, IsNil)

	err = env.SetConfig(cfg)

	c.Assert(err, NotNil)
	c.Check(
		err,
		ErrorMatches,
		`cannot change name from ".*" to "new-name"`)
	c.Check(env.Name(), Equals, originalName)
}

func (EnvironSuite) TestStateInfoFailsIfNoStateInstances(c *C) {
	env := makeEnviron(c)
	cleanup := setDummyStorage(c, env)
	defer cleanup()
	_, _, err := env.StateInfo()
	c.Check(errors.IsNotFoundError(err), Equals, true)
}

func (EnvironSuite) TestStateInfo(c *C) {
	instanceID := "my-instance"
	// In the Azure provider, DNS name and instance ID are the same thing.
	patchWithPropertiesResponse(c, []gwacl.Deployment{{
		Name: instanceID,
		URL:  fmt.Sprintf("http://%s/", instanceID),
	}})
	env := makeEnviron(c)
	cleanup := setDummyStorage(c, env)
	defer cleanup()
	err := environs.SaveState(
		env.Storage(),
		&environs.BootstrapState{StateInstances: []instance.Id{instance.Id(instanceID)}})
	c.Assert(err, IsNil)

	stateInfo, apiInfo, err := env.StateInfo()
	c.Assert(err, IsNil)

	config := env.Config()
	statePortSuffix := fmt.Sprintf(":%d", config.StatePort())
	apiPortSuffix := fmt.Sprintf(":%d", config.APIPort())
	c.Check(stateInfo.Addrs, DeepEquals, []string{instanceID + statePortSuffix})
	c.Check(apiInfo.Addrs, DeepEquals, []string{instanceID + apiPortSuffix})
}

func (EnvironSuite) TestAttemptCreateServiceCreatesService(c *C) {
	responses := []gwacl.DispatcherResponse{
		gwacl.NewDispatcherResponse(nil, http.StatusOK, nil),
	}
	requests := gwacl.PatchManagementAPIResponses(responses)
	azure, err := gwacl.NewManagementAPI("subscription", "certfile.pem")
	c.Assert(err, IsNil)

	service, err := newHostedService(azure)
	c.Assert(err, IsNil)

	c.Assert(*requests, HasLen, 1)
	body := gwacl.CreateHostedService{}
	err = xml.Unmarshal((*requests)[0].Payload, &body)
	c.Assert(err, IsNil)
	c.Check(body.ServiceName, Equals, service.ServiceName)
}

func (EnvironSuite) TestAttemptCreateServiceReturnsNilIfNameNotUnique(c *C) {
	errorBody, err := xml.Marshal(gwacl.AzureError{
		error:      fmt.Errorf("Name already in use"),
		HTTPStatus: http.StatusConflict,
		Code:       "Conflict",
		Message:    "Name already in use",
	})
	c.Assert(err, IsNil)
	responses := []gwacl.DispatcherResponse{
		gwacl.NewDispatcherResponse(errorBody, http.StatusConflict, nil),
	}
	gwacl.PatchManagementAPIResponses(responses)
	azure, err := gwacl.NewManagementAPI("subscription", "certfile.pem")
	c.Assert(err, IsNil)

	service, err := attemptCreateService(azure)
	c.Check(err, IsNil)
	c.Check(service, IsNil)
}

func (EnvironSuite) TestAttemptCreateServicePropagatesOtherFailure(c *C) {
	responses := []gwacl.DispatcherResponse{
		gwacl.NewDispatcherResponse(nil, http.StatusNotFound, nil),
	}
	gwacl.PatchManagementAPIResponses(responses)
	azure, err := gwacl.NewManagementAPI("subscription", "certfile.pem")
	c.Assert(err, IsNil)

	_, err = newHostedService(azure)
	c.Assert(err, NotNil)
	c.Check(err, ErrorMatches, ".*Not Found.*")
}

func (EnvironSuite) TestExtractDeploymentHostnameExtractsHost(c *C) {
	// Example taken from Azure documentation:
	// http://msdn.microsoft.com/en-us/library/windowsazure/ee460804.aspx
	instanceURL := "http://MyService.cloudapp.net"
	hostname, err := extractDeploymentHostname(instanceURL)
	c.Assert(err, IsNil)
	c.Check(hostname, Equals, "MyService.cloudapp.net")
}

func (EnvironSuite) TestNewHostedServiceCreatesService(c *C) {
	c.Fail() // TEST THIS
}

func (EnvironSuite) TestNewHostedServiceRetriesIfNotUnique(c *C) {
	c.Fail() // TEST THIS
}

func (EnvironSuite) TestNewHostedServiceFailsIfUnableToFindUniqueName(c *C) {
	c.Fail() // TEST THIS
}

func (EnvironSuite) TestExtractDeploymentHostnamePropagatesError(c *C) {
	_, err := extractDeploymentHostname(":x:THIS BREAKS:x:")
	c.Check(err, NotNil)
}

func (EnvironSuite) TestSetServiceDNSNameReadsDeploymentAndUpdatesService(c *C) {
	serviceName := "fub"
	deploymentName := "default"
	instanceHostname := "foobar.cloudapp.net"
	deploymentBody, err := xml.Marshal(gwacl.Deployment{
		Name: deploymentName,
		URL:  fmt.Sprintf("http://%s", instanceHostname),
	})
	c.Assert(err, IsNil)
	// setServiceDNSName reads the Deployment to obtain the instance URL,
	// then updates the Hosted Service by setting its label to the DNS
	// name of the instance.
	responses := []gwacl.DispatcherResponse{
		gwacl.NewDispatcherResponse(deploymentBody, http.StatusOK, nil),
		gwacl.NewDispatcherResponse(nil, http.StatusOK, nil),
	}
	requests := gwacl.PatchManagementAPIResponses(responses)

	azure, err := gwacl.NewManagementAPI("subscription", "certfile.pem")
	c.Assert(err, IsNil)
	err = setServiceDNSName(azure, serviceName, deploymentName)
	c.Assert(err, IsNil)

	c.Assert(*requests, HasLen, 2)
	getDeploymentReq := (*requests)[0]
	updateServiceReq := (*requests)[1]

	c.Check(getDeploymentReq.URL, Equals, fmt.Sprintf(
		"https://management.core.windows.net/%s/services/hostedservices/%s/deployments/%s",
		"subscription", serviceName, deploymentName))
	updateServiceBody := &gwacl.UpdateHostedService{}
	err = xml.Unmarshal(updateServiceReq.Payload, updateServiceBody)
	c.Assert(err, IsNil)
	newLabel, err := base64.StdEncoding.DecodeString(updateServiceBody.Label)
	c.Check(string(newLabel), Equals, instanceHostname)
}

func (EnvironSuite) TestMakeProvisionalDeploymentLabelIsConsistent(c *C) {
	c.Check(makeProvisionalDeploymentLabel("foo"), Equals, makeProvisionalDeploymentLabel("foo"))
}

func (EnvironSuite) TestMakeProvisionalDeploymentLabelIncludesName(c *C) {
	c.Check(makeProvisionalDeploymentLabel("splyz"), Matches, ".*splyz.*")
}

func (EnvironSuite) TestIsProvisionalDeploymentLabelRecognizesProvisionalLabel(c *C) {
	c.Check(isProvisionalDeploymentLabel(makeProvisionalDeploymentLabel("x")), Equals, true)
}

func (EnvironSuite) TestIsProvisionalDeploymentLabelRecognizesPermanentLabel(c *C) {
	c.Check(isProvisionalDeploymentLabel("label"), Equals, false)
}
