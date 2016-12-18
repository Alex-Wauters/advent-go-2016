package main

import "fmt"

func main() {
	calculate(40)     //part 1
	calculate(400000) //part 2
}

func calculate(rows int) {
	maze := make([]string, rows)
	maze[0] = ".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^."
	count := 48
	for i := 1; i < len(maze); i++ {
		for k := 0; k < len(maze[0]); k++ {
			var safe bool
			switch k {
			case 0:
				safe = isSafe(true, isSafeTile(maze[i-1][k]), isSafeTile(maze[i-1][k+1]))
			case len(maze[0]) - 1:
				safe = isSafe(isSafeTile(maze[i-1][k-1]), isSafeTile(maze[i-1][k]), true)
			default:
				safe = isSafe(isSafeTile(maze[i-1][k-1]), isSafeTile(maze[i-1][k]), isSafeTile(maze[i-1][k+1]))
			}
			if safe {
				count++
				maze[i] = maze[i] + "."
			} else {
				maze[i] = maze[i] + "^"
			}
		}
	}
	fmt.Printf("There are %v safe tiles \n", count)
}

func isSafe(left, middle, right bool) bool {
	switch {
	case !left && !middle && right:
		return false
	case left && !middle && !right:
		return false
	case !left && middle && right:
		return false
	case left && middle && !right:
		return false
	default:
		return true
	}
}

func isSafeTile(tile byte) bool {
	return tile == byte('.')
}
