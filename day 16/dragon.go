package main

import (
	"fmt"
	"time"
)

func main() {
	calculate(272, []bool{true, false, false, false, true, true, true, false, false, true, true, true, true, false, false, false, false})
	calculate(35651584, []bool{true, false, false, false, true, true, true, false, false, true, true, true, true, false, false, false, false})
}

func calculate(discLen int, input []bool) {
	defer track(time.Now(), "Calculation")
	fill := data(input)
	for len(fill) < discLen {
		fill = data(fill)
	}
	sum := checksum(fill[:discLen])
	fmt.Printf("The checksum is %s", toString(sum))
}

func toString(sum []bool) string {
	result := ""
	for _, char := range sum {
		if char {
			result = result + "1"
		} else {
			result = result + "0"
		}
	}
	return result
}

func data(a []bool) []bool {
	b := make([]bool, len(a))
	for i := len(b)/2 - 1; i >= 0; i-- {
		opp := len(b) - 1 - i
		b[i], b[opp] = !a[opp], !a[i]
	}
	if len(a)%2 != 0 {
		b[len(a)/2] = !a[len(a)/2]
	}
	return append(a, append([]bool{false}, b...)...)
}

func checksum(a []bool) []bool {
	sum := make([]bool, len(a)/2)
	for i := 0; i < len(a); i = i + 2 {
		sum[i/2] = (a[i] == a[i+1])
	}
	if len(sum)%2 != 0 {
		return sum
	} else {
		return checksum(sum)
	}
}

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
