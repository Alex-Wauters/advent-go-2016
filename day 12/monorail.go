package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

var register map[string]int
var copyRE, incRE, jumpRE, decRE *regexp.Regexp

func main() {
	instructions := readInput()
	register = make(map[string]int)
	calculate(instructions, "Part 1")
	register = make(map[string]int)
	register["c"] = 1
	calculate(instructions, "Part 2")
}

func calculate(instructions []func(int) int, part string) {
	defer track(time.Now(), part)
	i := 0
	for i < len(instructions) {
		i = instructions[i](i)
	}
	fmt.Printf("%s: register a contains: %v \n", part, register["a"])
}
func readInput() (instructions []func(int) int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	copyRE = regexp.MustCompile(`cpy (-?[abcd]|\d+) (a|b|c|d)`)
	incRE = regexp.MustCompile(`inc ([abcd])`)
	decRE = regexp.MustCompile(`dec ([abcd])`)
	jumpRE = regexp.MustCompile(`jnz (-?[abcd]|\d+) (-?\d+)`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, toInstruction(scanner.Text()))
	}
	return
}
func copyAmount(r string, amount int) func(int) int {
	return func(instr int) int {
		register[r] = amount
		return instr + 1
	}
}
func copyRegister(destination, source string) func(int) int {
	return func(instr int) int {
		register[destination] = register[source]
		return instr + 1
	}
}
func inc(r string) func(int) int {
	return func(instr int) int {
		register[r]++
		return instr + 1
	}
}
func dec(r string) func(int) int {
	return func(instr int) int {
		register[r]--
		return instr + 1
	}
}
func jump(r string, jumpAmount int) func(int) int {
	return func(instr int) int {
		var compare int
		if isInt(r) {
			compare = toInt(r)
		} else {
			compare = register[r]
		}
		if compare != 0 {
			return instr + jumpAmount
		} else {
			return instr + 1
		}
	}
}
func toInstruction(line string) func(int) int {
	cmd := copyRE.FindStringSubmatch(line)
	if len(cmd) > 0 {
		if isInt(cmd[1]) {
			return copyAmount(cmd[2], toInt(cmd[1]))
		} else {
			return copyRegister(cmd[2], cmd[1])
		}
	} else if cmd = incRE.FindStringSubmatch(line); len(cmd) > 0 {
		return inc(cmd[1])
	} else if cmd = jumpRE.FindStringSubmatch(line); len(cmd) > 0 {
		return jump(cmd[1], toInt(cmd[2]))
	} else if cmd = decRE.FindStringSubmatch(line); len(cmd) > 0 {
		return dec(cmd[1])
	}
	panic("Could not parse " + line)
}
func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return
}
func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
func track(start time.Time, name string) {
	fmt.Printf("%s took %s \n", name, time.Since(start))
}
