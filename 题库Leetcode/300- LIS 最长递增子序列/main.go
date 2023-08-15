package main

// https://leetcode.cn/problems/longest-increasing-subsequence/
// 300. 最长递增子序列

// 二刷
// 也可以基于贪心+二分，贪心的思路就是，维系一个 tail 数组
// tail[i] 表示长度为 i+1 的递增子序列最后一位数最小是多少
// 然后一次遍历，基于二分去更新 tail 数组，如果遇到非常大的，比 tail[len(tail)-1] 还大，则直接添加到尾部
// 这样能够保证，最终答案：最长严格递增子序列的长度就是 len(tail)
func lengthOfLIS(nums []int) int {
	tail := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
			continue
		}
		l, r := -1, len(tail)
		for l+1 != r {
			mid := l + (r-l)/2
			if tail[mid] < nums[i] {
				l = mid
			} else {
				r = mid
			}
		}
		// l 落在最后一个小于 nums[i] 的位置，r 落在第一个大于等于 nums[i] 的位置
		// 如果 nums[i] 不是特别大，直接更新 r 下标对应的值
		tail[r] = nums[i]
	}
	return len(tail)
}

// // 二刷
// // 要求，遍历到最后时，确定最长的严格递增子序列的长度是多少
// // 设 dp[i] 为以 nums[i] 为结尾时，最长的严格递增子序列长度是多少
// // dp[i] = max(dp[i], dp[j]+1), if nums[i] > nums[j], j < i
// func lengthOfLIS(nums []int) int {
//     dp := make([]int, len(nums))
//     res := 1
//     for i := range dp {
//         dp[i] = 1
//     }
//     for i := 1; i < len(nums); i++ {
//         for j := 0; j < i; j++ {
//             if nums[i] > nums[j] {
//                 dp[i] = max(dp[i], dp[j]+1)
//             }
//         }
//         res = max(res, dp[i])
//     }
//     return res
// }
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// // 第二次做这道题，先考虑常规dp
// // 找最长递增子序列的长度，子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序
// // 定义 dp[i] 为 以 nums[i] 为结尾时，最长递增子序列的长度
// // dp[i] = max{dp[i], 前面所有比nums[i]小的num结尾的长度(dp[j])+1}
// func lengthOfLIS(nums []int) int {
// 	res := 1
// 	// 初始化，所有dp[i]初始都为单独一个数，长度为1
// 	// 这里一定要全部初始化，而不能值初始化dp[0]=1，否则dp[i]=dp[j]+1时可能出现误差
// 	dp := make([]int, len(nums))
// 	for i := range nums {
// 		dp[i] = 1
// 	}
// 	for i := 1; i < len(nums); i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		res = max(dp[i], res)
// 	}
// 	return res
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 再考虑贪心+dp+二分
// // 定义 tail[i] 为长度为i+1的情况下，末尾的数为多少
// // 初始化 tail[0]=nums[0]，表示长度为1的情况下，末尾的数为nums[0]
// // 一次遍历过去，遇到比tail[end]大的数，直接append到tail末尾，成为新的end
// // 如果不比end大，则二分查找tail数组，找到刚好小于它的，替换掉这个数右侧的数字
// // 当然，tail数组本身并不代表LIS，因为中间的数部分是最新的，显然无法与其右侧的数构成LIS
// // 它算是一个严格上升的状态数组，我们能够基于这个状态数组，求出LIS的长度
// // 这里可以自行论证单调性，最终求出来tail数组的长度即为最长递增子序列的长度
// // 比如 1 2 5 3 4
// func lengthOfLIS(nums []int) int {
// 	tail := make([]int, 0)
// 	tail = append(tail, nums[0])
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] > tail[len(tail)-1] {
// 			tail = append(tail, nums[i])
// 			continue
// 		}
// 		l, r := -1, len(tail)
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if tail[mid] < nums[i] {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		tail[r] = nums[i]
// 	}
// 	return len(tail)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 最后总结一下，该题的两种方法
// // 基础 dp，dp[i] 表示以 nums[i] 为结尾时的最长长度递增子序列
// // 贪心+二分 dp, dp[i] 表示 长度为 i+1 的所有上升子序列的结尾的最小值

// // 最长长度 严格递增 子序列
// // dp？
// // 但这里的状态分析却不好以最后一个状态开始考虑？如何解决？(看到了下面这一段分析，感觉确实不错)
// // 首先考虑题目问什么，就把什么定义成状态。题目问最长上升子序列的长度，其实可以把「子序列的长度」定义成状态，但是发现「状态转移」不好做。
// // 基于「动态规划」的状态设计需要满足「无后效性」的设计思想，可以将状态定义为「以 nums[i] 结尾 的「上升子序列」的长度」。
// // 「无后效性」的设计思想：让不确定的因素确定下来，以保证求解的过程形成一个逻辑上的有向无环图。这题不确定的因素是某个元素是否被选中，
// // 而我们设计状态的时候，让 nums[i] 必需被选中，这一点是「让不确定的因素确定下来」，也是我们这样设计状态的原因。
// // 链接：https://leetcode.cn/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/

// // dp[i] 是 以 nums[i] 为结尾时的最长长度递增子序列
// // 初始化，每个 dp[i] 默认都是 1，因为所有 nums[i] 都能以自身元素作为子序列
// // dp[i] = {前面的nums[j]小于当前的nums[i],则dp[i]=max(dp[j]+1,dp[i])}
// func lengthOfLIS(nums []int) int {
// 	ans := 1
// 	n := len(nums)
// 	// 初始化
// 	dp := make([]int, n, n)
// 	for i := 0; i < n; i++ {
// 		dp[i] = 1
// 	}
// 	for i := 1; i < n; i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] {
// 				dp[i] = max(dp[j]+1, dp[i])
// 				ans = max(dp[i], ans)
// 			}
// 		}
// 	}
// 	return ans
// }

// // 看到还有 贪心+二分 的解法，但是思路看了好久，觉得自己真的想不到，就算做了笔记，感觉很可能以后还是想不到
// // 但还是做一些笔记和思路解析，尽可能让自己理解
// // 记住 最长递增子序列（Longest Increasing Subsequence，简写 LIS）是非常经典的一个算法问题

// // 换一个 状态定义，虽然不知道为什么能想到要换一个 状态定义
// // 依然着眼于某个上升子序列的 **结尾的元素**，如果 **已经得到的上升子序列的结尾的数越小，那么遍历的时候后面接上一个数，会有更大的可能构成一个长度更长的上升子序列**
// // 既然结尾越小越好，我们可以记录 在长度固定的情况下，结尾最小的那个元素的数值
// //
// // 因此定义新状态 tail[i] 表示 长度为 i+1 的所有上升子序列的结尾的最小值
// // 注意区分，基本dp方法中， dp[i]  表示 以 nums[i] 为结尾时的最长长度递增子序列
// //
// // 但这里我们还需要注意，tail 数组并不一定是 实际上的 LIS，它只是用来求解 LIS 问题的状态数组
// // tail[0] 表示 长度为 1 的所有上升子序列中，结尾最小的元素的数值
// // 以题目示例 [10,9,2,5,3,7,101,18] 为例, tail[1] = 3，因为长度为 2 的所有上升子序列中，结尾最小的是 [2,3]
// // 显然的，tail 数组也是一个严格上升数组(需证明，但不会，见 https://leetcode.cn/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/ 解析)
// // 因此只需维护状态数组 tail 的定义，其长度就是 LIS 的长度
// // 初始化： tail[0] = nums[0]
// // 后续的 tail 数组的维护，基于其定义，我们需要保证，每次新遍历到一个数后，tail 数组中各个数仍是是对应长度下 结尾最小的数
// // 如何维护：
// // 1. 新进一个数，比 tail 数组尾部的数大，则直接 append 在最后
// // 2. 新进一个数，比 tail 数组尾部的数小，则基于二分，找到第一个大于等于它的数，将该数替换掉
// // 一个新员工一个老员工价值相当，老员工就可以走了，因为新员工被榨取的剩余空间更多
// func lengthOfLIS(nums []int) int {
// 	n := len(nums)
// 	tail := append([]int{}, nums[0])
// 	for i := 1; i < n; i++ {
// 		m := len(tail)
// 		if nums[i] > tail[m-1] {
// 			tail = append(tail, nums[i])
// 			continue
// 		}
// 		l, r := -1, m
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if tail[mid] < nums[i] {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// 此时 l 落在 tail 数组中 最后一个小于 nums[i] 的位置，r 落在 tail 数组中 第一个大于等于 nums[i] 的位置
// 		// 由于 nums[i] 小于等于 tail 数组最后一个数，所以 r 的落下位置必定在于[0,m-1] 中，不会越界
// 		tail[r] = nums[i]
// 	}
// 	return len(tail)
// }

// // 过程也可以简单参考 下面的题解
// // labuladong 蜘蛛纸牌理解，略过了证明，可看图理解过程，但无法得到思路的由来
// // https://leetcode.cn/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-she-ji-fang-fa-zhi-pai-you-xi-jia/

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
