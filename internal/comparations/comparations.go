package comparations

import "slices"

func HasAny(values, valuesToCompare []string) bool {
	for _, value := range valuesToCompare {
		if slices.Contains(values, value) {
			return true
		}
	}
	return false
}
