package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// signals.go and signals_pt2.go are written quite verbosely, with multi-threading
// An attempt to minimize it
func main() {
	lines := readInput()
	var counts = make([]map[rune]int, len(lines[0]))
	for i := range counts {
		counts[i] = make(map[rune]int)
	}
	for _, line := range lines {
		for i, char := range line {
			counts[i][char]++
		}
	}
	password1, password2 := make([]string, len(lines[0])), make([]string, len(lines[0]))
	for i, col := range counts {
		maxCount, minCount := 0, 999
		var maxChar, minChar rune
		for k, v := range col {
			if v > maxCount {
				maxCount, maxChar = v, k
			}
			if v < minCount {
				minCount, minChar = v, k
			}
		}
		password1[i], password2[i] = string(maxChar), string(minChar)
	}
	fmt.Printf("Password1 is %v and password2 is %v", strings.Join(password1, ""), strings.Join(password2, ""))
}

func readInput() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
