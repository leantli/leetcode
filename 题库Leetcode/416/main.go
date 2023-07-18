package main

// https://leetcode.cn/problems/partition-equal-subset-sum/
// 416. 分割等和子集

// 等和子集
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum & 1) == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for i := target; i >= num; i-- {
			if dp[i] < dp[i-num]+num {
				dp[i] = dp[i-num] + num
			}
		}
	}
	return dp[target] == target
}
