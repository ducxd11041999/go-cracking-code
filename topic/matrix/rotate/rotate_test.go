package rotate

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	RotateMatrix(mat)
	fmt.Println(mat)
}

func TestRotateByFlip(t *testing.T) {
	mat := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}

	RotateMatrixByFlip(mat)
	fmt.Println(mat)
}
