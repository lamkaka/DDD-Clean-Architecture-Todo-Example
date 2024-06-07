// Defines a common error format for demo projects.
package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// Capitalized strings that indicate the type of error.
// errors package provides some common error codes, but each project can define error codes that better describe the error they are throwing
type ErrorCode string

const (
	ErrorUnauthorized     ErrorCode = "UNAUTHORIZED_ERROR"
	ErrorEntityNotFound   ErrorCode = "ENTITY_NOT_FOUND_ERROR"
	ErrorEntityValidation ErrorCode = "ENTITY_VALIDATION_ERROR"
	ErrorEntityConflict   ErrorCode = "ENTITY_CONFLICT_ERROR"
	ErrorUnexpected       ErrorCode = "UNEXPECTED_ERROR"
)

// Common error format for all demo projects
type Error interface {
	Code() ErrorCode
	Messages() []string
	Stacktrace() []string
	error
}

type _error struct {
	code            ErrorCode
	messages        []string
	programCounters []uintptr
}

func New(code ErrorCode, messages ...string) Error {
	return newError(code, 2, messages...)
}

func newError(code ErrorCode, skipStack int, messages ...string) Error {
	var programCounters = make([]uintptr, 8)
	runtime.Callers(skipStack+1, programCounters)
	return _error{
		code:            code,
		messages:        messages,
		programCounters: programCounters,
	}
}

func (err _error) Is(e error) bool {
	_err, ok := e.(_error)
	if !ok {
		return false
	}
	if _err.code != err.code {
		return false
	}
	return equalSlice(err.messages, _err.messages) && equalSlice(err.programCounters, _err.programCounters)
}

func equalSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Returns the error code
func (err _error) Code() ErrorCode {
	return err.code
}

// Returns the error messages
func (err _error) Messages() []string {
	return err.messages
}

// Returns a string summary of the error
func (err _error) Error() string {
	return string(err.code) + ": " + strings.Join(err.messages, ", ")
}

// Return stack trace in a string array, with each element being a description of the caller
func (err _error) Stacktrace() []string {
	var stacks []string
	frames := runtime.CallersFrames(err.programCounters[:])
	for {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s (%s:%v)", frame.Func.Name(), frame.File, frame.Line))
		if !more {
			break
		}
	}
	return stacks
}
