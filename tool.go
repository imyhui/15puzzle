package main

import (
	"math/rand"
	"time"
)

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

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
