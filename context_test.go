//
// Copyright 2022 FloatMe, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golflog_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	// nolint:gci // This is the correct order
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/floatme-corp/golflog"
	"github.com/floatme-corp/golflog/mocks"
)

func monkeyPatchFallback() (func(t *testing.T), *bytes.Buffer) {
	configurator := &mocks.Configurator{}

	defaultConfigurator := golflog.DefaultConfigurator
	golflog.DefaultConfigurator = configurator

	defaultOutput := golflog.DefaultFallbackOutput
	buf := new(bytes.Buffer)
	golflog.DefaultFallbackOutput = buf

	mockErr := errors.New("mock error")
	configurator.On("Verbosity", mock.AnythingOfType("int")).Return(mockErr).Once()

	cleanup := func(t *testing.T) {
		t.Helper()

		golflog.DefaultConfigurator = defaultConfigurator
		golflog.DefaultFallbackOutput = defaultOutput

		configurator.AssertExpectations(t)
	}

	return cleanup, buf
}

type AlwaysFromContextSuite struct {
	suite.Suite
}

func (suite *AlwaysFromContextSuite) TestNilContext() {
	//nolint: staticcheck // Specifically check for nil case
	log := golflog.AlwaysFromContext(nil)
	suite.NotNil(log)
}

func (suite *AlwaysFromContextSuite) TestFromContext() {
	ctx := context.TODO()
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	ctx = logr.NewContext(ctx, log)
	log2 := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log2)

	suite.Equal(log, log2)
}

func (suite *AlwaysFromContextSuite) TestFallback() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	ctx := context.TODO()
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(` "level"=0 "msg"="test"`+"\n", buf.String())
}

func TestAlwaysFromContext(t *testing.T) {
	suite.Run(t, new(AlwaysFromContextSuite))
}

type ContextWithNameSuite struct {
	suite.Suite
}

func (suite *ContextWithNameSuite) TestContextWithName() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	ctx := golflog.ContextWithName(context.TODO(), "test")
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(`test "level"=0 "msg"="test"`+"\n", buf.String())
}

func (suite *ContextWithNameSuite) TestWithName() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	_, log := golflog.WithName(context.TODO(), "test")
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(`test "level"=0 "msg"="test"`+"\n", buf.String())
}

func (suite *ContextWithNameSuite) TestMultipleContextWithName() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	ctx := golflog.ContextWithName(context.TODO(), "test")
	ctx = golflog.ContextWithName(ctx, "test")
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(`test/test "level"=0 "msg"="test"`+"\n", buf.String())
}

func TestContextWithName(t *testing.T) {
	suite.Run(t, new(ContextWithNameSuite))
}

type ContextWithValuesSuite struct {
	suite.Suite
}

func (suite *ContextWithValuesSuite) TestWithValues() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	_, log := golflog.WithValues(context.TODO(), "test", "test")
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(` "level"=0 "msg"="test" "test"="test"`+"\n", buf.String())
}

func (suite *ContextWithValuesSuite) TestContextWithValues() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	ctx := golflog.ContextWithValues(context.TODO(), "test", "test")
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(` "level"=0 "msg"="test" "test"="test"`+"\n", buf.String())
}

func (suite *ContextWithValuesSuite) TestMultipleContextWithValues() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	ctx := golflog.ContextWithValues(context.TODO(), "test", "test")
	ctx = golflog.ContextWithValues(ctx, "test2", "test2")
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(` "level"=0 "msg"="test" "test"="test" "test2"="test2"`+"\n", buf.String())
}

func TestContextWithValues(t *testing.T) {
	suite.Run(t, new(ContextWithValuesSuite))
}

type ContextWithNameAndValuesSuite struct {
	suite.Suite
}

func (suite *ContextWithNameAndValuesSuite) TestContextWithNameAndValues() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	ctx := golflog.ContextWithNameAndValues(context.TODO(), "test", "testkey", "testval")

	golflog.Info(ctx, "test")
	suite.Equal(
		`test "level"=0 "msg"="test" "testkey"="testval" "severity"="info"`+"\n",
		buf.String(),
	)
}

func TestContextWithNameAndValues(t *testing.T) {
	suite.Run(t, new(ContextWithNameAndValuesSuite))
}

type NewContextSuite struct {
	suite.Suite
}

func (suite *NewContextSuite) TestNewContext() {
	buf, logger := bufferLogger()

	ctx := golflog.NewContext(context.TODO(), logger)
	log := golflog.AlwaysFromContext(ctx)
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(`"level"=0 "msg"="test"`, buf.String())
}

func TestNewContextSuite(t *testing.T) {
	suite.Run(t, new(NewContextSuite))
}

type WithNameAndValuesSuite struct {
	suite.Suite
}

func (suite *WithNameAndValuesSuite) TestWithNameAndValues() {
	cleanup, buf := monkeyPatchFallback()
	defer cleanup(suite.T())

	_, log := golflog.WithNameAndValues(context.TODO(), "test", "testkey", "testval")
	suite.NotNil(log)

	log.Info("test")
	suite.Equal(`test "level"=0 "msg"="test" "testkey"="testval"`+"\n", buf.String())
}

func TestWithNameAndValuesSuite(t *testing.T) {
	suite.Run(t, new(WithNameAndValuesSuite))
}
