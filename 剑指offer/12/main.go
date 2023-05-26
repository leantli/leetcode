package main

// https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 12. 矩阵中的路径

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helper(board, []byte(word), j, i, 0) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, word []byte, x, y, index int) bool {
	// 先判断边界
	if x < 0 || x >= len(board[0]) || y < 0 || y >= len(board) {
		return false
	}
	// 不等则直接返回
	if board[y][x] != word[index] {
		return false
	}
	// 相等的话原地修改一下，标识已走过，后面再复原
	temp := board[y][x]
	board[y][x] = '-'
	index++
	if index == len(word) {
		return true
	}
	res := helper(board, word, x+1, y, index) ||
		helper(board, word, x, y+1, index) ||
		helper(board, word, x-1, y, index) ||
		helper(board, word, x, y-1, index)
	board[y][x] = temp
	return res
}
