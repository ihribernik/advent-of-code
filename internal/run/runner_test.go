package run

import (
	"errors"
	"testing"

	"github.com/ihribernik/aoc-cli/internal/registry"
)

type stubSolver struct {
	part1    int
	part2    int
	part1Err error
	part2Err error
}

type mutatingSolver struct {
	part2FirstValue string
}

type stubResolver struct {
	solver registry.Solver
	found  bool
}

func (r stubResolver) GetSolver(year int, day int) (registry.Solver, bool) {
	return r.solver, r.found
}

func registerSolver(t *testing.T, reg registry.Registry, year, day int, solver registry.Solver) {
	t.Helper()
	if err := reg.Register(year, day, solver); err != nil {
		t.Fatalf("register solver for year %d day %d: %v", year, day, err)
	}
}

func mustNewRunner(t *testing.T, resolver SolverResolver, loader LoadInputFunc) Runner {
	t.Helper()
	runner, err := NewRunner(resolver, loader)
	if err != nil {
		t.Fatalf("construct runner: %v", err)
	}
	return runner
}

func (s *mutatingSolver) SolvePart1(input []string) (int, error) {
	input[0] = "mutated"
	return 1, nil
}

func (s *mutatingSolver) SolvePart2(input []string) (int, error) {
	s.part2FirstValue = input[0]
	input[0] = "also mutated"
	return 2, nil
}

func (s stubSolver) SolvePart1(input []string) (int, error) {
	if s.part1Err != nil {
		return 0, s.part1Err
	}
	return s.part1, nil
}

func (s stubSolver) SolvePart2(input []string) (int, error) {
	if s.part2Err != nil {
		return 0, s.part2Err
	}
	return s.part2, nil
}

func TestRunnerExecuteSuccess(t *testing.T) {
	reg := registry.NewRegistry()
	registerSolver(t, reg, 2015, 6, stubSolver{part1: 123, part2: 456})
	r := mustNewRunner(t,
		reg,
		func(year int, day int) ([]string, error) {
			return []string{"input"}, nil
		},
	)

	result, err := r.Execute(2015, 6)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result.Part1 != 123 || result.Part2 != 456 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestRunnerExecuteIsolatesInputForEachPart(t *testing.T) {
	reg := registry.NewRegistry()
	solver := &mutatingSolver{}
	registerSolver(t, reg, 2015, 1, solver)
	loadedInput := []string{"original"}
	r := mustNewRunner(t,
		reg,
		func(year int, day int) ([]string, error) { return loadedInput, nil },
	)

	result, err := r.Execute(2015, 1)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if result.Part1 != 1 || result.Part2 != 2 {
		t.Fatalf("unexpected result: %+v", result)
	}
	if solver.part2FirstValue != "original" {
		t.Fatalf("expected part 2 to receive original input, got %q", solver.part2FirstValue)
	}
	if loadedInput[0] != "original" {
		t.Fatalf("expected loaded input to remain unchanged, got %q", loadedInput[0])
	}
}

func TestNewRunnerRejectsMissingDependencies(t *testing.T) {
	tests := []struct {
		name       string
		resolver   SolverResolver
		loader     LoadInputFunc
		dependency string
		message    string
	}{
		{
			name:       "solver resolver",
			loader:     func(int, int) ([]string, error) { return nil, nil },
			dependency: "solver resolver",
			message:    "initialize runner solver resolver: runner dependencies are not configured",
		},
		{
			name:       "input loader",
			resolver:   stubResolver{},
			dependency: "input loader",
			message:    "initialize runner input loader: runner dependencies are not configured",
		},
		{
			name:       "resolver reported first when both are missing",
			dependency: "solver resolver",
			message:    "initialize runner solver resolver: runner dependencies are not configured",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewRunner(tt.resolver, tt.loader)
			if r != nil {
				t.Fatalf("expected nil runner, got %v", r)
			}
			if !errors.Is(err, ErrRunnerNotConfigured) {
				t.Fatalf("expected ErrRunnerNotConfigured, got %v", err)
			}

			var configurationErr ConfigurationError
			if !errors.As(err, &configurationErr) {
				t.Fatalf("expected ConfigurationError, got %T", err)
			}
			if configurationErr.Dependency != tt.dependency {
				t.Fatalf("expected dependency %q, got %q", tt.dependency, configurationErr.Dependency)
			}
			if configurationErr.Err != ErrRunnerNotConfigured {
				t.Fatalf("expected sentinel cause, got %v", configurationErr.Err)
			}
			if err.Error() != tt.message {
				t.Fatalf("expected error message %q, got %q", tt.message, err.Error())
			}
		})
	}
}

