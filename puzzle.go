package main

import (
	"flag"
	"github.com/imyhui/15puzzle/puzzle"
)



func main() {

	flag.Parse()
	if flag.Arg(0) == "server" {
		puzzle.RunServer()
	} else {
		puzzle.RunShell()
	}
}
