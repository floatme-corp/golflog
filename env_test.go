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
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"

	"github.com/floatme-corp/golflog"
)

type NewLoggerFromEnvSuite struct {
	suite.Suite
}

func (suite *NewLoggerFromEnvSuite) TestDefault() {
	log, err := golflog.NewLoggerFromEnv("test")
	if suite.NoError(err) {
		suite.NotNil(log)
	}
}

func (suite *NewLoggerFromEnvSuite) TestDevelopment() {
	viper.Set(golflog.LogProduction, false)

	log, err := golflog.NewLoggerFromEnv("test")
	if suite.NoError(err) {
		suite.NotNil(log)
	}
}

func (suite *NewLoggerFromEnvSuite) TestProduction() {
	viper.Set(golflog.LogProduction, true)

	log, err := golflog.NewLoggerFromEnv("test")
	if suite.NoError(err) {
		suite.NotNil(log)
	}
}

func (suite *NewLoggerFromEnvSuite) TestUnknownImplementation() {
	viper.Set(golflog.LogImplementation, "test")

	_, err := golflog.NewLoggerFromEnv("test")
	if suite.Error(err) {
		suite.ErrorIs(err, golflog.ErrUnknownImplementation)
	}
}

func TestNewLoggerFromEnv(t *testing.T) {
	suite.Run(t, new(NewLoggerFromEnvSuite))
}
