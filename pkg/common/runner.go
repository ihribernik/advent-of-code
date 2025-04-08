package common

import (
	"fmt"
	"os"

	"path/filepath"
)

type Runner struct {
	Year int
	Day  int
}

func (runner Runner) Parse() (string, error) {
	filepath := filepath.Join("input", fmt.Sprintf("%d", runner.Year), fmt.Sprintf("day%02d.txt", runner.Day))
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
