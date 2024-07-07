package problem_solving

// AVeryBigSum calculates the sum of all elements in an array
// The function receives a slice of int64
// The function returns an int64 with the sum of all elements in the array
func AVeryBigSum(array []int64) int64 {
	var sum int64
	for _, value := range array {
		sum += value
	}

	return sum
}
