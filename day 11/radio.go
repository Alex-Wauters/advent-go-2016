package main

import (
	"fmt"
	"time"
)

const ITEMS = 5 //part 2: change to 7

type state struct {
	fl1, fl2, fl3, fl4 floor
	floorId            int
	current            *floor
	distance           int
}

type floor struct {
	chips, generators int
}

func main() {
	defer track(time.Now(), "Shortest path calculation")
	visited := make(map[string]int)
	floor1 := floor{ITEMS -2, ITEMS} // In my input state, 2 microchips were on floor 2
	c := make(chan state, 100000)
	c <- state{floor1, floor{2, 0}, floor{}, floor{}, 1, &floor1, 0}
	for state := range c {
		for _, route := range state.adjacent() {
			if route.isSolved() {
				fmt.Printf("Solved in %v steps \n", state.distance+1)
				return
			}
			if _, seen := visited[route.String()]; !seen {
				route.distance = state.distance + 1
				visited[route.String()] = route.distance
				c <- route
			}
		}
	}
}

//Adjacent states from valid moves from this state
func (s state) adjacent() (routes []state) {
	maxG, maxM := Max(s.current.generators), Max(s.current.chips) //Only bring up to 2 items up/down
	if s.floorId != 4 {                                           //Move items up in all combinations
		for g := 0; g <= maxG; g++ {
			for m := 0; m <= maxM; m++ {
				if !(g+m == 0 || g+m > 2 || (maxM+maxG >= 2 && g+m == 1)) {
					routes = append(routes, getNewState(s, true, m, g))
				}
			}
		}
	}
	if s.floorId != 1 { // Move items down in all combinations
		for g := 0; g <= maxG; g++ {
			for m := 0; m <= maxM; m++ {
				if !(g+m == 0 || g+m > 1) {
					routes = append(routes, getNewState(s, false, m, g))
				}
			}
		}
	}
	return
}

func (s state) isInvalid() bool {
	return s.fl1.isInvalid() || s.fl2.isInvalid() || s.fl3.isInvalid() || s.fl4.isInvalid()
}
func (s state) isSolved() bool {
	return s.fl1.isEmpty() && s.fl2.isEmpty() && s.fl3.isEmpty()
}
func (s state) String() string {
	return fmt.Sprintf("%s%s%s%s%v", s.fl1, s.fl2, s.fl3, s.fl4, s.floorId)
}

func (fl floor) isEmpty() bool {
	return fl.chips == 0 && fl.generators == 0
}
func (fl floor) isInvalid() bool {
	return fl.chips != 0 && fl.generators > fl.chips
}
func (fl floor) String() string {
	return fmt.Sprintf("%v%v", fl.chips, fl.generators)
}

func getNewState(s state, up bool, c, g int) state {
	switch s.floorId {
	case 1:
		newFloor := floor{s.fl2.chips + c, s.fl2.generators + g}
		return state{floor{s.fl1.chips - c, s.fl1.generators - g}, newFloor, s.fl3, s.fl4, 2, &newFloor, 0}
	case 2:
		if up {
			newFloor := floor{s.fl3.chips + c, s.fl3.generators + g}
			return state{s.fl1, floor{s.fl2.chips - c, s.fl2.generators - g}, newFloor, s.fl4, 3, &newFloor, 0}
		} else {
			newFloor := floor{s.fl1.chips + c, s.fl1.generators + g}
			return state{newFloor, floor{s.fl2.chips - c, s.fl2.generators - g}, s.fl3, s.fl4, 1, &newFloor, 0}
		}
	case 3:
		if up {
			newFloor := floor{s.fl4.chips + c, s.fl4.generators + g}
			return state{s.fl1, s.fl2, floor{s.fl3.chips - c, s.fl3.generators - g}, newFloor, 4, &newFloor, 0}
		} else {
			newFloor := floor{s.fl2.chips + c, s.fl2.generators + g}
			return state{s.fl1, newFloor, floor{s.fl3.chips - c, s.fl3.generators - g}, s.fl4, 2, &newFloor, 0}
		}
	case 4:
		newFloor := floor{s.fl3.chips + c, s.fl3.generators + g}
		return state{s.fl1, s.fl2, newFloor, floor{s.fl1.chips - c, s.fl1.generators - g}, 3, &newFloor, 0}
	}
	panic("Invalid state")
}

func Max(x int) int {
	if x > 2 {
		return 2
	}
	return x
}

func track(start time.Time, name string) {
	fmt.Printf("%s took %s \n", name, time.Since(start))
}
