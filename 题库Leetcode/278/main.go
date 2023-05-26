package main

// https://leetcode.cn/problems/first-bad-version/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 278. 第一个错误的版本

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 类似猜数字，但我们需要找到一个分界点，左边是正确的，右边是错误的
// 因此我们使用万金油模板比较方便
func firstBadVersion(n int) int {
	l, r := -1, n+1
	for l+1 != r {
		mid := l + (r-l)/2
		if !isBadVersion(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func isBadVersion(version int) bool {
	return true
}
