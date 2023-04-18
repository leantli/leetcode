package main

// https://leetcode.cn/problems/max-sum-of-rectangle-no-larger-than-k/
// 363. 矩形区域不超过 K 的最大数值和

// // 刚才的压缩+一维dp失败了，我还是先尝试一下暴力解
// // 暴力解---枚举矩阵，求矩阵和判断是否不超过k且最大
// // 1. 矩阵前缀和
// // 2. 枚举矩阵，计算其矩阵和，判断是否满足条件
// // 矩阵前缀和---sum[i][j] = sum[i-1][j]+sum[i][j-1]-sum[i-1][j-1]+matrix[i][j]
// // 求子矩阵 i,j-->x,y 两个坐标作为矩阵右下角时 curSum = sum[y][x] - sum[y][j] - sum[i][x] + sum[i][j] (这里考虑大矩阵套小矩阵应该怎么求右下矩阵和的图形)
// // 可怕的是，提交后 164ms 击败91.18%，内存 6.3M，击败 82.35%
// // 时间复杂度 n^2*m^2，如果测试用例再多点，想来肯定是会超时的
// func maxSumSubmatrix(matrix [][]int, k int) int {
// 	n, m := len(matrix), len(matrix[0])
// 	// 1. 初始化矩阵前缀和
// 	sum := make([][]int, n+1)
// 	for i := range sum {
// 		sum[i] = make([]int, m+1)
// 	}
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= m; j++ {
// 			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
// 			if sum[i][j] == k {
// 				return k
// 			}
// 		}
// 	}
// 	res := math.MinInt
// 	// 2. 枚举子矩阵，求子矩阵的和
// 	for i := 0; i <= n; i++ {
// 		for j := 0; j <= m; j++ {
// 			for y := i + 1; y <= n; y++ {
// 				for x := j + 1; x <= m; x++ {
// 					curSum := sum[y][x] - sum[y][j] - sum[i][x] + sum[i][j]
// 					if curSum <= k && curSum > res {
// 						res = curSum
// 					}
// 					if res == k {
// 						return res
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// 题解参考
// https://leetcode.cn/problems/max-sum-of-rectangle-no-larger-than-k/solution/tong-ge-lai-shua-ti-la-yi-ti-wu-jie-bao-9opdb/

// // 有点像是 53 的矩阵进阶题
// // 附加了二维数组，以及一个不超过k的限制
// // 参考着 17.24 的思路写了一下，发现过不了，因为不超过k这个限制，
// // 在一定程度上会限制到dp的最优子结构，因此无法基于这种压缩进行dp
// // 当得到的res大于k时，可以通过双循环暴力解决，但在一定程度上还是影响到整体时间复杂度
// func maxSumSubmatrix(matrix [][]int, limit int) int {
// 	n, m := len(matrix), len(matrix[0])
// 	res := math.MinInt
// 	// 这里的 i 和 j，分别是矩阵的最顶部和最底部的行号
// 	for i := 0; i < n; i++ {
// 		// 基于 i~j 行中，每一列的数值，累加到一个一维数组
// 		// 便于方便的使用最大子数组和的 dp
// 		sum := make([]int, m)
// 		for j := i; j < n; j++ {
// 			dp := 0
// 			for k := 0; k < m; k++ {
// 				sum[k] += matrix[j][k]
// 				dp += sum[k]
// 				fmt.Printf("dp:%d, res:%d, k:%d\n", dp, res, k)
// 				if dp > res {
// 					res = dp
// 				}
// 				if dp < 0 {
// 					dp = 0
// 				}
// 			}
// 		}
// 	}
// 	return res
// }
