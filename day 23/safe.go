package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	COPY = iota // = 0, and subsequent const one higher
	DEC
	INC
	JUMP
	TOGGLE
)

var register map[string]int

func main() {
	register = make(map[string]int)
	register["a"] = 7
	calculate(readInput(), "Part 1")
	register = make(map[string]int)
	register["a"] = 12
	calculate(readInput(), "Part 2") //Runs fine without optimizations
}

func calculate(instructions []instruction, part string) {
	defer track(time.Now(), part)
	i := 0
	for i < len(instructions) {
		i = instructions[i].Do(i, instructions)
	}
	fmt.Printf("%s: register a contains: %v \n", part, register["a"])
}

type instruction struct {
	command    int
	arg1, arg2 string
}

func (i instruction) Do(p int, instructions []instruction) int {
	switch i.command {
	case COPY:
		_, err := strconv.Atoi(i.arg2)
		if err == nil {
			return p + 1
		}
		val, err := strconv.Atoi(i.arg1)
		if err != nil {
			register[i.arg2] = register[i.arg1]
		} else {
			register[i.arg2] = val
		}
		return p + 1
	case INC:
		register[i.arg1]++
		return p + 1
	case DEC:
		register[i.arg1]--
		return p + 1
	case JUMP:
		val, err := strconv.Atoi(i.arg1)
		if err != nil {
			val = register[i.arg1]
		}
		if val != 0 {
			jmpAmount, err := strconv.Atoi(i.arg2)
			if err != nil {
				jmpAmount = register[i.arg2]
			}
			return p + jmpAmount
		} else {
			return p + 1
		}
	case TOGGLE:
		val, err := strconv.Atoi(i.arg1)
		if err != nil {
			val = register[i.arg1]
		}
		target := p + val
		if target < 0 || target >= len(instructions) {
			return p + 1
		}
		if instructions[target].arg2 != "" {
			if instructions[target].command == JUMP {
				instructions[target].command = COPY
			} else {
				instructions[target].command = JUMP
			}
		} else {
			if instructions[target].command == INC {
				instructions[target].command = DEC
			} else {
				instructions[target].command = INC
			}
		}
		return p + 1
	}
	panic("Could not find command")
}

func readInput() (instructions []instruction) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	copyRE := regexp.MustCompile(`cpy (-?[abcd\d]+) (a|b|c|d)`)
	incRE := regexp.MustCompile(`inc ([abcd])`)
	decRE := regexp.MustCompile(`dec ([abcd])`)
	jumpRE := regexp.MustCompile(`jnz (-?[abcd\d]+) (-?[abcd\d+])`)
	tglRE := regexp.MustCompile(`tgl ([abcd])`)
	scanner := bufio.NewScanner(file)
	toInstruction := func(line string) instruction {
		cmd := copyRE.FindStringSubmatch(line)
		if len(cmd) > 0 {
			return instruction{COPY, cmd[1], cmd[2]}
		} else if cmd = incRE.FindStringSubmatch(line); len(cmd) > 0 {
			return instruction{INC, cmd[1], ""}
		} else if cmd = jumpRE.FindStringSubmatch(line); len(cmd) > 0 {
			return instruction{JUMP, cmd[1], cmd[2]}
		} else if cmd = decRE.FindStringSubmatch(line); len(cmd) > 0 {
			return instruction{DEC, cmd[1], ""}
		} else if cmd = tglRE.FindStringSubmatch(line); len(cmd) > 0 {
			return instruction{TOGGLE, cmd[1], ""}
		}
		panic("Could not parse " + line)
	}
	for scanner.Scan() {
		instructions = append(instructions, toInstruction(scanner.Text()))
	}
	return
}
func track(start time.Time, name string) {
	fmt.Printf("%s took %s \n", name, time.Since(start))
}
