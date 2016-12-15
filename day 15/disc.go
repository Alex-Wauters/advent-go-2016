package main

import "fmt"

func main() {
	discs := []disc{{5, 17}, {8, 19}, {1, 7}, {7, 13}, {1, 5}, {0, 3}} //Part 2: Add ,{0,11}

Time:
	for t := 0; ; t++ {
		k := t + 1
		for _, disc := range discs {
			if !disc.isOpen(k) {
				continue Time
			}
			k++
		}
		fmt.Printf("Retrieved capsule at t=%v \n", t)
		return
	}
}

type disc struct {
	begin, total int
}

func (d disc) isOpen(t int) bool {
	return (d.begin+t)%d.total == 0
}
