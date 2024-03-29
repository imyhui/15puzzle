package puzzle

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
	var cnt int
	// 计算逆序数和
	for i := 0; i < SQUARE; i++ {
		num := s.board[i]
		if num != 0 {
			for j := i; j < SQUARE; j++ {
				if s.board[j] < num && s.board[j] != 0 {
					cnt++
				}
			}
		} else {
			cnt += i/WIDTH + 1
		}
	}
	return cnt%2 == 0
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

// 计算该局面f(n),h(n) 为曼哈顿距离, *4/3 可大幅度缩短时间,但会使解变长
func (s State) Score() int {
	mhd := 0

	for i := 0; i < SQUARE; i++ {
		if s.board[i] != 0 {
			mhd += mhDis[s.board[i]-1][i]
		}
	}
	return s.depth + 4*mhd/3
}

// 获取本状态可移动的列表
func (s State) NextSteps() [4]int {
	var lastStep int
	nextSteps := [4]int{}
	if s.depth > 0 {
		lastStep = s.steps[s.depth-1]
	}

	if s.block%WIDTH > 0 && lastStep != RIGHT {
		nextSteps[0] = LEFT
	}
	if s.block%WIDTH < (WIDTH-1) && lastStep != LEFT {
		nextSteps[1] = RIGHT
	}
	if s.block/WIDTH > 0 && lastStep != DOWN {
		nextSteps[2] = UP
	}
	if s.block/WIDTH < (WIDTH-1) && lastStep != UP {
		nextSteps[3] = DOWN
	}
	return nextSteps
}

// 获取经过step移动后的状态
func (s State) NextState(step int) State {
	next := State{
		steps: s.steps,
		depth: s.depth + 1,
		block: s.block + step,
	}
	// fixme: 传拷贝 否则会改变原切片
	next.board = make([]int, 16)
	copy(next.board, s.board)

	next.steps[s.depth] = step
	next.board[s.block] = next.board[next.block]
	next.board[next.block] = 0
	return next
}

// 对当前状态求解
func (s State) Solution() (string, []int) {
	var result string
	if ok, steps := solve(s); ok {
		for _, step := range steps {
			result += move[step]
		}
		return result, steps
	}
	return "", []int{}
}
