package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rooms := readInput()
	part1(rooms)
	part2(rooms)
}

func part1(rooms []Room) {
	sum := 0
	for _, room := range rooms {
		if room.IsCorrect() {
			sum = sum + room.sector
		}
	}
	fmt.Printf("Sector sum of correct rooms: %v", sum)
}

type Room struct {
	sector         int
	top5, checksum, encrypted string
}

func (r Room) IsCorrect() bool {
	return r.top5 == r.checksum
}

func readInput() []Room {
	file, _ := os.Open("input.txt")
	defer file.Close()
	result := make([]Room, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := toRoom(scanner.Text())
		result = append(result, t)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func toRoom(line string) Room {
	f := strings.Split(line, "-")
	top := top5Chars(f[0 : len(f)-1])
	sector, hash := sectorAndHash(f[len(f)-1])
	return Room{sector, top, hash,strings.Join(f[0: len(f)-1]," ")}
}

var re = regexp.MustCompile(`(?P<sector>\d{3})\[(?P<hash>\w{5})\]`)

func sectorAndHash(line string) (int, string) {
	matches := re.FindAllStringSubmatch(line, -1)
	return toInt(matches[0][1]), matches[0][2]
}

func toInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

func top5Chars(line []string) string {
	chars := make(map[string]int)
	for _, word := range line {
		for _, char := range word {
			chars[string(char)]++
		}
	}
	p := make(pairList, len(chars))
	i := 0
	for k, v := range chars {
		p[i] = pair{k, v}
		i++
	}
	sort.Sort(p)
	return fmt.Sprintf("%v%v%v%v%v", p[0], p[1], p[2], p[3], p[4])
}

type pair struct {
	Key   string
	Value int
}

func (p pair) String() string {
	return p.Key
}

type pairList []pair

func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairList) Len() int      { return len(p) }
func (p pairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key < p[j].Key
	} else {
		return p[i].Value > p[j].Value
	}
}
