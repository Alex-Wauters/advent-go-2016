# Advent of Code 2016 in Go
Day 23: Safe Cracking
http://adventofcode.com/2016/day/23

Day 23 expands the solution of day 12. For day 12 I used a slice of functions. In this puzzle, a function may be morphed into another function, so I switched to structs which remember the arguments and the variable command instruction.
Part 1 takes about 8 ms.


The description for part 2 seems to hint that unless you optimize the 'inc / dec / jump' loops into multiply loops, the result would take too long to compute.
I let the part 1 implementation run anyway, and it finished after 9m33secs while I was still working on the optimization implementation.


## Installing Go
To install Go, see https://golang.org/

For an interactive tour of go's syntax and features, visit https://tour.golang.org

As for an IDE, I use the go plug-in for IntelliJ IDEA.
Your favorite text editor may have a plug-in:
https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

##Run From Source
`go run safe.go`


![lets go](http://i.imgur.com/sDBaVEy.png)


