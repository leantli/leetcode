package main

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 62. 圆圈中最后剩下的数字

// 逆着去补位
// 我们知道最后的剩下的数，此时在数组中必定是 0 下标
// 逆折往前推 idx = (idx+m)%i (i 为这一轮数组的长度)
// 如此到 i == n 的轮次时，我们得到该长度下 原始数 在初始数组中的下标
// 而在初始数组中，下标==数
func lastRemaining(n int, m int) int {
	// 逆着推，最后一个数，最后时在数组中下标为 0
	var idx int
	// i 代表当前数组的长度, 从 2 开始，逐步推到初始长度 n
	for i := 2; i <= n; i++ {
		idx = (idx + m) % i
	}
	return idx
}

// // 超时。。。
// // 模拟做法好吧
// func lastRemaining(n int, m int) int {
// 	nn := make([]int, n)
// 	for i := range nn {
// 		nn[i] = i
// 	}
// 	var idx int
// 	for len(nn) != 1 {
// 		idx = (idx + m - 1) % len(nn)
// 		nn = append(nn[:idx], nn[idx+1:]...)
// 	}
// 	return nn[0]
// }
