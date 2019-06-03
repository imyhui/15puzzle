package main

import "fmt"

// 每一个局面算一个状态
type State struct {
	board []int
	block int

	steps [MAXSOLVE]int
	depth int
}

func (s State) Block() int {
	for i := 0; i < SQUARE; i++ {
		if s.board[i] == 0 {
			return i
		}
	}
	return -1
}

func (s State) Show() {
	board := s.board
	fmt.Println("------------")
	for i := 0; i < WIDTH; i++ {
		fmt.Printf("|%2v|%2v|%2v|%2v|\n", board[i*WIDTH], board[i*WIDTH+1], board[i*WIDTH+2], board[i*WIDTH+3])
	}
	fmt.Println("-------------")
}
