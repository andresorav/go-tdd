package arrays_slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumAll(t *testing.T) {
	t.Parallel()

	t.Run("sum all slices", func(t *testing.T) {
		t.Parallel()

		res := SumAll([]int{1, 6}, []int{7, 1}, []int{5, 4})
		expected := []int{7, 8, 9}

		assert.Equal(t, expected, res)
	})

	t.Run("sum with empty slices", func(t *testing.T) {
		t.Parallel()

		res := SumAll([]int{}, []int{1, 2}, []int{})
		expected := []int{0, 3, 0}

		assert.Equal(t, expected, res)
	})
}

func ExampleSumAll() {
	fmt.Println(SumAll([]int{1, 2}, []int{5, 6}))
	// Output: [3 11]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 7})
	}
}
