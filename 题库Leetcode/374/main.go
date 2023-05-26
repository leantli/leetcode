package main

// https://leetcode.cn/problems/guess-number-higher-or-lower/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 374. 猜数字大小

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// 准确值模板
func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		switch guess(mid) {
		case 0:
			return mid
		case 1:
			l = mid + 1
		case -1:
			r = mid - 1
		}
	}
	return -1
}

func guess(num int) int {
	return -1
}
