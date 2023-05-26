package main

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 57-2. 和为s的连续正数序列

// 首先关注，需要连续序列
// 感觉就双指针枚举？分别指向下标，不过这样感觉也要 O(n^2)，耗时太长了
// func findContinuousSequence(target int) [][]int {
// 	res := make([][]int, 0)
// 	for l := 1; l < target; l++ {
// 		sum := l
// 		temp := make([]int, 0)
// 		temp = append(temp, l)
// 		for r := l + 1; r < target; r++ {
// 			sum += r
// 			temp = append(temp, r)
// 			if sum == target {
// 				res = append(res, temp)
// 			}
// 			if sum > target {
// 				break
// 			}
// 		}
// 	}
// 	return res
// }

// 没有第一时间想到滑动窗口，需要再加强一下滑动窗口的触发点记忆，连续，有序，求和
// 先维持一个窗口，当 sum 小于 target 时就右移 r
// 大于 target 时 右移 l
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	var sum int
	list := make([]int, 0)
	for i := 1; i < target; i++ {
		list = append(list, i)
		sum += i
		for sum > target {
			sum -= list[0]
			list = list[1:]
		}
		if sum == target {
			res = append(res, append([]int{}, list...))
		}
	}
	return res
}
