package registry

import (
	"errors"
	"fmt"
)

var (
	// ErrNilRegistry indicates that a registry dependency is nil.
	ErrNilRegistry = errors.New("nil registry")
	// ErrNilSolver indicates that a nil solver was registered.
	ErrNilSolver = errors.New("nil solver")
	// ErrSolverAlreadyRegistered indicates that a solver key is already in use.
	ErrSolverAlreadyRegistered = errors.New("solver already registered")
)

// RegistrationError describes a failed solver registration.
type RegistrationError struct {
	Year int
	Day  int
	Err  error
}

func (e *RegistrationError) Error() string {
	return fmt.Sprintf("register solver for year %d day %02d: %v", e.Year, e.Day, e.Err)
}

func (e *RegistrationError) Unwrap() error { return e.Err }
