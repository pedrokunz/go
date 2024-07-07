package problem_solving

import "fmt"

// PlusMinus calculates the fractions of its elements that are positive, negative, and are zeros
// The function receives a slice of int32, and prints the ratio of positive, negative and zero values in the array
// with 6 decimal places
func PlusMinus(array []int32) string {
	var (
		positive, negative, zero int
	)

	for _, value := range array {
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}

	format := "%.6f\n%.6f\n%.6f\n"
	total := len(array)
	return fmt.Sprintf(
		format,
		float64(positive)/float64(total),
		float64(negative)/float64(total),
		float64(zero)/float64(total))
}
