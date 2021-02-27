package arrays_slices

// SumAll returns new numbers slice containing the totals for each numbers slice
func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}
	return sums
}
