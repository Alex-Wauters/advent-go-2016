package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	nodes := parseInput()
	fmt.Println("Part 1: Valid pairs:", validPairs(nodes))
	grid := make([][]node, 30)
	for i := 0; i < 30; i++ {
		grid[i] = make([]node, 0)
	}
	for _, node := range nodes {
		grid[node.y] = append(grid[node.y], node)
	}
	fmt.Println("0 1 2 3 4 5 6 7 8 9 101112131415161718192021222324252627282930313233")
	for i := 0; i < 30; i++ {
		for _, n := range grid[i] {
			fmt.Print(n)
		}
		fmt.Println()
	}
}

func validPairs(n []node) int {
	valid := make(map[string]bool)
	for _, a := range n {
		for _, b := range n {
			if a != b && a.used <= b.avail && a.used != 0 {
				valid[pairString(a, b)] = true
			}
		}
	}
	return len(valid)
}

type node struct {
	x, y, size, used, avail, percent int
}

func (n node) String() string {
	if n.used > 100 {
		return "# "
	}
	if n.used == 0 {
		return "_ "
	}
	if n.x == 33 && n.y == 0 {
		return "G "
	}
	return ". "
}

func pairString(a, b node) string {
	switch {
	case a.x < b.x:
		return fmt.Sprintf("%v,%v-%v,%v", a.x, a.y, b.x, b.y)
	case a.x == b.x:
		if a.y < b.y {
			return fmt.Sprintf("%v,%v-%v,%v", a.x, a.y, b.x, b.y)
		}
		fallthrough
	default:
		return fmt.Sprintf("%v,%v-%v,%v", b.x, b.y, a.x, a.y)
	}
}

func parseInput() (nodes []node) {
	nodeRE := regexp.MustCompile(`/dev/grid/node-x(\d*)-y(\d*)[\s]*(\d*)T[\s]*(\d*)T[\s]*(\d*)T[\s]*(\d*)%`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Scan() //Skip the first two lines
	for scanner.Scan() {
		n := nodeRE.FindStringSubmatch(scanner.Text())
		nodes = append(nodes, node{toInt(n[1]), toInt(n[2]), toInt(n[3]), toInt(n[4]), toInt(n[5]), toInt(n[6])})
	}
	return
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
