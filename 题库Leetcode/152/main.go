package main

// https://leetcode.cn/problems/maximum-product-subarray/
// 152. 乘积最大子数组

// 连续子数组，最大乘积
// 秒想到不定长滑窗和dp
// 但是这个滑窗的扩和缩，想不到对应的性质
// 因此还是采用dp比较合适
// dp[i] 定义为 以nums[i]为结尾的连续子数组的最大乘积
// dp[i] = max(nums[i], dp[i-1]*nums[i])
// 但是这个转移方程有个问题，遇到 5 -5 -10 呢
// 此时得到的 dp 为 5 -5 50， 但其实更好的应该是 5 -25 250
// 那么符号要如何考虑呢，即，相乘时，两方有一方为负数时，越低越好
// 两方都为负数或整数时，越大越好？
// 但是这样也不好考虑，因为这样只能考虑到两个数，但是可能存在 -2 10 10 -900 这类
// 按刚才的思路就肯定凉了
// 想来想去没想到好办法，看了官解....
// 发现是同时维护两个状态数组，一个求 max，一个求min
// 并且我们只要最大值，min只是用来保存是否存在一个负负得正的情况
func maxProduct(nums []int) int {
	res, big, small := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		b, s := big, small
		big = max(nums[i], b*nums[i], s*nums[i])
		small = min(nums[i], b*nums[i], s*nums[i])
		res = max(res, big)
	}
	return res
}

// 求最大值
func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

// 求最小值
func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
