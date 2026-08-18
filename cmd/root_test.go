package cmd

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/container"
)

func TestExecuteReturnsMappedStartupError(t *testing.T) {
	newContainer := container.New
	originalNewContainer := newContainer
	t.Cleanup(func() { newContainer = originalNewContainer })

	cause := errors.New("registration failed")
	newContainer = func() (container.Container, error) { return nil, cause }

	err := Execute()
	if err == nil {
		t.Fatal("expected startup error")
	}
	if err.Error() != "cannot initialize application" {
		t.Fatalf("unexpected startup message: %q", err.Error())
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected startup error to wrap %v", cause)
	}
}
