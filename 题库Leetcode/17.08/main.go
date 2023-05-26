package main

import "sort"

// https://leetcode.cn/problems/circus-tower-lcci/
// 面试题 17.08. 马戏团人塔

// // 上面的人比下面的人要矮且轻，显然 h 和 w 这两个值都需要严格小于才可以
// // 这里我们可以从上往下看，显然是一个单调递增子序列的情况，类似 T300 LIS，此时即可寻找最长递增子序列即可
// // 但是这里有区别，T300的子序列，是已经确定好了的子序列，其状态转移方程中，dp[i] 仅需要考虑 j,, j < i 这部分
// // 而目前的 h/w 数组，显然是可以随意组合的，但我们不应该用 dfs 进行组合枚举，时间复杂度过高
// // 所以我们可以考虑如何将现在的 h/w 数组，转化为 T300 这样可以直接 dp 的子序列
// // 这里我们可以先依据 h 排序, 此时 数组中第 i 人只需考虑前面的人的 w 是不是比自己轻即可，因为 i 后面的人的 h，
// // 肯定比 i 高或等高，是不满足条件的
// // 接下来，我们还要考虑一个点，现在 h 是已经排序好的，我们都是根据 w 进行 LIS dp，但是等高情况下，我们要注意
// // 此时 i 的 w 比前面的人的 w 重，但是他们的 h 是相等的，这里也不能取，因此我们可以在排序的时候把等高的 w，降序排列
// // 这样的话，LIS dp 时就不会将这种情况考虑进来
// // 当然，也可以正常根据 w 进行 LIS dp，但是判断状态转移条件时，就多一个是否等高的判断即可
// // 做完发现超时，说明这道题不能用 dp 常规解，而需要基于 贪心+二分，这样的话我觉得这道题应该定为困难而不是中等
// func bestSeqAtIndex(height []int, weight []int) int {
// 	// 初始化人的数据，将体重和身高绑定，便于排序后仍能对上
// 	members := make([][]int, len(height))
// 	for i := range height {
// 		members[i] = []int{height[i], weight[i]}
// 	}
// 	// 排序，削弱问题，便于套用 T300 LIS 的思路
// 	// 第一排列以 h 升序为主，第二排序以等 h，则 w 降序
// 	sort.Slice(members, func(i, j int) bool {
// 		return members[i][0] < members[j][0] || (members[i][0] == members[j][0]) && members[i][1] > members[j][1]
// 	})
// 	// LIS dp 定义 dp[i] 表示，当 member[i] 是人塔的最底部的人时，能叠起来的最多人数
// 	// 初始化 dp 数组, 每个人至少都能自己一个人作为一个人塔
// 	dp := make([]int, len(height))
// 	for i := range height {
// 		dp[i] = 1
// 	}
// 	res := 1
// 	for i := 1; i < len(members); i++ {
// 		for j := 0; j < i; j++ {
// 			if members[j][1] < members[i][1] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		res = max(res, dp[i])
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 贪心+二分，排序削弱前的操作是一致的
func bestSeqAtIndex(height []int, weight []int) int {
	// 初始化人的数据，将体重和身高绑定，便于排序后仍能对上
	members := make([][]int, len(height))
	for i := range height {
		members[i] = []int{height[i], weight[i]}
	}
	// 排序，削弱问题，便于套用 T300 LIS 的思路
	// 第一排列以 h 升序为主，第二排序以等 h，则 w 降序
	sort.Slice(members, func(i, j int) bool {
		return members[i][0] < members[j][0] || (members[i][0] == members[j][0]) && members[i][1] > members[j][1]
	})
	// 创建一个 buttom 数组，buttom[i] 表示 i+1 长度的人塔中，最顶部的人的重量是多少
	// 这里基于贪心，我们要不断迭代 i+1 这个长度最顶部的人的重量，越小越好；这样他下面的人的可选余地更多(重量比他大的数量更多)
	// 便于后续延长 buttom 的长度，得到更长的人塔
	buttom := make([]int, 0)
	for _, v := range members {
		// 这里我们要找到当前 v[1] 重量要插入的位置，
		l, r := -1, len(buttom)
		for l+1 != r {
			mid := l + (r-l)/2
			if buttom[mid] < v[1] {
				l = mid
			} else {
				r = mid
			}
		}
		if r == len(buttom) {
			buttom = append(buttom, v[1])
		} else {
			buttom[r] = v[1]
		}
	}
	return len(buttom)
}
