package main

// https://leetcode.cn/problems/unique-binary-search-trees/description/
// 96. 不同的二叉搜索树

// 模拟一下规律，首先单个节点，肯定是一种；两个节点，1-2 和 2-1
// 三个节点则如题目图片所示，分别以 1、2、3 为根结点，其中以 1 为根结点时，其左子节点的组成为 0 种，右子结点的组成为 2 种(2 个节点)
// 以 2 为根结点时，左子节点的组成为 1 种(1 个节点)，右子结点的组成为 1 种(1 个节点)，以 3 为根结点时，左子节点的组成为 2 种，右子结点的组成为 0 种
// 此时我们定义 dp[i] 为 i 个节点组成的互不相同的二叉搜索树的数量
// 显然 dp[3] = dp[0]*dp[2] + dp[1] * dp[1] + dp[2] * dp[0] (分别对应着根结点为 1、2、3 的情况), dp[0] 初始化为 1
// dp[2] = dp[0] * dp[1] + dp[1] * dp[0], dp[1] = 1
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}