func TestRunnerExecuteNilSolver(t *testing.T) {
	inputLoaded := false
	r := mustNewRunner(t,
		stubResolver{found: true},
		func(year int, day int) ([]string, error) {
			inputLoaded = true
			return []string{}, nil
		},
	)

	_, err := r.Execute(2015, 1)
	var solverErr *ErrSolverNotFound
	if !errors.As(err, &solverErr) {
		t.Fatalf("expected ErrSolverNotFound, got %v", err)
	}
	if inputLoaded {
		t.Fatal("expected input loader not to be called")
	}
}

func TestRunnerExecuteSolverNotFound(t *testing.T) {
	r := mustNewRunner(t,
		registry.NewRegistry(),
		func(year int, day int) ([]string, error) { return []string{}, nil },
	)

	_, err := r.Execute(2015, 2)
	var solverErr *ErrSolverNotFound
	if !errors.As(err, &solverErr) {
		t.Fatalf("expected ErrSolverNotFound, got %v", err)
	}
	if solverErr.Year != 2015 || solverErr.Day != 2 {
		t.Fatalf("unexpected solver error payload: %+v", solverErr)
	}
}

func TestRunnerExecuteGetInputError(t *testing.T) {
	wantErr := errors.New("input missing")
	reg := registry.NewRegistry()
	registerSolver(t, reg, 2015, 1, stubSolver{})
	r := mustNewRunner(t,
		reg,
		func(year int, day int) ([]string, error) { return nil, wantErr },
	)

	_, err := r.Execute(2015, 1)
	var inputErr *ErrGetInput
	if !errors.As(err, &inputErr) {
		t.Fatalf("expected ErrGetInput, got %v", err)
	}
	if inputErr.Year != 2015 || inputErr.Day != 1 {
		t.Fatalf("unexpected input error payload: %+v", inputErr)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}

func TestRunnerExecuteSolvePart1Error(t *testing.T) {
	wantErr := errors.New("part1 failed")
	reg := registry.NewRegistry()
	registerSolver(t, reg, 2015, 1, stubSolver{part1Err: wantErr})
	r := mustNewRunner(t,
		reg,
		func(year int, day int) ([]string, error) { return []string{"ok"}, nil },
	)

	_, err := r.Execute(2015, 1)
	var solveErr *ErrSolvePart
	if !errors.As(err, &solveErr) {
		t.Fatalf("expected ErrSolvePart, got %v", err)
	}
	if solveErr.Part != 1 {
		t.Fatalf("expected part 1, got %d", solveErr.Part)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}

func TestRunnerExecuteSolvePart2Error(t *testing.T) {
	wantErr := errors.New("part2 failed")
	reg := registry.NewRegistry()
	registerSolver(t, reg, 2015, 1, stubSolver{part1: 10, part2Err: wantErr})
	r := mustNewRunner(t,
		reg,
		func(year int, day int) ([]string, error) { return []string{"ok"}, nil },
	)

	_, err := r.Execute(2015, 1)
	var solveErr *ErrSolvePart
	if !errors.As(err, &solveErr) {
		t.Fatalf("expected ErrSolvePart, got %v", err)
	}
	if solveErr.Part != 2 {
		t.Fatalf("expected part 2, got %d", solveErr.Part)
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected wrapped error %v, got %v", wantErr, err)
	}
}
