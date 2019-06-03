package main

import (
	"fmt"
	"math/rand"
	"time"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// 生成一个初始状态
func generate() State {
	var s State
	var board []int
	for i := 0; i < SQUARE; i++ {
		board = append(board, i)
	}
	// Fisher-Yates Shuffle
	rand.Seed(time.Now().Unix())
	for i := len(board) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		board[i], board[num] = board[num], board[i]
	}

	s.board = board
	s.block = s.Block()
	return s
}

func calMhDis() {
	for i := 0; i < SQUARE; i++ {
		for j := 0; j < SQUARE; j++ {
			mhDis[i][j] = abs(i/WIDTH-j/WIDTH) + abs(i%WIDTH-j%WIDTH)
		}
	}
}
func printMhDis() {
	for i := 0; i < SQUARE; i++ {
		for j := 0; j < SQUARE; j++ {
			fmt.Printf("%v ", mhDis[i][j])
		}
		fmt.Printf("\n")
	}
}
func solve(s State) (bool, []int) {
	if s.Solved() {
		return true, s.steps[:]
	}

	return false, []int{}
}
