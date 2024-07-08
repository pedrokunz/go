package problem_solving

import "time"

// TimeConversion converts a time in 12-hour format to 24-hour format
// The function receives a string with the time in 12-hour format (hh:mm:ssAM or hh:mm:ssPM),
// and returns a string with the time in 24-hour format
func TimeConversion(s string) string {
	timeParsed, _ := time.Parse("03:04:05PM", s)
	return timeParsed.Format("15:04:05")
}
