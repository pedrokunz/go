package problem_solving

// SimpleArraySum calculates the sum of all elements in an array
// The function receives a slice of int32
// The function returns an int32 with the sum of all elements in the array
func SimpleArraySum(array []int32) int32 {
	var sum int32
	for _, value := range array {
		sum += value
	}

	return sum
}
