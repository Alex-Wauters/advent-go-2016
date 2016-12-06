package main

import (
	"testing"
)

var triangleTests = []struct {
	input triangle
	result bool
}{
	{triangle{5,10,25}, false},
	{triangle{5,10,2}, false},
	{triangle{2,6,5}, true},

}

func TestParser(t *testing.T) {
	for _, pair := range triangleTests {
		t.Run("Triangle test", func(t *testing.T) {
			possible := pair.input.possible()
			if pair.result != possible {
				t.Error("For", pair.input,
					"expected", pair.result,
					"got", possible)
			}
		})
	}
}
