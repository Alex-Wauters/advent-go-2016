# Day 5 of Advent of Code 2016 in Go
Day 5: How About a Nice Game of Chess?

## How To Run
The `dist` directory contains the windows and linux (x86 and x64) binaries.
Use the following command line arguments:

`chess -input yourinput -part 1`

See `chess -help` for the default values

## Installing Go
See https://golang.org/

For an interactive tour of go, visit https://tour.golang.org

As for an IDE, I use the golang plug-in for IntelliJ IDEA.
Your favorite text editor may have a go plug-in:
https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

##To Run From Source
`go run chess.go chess_pt2.go`

##To Test
`go test`

##Build yourself

Set `GOOS` and `GOARCH` env vars if you'd like to build for a different environment such as windows/386 or linux/amd64

`go build`


![lets go](http://i.imgur.com/sDBaVEy.png)
