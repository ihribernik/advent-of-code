package common

import (
	"fmt"
	"os"
	"strings"

	"path/filepath"
)

type Runner struct {
	Year int
	Day  int
}

func (runner Runner) Input() ([]string, error) {
	filepath := filepath.Join("input", fmt.Sprintf("%d", runner.Year), fmt.Sprintf("day%02d.txt", runner.Day))
	data, err := os.ReadFile(filepath)
	if err != nil {
		return []string{}, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	lineas := strings.Split(content, "\n")

	return lineas, nil
}
