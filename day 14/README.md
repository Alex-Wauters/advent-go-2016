# Advent of Code 2016 in Go
Day 14: One-Time Pad
http://adventofcode.com/2016/day/14

Part 1: Used a separate goroutine for finding hashes and validating them, which is simple and performant for part 1 (~1.2 sec on my machine).

For part 2, the approach of part 1 means the expensive stretchedHash is generated anew by each routine. The sequential single threaded process works pretty decent, around 14 sec on my machine.
Can be optimized by generating a rainbow table with multiple goroutines and validating the candidate keys afterwards.

## Installing Go
To install Go, see https://golang.org/

For an interactive tour of go's syntax and features, visit https://tour.golang.org

As for an IDE, I use the go plug-in for IntelliJ IDEA (also available for the free Community edition).
Your favorite text editor may have a plug-in:
https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

##Run From Source
Part 1:
`go run pad_pt1.go`
Part 2:
`go run pad_pt2.go`

![lets go](http://i.imgur.com/sDBaVEy.png)


