package main

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 10-2 青蛙跳台阶

// 动态规划
// 第 n 级的台阶的跳法数量 = (n-1)级的台阶跳法数量 + (n-2)级台阶跳法数量
// 等价于斐波那契
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	res, pre, post := 0, 1, 1
	for i := 1; i < n; i++ {
		res = (pre + post) % (1e9 + 7)
		pre = post
		post = res
	}
	return res
}
