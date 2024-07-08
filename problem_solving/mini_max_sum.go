package problem_solving

// MiniMaxSum calculates the minimum and maximum values that can be obtained by summing the lesser and greater values of an array
// The function receives a slice of int32 with n integers, and returns two integers, the minimum and maximum values that can be obtained
func MiniMaxSum(array []int32) (int32, int32) {
	var sum int32
	minimumValue := array[0]
	maximumValue := array[0]

	for _, value := range array {
		if value < minimumValue {
			minimumValue = value
		}

		if value > maximumValue {
			maximumValue = value
		}

		sum += value
	}

	minimumSum := sum - maximumValue
	maximumSum := sum - minimumValue

	return minimumSum, maximumSum
}
