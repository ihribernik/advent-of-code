package y2015

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/ihribernik/aoc-cli/internal/puzzle"
)

type Day04 struct{}

// SolvePart1 implements solutions.Solver.
func (d Day04) SolvePart1(input []string) (int, error) {
	if len(input) == 0 {
		return 0, newInputError(4, 0, "", errors.New("missing secret key"))
	}

	answer, err := findHashWithPrefix(input[0], "00000", math.MaxInt32)
	if err != nil {
		return 0, fmt.Errorf("2015 day 04 part 1: %w", err)
	}
	return answer, nil
}

// SolvePart2 implements solutions.Solver.
func (d Day04) SolvePart2(input []string) (int, error) {
	if len(input) == 0 {
		return 0, newInputError(4, 0, "", errors.New("missing secret key"))
	}

	answer, err := findHashWithPrefix(input[0], "000000", math.MaxInt32)
	if err != nil {
		return 0, fmt.Errorf("2015 day 04 part 2: %w", err)
	}
	return answer, nil
}

func findHashWithPrefix(secret string, prefix string, limit int) (int, error) {
	for i := range limit {
		currentInput := fmt.Sprintf("%s%v", secret, i)
		hashedString := fmt.Sprintf("%x", md5.Sum([]byte(currentInput)))
		if strings.HasPrefix(hashedString, prefix) {
			return i, nil
		}
	}
	return 0, puzzle.ErrAnswerNotFound
}
