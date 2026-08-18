package container

import "github.com/ihribernik/aoc-cli/internal/run"

type Container interface {
	GetRunner() run.Runner
}
