package problem_solving

func AppleAndOrange(
	startPoint,
	endPoint,
	appleTreeLocation,
	orangeTreeLocation int32,
	apples,
	oranges []int32,
) (int, int) {
	var appleCount, orangeCount int

	for _, apple := range apples {
		if appleTreeLocation+apple >= startPoint &&
			appleTreeLocation+apple <= endPoint {
			appleCount++
		}
	}

	for _, orange := range oranges {
		if orangeTreeLocation+orange >= startPoint &&
			orangeTreeLocation+orange <= endPoint {
			orangeCount++
		}
	}

	return appleCount, orangeCount
}
