package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strings"
)

func main() {
	input := flag.String("input", "ffykfhsq", "input string")
	part := flag.Int("part", 1, "part 1 or 2?")
	flag.Parse()
	if *part == 1 {
		part1(*input)
	} else {
		part2(*input)
	}
}

func part1(input string) {
	c, quit := make(chan string, 8), make(chan bool)
	go dispatch(input, c, quit)
	p := make([]string, 8, 8)
	p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7] = <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c
	fmt.Printf("The password is %v", strings.Join(p, ""))
	close(quit)
}

func dispatch(prefix string, c chan string, quit <-chan bool) {
	i := 0
	routines := 20
	for ; i < routines; i++ {
		go findGoodHash(prefix, i, routines, c, quit)
	}
}

func findGoodHash(prefix string, i, increment int, c chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			hash := hash(prefix, i)
			if strings.HasPrefix(hash, "00000") {
				c <- string(hash[5])
			}
			i += increment
		}
	}
}

func hash(prefix string, i int) string {
	md5HashInBytes := md5.Sum([]byte(fmt.Sprintf("%v%v", prefix, i)))
	return hex.EncodeToString(md5HashInBytes[:])
}
