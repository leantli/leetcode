package main

// https://leetcode.cn/problems/jump-game/description/
// 55. 跳跃游戏

// 只问我们能不能达到最后一个下标
// 那我们可以遍历每一个 nums[i]，并不断更新当前能够到达的下标位置
// 当能够到达的下标位置大于等于 len(nums)-1 时，直接返回 true 即可
// 如果遍历到能够到达的最后一个下标位置，还未等于 len(nums)-1，则返回 false
func canJump(nums []int) bool {
	var right int
	for i := 0; i <= right; i++ {
		if i+nums[i] > right {
			right = i + nums[i]
		}
		if right >= len(nums)-1 {
			return true
		}
	}
	return false
}
