# golflog [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci]

`golflog` is a logging utility package built around `go-logr`. It's main use
is a higher level api for building context based logging iteratively and
stored in `context.Context`.

## Installation

	go get -u github.com/floatme-corp/golflog

## Use

As close to your application entry point or initial context creation, create
a `Configurator` and `logr.Logger`, then set that logger in the
`context.Context`:
```golang
import (
    "context"
    "fmt"

    "github.com/floatme-corp/golflog"
)

func main() {
    configurator := golflog.NewZapProductionConfigurator()
    verbosity := 0

    logger, err := golflog.NewLogger(configurator, "RootLoggerName", verbosity)
    if err != nil {
        panic(fmt.Sprintf("runtime error: failed to setup logging: %s", err))
    }

    ctx := golflog.NewContext(context.Background(), logger)

    ...
}
```
Later if your application has another section, such as a `Queue` handler, you
can set that in the logger name:
```golang
func HandleQueue(ctx context.Context) {
    ctx = golflog.ContextWithName(ctx, "Queue")
}
```
All messages logged from that point forward will have the name added to the
existing name: `RootLoggerName.Queue`.

Additional values can be setup as well for future logging:
```golang
func HandleUser(ctx context.Context, userID string) {
    ctx = golflog.ContextWithValues(ctx, "user_id", userID)
}
```

The name and values can be setup together in one shot:
```golang
func HandleUser(ctx context.Context, userID string) {
    ctx = golflog.ContextWithNameAndValues(ctx, "Queue", "user_id", userID)
}
```

Functions are guaranteed to be able to get a logger from any context:
```golang
func randoFunc(ctx context.Context) {
    log := golflog.AlwaysFromContext(ctx)
    log.Info("my log message")
}
```
If the context does not have a logger associated with it `golflog` will
create a fallback logger with the default configuration. If that fails
it will fallback to logging via `fmt.Fprintln` to `os.Stdout`

`golflog` provides helpers to create a log and a context with a name, values,
or both:
```golang
func randoFunc(ctx context.Context) {
    // for tracing logs from this function
    ctx, log := golflog.WithName(ctx, "randoFunc")
    log.Info("my log message") // logs `randoFunc "msg"="my log message"`
}

// assume you have the important value `foo`
func randoFunc(ctx context.Context, importantValue string) {
    // for always seeing important value in future logs with this context
    ctx, log := golflog.WithValues(ctx, "important-value", importantValue)
    log.Info("my log message")
    // logs `"msg"="my log message" "important-value"="foo"`
}

func randoFunc(ctx context.Context, importantValue string) {
    ctx, log := golflog.WithNameAndValues(
        ctx,
        "randoFunc",
        "important-value", importantValue,
    )
    log.Info("my log message")
    // logs `randoFunc "msg"="my log message" "important-value"="foo"`
}
```

### Logging Convenience Helpers

`Info` and `Error` convenience helpers are provided to log or report an error
to the logger in the context. A `severity` key is added to each to invocation
automatically to assist log aggregation services.

```golang
func randoFunc(ctx context.Context) {
    // Same as:
    // log := golflog.AlwaysFromContext(ctx)
    // log.Info("message", "severity", "info", "key", "value")
    golflog.Info(ctx, "message", "key", "value")

    ...

    // Same as:
    // log := golflog.AlwaysFromContext(ctx)
    // log.Error(err, "message", "severity", "error", "key", "value")
    golflog.Error(ctx, err, "message", "key", "value")
}
```

A `Wrap` helper will log the message at the info level and return a new error
wrapping the `err` with `message`. A `WarnWrap` helper is also provided to log
the message with `severity=warning` prepended to the list of key/value pairs
before returning the wrapped error.

```golang
func randoFunc(ctx context.Context) {
    ...

    if err != nil {
        // Same as:
        // golflog.Error(ctx, err, "message", "key", "value")
        // return fmt.Errorf("%s: %w", "message", "err")
        return golflog.Wrap(ctx, err, "message", "key", "value")
    }
    ...

    thingID := "1234"
    if err := findThing(thingID); err != nil {
        // Suppose you expect this error
        if err.As(ErrThingNotFound) {
            // Same as:
            // golflog.Warn(ctx, "message", "thing_id", thingID)
            // return fmt.Errorf("%s: %w", "message", "err")
            return golflog.WarnWrap(ctx, err, "Could not find thing!", "thing_id", thingID)
        }
    }
}
```

A `V` helper will return a logger from the context at the level requested.

```golang
func randoFunc(ctx context.Context) {
    ...

    // Will be logged at level 1
    golflog.V(ctx, 1).Info("message", "key", "value")
}
```

A `Warn` helper will automatically prepend `severity=warning` to the list of
key/value pairs and logs out the message. This helper should be used sparingly,
and instead logging levels should be used. It is included to make it easier
for log aggregation services to key off.

```golang
func randoFunc(ctx context.Context) {
    ...

    // The same as
    // golflog.Info(ctx, "message", "severity", "warning", "key", "value")
    golflog.Warn(ctx, "message", "key", "value")
}
```

`Warning` is an alias of `Warn`.

```golang
func randoFunc(ctx context.Context) {
    ...

    // The same as
    // golflog.Info(ctx, "message", "severity", "warning", "key", "value")
    golflog.Warning(ctx, "message", "key", "value")
}
```

A `Debug` helper will automatically prepend `severity=debug` to the list of
key/value pairs and logs out the message at a level 1 higher than the level
of the logger in the `context.Context`. This helper should be used sparingly,
and instead logging levels should be used. It is included to make it easier
for log aggregation services to key off.

```golang
func randoFunc(ctx context.Context) {
    ...

    // The same as
    // golflog.V(ctx, 1).Info("message", "severity", "warning", "key", "value")
    golflog.Debug(ctx, "message", "key", "value")
}
```

### Env setup

An alternative to calling `golflog.NewLogger` with the parameters, is to call
`NewLoggerFromEnv` and give it only the root name:
```golang
import (
    "context"
    "fmt"

    "github.com/floatme-corp/golflog
)

func main() {
    logger, err := golflog.NewLoggerFromEnv("RootLoggerName")
    if err != nil {
        panic(fmt.Sprintf("runtime error: failed to setup logging: %s", err))
    }

    ctx := golflog.NewContext(context.Background(), logger)

    ...
}
```
`NewLoggerFromEnv` uses the environment variables `LOG_PRODUCTION`,
`LOG_IMPLEMENTATION`, and `LOG_VERBOSITY` to configure the logger. If they
do not exist, it will default to configuring a production logger, with
a `zap` `Configurator` at `0` verbosity (normal / info level).

-------------------------------------------------------------------------------

Released under the [Apache 2.0 License].

[Apache 2.0 License]: LICENSE
[doc-img]: https://pkg.go.dev/badge/github.com/floatme-corp/golflog
[doc]: https://pkg.go.dev/github.com/floatme-corp/golflog
[ci-img]: https://github.com/floatme-corp/golflog/actions/workflows/test.yaml/badge.svg
[ci]: https://github.com/floatme-corp/golflog/actions/workflows/test.yaml
