package main

import (
	"fmt"
	"math/rand"
)

// 每一个局面算一个状态
type State struct {
	board []int
	block int

	steps [MAXSOLVE]int
	depth int
}

// 找到滑块位置
func (s State) Block() int {
	for i := 0; i < SQUARE; i++ {
		if s.board[i] == 0 {
			return i
		}
	}
	return -1
}

// 命令行打印当前棋盘
func (s State) Show() {
	board := s.board
	fmt.Println("------------")
	for i := 0; i < WIDTH; i++ {
		fmt.Printf("|%2v|%2v|%2v|%2v|\n", board[i*WIDTH], board[i*WIDTH+1], board[i*WIDTH+2], board[i*WIDTH+3])
	}
	fmt.Println("-------------")
}

// 判断是否可解
func (s State) SolveAble() bool {
	var sum, row int
	for i := 0; i < SQUARE; i++ {
		num := s.board[i]
		// 记录滑块行号
		if num == 0 {
			row = i/WIDTH + 1
			continue
		}
		// 计算逆序数和
		for j := i; j < SQUARE; j++ {
			if s.board[j] < num && s.board[j] != 0 {
				sum++
			}
		}
	}
	return (sum+row)%2 == 0
}

//判断是否已解
func (s State) Solved() bool {
	if s.board[SQUARE-1] != 0 {
		return false
	}
	for i := 0; i < SQUARE-1; i++ {
		if s.board[i] != i+1 {
			return false
		}
	}
	return true
}

// 调整使局面可解
func (s State) Adjust() {
	if s.SolveAble() {
		return
	}
	//交换任意两个不为0的数字 原理: 对换改变排列逆序数
	var fst, scd int
	fst = rand.Intn(SQUARE)
	if s.board[fst] == 0 {
		fst = (fst + 1) % SQUARE
	}

	for {
		scd = rand.Intn(SQUARE)
		if s.board[scd] == 0 {
			scd = (scd + 1) % SQUARE
		}
		if scd != fst {
			break
		}
	}
	s.board[fst], s.board[scd] = s.board[scd], s.board[fst]
}
