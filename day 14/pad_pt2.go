package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const PREFIX = "ihaygndm"

func main() {
	defer track(time.Now(), "Part 2")
	keys := make(chan candidate, 64)
	go findCandidate(keys)
	pad := []candidate{}
	for len(pad) < 64 {
		key := <-keys
		pad = append(pad, key)
	}

	fmt.Printf("The index of the 64th pad is %v", pad[63].index)
}

type candidate struct {
	index       int
	char        byte
	remaining   int
	isValidated bool
}

func findCandidate(keys chan<- candidate) {
	var hash string
	candidates := []*candidate{}
Hash:
	for i := 0; ; i++ {
		hash = stretchedHash(PREFIX, i) // Generate hash once for both finding and validating candidates

		for _, candidate := range candidates { // Validate candidates
			if !candidate.isValidated && candidate.remaining > 0 {
				if isValidCandidate(candidate, hash) {
					candidate.isValidated = true
				} else {
					candidate.remaining--
				}
			}
		}
		if len(candidates) > 0 && candidates[0].isValidated { // Remove a validated candidate and send to main
			keys <- *candidates[0]
			candidates = candidates[1:]
		}
		if len(candidates) > 0 && candidates[0].remaining < 1 { //Remove candidate with no validation key
			candidates = candidates[1:]
		}
		for k := 0; k < len(hash)-2; k++ { // Generate new candidates
			if hash[k] == hash[k+1] && hash[k+1] == hash[k+2] {
				candidates = append(candidates, &candidate{i, hash[k], 1000, false})
				continue Hash
			}
		}
	}
}

func isValidCandidate(c *candidate, hash string) bool {
	for k := 0; k < len(hash)-4; k++ {
		if hash[k] == c.char && hash[k] == hash[k+1] && hash[k+1] == hash[k+2] && hash[k+2] == hash[k+3] && hash[k+3] == hash[k+4] {
			return true
		}
	}
	return false
}

func hash(prefix string, i int) string {
	md5HashInBytes := md5.Sum([]byte(fmt.Sprintf("%v%v", prefix, i)))
	return hex.EncodeToString(md5HashInBytes[:])
}

func stretchedHash(prefix string, i int) string {
	hash := hash(prefix, i)
	for k := 0; k < 2016; k++ {
		hash = hashString(hash)
	}
	return hash
}

func hashString(plain string) string {
	md5HashInBytes := md5.Sum([]byte(plain))
	return hex.EncodeToString(md5HashInBytes[:])
}

func track(start time.Time, name string) {
	fmt.Printf("%s took %s \n", name, time.Since(start))
}
