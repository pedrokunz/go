package problem_solving

// DiagonalDifference calculates the absolute difference between the sums of the diagonals of a square matrix
// The function receives a slice of slices of int32, representing a square matrix,
// and returns an int32 with the absolute difference.
// The left diagonal is the diagonal that goes from the top left to the bottom right of the matrix.
// The right diagonal is the diagonal that goes from the bottom left to the top right of the matrix.
func DiagonalDifference(matrix [][]int32) int32 {
	var (
		leftDiagonalSum, rightDiagonalSum int32
	)

	for i := 0; i < len(matrix); i++ {
		leftDiagonalSum += matrix[i][i]                // left diagonal -> [0][0], [1][1], [2][2]... -> top left to bottom right
		rightDiagonalSum += matrix[i][len(matrix)-1-i] // right diagonal -> [0][2], [1][1], [2][0]... -> bottom left to top right
	}

	return leftDiagonalSum - rightDiagonalSum
}
