package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	ZERO   = 0
	OPEN   = 8
	WALL   = 9
	WIDTH  = 179
	HEIGHT = 39
)

var maze [][]int
var start state

func main() {
	readInput()
	solve("Part 1", 8)
	solve("Part 2", 9)
}

func solve(task string, end int) {
	defer track(time.Now(), task)
	visited := make(map[string]int)
	q := make(chan state, 100000)
	q <- start
	for node := range q {
		for _, route := range node.Routes() {
			if len(route.visited) == end {
				fmt.Printf("Reached all destinations! Took %v steps\n", route.distance)
				return
			}
			if visited[route.String()] == 0 {
				q <- route
				visited[route.String()] = route.distance
			}
		}
	}
}

type state struct {
	x, y     int
	visited  string
	distance int
}

func (s state) Routes() (routes []state) {
	if s.x < WIDTH-1 {
		visited, canVisit := Visit(s.x+1, s.y, s.visited)
		if canVisit {
			routes = append(routes, state{s.x + 1, s.y, visited, s.distance + 1})
		}
	}
	if s.x > 0 {
		visited, canVisit := Visit(s.x-1, s.y, s.visited)
		if canVisit {
			routes = append(routes, state{s.x - 1, s.y, visited, s.distance + 1})
		}
	}
	if s.y < HEIGHT-1 {
		visited, canVisit := Visit(s.x, s.y+1, s.visited)
		if canVisit {
			routes = append(routes, state{s.x, s.y + 1, visited, s.distance + 1})
		}
	}
	if s.y > 0 {
		visited, canVisit := Visit(s.x, s.y-1, s.visited)
		if canVisit {
			routes = append(routes, state{s.x, s.y - 1, visited, s.distance + 1})
		}
	}
	return
}

func (s state) String() string {
	return fmt.Sprintf("%v,%v,%v", s.x, s.y, s.visited)
}

func Visit(x, y int, visited string) (string, bool) {
	spot := maze[y][x]
	if spot == WALL {
		return visited, false
	}
	if spot == OPEN {
		return visited, true
	}
	if len(visited) == 8 && spot == ZERO { //Part 2 addition
		return visited + strconv.Itoa(spot), true
	}
	if strings.Contains(visited, strconv.Itoa(spot)) {
		return visited, true
	} else {
		return SortString(visited + strconv.Itoa(spot)), true
	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func readInput() {
	maze = make([][]int, HEIGHT)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		maze[i] = make([]int, WIDTH)
		for k, r := range scanner.Text() {
			parsed, err := strconv.Atoi(string(r))
			if err != nil {
				if r == '#' {
					maze[i][k] = WALL
				} else {
					maze[i][k] = OPEN
				}
			} else {
				maze[i][k] = parsed
				if parsed == ZERO {
					start = state{k, i, "0", 0}
				}
			}
		}
		i++
	}
}

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
