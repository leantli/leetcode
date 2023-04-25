package main

// https://leetcode.cn/problems/counting-bits/
// 338. 比特位计数

// 这里的动态规划是什么样的，其实只是基于二进制的规律
// 假定 dp[i] 表示 i 的二进制有多少个 1
// 8 的二进制是 1000，4 的二进制是 100，显然 dp[8]=dp[4], 8>>1=4
// 7 的二进制是 111，3 的二进制是 11, 显然，dp[7] = dp[3] + 1, 7>>1=3, 多了一个1只要看新增的一位是不是 1 即可
// 那么整合起来就是 dp[i] = dp[i>>1] + (i&1)
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

// // 6，除了暴力之外我想不到其他方法，但是暴力显然太拉了，我选择直接看官解
// // 看完首先是要了解到这个位运算技巧 --- Brian Kernighan 算法
// // 对任意整数 x，令 x=x&(x-1)，该运算会将 x 的二进制的最后一个 1 变为 0
// // 重复该操作，直到 x 变成 0，操作次数就是 x 二进制中 1 的数量
// // 该计算方法，每个数只需要 logn 的时间即可，因此总的时间复杂度为 nlogn
// func countBits(n int) []int {
// 	res := make([]int, n+1)
// 	for i := 0; i <= n; i++ {
// 		var cnt int
// 		cur := i
// 		for cur != 0 {
// 			cur &= cur - 1
// 			cnt++
// 		}
// 		res[i] = cnt
// 	}
// 	return res
// }
