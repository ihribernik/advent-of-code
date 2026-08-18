package cmd

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/inputs"
	runusecase "github.com/ihribernik/aoc-cli/internal/run"
)

func TestMapRunError(t *testing.T) {
	inputCause := errors.New("permission denied")
	solveCause := errors.New("boom")
	fallbackCause := errors.New("unexpected")

	tests := []struct {
		name    string
		year    int
		day     int
		err     error
		message string
		wantIs  error
	}{
		{
			name:    "solver not found",
			year:    2015,
			day:     9,
			err:     &runusecase.ErrSolverNotFound{Year: 2015, Day: 9},
			message: "no solver registered for year 2015 day 09",
		},
		{
			name:    "input file empty",
			year:    2015,
			day:     2,
			err:     &runusecase.ErrGetInput{Year: 2015, Day: 2, Err: inputs.ErrEmptyInput},
			message: "input file is empty for year 2015 day 02; download the puzzle input from Advent of Code",
			wantIs:  inputs.ErrEmptyInput,
		},
		{
			name:    "input file missing",
			year:    2015,
			day:     2,
			err:     &runusecase.ErrGetInput{Year: 2015, Day: 2, Err: fmt.Errorf("open x: %w", os.ErrNotExist)},
			message: "input file not found for year 2015 day 02",
			wantIs:  os.ErrNotExist,
		},
		{
			name:    "input other error",
			year:    2015,
			day:     2,
			err:     &runusecase.ErrGetInput{Year: 2015, Day: 2, Err: inputCause},
			message: "cannot load input for year 2015 day 02",
			wantIs:  inputCause,
		},
		{
			name:    "solve part",
			year:    2015,
			day:     2,
			err:     &runusecase.ErrSolvePart{Part: 2, Err: solveCause},
			message: "failed while solving part 2",
			wantIs:  solveCause,
		},
		{
			name:    "fallback",
			year:    2015,
			day:     2,
			err:     fallbackCause,
			message: "run failed for year 2015 day 02",
			wantIs:  fallbackCause,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := mapRunError(tc.year, tc.day, tc.err)
			if got.Error() != tc.message {
				t.Fatalf("expected %q, got %q", tc.message, got.Error())
			}
			if !errors.Is(got, tc.err) {
				t.Fatalf("expected mapped error to wrap original error %v", tc.err)
			}
			if tc.wantIs != nil && !errors.Is(got, tc.wantIs) {
				t.Fatalf("expected mapped error to wrap %v", tc.wantIs)
			}
		})
	}
}

func TestMapStartupError(t *testing.T) {
	cause := errors.New("registration failed")
	err := mapStartupError(cause)

	if err.Error() != "cannot initialize application" {
		t.Fatalf("unexpected startup message: %q", err.Error())
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected startup error to wrap %v", cause)
	}
}
