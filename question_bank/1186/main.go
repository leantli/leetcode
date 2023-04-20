package main

import "math"

// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
// 1186. 删除一次得到子数组最大和

// 返回非空子数组最大元素和
// 选出一个子数组，该子数组是所有子数组中最大的(可以删一个，相当于允许间隔一个数？)
// 相比较 53， 多了一个删除操作，多了这个操作后，我们需要区分 未删除/删除 两个子状态
// 如果不考虑删除，其 dp 和 53 一致， dp[i] = max(dp[i-1] + nums[i], nums[i])
// 如果考虑删除，deleted[i]要么是删除当前数( deleted[i]=dp[i-1] )
// 要么就是之前已经删除过了，只要加上当前数 ( deleted[i-1]+nums[i] )
// 因此 deleted[i] = max(deleted[i-1]+nums[i], dp[i-1])
// 此时我们又关注到，虽然有两个状态，但是每次都只需要 i 和 i-1，因此我们可以简单用两个局部变量替代，减少空间复杂度
func maximumSum(arr []int) int {
	// 初始化
	dp, deleted, res := arr[0], 0, arr[0]
	// 状态更新时因为 deleted 依赖 dp 之前的值，所以先更新 deleted ，再更新 dp
	for _, num := range arr[1:] {
		deleted = max(deleted+num, dp)
		dp = max(dp+num, num)
		res = max(res, dp, deleted)
	}
	return res
}

func max(arr ...int) int {
	res := math.MinInt
	for _, num := range arr {
		if res < num {
			res = num
		}
	}
	return res
}
