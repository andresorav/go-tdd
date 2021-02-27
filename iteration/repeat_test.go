package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	res := Repeat("a", 5)
	expected := "aaaaa"

	assert.Equal(t, expected, res)

	res = Repeat("word", 1)
	expected = "word"

	assert.Equal(t, expected, res)
}

func ExampleRepeat() {
	fmt.Println(Repeat("cool", 3))
	// Output: coolcoolcool
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("word", 5)
	}
}
