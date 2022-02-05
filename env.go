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

package golflog

import (
	"errors"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
)

const (
	LogProduction     = "LOG_PRODUCTION"
	LogImplementation = "LOG_IMPLEMENTATION"
	LogVerbosity      = "LOG_VERBOSITY"
)

var ErrUnknownImplementation = errors.New("unknown implementation")

// NewLoggerFromEnv uses the environment variables `LOG_PRODUCTION`, `LOG_IMPLEMENTATION`,
// and `LOG_VERBOSITY` to configure the logger. If they do not exist, it will default to
// configuring a production logger, with a `zap` `Configurator` at `0` verbosity
// (normal / info level).
func NewLoggerFromEnv(rootname string) (logr.Logger, error) {
	viper.SetDefault(LogProduction, true)
	viper.SetDefault(LogImplementation, "zapr")
	viper.SetDefault(LogVerbosity, 0)

	production := viper.GetBool(LogProduction)
	implementation := viper.GetString(LogImplementation)

	var configurator Configurator

	switch implementation {
	case "zapr":
		if production {
			configurator = NewZapProductionConfigurator()
		} else {
			configurator = NewZapDevelopmentConfigurator()
		}
	default:
		return logr.Logger{}, fmt.Errorf(
			"cannot create %s logger: %w",
			implementation,
			ErrUnknownImplementation,
		)
	}

	return NewLogger(
		configurator,
		rootname,
		viper.GetInt(LogVerbosity),
	)
}
