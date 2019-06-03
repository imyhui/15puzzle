package main

import (
	"flag"
)

const (
	WIDTH    = 4
	SQUARE   = WIDTH * WIDTH
	MAXSOLVE = 70
	LEFT     = -1
	RIGHT    = 1
	UP       = -WIDTH
	DOWN     = WIDTH
	NONE     = 0
)

var (
	mhDis [SQUARE][SQUARE]int
	move  = map[int]string{-1: "L", 1: "R", -WIDTH: "U", WIDTH: "D"}
)

func init() {
	calMhDis()
}

func main() {

	flag.Parse()
	if flag.Arg(0) == "server" {
		runServer()
	} else {
		runShell()
	}

}
