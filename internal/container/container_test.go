package container

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/run"
)

func TestNew(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("unexpected construction error: %v", err)
	}
	if c == nil || c.GetRunner() == nil {
		t.Fatal("expected a configured container")
	}
}

func TestNewReturnsRegistrationError(t *testing.T) {
	cause := errors.New("duplicate solver")
	c, err := newContainer(func(registry.Registry) error { return cause }, run.NewRunner)

	if c != nil {
		t.Fatalf("expected nil container, got %v", c)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected construction error to wrap %v, got %v", cause, err)
	}
	if err.Error() != "initialize solver registry: duplicate solver" {
		t.Fatalf("unexpected construction message: %q", err.Error())
	}
}

func TestNewReturnsRunnerInitializationError(t *testing.T) {
	cause := run.ConfigurationError{
		Dependency: "input loader",
		Err:        run.ErrRunnerNotConfigured,
	}
	c, err := newContainer(
		func(registry.Registry) error { return nil },
		func(run.SolverResolver, run.LoadInputFunc) (run.Runner, error) { return nil, cause },
	)

	if c != nil {
		t.Fatalf("expected nil container, got %v", c)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected construction error to wrap %v, got %v", cause, err)
	}
	var configurationErr run.ConfigurationError
	if !errors.As(err, &configurationErr) {
		t.Fatalf("expected ConfigurationError, got %T", err)
	}
	if configurationErr.Dependency != "input loader" {
		t.Fatalf("unexpected dependency: %q", configurationErr.Dependency)
	}
	if !errors.Is(err, run.ErrRunnerNotConfigured) {
		t.Fatalf("expected ErrRunnerNotConfigured, got %v", err)
	}
	const wantMessage = "initialize runner: initialize runner input loader: runner dependencies are not configured"
	if err.Error() != wantMessage {
		t.Fatalf("expected error message %q, got %q", wantMessage, err.Error())
	}
}
