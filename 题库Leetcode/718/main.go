package main

// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
// 718. 最长重复子数组

// 其实我们也可以注意到，dp[i][j] 只由 dp[i-1][j-1] 和 nums1[i-1] == nums2[j-1] 决定
// 我们完全不需要维系一个完整的二维数组，采用滚动数组进行优化降维，变成 dp[j] = dp[j-1], if nums1[i-1] == nums2[j-1]
// 并且由于 j 依赖于 j-1，我们需要从后往前遍历，避免覆盖，而 i 的遍历仍然从 1 到 n，因为每次 i 都是依赖于 i-1 的
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	var res int
	dp := make([]int, m+1)
	for i := 1; i <= n; i++ {
		for j := m; j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}
			if res < dp[j] {
				res = dp[j]
			}
		}
	}
	return res
}

// // 显然暴力的时间复杂度比较高，每次从 nums1 选取出对应的数后，要在 nums2 中找到相等的数，再两个指针共同右移获取最长公共长度
// // 考虑最终目的是求 两个数组公共的、长度最长的子数组长度
// // 定义 dp[i][j] 表示以 nums1[i-1] 为结尾的子数组和 nums2[j-1] 为结尾的子数组，最长的公共子数组长度是多少
// // dp[i][j] = dp[i-1][j-1]+1, if nums1[i-1] == nums2[j-1]
// // 而显然，并不一定是以 dp[n][m] 为结尾的最长公共子数组长度就是答案，任何一个 dp[i][j] 都有可能，因此我们需要用另一个变量存储结果
// func findLength(nums1 []int, nums2 []int) int {
//     var res int
//     n, m := len(nums1),len(nums2)
//     dp := make([][]int, n + 1)
//     for i := range dp {
//         dp[i] = make([]int, m + 1)
//     }
//     for i := 1; i <= n; i++ {
//         for j := 1; j <= m; j++ {
//             if nums1[i-1] == nums2[j-1] {
//                 dp[i][j] = dp[i-1][j-1] + 1
//             }
//             if dp[i][j] > res {
//                 res = dp[i][j]
//             }
//         }
//     }
//     return res
// }

// // 暴力的话，就是遍历 nums1，取nums1[i] 在 nums[2] 中遍历，遇到相等的之后，左右指针共同右移，记录最长的即可
// func findLength(nums1 []int, nums2 []int) int {
// 	var res int
// 	for i := 0; i < len(nums1); i++ {
// 		for j := 0; j < len(nums2); j++ {
// 			var cnt int
// 			if nums1[i] == nums2[j] {
// 				tempI, tempJ := i, j
// 				for tempI < len(nums1) && tempJ < len(nums2) && nums1[tempI] == nums2[tempJ] {
// 					tempI++
// 					tempJ++
// 					cnt++
// 				}
// 				if res < cnt {
// 					res = cnt
// 				}
// 			}
// 		}
// 	}
// 	return res
// }
