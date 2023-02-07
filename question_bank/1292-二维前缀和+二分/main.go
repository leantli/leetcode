package main

// https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/
// 1292. 元素和小于等于阈值的正方形的最大边长

// 看起来是基于 二分 确定边长，然后 dfs/枚举 去看是否存在满足阈值，才方便继续二分
// 但是这里计算正方形和，显然要用到前缀和
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	// 二维数组初始化
	sum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([]int, n+1)
	}
	// 先计算好二维前缀和
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// sum[i-1][j-1] 这部分区域同时被 sum[i-1][j] 和 sum[i][j-1] 包含，因此重复加了一次，所以要减去
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	// 前缀和计算完毕，接下来二分取正方形边长，并在后面计算是否满足阈值内
	// 这里注意取值，正方形的边长取值明显为 [0,短边]，下面开区间二分
	l, r := -1, min(m, n)+1
	for l+1 != r {
		mid := l + (r-l)/2
		// 当存在满足阈值内，则 l 向右，表示 mid 可尝试取更大
		if calCount(sum, mid, threshold) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 计算指定正方形边长下，是否存在矩阵满足阈值，满足则返回 true
func calCount(sum [][]int, side, threshold int) bool {
	for i := side; i < len(sum); i++ {
		for j := side; j < len(sum[0]); j++ {
			// 求边长为 side 的矩阵和
			res := sum[i][j] - sum[i-side][j] - sum[i][j-side] + sum[i-side][j-side]
			if res <= threshold {
				return true
			}
		}
	}
	return false
}
