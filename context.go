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
	"context"
	"fmt"
	"io"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
)

// DefaultConfigurator is the default configurator for fallback logging.
//nolint: gochecknoglobals // Allow for testing
var DefaultConfigurator Configurator = NewZapProductionConfigurator()

// DefaultFallbackOutput is the default io stream jfor fallback logging.
//nolint: gochecknoglobals // Allow for testing
var DefaultFallbackOutput io.Writer = os.Stdout

// AlwaysFromContext retrieves the logger from the context, if it is not
// found a new production logger at `MaxLevel` will be returned. If that
// fails a `fmt.Println` based logger is returned.
// nolint:contextcheck // ALWAYS!
func AlwaysFromContext(ctx context.Context) logr.Logger {
	if ctx == nil {
		ctx = context.Background()
	}

	log, err := logr.FromContext(ctx)
	if err == nil {
		return log
	}

	log, err = NewLogger(DefaultConfigurator, "fallback", MaxLevel)
	if err == nil {
		return log
	}

	fallback := funcr.New(
		func(prefix, args string) {
			fmt.Fprintln(DefaultFallbackOutput, prefix, args)
		},
		funcr.Options{
			Verbosity: MaxLevel,
		},
	)

	return fallback
}

// ContextWithName returns a context with the name set in its logger.
func ContextWithName(ctx context.Context, name string) context.Context {
	return NewContext(ctx, AlwaysFromContext(ctx).WithName(name))
}

// ContextWithValues returns a context with the values set in its logger.
func ContextWithValues(ctx context.Context, keysAndValues ...interface{}) context.Context {
	return NewContext(ctx, AlwaysFromContext(ctx).WithValues(keysAndValues...))
}

// ContextWithNameAndValues returns a context with the name and values set in its logger.
func ContextWithNameAndValues(
	ctx context.Context,
	name string,
	keysAndValues ...interface{},
) context.Context {
	return ContextWithName(ContextWithValues(ctx, keysAndValues...), name)
}

// WithName returns a context and logger with the given `name` set in the context.
func WithName(ctx context.Context, name string) (context.Context, logr.Logger) {
	newCtx := ContextWithName(ctx, name)

	return newCtx, AlwaysFromContext(newCtx)
}

// WithValues returns a context and logger with the given `values` set in the context.
func WithValues(ctx context.Context, keysAndValues ...interface{}) (context.Context, logr.Logger) {
	newCtx := ContextWithValues(ctx, keysAndValues...)

	return newCtx, AlwaysFromContext(newCtx)
}

// WithNameAndValues returns a context and logger with the given `name` and `values` set in the
// context.
func WithNameAndValues(
	ctx context.Context,
	name string,
	keysAndValues ...interface{},
) (context.Context, logr.Logger) {
	newCtx := ContextWithNameAndValues(ctx, name, keysAndValues...)

	return newCtx, AlwaysFromContext(newCtx)
}

// NewContext returns a context with the specified logger set in it.
func NewContext(ctx context.Context, log logr.Logger) context.Context {
	return logr.NewContext(ctx, log)
}
