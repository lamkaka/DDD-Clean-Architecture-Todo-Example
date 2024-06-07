// Provides loggers to unify all logging formats
package log

import (
	"context"
	stderrors "errors"
	"fmt"
	"libs/errors"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

// Each method will output a log line with the specified log level.
// The arguments can either be an error,
// or in the form of format, ...interpolationValues, similar to the arguments for fmt.Sprintf.
// Other types of value will be converted to string before outputting.
type Logger interface {
	// Logs in panic level.
	// Use this for logging errors that the system cannot recover from.
	// e.g. A pointer iss unexpectedly nil.
	Panic(ctx context.Context, args ...interface{})

	// Logs in error level.
	// Use this for logging abnormal operation that is an error of the system.
	// e.g. A request to a downstream service has failed.
	Error(ctx context.Context, args ...interface{})

	// Logs in warn level.
	// Use this for logging abnormal operation that may not be an error of the system, possibly due to user error.
	// e.g. A user has requested a resource that does not exist.
	Warn(ctx context.Context, args ...interface{})

	// Logs in info level.
	// Use this for logging normal operation.
	// e.g. A user has successfully requested a resource.
	Info(ctx context.Context, args ...interface{})

	// Logs in debug level.
	// Use this for low level logging information for debugging.
	// e.g. A call to database took 10 seconds.
	Debug(ctx context.Context, args ...interface{})

	// Logs in trace level.
	// Use this for very low level logging information for debugging.
	Trace(ctx context.Context, args ...interface{})

	// Creates a child logger with a different context name but same config
	Child(contextName string) Logger
}

type logger struct {
	logger      *zerolog.Logger
	contextName string
}

func errorWrapper(args ...any) error {
	if len(args) == 0 {
		return nil
	}
	switch e := args[0].(type) {
	case error:
		return e
	case string:
		return fmt.Errorf(e, args[1:]...)
	default:
		return fmt.Errorf("%v", args...)
	}
}

func (l logger) Panic(ctx context.Context, args ...interface{}) {
	ev := prepare(l.logger.Panic(), withCallerInfo(0), withContext(l.contextName))
	log(ev, errorWrapper(args...))
}

func (l logger) Error(ctx context.Context, args ...interface{}) {
	ev := prepare(l.logger.Error(), withCallerInfo(0), withContext(l.contextName))
	log(ev, errorWrapper(args...))
}

func (l logger) Warn(ctx context.Context, args ...interface{}) {
	ev := prepare(l.logger.Warn(), withCallerInfo(0), withContext(l.contextName))
	log(ev, args...)
}

func (l logger) Info(ctx context.Context, args ...interface{}) {
	ev := prepare(l.logger.Info(), withCallerInfo(0), withContext(l.contextName))
	log(ev, args...)
}

func (l logger) Debug(ctx context.Context, args ...interface{}) {
	ev := prepare(l.logger.Debug(), withCallerInfo(0), withContext(l.contextName))
	log(ev, args...)
}

func (l logger) Trace(ctx context.Context, args ...interface{}) {
	ev := prepare(l.logger.Trace(), withCallerInfo(0), withContext(l.contextName))
	log(ev, args...)
}

func (l logger) Child(contextName string) Logger {
	child := l.logger.With().Logger()
	return logger{&child, contextName}
}

func prepare(ev *zerolog.Event, modifiers ...func(*zerolog.Event)) *zerolog.Event {
	for _, m := range modifiers {
		m(ev)
	}
	return ev
}

func withCallerInfo(skip int) func(*zerolog.Event) {
	return func(ev *zerolog.Event) {
		ev.Interface("caller", getCallerInfo(skip+1))
	}
}

func withContext(context string) func(*zerolog.Event) {
	return func(ev *zerolog.Event) {
		ev.Str("context", context)
	}
}

func log(ev *zerolog.Event, args ...interface{}) {
	if len(args) == 0 {
		ev.Msg("Cannot log message: no arguments provided")
		return
	}

	switch args[0].(type) {
	case string:
		format := args[0].(string)
		ev.Msgf(format, args[1:]...)
	case error:
		err := args[0].(error)
		ev.Stack().Err(err).Send()
	default:
		ev.Msgf("%v", args[0])
	}
}

func New(config Config, contextName string) (Logger, error) {
	level, err := zerolog.ParseLevel(strings.ToLower(string(config.Level)))
	if err != nil {
		return nil, err
	}

	zerolog.ErrorFieldName = zerolog.MessageFieldName
	zerolog.ErrorStackMarshaler = marshalErrorStack

	zerolog.SetGlobalLevel(level)
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &logger{contextName: contextName, logger: &l}, nil
}

func getCallerInfo(skip int) map[string]string {
	pc, file, line, ok := runtime.Caller(3 + skip)
	if !ok {
		return map[string]string{}
	}
	runtime.FuncForPC(pc)
	return map[string]string{
		"function": runtime.FuncForPC(pc).Name(),
		"file":     fmt.Sprintf("%s:%v", file, line),
	}
}

func marshalErrorStack(err error) any {
	var demoErr errors.Error
	if stderrors.As(err, &demoErr) {
		return demoErr.Stacktrace()
	}
	return stacktrace(5)
}

// stacktrace print the callstack of the caller
// skip is the number of stack frames to skip, 0 means stacktrace() itself
func stacktrace(skip int) []string {
	var programCounters = make([]uintptr, 8)
	runtime.Callers(skip+1, programCounters)
	frames := runtime.CallersFrames(programCounters)
	var stacks []string
	for {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s (%s:%v)", frame.Func.Name(), frame.File, frame.Line))
		if !more {
			break
		}
	}
	return stacks
}
