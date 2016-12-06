package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSectorHash(t *testing.T) {
	tests := []struct {
		input  string
		sector int
		hash   string
	}{
		{"395[idjmx]", 395, "idjmx"},
	}

	for _, test := range tests {
		es, eh := sectorAndHash(test.input)
		if es != test.sector || eh != test.hash {
			t.Error("For", test.input,
				"expected", fmt.Sprintf("%v %s", test.sector, test.hash),
				"got", fmt.Sprintf("%v %s", es, eh))
		}
	}
}

func TestToRoom(t *testing.T) {
	line := "aaaaa-bbb-z-y-x-123[abxyz]"
	room := toRoom(line)
	expected := Room{123, "abxyz", "abxyz", "aaaaa bbb z y x"}
	if !reflect.DeepEqual(room, expected) {
		t.Error("For", line,
			"expected", expected,
			"got", room)
	}

	if !room.IsCorrect() {
		t.Error("Room should be correct")
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		room     Room
		expected string
	}{
		{Room{1, "abcde", "abcde", "abcde"}, "bcdef"},
		{Room{2, "","", "xyz b"}, "zab d"},
		{Room{26, "","", "wut sup dawg"}, "wut sup dawg"},
		{Room{27, "","", "derp a"}, "efsq b"},
	}
	fmt.Print(string(rune(26)))
	for _, test := range tests {
		decrypted := test.room.Decrypt()
		if decrypted != test.expected {
			t.Error("Expected", test.expected,
				"received", decrypted)
		}
	}
}
