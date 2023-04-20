package main

import "math"

// https://leetcode.cn/problems/pizza-with-3n-slices/
// 1388. 3n 块披萨

// 有点类似打家劫舍2？但是如果一样的话显然不会是困难题，仔细看后发现有个比较大的变动
// 拿完一次披萨后，对应数及其相邻的数都要从数组删去，此时原数组中部分数相邻的数发生变化
// 因此 dp[i] = max(dp[i-1],dp[i-2]+slices[i]) 不成立，因为 i, i-1, i-2 指向的值可能发生变化
// 也就是说，打家劫舍2 可以取到 n/2 个数，而本题，只能取到 n/3 个数
// 接下来先明确本题目标：从3n块披萨中，选择n/3块的最大和，并且这n/3块不相邻
// 接着将目标分解成小问题，从前 i 块披萨中选取 j 块最大和，并且这 j 块不相邻
// 因此可以定义 dp[i][j] 表示从前 i 块披萨中选出 j 块不相邻的披萨达到最大和
// 再继续考虑这个 dp 的转移方程
// 如果我们要选第 i 块披萨，则不能选第i-1块披萨，并且之前要已经选出 j-1 个披萨 dp[i][j] = dp[i-2][j-1] + slices[i]
// 如果我们不选第 i 块披萨，则之前要已经选过 j 个披萨，dp[i][j] = dp[i-1][j]
// 因此 dp[i][j] = max( dp[i-2][j-1] + slices[i], dp[i-1][j] )
func maxSizeSlices(slices []int) int {
	n := len(slices)
	return max(getMax(slices[1:]), getMax(slices[:n-1]))
}

func getMax(slices []int) int {
	// dp数组初始化以及边界条件初始化
	n, m := len(slices), int(math.Ceil(float64(len(slices))/3.0))
	dp := make([][]int, n+1)
	dp[0] = make([]int, m+1)
	dp[1] = make([]int, m+1)
	dp[1][1] = slices[0]
	// 注意，这里是前 i 个披萨，因此再取 slices 时，需要 i-1，和上面推出来的转移方程本质无区别
	for i := 2; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i-2][j-1]+slices[i-1], dp[i-1][j])
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
