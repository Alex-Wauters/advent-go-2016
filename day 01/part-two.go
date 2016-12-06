package main

import "fmt"

func part2(path string) {
	location := realHQ(parseInstructions(path))
	fmt.Printf("the REAL hq is located at %v with distance: %v", location, location.distance())
}

func (p point) String() string {
	return fmt.Sprintf("%v %v", p.x, p.y)
}

func realHQ(instr []instruction) point {
	location := point{0, 0, N}
	visited := make(map[string]bool)
	for _, instruction := range instr {
		oldLoc := location
		updateLocation(&location, instruction)
		for _, loc := range locationsBetween(oldLoc, location) {
			if visited[loc.String()] {
				return loc
			}
			visited[loc.String()] = true
		}
	}
	panic("No location visited twice")
}

func locationsBetween(p, q point) []point {
	result := make([]point, 0)
	if p.x != q.x {
		if p.x > q.x {
			for i := p.x -1; i > q.x; i-- {
				result = append(result, point{x:i,y:p.y})
			}
		} else {
			for i := p.x + 1; i < q.x; i++ {
				result = append(result, point{x:i,y: p.y})
			}
		}
	} else {
		if p.y > q.y {
			for i := p.y -1; i > q.y; i-- {
				result = append(result, point{x:p.x,y:i})
			}
		} else {
			for i := p.y + 1; i < q.y ; i++ {
				result = append(result, point{x:p.x,y:i})
			}
		}
	}
	result = append(result, q)
	return result
}
