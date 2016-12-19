package main

import (
	"fmt"
	"time"
)

const MAX_ELVES = 3004953

func main() {
	part1()
	part2()
}

func part1() {
	defer track(time.Now(), "part 1")
	elves := make([]int, MAX_ELVES)
	for i := range elves {
		elves[i] = 1 // Each elf has a present
	}
	for {
	Round:
		for i := range elves {
			if elves[i] == 0 {
				continue Round
			}
			for k := i + 1; k < MAX_ELVES; k++ {
				if elves[k] != 0 {
					elves[k] = 0
					continue Round
				}
			} //Reached end of circle
			for k := 0; k < i; k++ {
				if elves[k] != 0 {
					elves[k] = 0
					continue Round
				}
			} //Looped all the way around
			fmt.Printf("Part 1: The master thief is elf # %v \n", i+1)
			return
		}
	}
}

func part2() {
	defer track(time.Now(), "part 2")
	elves := make([]int, MAX_ELVES)
	for i := 0; i < MAX_ELVES; i++ {
		elves[i] = i + 1 // Elf at position i has number i+1
	}
	var next, opp int
	for {
		opp = (next + len(elves)/2) % len(elves)
		elves = append(elves[:opp], elves[opp+1:]...) // Cut the opposite elf from the circle
		if len(elves) == 1 {
			fmt.Printf("Part 2: The master thief is elf # %v \n", elves[0])
			return
		}
		if opp > next { // Don't increase next if the element cut took place before it
			next = next + 1
		}
		if next >= len(elves) {
			next = 0 // If at end of circle, restart circle
		}
	}
}

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
