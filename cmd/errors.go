package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ihribernik/aoc-cli/internal/inputs"
	runusecase "github.com/ihribernik/aoc-cli/internal/run"
)

type cliError struct {
	message string
	cause   error
}

func (e *cliError) Error() string { return e.message }

func (e *cliError) Unwrap() error { return e.cause }

func newCLIError(message string, cause error) error {
	return &cliError{message: message, cause: cause}
}

func mapStartupError(err error) error {
	return newCLIError("cannot initialize application", err)
}

func mapRunError(year int, day int, err error) error {
	var solverErr *runusecase.ErrSolverNotFound
	if errors.As(err, &solverErr) {
		return newCLIError(
			fmt.Sprintf("no solver registered for year %d day %02d", solverErr.Year, solverErr.Day),
			err,
		)
	}

	var inputErr *runusecase.ErrGetInput
	if errors.As(err, &inputErr) {
		if errors.Is(inputErr, inputs.ErrEmptyInput) {
			return newCLIError(
				fmt.Sprintf("input file is empty for year %d day %02d; download the puzzle input from Advent of Code", inputErr.Year, inputErr.Day),
				err,
			)
		}
		if errors.Is(inputErr, os.ErrNotExist) {
			return newCLIError(
				fmt.Sprintf("input file not found for year %d day %02d", inputErr.Year, inputErr.Day),
				err,
			)
		}
		return newCLIError(
			fmt.Sprintf("cannot load input for year %d day %02d", inputErr.Year, inputErr.Day),
			err,
		)
	}

	var solveErr *runusecase.ErrSolvePart
	if errors.As(err, &solveErr) {
		return newCLIError(fmt.Sprintf("failed while solving part %d", solveErr.Part), err)
	}

	return newCLIError(fmt.Sprintf("run failed for year %d day %02d", year, day), err)
}
