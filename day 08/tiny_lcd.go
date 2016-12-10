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
	wide = 50
	tall = 6
)

type lcd [][]bool

var rotateExp, rectExp *regexp.Regexp

func main() {
	defer track(time.Now(), "lcd pt 1 and 2")
	rotateExp = regexp.MustCompile(`rotate (?P<Entity>column|row) [xy]=(?P<Id>\d+) by (?P<Amount>\d+)`) //entity = 1, id = 2, amount = 3
	rectExp = regexp.MustCompile(`rect (?P<x>\d+)x(?P<y>\d+)`)                                          //x =1, y = 2
	lcd := matrix(wide, tall)
	commands := readInput()
	for _, command := range commands {
		command(lcd)
	}
	fmt.Printf("There are %v pixels lit up \n", lcd.count())
	for _, row := range lcd {
		for _, px := range row {
			if px {
				fmt.Print("X")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Printf("\n")
	}
}

func matrix(x, y int) (result lcd) {
	result = make(lcd, y)
	for i := range result {
		result[i] = make([]bool, x)
	}
	return
}

func (lcd lcd) count() (count int) {
	for x, row := range lcd {
		for y := range row {
			if lcd[x][y] {
				count++
			}
		}
	}
	return
}

//In-place rotation of a row of booleans
func rotateRow(a []bool, r int) {
	n := len(a) - r
	i := 0
	j := n
	for i != j {
		a[i], a[j] = a[j], a[i]
		i++
		j++
		if j == len(a) {
			j = n
		} else if i == n {
			n = j
		}
	}
}

//Creating a new []bool slice with the column values won't update the original matrix
//To update a column, use the original matrix instead of a new slice with column values
func rotateCol(a lcd, col, r int) {
	n := len(a) - r
	i := 0
	j := n
	for i != j {
		a[i][col], a[j][col] = a[j][col], a[i][col]
		i++
		j++
		if j == len(a) {
			j = n
		} else if i == n {
			n = j
		}
	}
}

func fill(lcd lcd, wide int, tall int) {
	for row := 0; row < tall; row++ {
		for col := 0; col < wide; col++ {
			lcd[row][col] = true
		}
	}
}

func command(line string) func(lcd) {
	cmd := rotateExp.FindStringSubmatch(line)
	if len(cmd) > 0 {
		return func(lcd lcd) {
			if cmd[1] == "column" {
				rotateCol(lcd, toInt(cmd[2]), toInt(cmd[3]))
			} else {
				rotateRow(lcd[toInt(cmd[2])], toInt(cmd[3]))
			}
		}
	} else if cmd = rectExp.FindStringSubmatch(line); len(cmd) > 0 {
		return func(lcd lcd) {
			fill(lcd, toInt(cmd[1]), toInt(cmd[2]))
		}
	}
	panic("Could not parse line: " + line)
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func readInput() (commands []func(lcd)) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, command(scanner.Text()))
	}
	return
}

func track(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s \n", name, elapsed)
}
