package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readInput()
	decodeInput(input, mostCommonChar)  //Part 1
	decodeInput(input, leastCommonChar) //Part 2
}

func decodeInput(input [][]rune, decodeFn func([]rune, int, chan<- PuzzlePiece)) {
	c := make(chan PuzzlePiece)
	message := Message{0, len(input), make([]string, len(input))}
	for i, letters := range input {
		go decodeFn(letters, i, c)
	}
	for !message.IsComplete() {
		piece := <-c
		message.Assign(piece)
	}
	fmt.Printf("The message is %v", strings.Join(message.letters, ""))
}

type Message struct {
	curLength, maxLength int
	letters              []string
}

func (m *Message) Assign(p PuzzlePiece) {
	m.letters[p.position] = string(p.char)
	m.curLength++
}

func (m *Message) IsComplete() bool {
	return m.curLength == m.maxLength
}

type PuzzlePiece struct {
	position int
	char     rune
}

func mostCommonChar(chars []rune, position int, c chan<- PuzzlePiece) {
	count := make(map[rune]int)
	for _, char := range chars {
		count[char]++
	}
	var maxCount int
	var maxChar rune
	for k, v := range count {
		if v > maxCount {
			maxCount, maxChar = v, k
		}
	}
	c <- PuzzlePiece{position, maxChar}
}

func readInput() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result [][]rune
	initialized := false
	for scanner.Scan() {
		t := scanner.Text()
		if !initialized {
			result = setup(len(t))
			initialized = true
		}
		for i, char := range t {
			result[i] = append(result[i], char)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func setup(length int) [][]rune {
	result := make([][]rune, length)
	for i := range result {
		result[i] = make([]rune, 0)
	}
	return result
}
