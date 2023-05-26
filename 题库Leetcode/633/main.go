package main

// https://leetcode.cn/problems/sum-of-square-numbers/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 633. 平方数之和

// 我想到的思路是这样，先对 c 开一次根，向下取整
// 得到的这个数，作为边界 r，l 初始为 1，再双指针逐渐移动判断是否可等
// 整个时间复杂度应该是 logc + c/2 ?
func judgeSquareSum(c int) bool {
	// 这里要注意 l 和 r 的取值，c 的开根范围可能是 [1,c]，因此使用万用模板时，l=1-1,r=c+1
	l, r := 0, c+1
	for l+1 != r {
		mid := l + (r-l)/2
		if mid*mid <= c {
			l = mid
		} else {
			r = mid
		}
	}
	// 此时 l 会是 c 开根号向下取整
	// 初始化 l 和 r 作为双指针
	l, r = 0, l
	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		} else if sum > c {
			r--
		} else {
			l++
		}
	}
	return false
}
