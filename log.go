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

//go:generate mockery --name Configurator --case underscore
//go:generate mockery --name BuildInfo --case underscore

package golflog

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
)

const (
	// MaxLevel is the maximum logging level supported. All verbosity values higher will be
	// clamped to this value.
	MaxLevel = 127

	// MinLevel is the minimum logging level supported. All verbosity values lower will be
	// clamped to this value.
	MinLevel = 0
)

// BuildInfo allows the application to add build info to the logger by default. Any non empty
// value returned will be set in the logger when calling `NewLoggerWithBuildInfo`.
type BuildInfo interface {
	Version() string
	Commit() string
	Date() string
	Time() string
}

// Configurator allows a pluggable logging implementation.
type Configurator interface {
	// Build and return a `logr.Logger` instance.
	Build() (logr.Logger, error)

	// Verbosity sets the desired verbosity level. NOTE this will be called prior to `Build`.
	Verbosity(verbosity int) error
}

// NewLogger sets the log verbosity, configures a new logger from `configurator`, and sets the
// initial name of the logger.
func NewLogger(
	configurator Configurator,
	rootName string,
	verbosity int,
) (logr.Logger, error) {
	switch {
	case verbosity > MaxLevel:
		verbosity = MaxLevel

	case verbosity < MinLevel:
		verbosity = MinLevel
	}

	if err := configurator.Verbosity(verbosity); err != nil {
		return logr.Logger{}, fmt.Errorf("failed to set verbosity to %d: %w", verbosity, err)
	}

	log, err := configurator.Build()
	if err != nil {
		return logr.Logger{}, fmt.Errorf("failed to get logger: %w", err)
	}

	return log.WithName(rootName), nil
}

// NewLogger sets the log verbosity, configures a new logger from `configurator`, and sets the
// initial name of the logger. If `buildInfo` is not `nil`, non-empty values from the interface
// will be set as values on the resulting `logr.Logger`.
func NewLoggerWithBuildInfo(
	configurator Configurator,
	buildInfo BuildInfo,
	rootName string,
	verbosity int,
) (logr.Logger, error) {
	log, err := NewLogger(configurator, rootName, verbosity)
	if err != nil {
		return logr.Logger{}, fmt.Errorf("failed to get logger: %w", err)
	}

	if buildInfo != nil {
		version := buildInfo.Version()
		if version != "" {
			log = log.WithValues("build_version", version)
		}

		commit := buildInfo.Commit()
		if commit != "" {
			log = log.WithValues("build_commit", commit)
		}

		date := buildInfo.Date()
		if date != "" {
			log = log.WithValues("build_date", date)
		}

		time := buildInfo.Time()
		if time != "" {
			log = log.WithValues("build_time", time)
		}
	}

	return log, nil
}

// Wrap logs the error on the given logger and returns it wrapped by message.
func Wrap(
	ctx context.Context,
	err error,
	message string,
	keysAndValues ...interface{},
) error {
	logger := AlwaysFromContext(ctx)
	logger.WithCallDepth(1).Error(err, message, keysAndValues...)

	return fmt.Errorf("%s: %w", message, err)
}

// Gets a logger from the given context and logs message and optional values.
func Info(
	ctx context.Context,
	message string,
	keysAndValues ...interface{},
) {
	logger := AlwaysFromContext(ctx)
	logger.WithCallDepth(1).Info(message, keysAndValues...)
}
