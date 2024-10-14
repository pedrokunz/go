package problem_solving

/*
	Rotating arrays is a common operation in various programming and algorithmic contexts. Here are some common use cases:
	Data Shuffling:
		Rotating arrays can be used to shuffle data in a controlled manner, which is useful in simulations and games.
	Circular Buffers:
		In systems programming, circular buffers (or ring buffers) use array rotation to manage buffer overflows and underflows efficiently.
	Image Processing:
		Rotating pixel arrays is a fundamental operation in image processing for tasks like rotating images or applying certain filters.
	Cryptography:
		Some encryption algorithms use array rotation as part of their transformation steps to obfuscate data.
	Load Balancing:
		In distributed systems, rotating arrays can help in load balancing tasks by cyclically distributing tasks or resources.
	Scheduling Algorithms:
		Rotating arrays can be used in round-robin scheduling algorithms to manage tasks or processes in a cyclic order.
	String Manipulation:
		Rotating character arrays (strings) can be used in various string manipulation tasks, such as generating cyclic permutations of a string.
	Data Analysis:
		In data analysis, rotating arrays can help in creating lagged versions of time series data for correlation and regression analysis.
	Game Development:
		In game development, rotating arrays can be used to manage game states, player turns, or cyclic events.
	Algorithm Optimization:
		Some algorithms, like certain sorting or searching algorithms, can be optimized using array rotation techniques to reduce time complexity.
*/

// RotateArray rotates an array to the left or right by a given number of rotations.
// Example: RotateArray([1, 2, 3, 4, 5], 2, "left") => [3, 4, 5, 1, 2]
// Example: RotateArray([1, 2, 3, 4, 5], -2, "left") => [4, 5, 1, 2, 3]
// Example: RotateArray([1, 2, 3, 4, 5], 2, "right") => [4, 5, 1, 2, 3]
// Example: RotateArray([1, 2, 3, 4, 5], -2, "right") => [3, 4, 5, 1, 2]
// Time complexity: O(n) where n is the size of the array
func RotateArray(array []int, rotations int, direction string) []int {
	if direction != "right" && direction != "left" {
		return array
	}

	// Get the size of the array O(1)
	arraySize := len(array)
	// Create a new array with the same size as the original array O(n)
	rotatedArray := make([]int, arraySize)

	rotations = normalizeRotations(rotations, arraySize, direction)

	// Iterate over the original array O(n)
	for index := 0; index < arraySize; index++ {
		// Calculate the new index after rotating the array
		// Example:
		// (0 + 2) % 5 = 2 => 2 % 5 => 2
		// (1 + 2) % 5 = 3 => 3 % 5 => 3
		newIndex := (index + rotations) % arraySize

		// Assign the value of the original array to the new array O(1)
		rotatedArray[index] = array[newIndex]
	}

	return rotatedArray
}

// normalizeRotations normalizes the number of rotations to be within the bounds of the array size.
// Examples:
// normalizeRotations(-6, 5, "left") => 4
// normalizeRotations(-2, 5, "left") => 3
// normalizeRotations(6, 5, "left") => 1
// normalizeRotations(6, 5, "right") => 4
// Time complexity: O(1)
func normalizeRotations(rotations int, arraySize int, direction string) int {
	// Handle negative rotations
	if rotations < 0 {
		// Convert negative rotations to positive rotations
		// Example:
		// -6 % 5 = -1 => -1 + 5 = 4 => 4 % 5 => 4
		// -2 % 5 = -2 => -2 + 5 = 3 => 3 % 5 => 3
		rotations = (rotations%arraySize + arraySize) % arraySize
		direction = switchRotateArrayDirection(direction)
	}

	// Normalize rotations to be within the bounds of the array size
	rotations = rotations % arraySize

	if direction == "right" {
		// Convert right rotations to equivalent left rotations
		rotations = arraySize - rotations
	}
	return rotations
}

func switchRotateArrayDirection(direction string) string {
	if direction == "left" {
		return "right"
	}
	return "left"
}
