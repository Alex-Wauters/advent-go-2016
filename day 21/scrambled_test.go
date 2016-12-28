package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestScramble(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e"}
	instructions := []func([]string) []string{
		func(w []string) []string { return swapPosition(w, 4, 0) },
		func(w []string) []string { return swapLetters(w, "d", "b") },
		func(w []string) []string { return reverse(w, 0, 4) },
		func(w []string) []string { return rotate(w, -1) },
		func(w []string) []string { return move(w, 1, 4) },
		func(w []string) []string { return move(w, 3, 0) },
		func(w []string) []string { return rotatePos(w, "b") },
		func(w []string) []string { return rotatePos(w, "d") },
	}
	for _, instruction := range instructions {
		input = instruction(input)
		fmt.Println(input)
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input      []string
		start, end int
		expected   string
	}{
		{[]string{"a", "b", "c", "d", "e"}, 0, 4, "edcba"},
		{[]string{"a", "b", "c", "d", "e"}, 0, 3, "dcbae"},
		{[]string{"a", "b", "c", "d", "e"}, 1, 3, "adcbe"},
	}
	for _, test := range tests {
		result := strings.Join(reverse(test.input, test.start, test.end), "")
		if result != test.expected {
			t.Error("Expected", test.expected, "got", result)
		}
	}
}

func TestMovePosition(t *testing.T) {
	tests := []struct {
		input      []string
		start, end int
		expected   string
	}{
		{[]string{"a", "b", "c", "d", "e"}, 0, 4, "bcdea"},
		{[]string{"a", "b", "c", "d", "e"}, 2, 3, "abdce"},
		{[]string{"a", "b", "c", "d", "e"}, 1, 3, "acdbe"},
		{[]string{"a", "b", "c", "d", "e"}, 4, 0, "eabcd"},
		{[]string{"a", "b", "c", "d", "e"}, 3, 2, "abdce"},
	}
	for _, test := range tests {
		result := strings.Join(move(test.input, test.start, test.end), "")
		if result != test.expected {
			t.Error("Expected", test.expected, "got", result)
		}
	}
}

func TestRotatePos(t *testing.T) {
	tests := []struct {
		input     []string
		indicator string
		expected  string
	}{
		{[]string{"a", "b", "c", "d", "e"}, "e", "eabcd"},
		{[]string{"a", "b", "c", "d", "e"}, "a", "eabcd"},
		{[]string{"a", "b", "c", "d", "e"}, "b", "deabc"},
	}
	for _, test := range tests {
		result := strings.Join(rotatePos(test.input, test.indicator), "")
		if result != test.expected {
			t.Error("Expected", test.expected, "got", result)
		}
	}
}
