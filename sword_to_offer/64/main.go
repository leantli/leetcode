package main

// https://leetcode.cn/problems/qiu-12n-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 求 1+2+ ... + n

// for 循环用 递归 替代
// if 用 短路与替代

func sumNums(n int) int {
	var ans int
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
