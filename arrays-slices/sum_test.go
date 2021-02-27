package arrays_slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Sum(numbers)
	expected := 45

	assert.Equal(t, expected, res)
}

func ExampleSum() {
	numbers := []int{3, 5, 8}
	fmt.Println(Sum(numbers))
	// Output: 16
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
