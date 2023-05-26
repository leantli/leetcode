package main

// https://leetcode.cn/problems/diet-plan-performance/
// 1176. 健身计划评估

// 定长滑动窗口，窗口长度固定为 k
// 窗口始终维护一个 T 值
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	var T int
	var score int
	for _, calory := range calories[:k] {
		T += calory
	}
	if T < lower {
		score -= 1
	} else if T > upper {
		score += 1
	}
	for i := k; i < len(calories); i++ {
		T += calories[i] - calories[i-k]
		if T < lower {
			score -= 1
		} else if T > upper {
			score += 1
		}
	}
	return score
}
