package main

// https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 29. 顺时针打印矩阵

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	l, t := 0, 0
	r, d := len(matrix[0]), len(matrix)
	length := r * d
	r--
	d--
	res := make([]int, 0, length)
	count := 0
	for {
		for x := l; x <= r; x++ {
			res = append(res, matrix[t][x])
			count++
		}
		if count == length {
			break
		}
		t++
		for y := t; y <= d; y++ {
			res = append(res, matrix[y][r])
			count++
		}
		if count == length {
			break
		}
		r--
		for x := r; x >= l; x-- {
			res = append(res, matrix[d][x])
			count++
		}
		if count == length {
			break
		}
		d--
		for y := d; y >= t; y-- {
			res = append(res, matrix[y][l])
			count++
		}
		if count == length {
			break
		}
		l++
	}
	return res
}
