// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package machinemanager_test

import (
	"testing"

	coretesting "github.com/juju/juju/testing"
)

//go:generate go run github.com/golang/mock/mockgen -package mocks -destination mocks/leadership_mock.go github.com/juju/juju/apiserver/facades/client/machinemanager Leadership
//go:generate go run github.com/golang/mock/mockgen -package mocks -destination mocks/upgradeseries_mock.go github.com/juju/juju/apiserver/facades/client/machinemanager Authorizer,UpgradeSeries,UpgradeSeriesState,UpgradeSeriesValidator
//go:generate go run github.com/golang/mock/mockgen -package mocks -destination mocks/types_mock.go github.com/juju/juju/apiserver/facades/client/machinemanager Backend,Pool,Machine,Application,Unit,Charm,CharmhubClient

func TestPackage(t *testing.T) {
	// TODO(wallyworld) - needed until instance config tests converted to gomock
	coretesting.MgoTestPackage(t)
}
