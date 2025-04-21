package y2015

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

type Day04 struct{}

// SolvePart1 implements solutions.Solver.
func (d Day04) SolvePart1(input []string) (int, error) {
	for i := range math.MaxInt32 {
		currentInput := fmt.Sprintf("%s%v", input[0], i)
		hashedString := fmt.Sprintf("%x", md5.Sum([]byte(currentInput)))
		if strings.HasPrefix(hashedString, "00000") {
			return i, nil
		}
	}
	return 0, fmt.Errorf("cannot guess the name")
}

// SolvePart2 implements solutions.Solver.
func (d Day04) SolvePart2(input []string) (int, error) {
	for i := range math.MaxInt32 {
		currentInput := fmt.Sprintf("%s%v", input[0], i)
		hashedString := fmt.Sprintf("%x", md5.Sum([]byte(currentInput)))
		if strings.HasPrefix(hashedString, "000000") {
			return i, nil
		}
	}
	return 0, fmt.Errorf("cannot guess the name")
}

func init() {
	solutions.Register(2015, 4, Day04{})
}
