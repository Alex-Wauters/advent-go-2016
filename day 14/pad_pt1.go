package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const PREFIX = "ihaygndm"

func main() {
	defer track(time.Now(), "Part 1")
	keys, candidates := make(chan candidate, 64), make(chan candidate, 100)
	go findCandidate(candidates)
	go validateCandidate(candidates, keys)
	pad := []candidate{}
	for len(pad) < 64 {
		key := <-keys
		pad = append(pad, key)
	}
	fmt.Printf("The index of the 64th pad is %v", pad[63].index)
}

type candidate struct {
	index int
	char  byte
}

func findCandidate(candidates chan<- candidate) {
	var candHash string
Hash:
	for i := 0; ; i++ {
		candHash = hash(PREFIX, i)
		for k := 0; k < len(candHash)-2; k++ {
			if candHash[k] == candHash[k+1] && candHash[k+1] == candHash[k+2] {
				candidates <- candidate{i, candHash[k]}
				continue Hash
			}
		}
	}
}

func validateCandidate(candidates <-chan candidate, keys chan<- candidate) {
	var validator string
Validate:
	for c := range candidates {
		for i := c.index + 1; i < c.index+1001; i++ {
			validator = hash(PREFIX, i)
			for k := 0; k < len(validator)-4; k++ {
				if validator[k] == c.char && validator[k] == validator[k+1] && validator[k+1] == validator[k+2] && validator[k+2] == validator[k+3] && validator[k+3] == validator[k+4] {
					keys <- c
					continue Validate
				}
			}
		}
	}
}

func hash(prefix string, i int) string {
	md5HashInBytes := md5.Sum([]byte(fmt.Sprintf("%v%v", prefix, i)))
	return hex.EncodeToString(md5HashInBytes[:])
}

func track(start time.Time, name string) {
	fmt.Printf("%s took %s \n", name, time.Since(start))
}
