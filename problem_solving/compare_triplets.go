package problem_solving

// CompareTriplets compares two triplets and returns the score of each one
// The function receives two slices of int32, each one with three elements
// The function returns a slice of int32 with two elements, the first one is the score of the first triplet
// and the second one is the score of the second triplet
// The score is calculated by comparing the elements of the two triplets, element by element
// If the element of the first triplet is equal to the element of the second triplet,
// no score is incremented
// If the element of the first triplet is greater than the element of the second triplet,
// the score of the first triplet is incremented by one
// If the element of the first triplet is less than the element of the second triplet,
// the score of the second triplet is incremented by one
func CompareTriplets(tripletA []int32, tripletB []int32) []int32 {
	var (
		result         []int32
		aScore, bScore int32
	)

	for i := 0; i < len(tripletA); i++ {
		if tripletA[i] == tripletB[i] {
			continue
		} else if tripletA[i] > tripletB[i] {
			aScore++
		} else if tripletA[i] < tripletB[i] {
			bScore++
		}
	}

	result = append(result, aScore, bScore)

	return result
}
