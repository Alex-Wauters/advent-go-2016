package main

func leastCommonChar(chars []rune, position int, c chan<- PuzzlePiece) {
	count := make(map[rune]int)
	for _, char := range chars {
		count[char]++
	}
	minCount := 999
	var minChar rune
	for k, v := range count {
		if v < minCount {
			minCount, minChar = v, k
		}
	}
	c <- PuzzlePiece{position, minChar}
}
