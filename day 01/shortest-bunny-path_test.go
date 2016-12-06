package main

import (
	"testing"
	"reflect"
)

var parseTests = []struct {
	input  string
	result []instruction
}{
	{"L1", []instruction{{"L", 1}}},
	{"R23, L143", []instruction{{"R", 23}, {"L",143}}},
}

func TestParser(t *testing.T) {
	for _, pair := range parseTests {
		t.Run(pair.input, func(t *testing.T) {
			parsing := parseInstructions(pair.input)
			if !reflect.DeepEqual(parsing, pair.result) {
				t.Error("For", pair.input,
				"expected", pair.result,
				"got", parsing)
			}
		})
	}
}
