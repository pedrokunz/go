package problem_solving

import "fmt"

// GroupOpenHours Given a list of open hours for a store, create a digest version that groups runs of days with the same hours together.
func GroupOpenHours(openHours []map[string]string) []map[string]string {
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	dayToHoursMap := make(map[string]map[string]string)

	// Initialize all days to closed
	for _, day := range daysOfWeek {
		dayToHoursMap[day] = map[string]string{
			"open":  "",
			"close": "",
		}
	}

	// Fill in the provided open hours
	for _, hours := range openHours {
		day := hours["day"]
		dayToHoursMap[day] = hours
	}

	var result []map[string]string
	currentStartDay := ""
	currentOpenTime := ""
	currentCloseTime := ""

	for i, day := range daysOfWeek {
		openTime := dayToHoursMap[day]["open"]
		closeTime := dayToHoursMap[day]["close"]

		if openTime != currentOpenTime || closeTime != currentCloseTime {
			if currentStartDay != "" {
				if currentStartDay == daysOfWeek[i-1] {
					// Single day range
					result = append(result, map[string]string{
						"days":  currentStartDay,
						"open":  currentOpenTime,
						"close": currentCloseTime,
					})
				} else {
					// Multi-day range
					result = append(result, map[string]string{
						"days":  fmt.Sprintf("%s-%s", currentStartDay, daysOfWeek[i-1]),
						"open":  currentOpenTime,
						"close": currentCloseTime,
					})
				}
			}
			currentStartDay = day
			currentOpenTime = openTime
			currentCloseTime = closeTime
		}
	}

	// Add the last group
	if currentStartDay == "Sunday" {
		// Single day for Sunday
		result = append(result, map[string]string{
			"days":  "Sunday",
			"open":  currentOpenTime,
			"close": currentCloseTime,
		})
	} else {
		// Multi-day ending on Sunday
		result = append(result, map[string]string{
			"days":  fmt.Sprintf("%s-Sunday", currentStartDay),
			"open":  currentOpenTime,
			"close": currentCloseTime,
		})
	}

	return result
}
