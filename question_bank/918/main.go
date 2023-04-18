package main

// https://leetcode.cn/problems/maximum-sum-circular-subarray/
// 918. 环形子数组的最大和

// 看了佬的解答，发现自己思维还是太局限了
// 环形子数组的最大和，最终组成最大和的只有两种情况
// 1. 和 53 子数组最大和一样
// 2. 组成的子数组在首尾，中间某一部分不在子数组中 A A B B A A (此时A属于最终结果--组成最大和的子数组)
// 此时 1 我们可以按照53的方式正常来
// 2 我们可以考虑，首尾的子数组和最大，其实相当于中间这块子数组和为最小( B B )的总和最小
// 整个数组的总和是固定的，当中间这一块总和最小时，首尾的组成的子数组总和则最大
func maxSubarraySumCircular(nums []int) int {
	// dp[i] = dp[i-1] + nums[i], if dp[i-1]>0 求最大
	// 求最小 dp[i] = dp[i-1] + nums[i]. if dp[i-1] < 0
	big, small, sum := nums[0], nums[0], nums[0]
	biggest, smallest := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		big = max(big+nums[i], nums[i])
		small = min(small+nums[i], nums[i])
		biggest = max(big, biggest)
		smallest = min(small, smallest)
	}
	// 当最大和小于 0 时，说明全是负数，直接返回最大值即可
	if biggest < 0 {
		return biggest
	}
	return max(biggest, sum-smallest)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 环形的子数组
// // 那么dp状态转移时，应该从何处开始？
// // 其实就是每一个地方都能作为一个起点开始
// // dp[i] = dp[i-1]+nums[i], if dp[i-1] >=0
// // dp[i] = nums[i], if dp[i-1] <0
// // 不过存在不少重复计算，时间复杂度也到了 n^2 排除
// func maxSubarraySumCircular(nums []int) int {
// 	res := -30001
// 	n := len(nums)
// 	// 将数组的每个位置作为起始点开始dp
// 	for l := 0; l < n; l++ {
// 		sum, curMax := nums[l], nums[l]
// 		for r := l + 1; r < l+n; r++ {
// 			sum = max(sum+nums[r%n], nums[r%n])
// 			curMax = max(sum, curMax)
// 		}
// 		res = max(curMax, res)
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
