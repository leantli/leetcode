package main

// https://leetcode.cn/problems/jump-game-ii/
// 45. 跳跃游戏 II

// 每一个测试用例都能到达最后下标，计算最少多少步
// 计算出当前位置能够到达的最远距离，接着只要还没到当前能到达的最远距离，就不断更新下一个能到达的最远距离
// 如果到达了走一步能走到的最远距离时，则更新步数
// curMax := 0, nextMax := 0 + nums[0], 此时如果 len(nums) > 0，显然 i == curMax, 这个时候计算步数 cnt++
// 接着更新 curMax = nextMax, 然后还未到达新的 curMax，就正常继续更新 nextMax
// 由于到达最后一个下标就不需要再计算 cnt++ 了，因此遍历 i 时，直接遍历到 len(nums)-2 即可
func jump(nums []int) int {
	var curMax, nextMax, cnt int
	for i := 0; i < len(nums)-1; i++ {
		// nextMax 其实都是在计算下一步能覆盖的范围
		// 但只要 i 还小于 curMax，i 遍历的位置其实都是当前步数就可以到达的位置
		// 因此等 i==curMax，才对步数+1，获得最少步数
		nextMax = max(nextMax, i+nums[i])
		if i == curMax {
			cnt++
			curMax = nextMax
		}
	}
	return cnt
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
