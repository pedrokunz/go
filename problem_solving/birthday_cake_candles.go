package problem_solving

// BirthdayCakeCandles calculates the number of candles that can be blown out
// The function receives a slice of int32 with the heights of the candles,
// and returns an int32 with the number of candles that can be blown out
// The candles can only be blown out if they have the maximum height
func BirthdayCakeCandles(candles []int32) int32 {
	var maxHeight, count int32

	for _, candle := range candles {
		if candle > maxHeight {
			maxHeight = candle
			count = 1
		} else if candle == maxHeight {
			count++
		}
	}

	return count
}
