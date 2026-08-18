package container

import (
	"github.com/ihribernik/aoc-cli/internal/inputs"
	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/run"
	"github.com/ihribernik/aoc-cli/internal/solutions"
)

type container struct {
	Runner run.Runner
}

func New() (Container, error) {
	return newContainer(solutions.RegisterAll, run.NewRunner)
}

func newContainer(registerAll func(registry.Registry) error, newRunner func(run.SolverResolver, run.LoadInputFunc) (run.Runner, error)) (Container, error) {
	reg := registry.NewRegistry()
	if err := registerAll(reg); err != nil {
		return nil, err
	}

	runner, err := newRunner(reg, inputs.GetInput)
	if err != nil {
		return nil, err
	}

	return &container{Runner: runner}, nil
}

func (c *container) GetRunner() run.Runner {
	return c.Runner
}
