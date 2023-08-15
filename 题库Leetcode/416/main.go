package main

// https://leetcode.cn/problems/partition-equal-subset-sum/
// 416. 分割等和子集

// // 首先和一定要为偶数，否则必定无法分成两个子集元素和相等
// // 其次，sort 之后双指针？应该不行，因为是求子集而不是两个数，那么回溯寻找子集和为 sum/2 ？时间复杂度显然会过高
// // 此时考虑是否能够使用 dp 背包思想，每个数，就是一个物品，容量最大为 sum/2，判断取与不取每个数，最后能否满足放进背包的数值和为 sum/2
// // 常规 01 背包思路
// func canPartition(nums []int) bool {
// 	var sum int
// 	for i := range nums {
// 		sum += nums[i]
// 	}
// 	// 如果和不为偶数，直接返回
// 	if (sum & 1) == 1 {
// 		return false
// 	}
// 	dp := make([][]int, len(nums)) // dp 的行，为物品的重量, nums[i]
// 	for i := range dp {
// 		dp[i] = make([]int, sum/2+1) // dp 的列，为背包的最大负重，为 sum/2
// 	}
// 	// 首先第一列，负重为 0 时，不管什么物品，都一定不能放，所以 dp[i][0] 都为 0，无需管
// 	// 其次第一行，负重大于 nums[0] 的负重，都可以放 nums[0]，只要 nums[0] 不超过 sum/2
// 	if nums[0] <= sum/2 {
// 		for i := nums[0]; i <= sum/2; i++ {
// 			dp[0][i] = nums[0]
// 		}
// 	}
// 	// 此时为 01 背包，状态转移方程为 dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
// 	// 此处 nums[i] 既为 weight 也为 value，所以当 j >= nums[i] 时 dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
// 	// 否则 dp[i][j] = dp[i-1][j]
// 	for i := 1; i < len(nums); i++ {
// 		for j := 0; j <= sum/2; j++ {
// 			if j >= nums[i] {
// 				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
// 			} else {
// 				dp[i][j] = dp[i-1][j]
// 			}
// 		}
// 	}
// 	return dp[len(nums)-1][sum/2] == sum/2
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 01背包状态转移方程为：dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
// 显然我们能够优化成滚动数组，采用一位数组即可 dp[j] 表示负重最大为 j 的背包最多能装多少货物，使得价值最大
// 此时 dp[j] = max(dp[j], dp[j-weight[i]]+value[i])，此时 每次遍历 j，都从后往前，因为 dp[j] 依赖于前面的 dp[j]，从后往前不会导致语义出现问题
// 背包最大负重为 sum/2，并且 sum 不能为奇数，否则不能切分成两个和相等的子集
func canPartition(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if (sum & 1) == 1 {
		return false
	}
	dp := make([]int, sum/2+1)
	// 以下为滚动数组，当 j < nums[i] 时，显然 dp[j] = dp[j]，相当于复用上一层的 dp[j]，即未选取当前的物品 i
	for i := range nums {
		for j := sum / 2; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}
	return dp[sum/2] == sum/2
}
