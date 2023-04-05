package main

// https://leetcode.cn/problems/delete-columns-to-make-sorted-iii/
// 960. 删列造序 III

// 返回 answer 的最小可能值
// 执行answer删除操作后，数组中的元素仍然是按字典序排列
// 以下理解有问题！！正确的理解是，每个字符串共用同一个answer序列，删掉后，所有的字符串都还符合字典序
// 而不能像下面的思路一样，每个字符串单独求answer序列长度
// (X)显然是要我们求每个字符串的最长上升子序列，获得所有字符串中最长的上升子序列的长度
// (X)结果就是原长度减去LIS长度
// 那么此时我们状态转移方程虽然还是dp[i]=max(dp[i],dp[j]+1),j<i
// 但是我们要确保，每一个字符串的 str[j] 都小于等于 str[i]，否则不能进行状态转移
func minDeletionSize(strs []string) int {
	n := len(strs[0])
	dp := make([]int, n)
	for i := range strs[0] {
		dp[i] = 1
	}
	longest := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			var flag bool
			for _, str := range strs {
				if str[j] > str[i] {
					flag = true
					break
				}
			}
			if !flag {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		longest = max(dp[i], longest)
	}
	return n - longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
