package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ips := readInput()
	var countAbba, countSSL int
	for _, ip := range ips {
		if ip.supportsAbba() {
			countAbba++
		}
		if ip.supportsSSL() { //part 2
			countSSL++
		}
	}
	fmt.Printf("There are %v IPv7 addresses which support ABBA and %v which support SSL \n", countAbba, countSSL)
}

type ipv7 struct {
	supernet []string
	hypernet []string
}

func (ip ipv7) supportsAbba() bool {
	for _, hyp := range ip.hypernet {
		if isValidAbbaSegment(hyp) {
			return false
		}
	}
	for _, sequence := range ip.supernet {
		if isValidAbbaSegment(sequence) {
			return true
		}
	}
	return false
}

func isValidAbbaSegment(seg string) bool {
	for i := 0; i < len(seg)-3; i++ {
		if isAbba(seg[i : i+4]) {
			return true
		}
	}
	return false
}

func isAbba(chars string) bool {
	return chars[0] == chars[3] && chars[1] == chars[2] && chars[0] != chars[1]
}

func readInput() []ipv7 {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []ipv7
	for scanner.Scan() {
		lines = append(lines, toIp(scanner.Text()))
	}
	return lines
}

func toIp(s string) (ip ipv7) {
	ip = ipv7{make([]string, 0), make([]string, 0)}
	lastPos := 0
	for i, char := range s {
		if char == '[' {
			ip.supernet = append(ip.supernet, s[lastPos:i])
			lastPos = i + 1
		} else if char == ']' {
			ip.hypernet = append(ip.hypernet, s[lastPos:i])
			lastPos = i + 1
		} else if i == len(s)-1 {
			ip.supernet = append(ip.supernet, s[lastPos:])
		}
	}
	return
}
