package puzzle

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidInput indicates that puzzle input is invalid.
	ErrInvalidInput = errors.New("invalid puzzle input")
	// ErrAnswerNotFound indicates that a solver could not find an answer.
	ErrAnswerNotFound = errors.New("puzzle answer not found")
	// ErrNotImplemented indicates that a puzzle solution is not implemented.
	ErrNotImplemented = errors.New("puzzle solution not implemented")
)

// InputError describes invalid input at a puzzle-specific location.
type InputError struct {
	Year  int
	Day   int
	Line  int
	Input string
	Err   error
}

func (e *InputError) Error() string {
	location := fmt.Sprintf("invalid input for %d day %02d", e.Year, e.Day)
	if e.Line > 0 {
		location += fmt.Sprintf(" line %d", e.Line)
	}
	if e.Input != "" {
		location += fmt.Sprintf(" %q", e.Input)
	}
	if e.Err != nil {
		location += ": " + e.Err.Error()
	}
	return location
}

func (e *InputError) Is(target error) bool { return target == ErrInvalidInput }

func (e *InputError) Unwrap() error { return e.Err }
