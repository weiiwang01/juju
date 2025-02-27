// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package controller

import (
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/core/migration"
	"github.com/juju/juju/state"
)

type patcher interface {
	PatchValue(destination, source interface{})
}

func SetPreCheckResult(p patcher, err error) {
	p.PatchValue(&runMigrationPreChecks, func(*state.State, *state.State, *migration.TargetInfo, facade.Presence, map[string]string) error {
		return err
	})
}

func NewControllerAPIForTest(backend Backend) *ControllerAPI {
	return &ControllerAPI{state: backend}
}

var (
	NewControllerAPIv11 = newControllerAPIv11
)
