package main

// https://leetcode.cn/problems/two-sum/?favorite=2cktkvj
// 1. 两数之和

// 最最基础的方法，就是 n^2 的遍历去判断是否能得到 target
// 但显然我们不能这样
// 这里可以借助 map，在遍历时判断是否满足 sum = target，一轮遍历解决

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		// 如果 map 中存在值，与 num 相加能得到 target，则直接返回结果
		if _, ok := m[target-num]; ok {
			return []int{m[target-num], i}
		}
		m[num] = i
	}
	return []int{}
}
