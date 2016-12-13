package main

import (
	"fmt"
	"strconv"
)

const FAVORITE_NUMBER = 1362

func main() {
	part1()
	part2()
}

func part1() {
	visited := make(map[string]position)
	q := make(chan position, 10000)
	start := position{point{1, 1}, 0, true}
	visited["1,1"] = start
	q <- start
	for node := range q {
		for _, route := range node.routes() {
			if route.IsDestination() {
				fmt.Printf("It took %v steps to reach the destination \n", node.distance+1)
				return
			}
			if _, seen := visited[route.String()]; !seen {
				newNode := position{route, node.distance + 1, route.isOpenSpace()}
				visited[route.String()] = newNode
				if newNode.isOpenSpace() {
					q <- newNode
				}
			}
		}
	}
}

type position struct {
	point    //Embed the point struct in position
	distance int
	open     bool
}

//Adjacent walls or open spaces from this position
func (p position) routes() (routes []point) {
	for _, addX := range []int{-1, 1} {
		if newX := p.x + addX; newX >= 0 {
			routes = append(routes, point{newX, p.y})
		}
	}
	for _, addY := range []int{-1, 1} {
		if newY := p.y + addY; newY >= 0 {
			routes = append(routes, point{p.x, newY})
		}
	}
	return
}

type point struct {
	x, y int
}

func (p point) isOpenSpace() bool {
	isOpen := true
	binary := strconv.FormatInt(int64(p.x*p.x+3*p.x+2*p.x*p.y+p.y+p.y*p.y+FAVORITE_NUMBER), 2)
	for _, c := range binary {
		if c == '1' {
			isOpen = !isOpen
		}
	}
	return isOpen
}

func (p point) String() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func (p point) IsDestination() bool {
	return p.x == 31 && p.y == 39
}
