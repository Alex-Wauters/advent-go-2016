package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	pwdLength = 8
)

type Password struct {
	letters []string
	count   int
}

func (p *Password) IsComplete() bool {
	return p.count == pwdLength
}

func (p *Password) Assign(piece PuzzlePiece) {
	if p.letters[piece.position] == "" {
		p.letters[piece.position] = piece.char
		p.count++
	}
}

func part2(input string) {
	c := make(chan PuzzlePiece)
	quit := make(chan bool)
	go dispatchPt2(input, c, quit)
	password := Password{make([]string, 8, 8), 0}
	for !password.IsComplete() {
		candidate := <-c
		password.Assign(candidate)
	}
	fmt.Printf("The password is %v", strings.Join(password.letters, ""))
	close(quit)
}

func dispatchPt2(prefix string, c chan PuzzlePiece, quit <-chan bool) {
	i := 0
	routines := 20
	for ; i < routines; i++ {
		go findGoodHashPt2(prefix, i, routines, c, quit)
	}
}

type PuzzlePiece struct {
	position int
	char     string
}

func findGoodHashPt2(prefix string, i, increment int, c chan<- PuzzlePiece, quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			clue := hash(prefix, i)
			if strings.HasPrefix(clue, "00000") {
				position, err := strconv.Atoi(string(clue[5]))
				if err == nil && position < pwdLength {
					c <- PuzzlePiece{position, string(clue[6])}
				}
			}
			i += increment
		}
	}
}
