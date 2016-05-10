package env

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type ENVSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&ENVSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *ENVSuite) TestEnv(c *C) {
	envs := Get()

	c.Assert(envs["EK_TEST_PORT"], Equals, "8080")

	c.Assert(envs.GetS("EK_TEST_PORT"), Equals, "8080")
	c.Assert(envs.GetI("EK_TEST_PORT"), Equals, 8080)
	c.Assert(envs.GetF("EK_TEST_PORT"), Equals, 8080.0)

	c.Assert(envs.GetS("UNKNOWN_VARIABLE"), Equals, "")
	c.Assert(envs.GetI("UNKNOWN_VARIABLE"), Equals, -1)
	c.Assert(envs.GetF("UNKNOWN_VARIABLE"), Equals, -1.0)

	c.Assert(Which("cat"), Not(Equals), "")
	c.Assert(Which("catABCD1234"), Equals, "")
}