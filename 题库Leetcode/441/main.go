package main

// https://leetcode.cn/problems/arranging-coins/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 441. 排列硬币

// 对题目简而言之，即是 找到一个总行数，其总硬币数是小于 n 枚硬币的最大数量
// 这里我们可以简单模拟，但是简单模拟就是单纯的累加，等差求和，直到和大于 n，这样的方法在 n 比较大时可能耗时过长
// 这里我们可以看到，题目其实就是找到一个总行数大总硬币数是 小于 n 枚硬币的最大数量
// 很明显的一个 二分的 分界条件
// 而计算总行数的总硬币数，显然是通过等差数列求和公式即可
// 因为不是一个准确值，我们使用分界模板
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	l, r := 1, n
	for l+1 != r {
		mid := l + (r-l)/2
		if (1+mid)*mid/2 <= n {
			l = mid
		} else {
			r = mid
		}
	}
	// 此时 r 必定指向一个行数，该行数的总硬币数是 大于 n 枚硬币的 最小数量
	// 并且 l 指向的就是 完整阶梯行的总行数
	return l
}
