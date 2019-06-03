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

// 计算填充曼哈顿距离表
func calMhDis() {
	for i := 0; i < SQUARE; i++ {
		for j := 0; j < SQUARE; j++ {
			mhDis[i][j] = abs(i/WIDTH-j/WIDTH) + abs(i%WIDTH-j%WIDTH)
		}
	}
}

// 打印曼哈顿距离表
func printMhDis() {
	for i := 0; i < SQUARE; i++ {
		for j := 0; j < SQUARE; j++ {
			fmt.Printf("%v ", mhDis[i][j])
		}
		fmt.Printf("\n")
	}
}

// IDA* 解 puzzle
func solve(s State) (bool, []int) {
	if s.Solved() {
		return true, s.steps[:]
	}

	var stack []State
	min_depth := s.Score()
	//todo 每个深度一个go协程 并发解决
	for limit := min_depth; limit <= MAXSOLVE; limit++ {
		if limit < min_depth {
			limit = min_depth
		}
		min_depth = 1<<32 - 1

		stack = append(stack, s)
		for len(stack) > 0 {
			now := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextSteps := now.NextSteps()
			//fmt.Println(len(nextSteps))
			for _, step := range nextSteps {
				if step == NONE {
					continue
				}
				next := now.NextState(step)
				if next.Solved() {
					return true, next.steps[:]
				}
				//next.Show()

				// 剪枝
				score := next.Score()
				if score < limit {
					stack = append(stack, next)
				} else {
					if score < min_depth {
						min_depth = score
					}
				}
			}
		}
	}

	return false, []int{}
}
