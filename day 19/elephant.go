package main

import (
	"container/ring"
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
	r := ring.New(MAX_ELVES)
	var beforeOpposite *ring.Ring
	for i := 0; i < MAX_ELVES; i++ {
		r.Value = i + 1
		r = r.Next()
		if i == MAX_ELVES/2-2 {
			beforeOpposite = r
		}
	}
	active := r
	stay := true
	for remaining := MAX_ELVES; remaining > 1; remaining-- {
		if !stay {
			beforeOpposite = beforeOpposite.Next()
		}
		beforeOpposite.Unlink(1)
		stay = !stay
		active = active.Next()
	}
	fmt.Printf("Part 2: The master thief is elf # %v \n", active.Value)
}

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
