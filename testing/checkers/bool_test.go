// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package checkers_test

import (
	. "launchpad.net/gocheck"
	. "launchpad.net/juju-core/testing/checkers"
)

type BoolSuite struct{}

var _ = Suite(&BoolSuite{})

func (s *BoolSuite) TestIsTrue(c *C) {
	c.Assert(true, IsTrue)
	c.Assert(false, Not(IsTrue))
}

func (s *BoolSuite) TestIsFalse(c *C) {
	c.Assert(false, IsFalse)
	c.Assert(true, Not(IsFalse))
}
