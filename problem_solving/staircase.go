package problem_solving

import "strings"

// Staircase prints a staircase of a specific size
// The function receives an int32 with the size of the staircase, and prints it
// The staircase is right-aligned, composed of spaces and hashes (#)
func Staircase(size int32) string {
	var result strings.Builder
	for i := 1; i <= int(size); i++ {
		result.WriteString(strings.Repeat(" ", int(size)-i))
		result.WriteString(strings.Repeat("#", i))
		result.WriteString("\n")
	}

	return result.String()
}
