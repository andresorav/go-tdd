package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	res := Add(1, 2)
	expected := 3

	assert.Equal(t, expected, res)
}

func ExampleAdd() {
	fmt.Println(Add(5, 6))
	// Output: 11
}
