// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package network

import (
	"testing"

	jujutesting "github.com/juju/testing"
	gc "gopkg.in/check.v1"
)

//go:generate go run go.uber.org/mock/mockgen -typed -package network_test -destination discovery_mock_test.go github.com/juju/juju/core/network ConfigSource,ConfigSourceNIC,ConfigSourceAddr

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

// BaseSuite exposes base testing functionality to the network tests,
// including patching package-private functions/variables.
type BaseSuite struct {
	jujutesting.IsolationSuite
}

// PatchGOOS allows us to simulate the running OS.
func (s *BaseSuite) PatchGOOS(os string) {
	s.PatchValue(&goos, func() string { return os })
}

// PatchUnIPRouteShow allows us to simulate the return from running
// "ip route show" on the host.
func (s *BaseSuite) PatchRunIPRouteShow(run func() (string, error)) {
	s.PatchValue(&runIPRouteShow, run)
}
