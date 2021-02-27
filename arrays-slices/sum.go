package arrays_slices

// Sum returns the sum of all numbers in the array
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
