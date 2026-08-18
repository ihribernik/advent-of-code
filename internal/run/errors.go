package run

import (
	"errors"
	"fmt"
)

var (
	// ErrRunnerNotConfigured indicates that a required runner dependency is missing.
	ErrRunnerNotConfigured = errors.New("runner dependencies are not configured")
)

// ConfigurationError describes a missing dependency required to initialize a runner.
type ConfigurationError struct {
	Dependency string
	Err        error
}

func (e ConfigurationError) Error() string {
	return fmt.Sprintf("initialize runner %s: %v", e.Dependency, e.Err)
}

func (e ConfigurationError) Unwrap() error { return e.Err }

// ErrSolverNotFound describes a missing solver for a year and day.
type ErrSolverNotFound struct {
	Year int
	Day  int
}

func (e *ErrSolverNotFound) Error() string {
	return fmt.Sprintf("cannot find a solution for year %d day %d", e.Year, e.Day)
}

// ErrGetInput describes a failure to load puzzle input.
type ErrGetInput struct {
	Year int
	Day  int
	Err  error
}

func (e *ErrGetInput) Error() string {
	return fmt.Sprintf("get input for year %d day %d: %v", e.Year, e.Day, e.Err)
}

func (e *ErrGetInput) Unwrap() error { return e.Err }

// ErrSolvePart describes a failure while solving one puzzle part.
type ErrSolvePart struct {
	Part int
	Err  error
}

func (e *ErrSolvePart) Error() string {
	return fmt.Sprintf("solve part %d: %v", e.Part, e.Err)
}

func (e *ErrSolvePart) Unwrap() error { return e.Err }
