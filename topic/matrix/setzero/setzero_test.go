package setzero

import (
	"fmt"
	"testing"
)

func TestSetZero(t *testing.T) {
	mat := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
		[]int{9, 8, 0},
	}

	SetZero(mat)
	fmt.Println(mat)
}
