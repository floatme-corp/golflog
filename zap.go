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
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConfigurator struct {
	Zap zap.Config
}

func NewZapDevelopmentConfigurator() *ZapConfigurator {
	return &ZapConfigurator{
		Zap: zap.NewDevelopmentConfig(),
	}
}

func NewZapProductionConfigurator() *ZapConfigurator {
	return &ZapConfigurator{
		Zap: zap.NewProductionConfig(),
	}
}

// Build a new `zap` `logr.Logger`.
func (config *ZapConfigurator) Build() (logr.Logger, error) {
	log, err := config.Zap.Build()
	if err != nil {
		return logr.Logger{}, fmt.Errorf("failed to build zap logger: %w", err)
	}

	return zapr.NewLogger(log), nil
}

// Verbosity sets the desirect log verbosity in the `zap` config.
func (config *ZapConfigurator) Verbosity(verbosity int) error {
	config.Zap.Level = zap.NewAtomicLevelAt(zapcore.Level(verbosity) * -1)

	return nil
}

var _ Configurator = (*ZapConfigurator)(nil)
