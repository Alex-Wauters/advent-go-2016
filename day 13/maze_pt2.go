package main

import "fmt"

func part2() {
	visited := make(map[string]position)
	q := make(chan position, 10000)
	start := position{point{1, 1}, 0, true}
	visited["1,1"] = start
	q <- start
	traverse(q, visited)
	count := 0
	for _, node := range visited {
		if node.open {
			count++
		}
	}
	fmt.Printf("In 50 steps, one can reach at most %v distinct locations \n", count)
}

func traverse(q chan position, visited map[string]position) {
	for {
		select {
		case node := <-q:
			for _, route := range node.routes() {
				if _, seen := visited[route.String()]; !seen {
					newNode := position{route, node.distance + 1, route.isOpenSpace()}
					visited[route.String()] = newNode
					if newNode.isOpenSpace() && newNode.distance < 50 {
						q <- newNode
					}
				}
			}
		default: //No values left on channel
			return
		}
	}
}
