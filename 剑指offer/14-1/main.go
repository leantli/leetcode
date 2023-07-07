package main

// https://leetcode.cn/problems/jian-sheng-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 14-1 剪绳子

// 对于这类题，肯定是有较好的数学拆分解题方式
// 但是我记不得了，可能考虑用 dfs 去组合
// 但是这种方法很可能不行，可能内存爆了，耗时也较长
// 测试了一下，结果没问题，但是在 leetcode 最终跑测试用例的时候会超时
// func cuttingRope(n int) int {
// 	var res int
// 	// n 是当前剩下的绳子长度，cur 是当前剪下的绳子的长度乘积，next 是下一次要剪的绳子长度
// 	var dfs func(n, cur, next int)
// 	dfs = func(n, cur, next int) {
// 		if n == 0 {
// 			return
// 		}
// 		cur *= next
// 		if cur > res {
// 			res = cur
// 		}
// 		for i := 1; i <= n-next; i++ {
// 			dfs(n-next, cur, i)
// 		}
// 	}
// 	for i := 1; i < n; i++ {
// 		dfs(n, 1, i)
// 	}
// 	return res
// }

// // 刚 dfs 也没剪枝，显然当 n 大时太多重复了
// // 感觉也可以用 dp 搞，即不断计算组成 n 的最大乘积
// // 其中 dp[i] 表示将正整数 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积
// // 当 i≥2 时，假设对正整数 i 拆分出的第一个正整数是 j（1≤j<i），则有以下两种方案：
// // 1) 将 i 拆分成 j 和 i−j 的和，且 i−j 不再拆分成多个正整数，此时的乘积是 j×(i−j)；
// // 2) 将 i 拆分成 j 和 i−j 的和，且 i−j 继续拆分成多个正整数，此时的乘积是 j×dp[i−j]。
// // dp[n] = max{ {dp[1] ... dp[n-1]} * {n-1, ..., 1}, {1*(n-1), 2*(n-2), k*(n-k) }}
// // 该方式下减少了 dfs 一部分无用数据的存储
// // 跑完已经是时间上战胜 100% 了，但时间复杂度还是挺高的 O(n^2)
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	// 至少剪 1， 并且不能剪 n
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 再看看官解复习一下数学理论 https://leetcode.cn/problems/jian-sheng-zi-lcof/solution/jian-sheng-zi-by-leetcode-solution-xku9/
// 证明见链接，总结就是
// n >=4 时，则对 3 取余数
// 当 余数为 0 时，n = 3m，将 n 拆成 m 个 3，此时乘积最大
// 当 余数为 1 时，n = 3m + 1，将 n 拆成 m - 1 个 3，将多出的 4，拆成 2 * 2，因为 2 * 2 > 3 * 1
// 当 余数为 2 时，n = 3m + 2，将 n 拆成 m 个 3 和 1 个 2，此时乘积最大
// n < 4 时，最大乘积为 n - 1

// func cuttingRope(n int) int {
// 	if n < 4 {
// 		return n - 1
// 	}
// 	temp := n % 3
// 	if temp == 0 {
// 		return int(math.Pow(3, float64(n/3)))
// 	}
// 	if temp == 1 {
// 		return int(math.Pow(3, float64(n/3-1))) * 4
// 	}
// 	return int(math.Pow(3, float64(n/3))) * 2
// }
