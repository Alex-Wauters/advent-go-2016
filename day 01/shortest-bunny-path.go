package main

import (
	"fmt"
	"strconv"
	"strings"
)

type facing string

const (
	L string = "L"
	R string = "R"
	N facing = "N"
	S facing = "S"
	W facing = "W"
	E facing = "E"
)

func main() {
	path := "L3, R1, L4, L1, L2, R4, L3, L3, R2, R3, L5, R1, R3, L4, L1, L2, R2, R1, L4, L4, R2, L5, R3, R2, R1, L1, L2, R2, R2, L1, L1, R2, R1, L3, L5, R4, L3, R3, R3, L5, L190, L4, R4, R51, L4, R5, R5, R2, L1, L3, R1, R4, L3, R1, R3, L5, L4, R2, R5, R2, L1, L5, L1, L1, R78, L3, R2, L3, R5, L2, R2, R4, L1, L4, R1, R185, R3, L4, L1, L1, L3, R4, L4, L1, R5, L5, L1, R5, L1, R2, L5, L2, R4, R3, L2, R3, R1, L3, L5, L4, R3, L2, L4, L5, L4, R1, L1, R5, L2, R4, R2, R3, L1, L1, L4, L3, R4, L3, L5, R2, L5, L1, L1, R2, R3, L5, L3, L2, L1, L4, R4, R4, L2, R3, R1, L2, R1, L2, L2, R3, R3, L1, R4, L5, L3, R4, R4, R1, L2, L5, L3, R1, R4, L2, R5, R4, R2, L5, L3, R4, R1, L1, R5, L3, R1, R5, L2, R1, L5, L2, R2, L2, L3, R3, R3, R1"
	part1(path)
	part2(path)
}

func part1(path string) {
	location := destination(parseInstructions(path))
	fmt.Println(location.distance())

}

type instruction struct {
	dir  string
	dist int
}

func parseInstructions(path string) []instruction {
	raw := strings.FieldsFunc(path, func(r rune) bool {
		return r == ' ' || r == ','
	})
	result := make([]instruction, len(raw))
	for i, dir := range raw {
		length64, err := strconv.ParseInt(string(dir[1:]), 10, 32)
		if err != nil {
			panic(err)
		}
		result[i] = instruction{string(dir[0]), int(length64)}
	}
	return result
}

type point struct {
	x, y   int
	facing facing
}

func (p point) distance() int {
	return abs(p.x) + abs(p.y)
}

func destination(instr []instruction) point {
	location := point{0, 0, N}
	for _, instruction := range instr {
		updateLocation(&location, instruction)
	}
	return location
}

type walkfn func(int, int, int) (int, int)

func decreaseX(x0, y0 int, length int) (int, int) {
	return x0 - length, y0
}

func increaseX(x0, y0 int, length int) (int, int) {
	return x0 + length, y0
}
func decreaseY(x0, y0 int, length int) (int, int) {
	return x0, y0 - length
}
func increaseY(x0, y0 int, length int) (int, int) {
	return x0, y0 + length
}

func updateLocation(loc *point, instr instruction) {
	switch loc.facing {
	case N:
		loc.x, loc.y = walk(decreaseX, increaseX, instr.dir)(loc.x, loc.y, instr.dist)
		loc.facing = facingAfterTurn(W, E, instr.dir)
	case E:
		loc.x, loc.y = walk(increaseY, decreaseY, instr.dir)(loc.x, loc.y, instr.dist)
		loc.facing = facingAfterTurn(N, S, instr.dir)
	case S:
		loc.x, loc.y = walk(increaseX, decreaseX, instr.dir)(loc.x, loc.y, instr.dist)
		loc.facing = facingAfterTurn(E, W, instr.dir)
	case W:
		loc.x, loc.y = walk(decreaseY, increaseY, instr.dir)(loc.x, loc.y, instr.dist)
		loc.facing = facingAfterTurn(S, N, instr.dir)
	}
}

func walk(goLeft, goRight walkfn, dir string) walkfn {
	switch dir {
	case L:
		return goLeft
	case R:
		return goRight
	}
	panic("Could not recognize dir " + dir)
}

func facingAfterTurn(left, right facing, dir string) facing {
	switch dir {
	case L:
		return left
	case R:
		return right
	}
	panic("Could not recognize dir " + dir)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}