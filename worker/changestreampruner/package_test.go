// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package changestreampruner

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	coredatabase "github.com/juju/juju/core/database"
	dbtesting "github.com/juju/juju/database/testing"
	gc "gopkg.in/check.v1"
)

//go:generate go run github.com/golang/mock/mockgen -package changestreampruner -destination stream_mock_test.go github.com/juju/juju/worker/changestreampruner DBGetter,Logger
//go:generate go run github.com/golang/mock/mockgen -package changestreampruner -destination clock_mock_test.go github.com/juju/clock Clock,Timer
//go:generate go run github.com/golang/mock/mockgen -package changestreampruner -destination worker_mock_test.go github.com/juju/worker/v3 Worker

func TestPackage(t *testing.T) {
	gc.TestingT(t)
}

type baseSuite struct {
	dbtesting.DBSuite

	dbGetter *MockDBGetter
	clock    *MockClock
	timer    *MockTimer
	logger   *MockLogger
}

func (s *baseSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.dbGetter = NewMockDBGetter(ctrl)
	s.clock = NewMockClock(ctrl)
	s.timer = NewMockTimer(ctrl)
	s.logger = NewMockLogger(ctrl)

	return ctrl
}

func (s *baseSuite) expectControllerDBGet() {
	s.dbGetter.EXPECT().GetDB(coredatabase.ControllerNS).Return(s.TxnRunner(), nil).Times(2)
}

func (s *baseSuite) expectDBGet(namespace string, txnRunner coredatabase.TxnRunner) {
	s.dbGetter.EXPECT().GetDB(namespace).Return(txnRunner, nil)
}

func (s *baseSuite) expectAnyLogs(c *gc.C) {
	s.logger.EXPECT().Errorf(gomock.Any()).Do(c.Logf).AnyTimes()
	s.logger.EXPECT().Warningf(gomock.Any()).Do(c.Logf).AnyTimes()
	s.logger.EXPECT().Infof(gomock.Any(), gomock.Any()).Do(c.Logf).AnyTimes()
	s.logger.EXPECT().Debugf(gomock.Any()).Do(c.Logf).AnyTimes()
	s.logger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Do(c.Logf).AnyTimes()
	s.logger.EXPECT().IsTraceEnabled().Return(false).AnyTimes()
}

func (s *baseSuite) expectClock() {
	s.clock.EXPECT().Now().Return(time.Now()).AnyTimes()
}
