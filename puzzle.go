package main

import "fmt"

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
	mhdis [SQUARE][SQUARE]int
)

func main() {
	s := generate()
	s.Show()
	if s.SolveAble() {
		fmt.Println("该 puzzle 有解")
	} else {
		fmt.Println("该 puzzle 无解")
		s.Adjust()
		s.Show()
		if s.SolveAble() {
			fmt.Println("该 puzzle 有解了")
		} else {
			fmt.Println("该 puzzle 依然无解")
		}
	}
}
