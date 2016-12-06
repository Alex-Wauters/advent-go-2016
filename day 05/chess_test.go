package main

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	result := hash("abc", 3231929)
	if result != "00000155f8105dff7f56ee10fa9b9abd" {
		t.Error("Wrong hash")
	}
	c, done := make(chan string), make(chan bool)
	go findGoodHash("abc", 3231929, 1, c, done)
	letter := <-c
	close(done)
	if letter != "1" {
		t.Error("Did not receive correct letter from hash")
	}
}

func TestHashPt2(t *testing.T) {
	c, done := make(chan PuzzlePiece), make(chan bool)
	go findGoodHashPt2("abc", 3231929, 1, c, done)
	piece := <-c
	close(done)
	if !reflect.DeepEqual(piece, PuzzlePiece{1, "5"}) {
		t.Error("Did not receive correct letter from hash")
	}
}
