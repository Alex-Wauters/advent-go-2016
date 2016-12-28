package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	defer track(time.Now(), "Part 1 and 2")
	fw := readInput()
	sort.Sort(fw)
	part1(fw)
	part2(fw)
}

func part1(fw firewall) {
	i := uint64(0)
	for _, r := range fw {
		if r.inRange(i) {
			i = r.end + 1
		} else if i < r.begin {
			fmt.Println(i)
			return
		}
	}
}

func part2(fw firewall) {
	merged := true
	for merged {
		merged = false
		for j := len(fw) - 1; j > 0; j-- {
			if overlap(fw[j-1], fw[j]) {
				fw[j-1] = merge(fw[j-1], fw[j])
				merged = true
				fw = append(fw[:j], fw[j+1:]...)
			}
		}
	}
	blacklist := uint64(0)
	for r := 0; r < len(fw); r++ {
		blacklist = blacklist + fw[r].end - fw[r].begin + 1
	}
	fmt.Println(uint64(4294967296) - blacklist)
}

type rule struct {
	begin, end uint64
}

func (r rule) inRange(i uint64) bool {
	return r.begin <= i && i <= r.end
}

func overlap(r1, r2 rule) bool {
	return r2.begin <= r1.end+1
}

func merge(r1, r2 rule) rule {
	end := r2.end
	if r1.end > r2.end {
		end = r1.end
	}
	return rule{r1.begin, end}
}

func readInput() (fw firewall) {
	rangeRE := regexp.MustCompile(`(\d+)-(\d+)`)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := rangeRE.FindStringSubmatch(scanner.Text())
		fw = append(fw, rule{toInt(matches[1]), toInt(matches[2])})
	}
	return
}

func toInt(s string) uint64 {
	if s, err := strconv.ParseUint(s, 10, 64); err == nil {
		return uint64(s)
	} else {
		panic(err)
	}
}

type firewall []rule

func (f firewall) Len() int           { return len(f) }
func (f firewall) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f firewall) Less(i, j int) bool { return f[i].begin < f[j].begin }

func track(start time.Time, name string) {
	fmt.Printf(" %s took %s \n", name, time.Since(start))
}
