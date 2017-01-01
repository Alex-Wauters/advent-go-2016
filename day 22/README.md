# Advent of Code 2016 in Go
Day 22: Grid Computing
http://adventofcode.com/2016/day/22

Calculated part 2 by printing the grid and calculating the route. Printing the grid results in

`0 1 2 3 4 5 6 7 8 9 101112131415161718192021222324252627282930313233
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . G
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . # # # # # # # # # # # # # # # # # # # # # # # # # #
 . . . . . . . . . . . . . . . . . _ . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .`
 
 The empty disk (_) needs to get right before the Goal data (top row, col 32), by taking a turn around the large nodes which we can't empty (#).
 
 `12` + `25` (to col 32) + `4` (top row)
 
 Then a swap operation with the Goal data (`1`) will place the goal data at X=32 and the empty disk at X=33.
 Moving the disk back in front of the goal data takes `4 moves, + 1 swap`. After doing this again, G is at x=31
 Repeat this another 31 times to get the disk at X=0 and swap, results in an extra 31*5 moves. (=202)
 

## Installing Go
To install Go, see https://golang.org/

For an interactive tour of go's syntax and features, visit https://tour.golang.org

As for an IDE, I use the go plug-in for IntelliJ IDEA.
Your favorite text editor may have a plug-in:
https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

##Run From Source
`go run scrambled.go`


![lets go](http://i.imgur.com/sDBaVEy.png)


