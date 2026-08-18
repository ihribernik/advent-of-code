package matrix

import "fmt"

type Matrix struct {
}

func NewBooleanMatrix(x, y int) map[string]bool {
	matrix := map[string]bool{}
	for dx := range x {
		for dy := range y {
			key := fmt.Sprintf("%d,%d", dx, dy)
			matrix[key] = false
		}
	}
	return matrix
}

func NewNumericMatrix(x, y int) map[string]int {
	matrix := map[string]int{}
	for dx := range x {
		for dy := range y {
			key := fmt.Sprintf("%d,%d", dx, dy)
			matrix[key] = 0
		}
	}
	return matrix
}
