package main

import (
	"bufio"
	"log"
	"os"
)

func partTwo() {
	triangles := readInputVertically()
	countPossible(triangles)
}

func readInputVertically() []triangle {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	result := make([]triangle, 0)

	scanner := bufio.NewScanner(file)
	lines := [3]triangle{}
	counter := 0
	for scanner.Scan() {
		lines[counter] = convertToTriangle(scanner.Text()) //Reusing old convert

		if counter == 2 {
			result = append(result, triangle{lines[0].a, lines[1].a, lines[2].a})
			result = append(result, triangle{lines[0].b, lines[1].b, lines[2].b})
			result = append(result, triangle{lines[0].c, lines[1].c, lines[2].c})
			counter = 0
		} else {
			counter++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
