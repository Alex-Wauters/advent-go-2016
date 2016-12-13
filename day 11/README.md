# Advent of Code 2016 in Go
Day 11: Radioisotope Thermoelectric Generators
http://adventofcode.com/2016/day/11

This one was quite tricky. The problem is a modification of the jealous husbands problem:
https://en.wikipedia.org/wiki/Missionaries_and_cannibals_problem

Solved with breadth-first search, part 1 takes about 1.5s on my machine and part 2 takes about 23 sec.
To change to part 2, modify the `items` constant from `5` (5 couples of generators+chips) to `7`. 

## Installing Go
To install Go, see https://golang.org/

For an interactive tour of go's syntax and features, visit https://tour.golang.org

As for an IDE, I use the go plug-in for IntelliJ IDEA (also available for the free Community edition).
Your favorite text editor may have a plug-in:
https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

##Run From Source
`go run radio.go`


![lets go](http://i.imgur.com/sDBaVEy.png)


