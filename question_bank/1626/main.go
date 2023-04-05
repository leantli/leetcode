package main

// https://leetcode.cn/problems/best-team-with-no-conflicts/
// 1626. 无矛盾的最佳球队

// 输出最高的得分
// 要基于无矛盾球队寻找最高分的队
// 什么球队无矛盾，A的年龄和分数都要比B小，那么就是无矛盾球队
// 因此我们要找递增的子序列，是不是最长无所谓，这个不一定？
// 递增子序列的话，我们得想到 T300-LIS，但是这里有两个参数，分别是年龄和成绩
// 这里可以参考 T354 的思路，将问题削减弱化成 LIS
// 正常 LIS 只有一个参数考虑，并且下标 i 的子问题只包括 j, j<i，但本题显然无序
// 因此我们可以先排序一番，不管是基于 年龄还是成绩，都可以保证 下标 i 的子问题只包括 j, j<i
// 而第二个参数时，是在第一个参数相等的情况下考虑，当年龄相等时，成绩高低无所谓，因此成绩升序排列
// 当成绩相等时，年龄高低也无所谓，因此年龄升序排列
// 总的来说，第一个参数必须要升序排列，保证下标 i 的子问题为最优子结构
// 第二个参数，在第一个参数相等时没有特殊限制，因此正常升序，属于可取的子序列
// 接着其实也就可以正常按照另一个参数进行LIS了，当然，这里的两个参数都要组合起来一起判断更方便
// 以下我们根据分数做第一升序排序，同分数时，年率做降序排序
// func bestTeamScore(scores []int, ages []int) int {
// 	members := make([][]int, len(scores))
// 	for i := range scores {
// 		members[i] = []int{scores[i], ages[i]}
// 	}
// 	// 排序，根据分数升序排，同分数时，年龄无所谓，因此升序排
// 	sort.Slice(members, func(i, j int) bool {
// 		return members[i][0] < members[j][0] || (members[i][0] == members[j][0] && members[i][1] < members[j][1])
// 	})
// 	// dp[i] 表示以 第i个成员作为无矛盾球队中年龄最大成绩最高时，能取得的最高总分
// 	dp := make([]int, len(scores))
// 	// 初始化，所有成员只有1个人时，他们自己的分数就是dp[i]最高的
// 	for i := range scores {
// 		dp[i] = members[i][0]
// 	}
// 	res := dp[0]
// 	// 状态转移：
// 	for i := 1; i < len(scores); i++ {
// 		for j := 0; j < i; j++ {
// 			if members[i][1] >= members[j][1] {
// 				dp[i] = max(dp[i], dp[j]+members[i][0])
// 			}
// 		}
// 		res = max(res, dp[i])
// 	}
// 	return res
// }

// func bestTeamScore(scores []int, ages []int) int {
// 	members := make([][]int, len(scores))
// 	for i := range scores {
// 		members[i] = []int{ages[i], scores[i]}
// 	}
// 	// 排序，根据年龄升序排，同年龄时，分数无所谓，因此分数升序排
// 	sort.Slice(members, func(i, j int) bool {
// 		return members[i][0] < members[j][0] || (members[i][0] == members[j][0] && members[i][1] < members[j][1])
// 	})
// 	// dp[i] 表示以 第i个成员作为无矛盾球队中年龄最大成绩最高时，能取得的最高总分
// 	dp := make([]int, len(scores))
// 	// 初始化，所有成员只有1个人时，他们自己的分数就是dp[i]最高的
// 	for i := range scores {
// 		dp[i] = members[i][1]
// 	}
// 	res := dp[0]
// 	// 状态转移：
// 	for i := 1; i < len(scores); i++ {
// 		for j := 0; j < i; j++ {
// 			if members[i][1] >= members[j][1] {
// 				dp[i] = max(dp[i], dp[j]+members[i][1])
// 			}
// 		}
// 		res = max(res, dp[i])
// 	}
// 	return res
// }
//
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
