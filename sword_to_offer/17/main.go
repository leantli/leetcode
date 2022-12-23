package main

import "math"

// https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 17. 打印从1到最大的n位数

func printNumbers(n int) []int {
	max := int(math.Pow10(n))
	res := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		res = append(res, i)
	}
	return res
}
