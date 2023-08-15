package main

// https://leetcode.cn/problems/last-stone-weight-ii/description/
// 1049. 最后一块石头的重量 II

// 类似于切分成等和子集？如果 sum 为偶数，显然可以没有石头剩下，否则剩下的石头大小为 (sum-dp[sum/2])-dp[sum/2]
// 为什么是 (sum-dp[sum/2])-dp[sum/2]？因为本问题其实就是等同于将数组切分成和最相近的两个子集，最终求其差值
// 和最相近的话，显然较小的那个，和一定小于等于 sum/2，通过dp计算出来则为 dp[sum/2]，较大的子集和则为 sum-dp[sum/2]
func lastStoneWeightII(stones []int) int {
	// 01 背包优化滚动数组 dp[j] = max(dp[j], dp[j-weight[i]]+value[i])，本题 weight 和 value 同意
	// 背包容量 j，最大为 sum/2，物品即为 stones
	var sum int
	for i := range stones {
		sum += stones[i]
	}
	dp := make([]int, sum/2+1)
	for i := range stones {
		for j := sum / 2; j >= stones[i]; j-- {
			if dp[j-stones[i]]+stones[i] > dp[j] {
				dp[j] = dp[j-stones[i]] + stones[i]
			}
		}
	}
	return sum - dp[sum/2]*2
}
