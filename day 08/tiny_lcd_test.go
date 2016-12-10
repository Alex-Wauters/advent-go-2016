package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		row      []bool
		amount   int
		expected []bool
	}{
		{[]bool{false, false, true, false}, 1, []bool{false, false, false, true}},
		{[]bool{false, true, true, true}, 2, []bool{true, true, false, true}},
		{[]bool{false, true, true, false, true}, 3, []bool{true, false, true, false, true}},
	}
	for _, test := range tests {
		rotateRow(test.row, test.amount)
		if !reflect.DeepEqual(test.row, test.expected) {
			t.Error("For row", test.row, "with rotate amount:", test.amount, test.amount, "Expected", test.expected)
		}
	}

}

func TestFill(t *testing.T) {
	lcd := matrix(wide, tall)
	fill(lcd, 3, 2)
}

func TestPt1(t *testing.T) {
	lcd := matrix(7, 3)
	fill(lcd, 3, 2)
	if !(lcd[0][0] && lcd[0][2]) {
		t.Error("Was not filled up right")
	}
	rotateCol(lcd, 1, 1)
	rotateRow(lcd[0], 1)
	fmt.Println(lcd)
}
