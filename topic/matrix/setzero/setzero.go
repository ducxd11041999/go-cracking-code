package setzero

import "fmt"

func SetZero(mat [][]int) {
	m := len(mat)
	if m == 0 {
		return
	}
	n := len(mat[0])
	rows := make([]bool, m)
	cols := make([]bool, n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if mat[r][c] == 0 {
				rows[r] = true
				cols[c] = true
			}
		}
	}

	fmt.Println(rows)
	fmt.Println(cols)

	for i := 0; i < m; i++ {
		if rows[i] {
			nullifyRows(mat, i)
		}
	}

	for i := 0; i < n; i++ {
		if cols[i] {
			fmt.Println(i)
			nullifyCols(mat, i)
		}
	}

}

func nullifyRows(mat [][]int, row int) {
	for j := 0; j < len(mat[0]); j++ {
		mat[row][j] = 0
	}
}

func nullifyCols(mat [][]int, col int) {
	for j := 0; j < len(mat); j++ {
		mat[j][col] = 0
	}
}
