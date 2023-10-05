// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretusersupplied_test

import (
	"fmt"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
	"github.com/juju/names/v4"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/authentication"
	commonsecrets "github.com/juju/juju/apiserver/common/secrets"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	facademocks "github.com/juju/juju/apiserver/facade/mocks"
	"github.com/juju/juju/apiserver/facades/controller/secretusersupplied"
	"github.com/juju/juju/apiserver/facades/controller/secretusersupplied/mocks"
	"github.com/juju/juju/core/permission"
	coresecrets "github.com/juju/juju/core/secrets"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/secrets/provider"
	coretesting "github.com/juju/juju/testing"
)

type secretusersuppliedSuite struct {
	testing.IsolationSuite

	authorizer *facademocks.MockAuthorizer
	resources  *facademocks.MockResources
	authTag    names.Tag

	state         *mocks.MockSecretsState
	stringWatcher *mocks.MockStringsWatcher

	provider *mocks.MockSecretBackendProvider
	backend  *mocks.MockSecretsBackend

	facade *secretusersupplied.SecretUserSuppliedManager
}

var _ = gc.Suite(&secretusersuppliedSuite{})

func (s *secretusersuppliedSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.authorizer = facademocks.NewMockAuthorizer(ctrl)
	s.resources = facademocks.NewMockResources(ctrl)
	s.authTag = names.NewUserTag("foo")
	s.state = mocks.NewMockSecretsState(ctrl)
	s.stringWatcher = mocks.NewMockStringsWatcher(ctrl)

	s.provider = mocks.NewMockSecretBackendProvider(ctrl)
	s.backend = mocks.NewMockSecretsBackend(ctrl)
	s.PatchValue(&commonsecrets.GetProvider, func(string) (provider.SecretBackendProvider, error) { return s.provider, nil })

	s.authorizer.EXPECT().AuthController().Return(true)

	var err error
	s.facade, err = secretusersupplied.NewTestAPI(
		s.authorizer, s.resources, s.authTag,
		coretesting.ControllerTag.Id(), coretesting.ModelTag.Id(), s.state,
		func() (*provider.ModelBackendConfigInfo, error) {
			return &provider.ModelBackendConfigInfo{
				ActiveID: "backend-id",
				Configs: map[string]provider.ModelBackendConfig{
					"backend-id": {
						ControllerUUID: coretesting.ControllerTag.Id(),
						ModelUUID:      coretesting.ModelTag.Id(),
						ModelName:      "some-model",
						BackendConfig: provider.BackendConfig{
							BackendType: "active-type",
							Config:      map[string]interface{}{"foo": "active-type"},
						},
					},
				},
			}, nil
		},
	)
	c.Assert(err, jc.ErrorIsNil)
	return ctrl
}

func (s *secretusersuppliedSuite) TestWatchObsoleteRevisionsNeedPrune(c *gc.C) {
	defer s.setup(c).Finish()

	s.state.EXPECT().WatchObsoleteRevisionsNeedPrune([]names.Tag{names.NewModelTag(coretesting.ModelTag.Id())}).Return(s.stringWatcher, nil)
	s.resources.EXPECT().Register(s.stringWatcher).Return("watcher-id")
	stringChan := make(chan []string, 1)
	stringChan <- []string{"1", "2", "3"}
	s.stringWatcher.EXPECT().Changes().Return(stringChan)

	result, err := s.facade.WatchObsoleteRevisionsNeedPrune()
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, jc.DeepEquals, params.StringsWatchResult{
		StringsWatcherId: "watcher-id",
		Changes:          []string{"1", "2", "3"},
	})
}

func (s *secretusersuppliedSuite) TestDeleteRevisionsAutoPruneEnabled(c *gc.C) {
	defer s.setup(c).Finish()

	uri := coresecrets.NewURI()
	s.authorizer.EXPECT().HasPermission(permission.SuperuserAccess, coretesting.ControllerTag).Return(
		errors.WithType(apiservererrors.ErrPerm, authentication.ErrorEntityMissingPermission))
	s.authorizer.EXPECT().HasPermission(permission.AdminAccess, coretesting.ModelTag).Return(nil)
	s.state.EXPECT().GetSecret(uri).Return(&coresecrets.SecretMetadata{
		URI: uri, OwnerTag: coretesting.ModelTag.String(),
		AutoPrune: true,
	}, nil).Times(2)
	s.state.EXPECT().DeleteSecret(uri, []int{666}).Return([]coresecrets.ValueRef{{
		BackendID:  "backend-id",
		RevisionID: "rev-666",
	}}, nil)

	cfg := &provider.ModelBackendConfig{
		ControllerUUID: coretesting.ControllerTag.Id(),
		ModelUUID:      coretesting.ModelTag.Id(),
		ModelName:      "some-model",
		BackendConfig: provider.BackendConfig{
			BackendType: "active-type",
			Config:      map[string]interface{}{"foo": "active-type"},
		},
	}
	s.provider.EXPECT().NewBackend(cfg).Return(s.backend, nil)
	s.backend.EXPECT().DeleteContent(gomock.Any(), "rev-666").Return(nil)
	s.provider.EXPECT().CleanupSecrets(
		cfg, names.NewUserTag("foo"),
		provider.SecretRevisions{uri.ID: set.NewStrings("rev-666")},
	).Return(nil)

	results, err := s.facade.DeleteRevisions(
		params.DeleteSecretArgs{
			Args: []params.DeleteSecretArg{
				{
					URI:       uri.String(),
					Revisions: []int{666},
				},
			},
		},
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{{}},
	})
}

func (s *secretusersuppliedSuite) TestDeleteRevisionsAutoPruneDisabled(c *gc.C) {
	defer s.setup(c).Finish()

	uri := coresecrets.NewURI()
	s.authorizer.EXPECT().HasPermission(permission.SuperuserAccess, coretesting.ControllerTag).Return(
		errors.WithType(apiservererrors.ErrPerm, authentication.ErrorEntityMissingPermission))
	s.authorizer.EXPECT().HasPermission(permission.AdminAccess, coretesting.ModelTag).Return(nil)
	s.state.EXPECT().GetSecret(uri).Return(&coresecrets.SecretMetadata{
		URI: uri, OwnerTag: coretesting.ModelTag.String(),
		AutoPrune: false,
	}, nil).Times(2)

	results, err := s.facade.DeleteRevisions(
		params.DeleteSecretArgs{
			Args: []params.DeleteSecretArg{
				{
					URI:       uri.String(),
					Revisions: []int{666},
				},
			},
		},
	)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(results, jc.DeepEquals, params.ErrorResults{
		Results: []params.ErrorResult{
			{
				Error: &params.Error{
					Message: fmt.Sprintf("cannot delete auto-prune secret %q", uri.String()),
				},
			},
		},
	})
}
