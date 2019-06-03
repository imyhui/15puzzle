package main

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

func main() {
	s := generate()
	s.Show()
}
