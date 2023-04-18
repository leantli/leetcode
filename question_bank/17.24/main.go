package main

// https://leetcode.cn/problems/max-submatrix-lcci/
// 面试题 17.24. 最大子矩阵

// 看到的瞬间有种矩阵前缀和的感觉
// dp[i][j] = matric[i][j] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] ?
// 但是真搞起来感觉时间复杂度会很高？

// // 没想到其他的，看了题解，是对二维矩阵进行压缩，将二维压缩成一维
// // 不过写出来后耗时还是挺长的
// //  1. 对每列进行前缀和，
// //  2. 对二维矩阵进行降维
// //     2.1 两个 for， 压缩 i~j 行的列的值都到同一行
// //     2.2 此时再对每一行 求 最大子数组和
// func getMaxMatrix(matrix [][]int) []int {
// 	n, m := len(matrix), len(matrix[0])
// 	// 1. 对每列求前缀和
// 	sum := make([][]int, n+1)
// 	// 注意的是，前缀和是1~n，而不是0~n-1
// 	for i := range sum {
// 		sum[i] = make([]int, m)
// 	}
// 	for i := 1; i <= n; i++ {
// 		for j := 0; j < m; j++ {
// 			sum[i][j] = sum[i-1][j] + matrix[i-1][j]
// 		}
// 	}
// 	// fmt.Println(sum)
// 	// 2. 压缩 i~j 行的列的值到同一行，再求该行的最大子数组和
// 	// i表示矩阵最上方一行的前一行，用于做前缀和相减，求这个矩阵每列的值综合
// 	// j表示矩阵最下方的行号, 取值时需要-1
// 	r1, c1, r2, c2 := 0, 0, 0, 0
// 	maxium := matrix[0][0]
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j <= n; j++ {
// 			// 2.1 获取 i+1~j 行矩阵每列数值的总和
// 			nums := make([]int, m)
// 			for k := range nums {
// 				nums[k] = sum[j][k] - sum[i][k]
// 			}
// 			// fmt.Println(nums)
// 			// 2.2 对这个矩阵压缩后的一维数组求最大子数组和
// 			sum := nums[0]
// 			var l int
// 			for r := 1; r < m; r++ {
// 				if sum <= 0 {
// 					l = r
// 					sum = 0
// 				}
// 				sum += nums[r]
// 				if sum > maxium {
// 					maxium = sum
// 					r1, r2 = i, j-1
// 					c1, c2 = l, r
// 					// fmt.Printf("r1:%d, c1:%d, r2:%d, c2:%d\n", r1, c1, r2, c2)
// 				}
// 			}
// 		}
// 	}
// 	return []int{r1, c1, r2, c2}
// }

// 再参考一下佬写的，就是把求前缀和以及dp过程整合在了一起
func getMaxMatrix(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	r1, c1, r2, c2 := 0, 0, 0, 0
	maxSum := matrix[0][0]
	for i := 0; i < n; i++ {
		sum := make([]int, m)
		for j := i; j < n; j++ {
			dp := 0
			startCol := 0
			for k := 0; k < m; k++ {
				sum[k] += matrix[j][k]
				dp += sum[k]
				if dp > maxSum {
					maxSum = dp
					r1, c1, r2, c2 = i, startCol, j, k
				}
				if dp < 0 {
					dp = 0
					startCol = k + 1
				}
			}
		}
	}
	return []int{r1, c1, r2, c2}
}
