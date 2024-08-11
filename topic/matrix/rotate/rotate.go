package rotate

func RotateMatrix(mat [][]int) bool {

	le := len(mat)
	if le == 0 || le != len(mat[0]) {
		return false
	}

	for layer := 0; layer < le/2; layer++ {
		first := layer
		last := le - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			saveTop := mat[first][i]
			mat[first][i] = mat[last-offset][first]
			mat[last-offset][first] = mat[last][last-offset]
			mat[last][last-offset] = mat[i][last]
			mat[i][last] = saveTop
		}
	}

	return true
}

func RotateMatrixByFlip(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

		for j, k := 0, m-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
