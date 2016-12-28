package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer track(time.Now(), "Part 1 and 2")
	instructions, reverseInstr := toInstructions()
	input := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for _, instruction := range instructions { //part 1
		input = instruction(input)
	}
	fmt.Println(strings.Join(input, ""))
	cipher := []string{"f", "b", "g", "d", "c", "e", "a", "h"}
	for i := len(reverseInstr) - 1; i >= 0; i-- {
		cipher = reverseInstr[i](cipher)
	}
	fmt.Println(strings.Join(cipher, ""))
}

func move(word []string, a, b int) []string {
	letter := word[a]
	result := make([]string, len(word))
	cut := append(word[:a], word[a+1:]...)
	incr := 0
	for i := range result {
		if i == b {
			result[i] = letter
		} else {
			result[i] = cut[incr]
			incr++
		}
	}
	return result
}
func swapPosition(word []string, a, b int) []string {
	word[a], word[b] = word[b], word[a]
	return word
}
func swapLetters(word []string, a, b string) []string {
	for i, letter := range word {
		if letter == a {
			word[i] = b
		} else if letter == b {
			word[i] = a
		}
	}
	return word
}
func reverse(word []string, start, end int) []string {
	b := make([]string, len(word))
	remaining := 0
	for i := range b {
		if i == start {
			remaining = end - start + 1
		}
		if remaining > 0 {
			b[i] = word[start+remaining-1]
			remaining--
		} else {
			b[i] = word[i]
		}
	}
	return b
}
func rotate(word []string, steps int) []string {
	b := make([]string, len(word))
	for i := range word {
		pos := (i + steps) % len(word)
		if pos < 0 {
			pos = len(word) + pos
		}
		b[pos] = word[i]
	}
	return b
}
func rotatePos(word []string, indicator string) []string {
	index := 0
	for i, letter := range word {
		if letter == indicator {
			index = i
			break
		}
	}
	amount := 1 + index
	if index >= 4 {
		amount++
	}
	return rotate(word, amount)
}

type instruction func([]string) []string

func toInstructions() (instructions []instruction, reverseInstr []instruction) {
	swapLettersRE := regexp.MustCompile(`swap letter (\w) with letter (\w)`)
	moveRE := regexp.MustCompile(`move position (\d) to position (\d)`)
	reverseRE := regexp.MustCompile(`reverse positions (\d) through (\d)`)
	rotateRE := regexp.MustCompile(`rotate (left|right) (\d) step[s]*`)
	rotatePosRE := regexp.MustCompile(`rotate based on position of letter (\w)`)
	swapPositionRE := regexp.MustCompile(`swap position (\d) with position (\d)`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	toInstruction := func(line string) func([]string) []string {
		cmd := swapLettersRE.FindStringSubmatch(line)
		if len(cmd) > 0 {
			return func(w []string) []string { return swapLetters(w, cmd[1], cmd[2]) }
		} else if cmd = moveRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return move(w, toInt(cmd[1]), toInt(cmd[2])) }
		} else if cmd = reverseRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return reverse(w, toInt(cmd[1]), toInt(cmd[2])) }
		} else if cmd = rotateRE.FindStringSubmatch(line); len(cmd) > 0 {
			steps := toInt(cmd[2])
			if cmd[1] == "left" {
				steps = -steps
			}
			return func(w []string) []string { return rotate(w, steps) }
		} else if cmd = rotatePosRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return rotatePos(w, cmd[1]) }
		} else if cmd = swapPositionRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return swapPosition(w, toInt(cmd[1]), toInt(cmd[2])) }
		} else {
			panic("Could not parse: " + line)
		}
	}
	toReverseInstruction := func(line string) instruction {
		cmd := swapLettersRE.FindStringSubmatch(line)
		if len(cmd) > 0 {
			return func(w []string) []string { return swapLetters(w, cmd[2], cmd[1]) }
		} else if cmd = moveRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return move(w, toInt(cmd[2]), toInt(cmd[1])) }
		} else if cmd = reverseRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return reverse(w, toInt(cmd[1]), toInt(cmd[2])) }
		} else if cmd = rotateRE.FindStringSubmatch(line); len(cmd) > 0 {
			steps := toInt(cmd[2])
			if cmd[1] == "left" {
				steps = -steps
			}
			return func(w []string) []string { return rotate(w, -steps) }
		} else if cmd = rotatePosRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string {
				return findReverseRotation(w, func(w []string) []string { return rotatePos(w, cmd[1]) })
			}
		} else if cmd = swapPositionRE.FindStringSubmatch(line); len(cmd) > 0 {
			return func(w []string) []string { return swapPosition(w, toInt(cmd[2]), toInt(cmd[1])) }
		} else {
			panic("Could not parse: " + line)
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, toInstruction(line))
		reverseInstr = append(reverseInstr, toReverseInstruction(line))
	}
	return
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func findReverseRotation(w []string, reverse instruction) []string { //Part 2 addition
	plainText := make([]string, len(w))
	copy(plainText, w)
	cipher := strings.Join(w, "")
	for i := 0; i < len(w); i++ {
		rotated := rotate(plainText, i)
		if strings.Join(reverse(rotated), "") == cipher {
			return rotated
		}
	}
	panic(fmt.Sprintf("Could not find cipher for plaintext: %v", plainText))
}

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
