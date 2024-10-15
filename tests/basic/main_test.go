package basic

import (
	"testing"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(input)
	if actual != output {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}

	// assert.Equal(t, AddOne(1), 2, "AddOne(1) should return 2")
}
