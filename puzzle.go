package main

import (
	"fmt"
	"time"
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
	start := time.Now()

	s := generate()
	s.board = []int{
		5, 1, 10, 12,
		9, 11, 13, 0,
		4, 7, 14, 8,
		15, 3, 6, 2,
	}
	//s.board = []int{12,15,0,6,2,5,1,14,8,4,7,13,10,11,9,3}
	s.block = s.Block()

	s.Show()
	if s.SolveAble() {
		fmt.Println(s.board)
		fmt.Println("该 puzzle 有解")
		fmt.Printf("解为: %v\n", s.Solution())
	} else {
		fmt.Println("该 puzzle 无解")
		s.Adjust()
		s.Show()
		if s.SolveAble() {
			fmt.Println("该 puzzle 有解了")
			fmt.Println(s.board)
			fmt.Printf("解为: %v\n", s.Solution())
		} else {
			fmt.Println("该 puzzle 依然无解")
		}
	}
	cost := time.Since(start)
	fmt.Printf("\ncost=[%s]", cost)
}
