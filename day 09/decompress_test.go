package main

import "testing"

func TestDecompress(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		len      int
	}{
		{"ADVENT", "ADVENT", 6},
		{"A(1x5)BC", "ABBBBBC", 7},
		{"(3x3)XYZ", "XYZXYZXYZ", 9},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG", 11},
		{"(6x1)(1x3)A", "(1x3)A", 6},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY", 18},
		{"A(1x10)BC", "ABBBBBBBBBBC", 12},
	}
	for _, test := range tests {
		if dec := decompress(test.input); dec != test.expected || len(dec) != test.len {
			t.Error("For", test.input, "expected", test.expected, "got", dec)
		}
	}
}

func TestDecompress2(t *testing.T) {
	tests := []struct {
		input string
		len   int
	}{
		{"(3x3)XYZ", 9},
		{"X(8x2)(3x3)ABCY", 20},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}
	for _, test := range tests {
		if dec := decompress2(test.input); dec != uint64(test.len) {
			t.Error("For", test.input, "expected", test.len, "got", dec)
		}
	}
}
