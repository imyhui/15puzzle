package main

import (
	"fmt"
	"math/rand"
	"os"
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

// 写文件
func writeFile() {
	f, _ := os.OpenFile("solve.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0755)
	os.Stdout = f
}

// 打印求解过程
func printSteps(s State, steps []int) {
	start := State{
		board: s.board,
		block: s.block,
	}

	for _, step := range steps {
		if step == NONE {
			break
		}
		nBlock := start.block + step
		start.board[start.block], start.board[nBlock] = start.board[nBlock], 0
		start.block = nBlock
		start.Show()
	}
}

// 命令行模式
func runShell() {
	fmt.Println("正在求解...,详情见solve.txt")
	writeFile()
	start := time.Now()

	s := generate()

	var (
		res string
		stp []int
	)
	s.Show()
	if s.SolveAble() {
		fmt.Println(s.board)
		res, stp = s.Solution()
		fmt.Println("该 puzzle 有解")
		fmt.Printf("解为: %v\n", res)
	} else {
		fmt.Println("该 puzzle 无解")
		s.Adjust()
		s.Show()
		if s.SolveAble() {
			fmt.Println("该 puzzle 有解了")
			fmt.Println(s.board)
			res, stp = s.Solution()
			fmt.Printf("解为: %v\n", res)
		} else {
			fmt.Println("该 puzzle 依然无解")
		}
	}
	cost := time.Since(start)
	fmt.Printf("\ncost=[%s]\n", cost)
	printSteps(s, stp)
}

// 服务端模式
func runServer() {
	server()
}
