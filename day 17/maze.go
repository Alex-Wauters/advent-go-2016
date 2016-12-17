package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const PREFIX = "awrkjxxr"

func main() {
	part1()
	part2()
}

func part1() {
	defer track(time.Now(), "Part 1")
	q := make(chan state, 1000)
	start := state{point{0, 0}, ""}
	q <- start
	for node := range q {
		for _, route := range node.routes() {
			if route.isDestination() {
				fmt.Printf("Reached the Vault! Directions: %v\n", route.history)
				return
			}
			q <- route
		}
	}
}

var max int

func part2() {
	defer track(time.Now(), "Part 2")
	start := state{point{0, 0}, ""}
	solve(start)
	fmt.Printf("Longest path takes %v steps \n", max)
}

func solve(s state) {
	for _, route := range s.routes() {
		if route.isDestination() {
			if len(route.history) > max {
				max = len(route.history)
			}
		} else {
			solve(route)
		}
	}
}

//Adjacent valid states from this position
func (s state) routes() (routes []state) {
	doors := s.hasDoors()
	open := s.openDirections()
	for i := range doors {
		if doors[i] && open[i] {
			routes = append(routes, state{s.nextPosition(i), s.history + getDirection(i)})
		}
	}
	return
}

func getDirection(i int) string {
	switch i {
	case 0:
		return "U"
	case 1:
		return "D"
	case 2:
		return "L"
	case 3:
		return "R"
	}
	panic("Unknown direction")
}

type state struct {
	point
	history string
}

type point struct {
	x, y int
}

func (p point) nextPosition(i int) point {
	switch i {
	case 0:
		return point{p.x, p.y - 1}
	case 1:
		return point{p.x, p.y + 1}
	case 2:
		return point{p.x - 1, p.y}
	case 3:
		return point{p.x + 1, p.y}
	}
	panic("Unknown direction")
}

//Does this point have doors (open or closed) UP, DOWN, LEFT or RIGHT?
func (p point) hasDoors() (doors [4]bool) {
	doors[0] = p.y != 0
	doors[1] = p.y != 3
	doors[2] = p.x != 0
	doors[3] = p.x != 3
	return
}

func (p point) isDestination() bool {
	return p.x == 3 && p.y == 3
}

//Open directions from this state (if there were a door, is it open?)
func (s state) openDirections() (doors [4]bool) {
	h := hash(s)
	doors[0] = isOpenCharacter(h[0])
	doors[1] = isOpenCharacter(h[1])
	doors[2] = isOpenCharacter(h[2])
	doors[3] = isOpenCharacter(h[3])
	return
}

func isOpenCharacter(b byte) bool {
	return b == 'b' || b == 'c' || b == 'd' || b == 'e' || b == 'f'
}

func hash(s state) string {
	md5HashInBytes := md5.Sum([]byte(fmt.Sprintf("%v%v", PREFIX, s.history)))
	return hex.EncodeToString(md5HashInBytes[:])
}

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
