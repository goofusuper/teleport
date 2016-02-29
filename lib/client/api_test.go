/*
Copyright 2016 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"gopkg.in/check.v1"
	"testing"
)

// register test suite
type APITestSuite struct {
}

// bootstrap check
func TestClientAPI(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&APITestSuite{})

func (s *APITestSuite) SetUpSuite(c *check.C) {
}

func (s *APITestSuite) TestConfig(c *check.C) {
	var conf Config
	c.Assert(conf.ProxySpecified(), check.Equals, false)
	conf.ProxyHost = "example.org"
	c.Assert(conf.ProxySpecified(), check.Equals, true)
	c.Assert(conf.ProxyHostPort(12), check.Equals, "example.org:12")
}