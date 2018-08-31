package math

// returns det of matrix if specified algebra
func Det(matrix [][]byte, alg ByteAlgebra) byte {
	size := len(matrix)
	if size == 1 {
		return matrix[0][0]
	}
	if size == 2 {
		return det2(matrix, alg)
	}
	var det byte = 0
	for i := 0; i < size; i++ {
		element := matrix[0][i]
		minor := Det(cross(matrix, 0, i), alg)
		if i%2 == 0 {
			det = alg.Add(det, alg.Mul(element, minor))
		} else {
			det = alg.Sub(det, alg.Mul(element, minor))
		}
	}
	return det
}

// returns matrix with size n-1 crossing row and column specified
func cross(matrix [][]byte, row, col int) [][]byte {
	size := len(matrix)
	result := make([][]byte, size-1)

	var i, j, a, b int
	for i, a = 0, 0; i < size; i++ {
		if i == row {
			continue
		}
		result[a] = make([]byte, size-1)
		for j, b = 0, 0; j < size; j++ {
			if j == col {
				continue
			}
			result[a][b] = matrix[i][j]
			b++
		}
		a++
	}
	return result
}

// returns det of 2x2 matrix
func det2(matrix [][]byte, alg ByteAlgebra) byte {
	return alg.Sub(alg.Mul(matrix[0][0], matrix[1][1]), alg.Mul(matrix[1][0], matrix[0][1]))
}