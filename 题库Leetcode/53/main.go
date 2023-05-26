package main

// https://leetcode.cn/problems/maximum-subarray/
// 53. 最大子数组和

// // 滑窗法
// // 连续子数组，和最大
// // 滑窗？窗口性质--窗口内数组和为非负数
// // 当窗口内和为非负数时扩窗口，否则缩窗口
// // 但是感觉操作起来还挺麻烦的
// // 并且，这里显然是可以基于前缀和进一步缩减操作时间，
// // 不过时间复杂度上没有明显提升，就不额外写出了
// func maxSubArray(nums []int) int {
// 	var l, sum int
// 	r := 1
// 	sum, res := nums[0], nums[0]
// 	for r < len(nums) {
// 		for l < r && sum < 0 {
// 			sum -= nums[l]
// 			l++
// 		}
// 		sum += nums[r]
// 		res = max(res, sum)
// 		r++
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // dp
// // 感觉基于dp会更正常
// // dp[i] 定义为 以 nums[i] 结尾时，组成的最大子数组和为多少
// // dp[i] = dp[i-1]+nums[i], if dp[i-1]>0
// // dp[i] = num[i], if dp[i-1] < 0
// // 大于0则基于前数的和求当前和，否则不加上前面的
// // 并且我们可以发现，dp[i]仅取决于dp[i-1]，因此我们不必完全存储全部dp数组
// // 只需存储前一个状态即可
// func maxSubArray(nums []int) int {
// 	var sum int
// 	res := nums[0]
// 	for _, num := range nums {
// 		// if sum < 0 {
// 		// 	sum = 0
// 		// }
// 		// sum += num
// 		sum = max(sum+num, num)
// 		res = max(res, sum)
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
