package helloworld

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Parallel()

	t.Run("says Hello to default name", func(t *testing.T) {
		t.Parallel()

		res := Hello("")
		expected := "Hello, world"

		assert.Equal(t, expected, res)
	})

	t.Run("says Hello to defined name", func(t *testing.T) {
		t.Parallel()

		res := Hello("John")
		expected := "Hello, John"

		assert.Equal(t, expected, res)
	})
}

func ExampleHello() {
	fmt.Println(Hello("world"))
	// Output: Hello, world
}
