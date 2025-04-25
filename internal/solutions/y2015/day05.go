package y2015

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ihribernik/advent-of-code/internal/solutions"
)

type Day05 struct{}

// SolvePart1 implements solutions.Solver.
func (d Day05) SolvePart1(input []string) (solutions.Solution, error) {
	vowels := "aeiou"
	forgotedStringRegexp := regexp.MustCompile("(ab|cd|pq|xy)")

	result := 0

	for _, word := range input {
		threeVowels := 0
		twice := false

		for i, letter := range word {
			if strings.ContainsRune(vowels, letter) {
				threeVowels += 1
			}
			if i > 0 {
				if letter == rune(word[i-1]) {
					twice = true
				}
			}
		}

		if threeVowels >= 3 && twice && !forgotedStringRegexp.MatchString(word) {
			result += 1
		}

	}
	return solutions.Solution{Result: result}, nil

}

// --- Part Two ---
// Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice.
// None of the old rules apply, as they are all clearly ridiculous.

// Now, a nice string is one with all of the following properties:

// It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
// It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
// For example:

//qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
// xx yx x is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
// uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
// ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.
// How many strings are nice under these new rules?

// SolvePart2 implements solutions.Solver.
func (d Day05) SolvePart2(input []string) (solutions.Solution, error) {

	result := 0
	for _, word := range input {
		twice := false
		repeats := false
		for i, letter := range word {
			if i+1 >= len(word) {
				continue
			}
			letters := fmt.Sprintf("%v%v", string(letter), string(word[i+1]))
			lettersCount := strings.Count(word, letters)
			if lettersCount >= 2 {
				twice = true
			}
			if i+2 >= len(word) {
				continue
			}
			if letter == rune(word[i+2]) {
				repeats = true
			}
		}

		if twice && repeats {
			result += 1
		}
	}

	return solutions.Solution{Result: result}, nil
}

func init() {
	solutions.Register(2015, 05, Day05{})
}
