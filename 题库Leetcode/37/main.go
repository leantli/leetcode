package main

import "fmt"

// https://leetcode.cn/problems/sudoku-solver/description/
// 37. 解数独

func main() {
	fmt.Println('1')
	fmt.Println(byte(1))
	j := 10
	fmt.Println(byte(j))
}

func solveSudoku(board [][]byte) {
	var dfs func() bool
	nums := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	dfs = func() bool {
		// 每一次都遍历全部的 i，j，将未填的数字进行枚举填写(只要符合要求)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				// 已有数字，直接略过
				if board[i][j] != '.' {
					continue
				}
				// 没有数字，枚举进行填写
				for k := 1; k <= 9; k++ {
					// 放置数字成功则 dfs 进行下一个数字的枚举，成功则直接返回，失败则将当前数字撤去，准备枚举下一个数字
					if canPlace(i, j, nums[k-1], board) {
						board[i][j] = nums[k-1]
						if dfs() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		// 当成功通过两个循环没有返回 true 和 false，说明所有数都填完了，直接返回 trues
		return true
	}
	dfs()
}

func canPlace(row, col int, k byte, board [][]byte) bool {
	// 行上还没有当前的数
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	// 列上还没有当前的数
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}
	// 先求出当前位置所属方块的起始行列位置
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	// 在判断方块内不存在填入的数字 k
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
