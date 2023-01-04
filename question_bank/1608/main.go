package main

// https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 1608. 特殊数组的特征值

// 找到一个数 x，该数使得 nums 中的数刚好有 x 个数 大于等于 x
// 有确定值，因此使用准确值模板
func specialArray(nums []int) int {
	// 先确定 x 的初始范围
	l, r := 1, len(nums)
	for l <= r {
		mid := l + (r-l)/2
		switch judgeSpecial(nums, mid) {
		case 0:
			return mid
			// 当太少 num 大于 x 时，说明 x 太大了，需要变小
		case -1:
			r = mid - 1
			// 当太多 num 大于 x 时，说明 x 太小了，需要变大
		case 1:
			l = mid + 1
		}
	}
	return -1
}

// 判断 x 是不是 特征值，是则返回 0，大于 x 的数量过多则返回 1， 大于 x 的数量太少则返回 -1
func judgeSpecial(nums []int, x int) int {
	var count int
	for _, num := range nums {
		if num >= x {
			count++
		}
		if count > x {
			return 1
		}
	}
	if count == x {
		return 0
	}
	return -1
}
