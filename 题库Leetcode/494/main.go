package main

// https://leetcode.cn/problems/target-sum/description/
// 494. 目标和

// 其实这里也是将数组拆分成两个子集，相减后值为 target 则满足
// 分割等和子集是判断能不能分成两个相等的子集；最后一块石头2是求分成两个最相近子集后差值
// 这里则是求分成两个子集后差值为 target 的拆分方法有多少种
// add - sub = target ; add + sub = sum
// 将两式相加， 2 * add = target + sum, add = (target+sum)/2
// 此时我们则计算 add 达到背包容量 (target+sum)/2 的个数有多少个
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	// 特殊判断，target 越界肯定无法拆分
	if target > sum || target < -sum {
		return 0
	}
	// 当 target+sum /2 需要向下取整时无法拆分，说明所求的 left 子集和为小数，显然凑不出来
	if (target+sum)&1 == 1 {
		return 0
	}
	t := (target + sum) / 2
	dp := make([]int, t+1)
	// dp[j] = dp[j], if j < nums[i]
	// dp[j] = dp[j] + dp[j-nums[i]], if j >= nums[i]
	// 初始化， dp[0] 必须等于 1，否则后续无法状态转移
	dp[0] = 1
	for i := range nums {
		for j := t; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[t]
}

// // 回溯枚举？每个数值有两种操作方式，每个数有两种操作方式，加或者减，当所有数都操作完后，只要结果为 target，则计数+1；时间复杂度较高
// func findTargetSumWays(nums []int, target int) int {
// 	var cnt int
// 	var dfs func(index, sum int)
// 	dfs = func(index, sum int) {
// 		if index == len(nums) {
// 			if sum == target {
// 				cnt++
// 			}
// 			return
// 		}
// 		dfs(index+1, sum+nums[index])
// 		dfs(index+1, sum-nums[index])
// 	}
// 	dfs(0, 0)
// 	return cnt
// }
