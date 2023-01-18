package main

// https://leetcode.cn/problems/single-element-in-a-sorted-array/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 540. 有序数组中的单一元素

// 如果没有限制 logn 时间复杂度，可以遍历一遍进行位运算异或
// 但是限制了 logn 时间复杂度，感觉得用二分，但是我们需要找到一个和下标具备单调性关系的东西，否则无法正确地缩进
// 感觉是看 mid 左右值数量的多少? 也不是，缩进之后就不方便用这个判断了，还得额外做处理进行判断
// 这里可以注意正常来说，成对出现的数，第一个数是偶数位，而从单个数出现后，成对数的首个就不是偶数位了
// 这里考虑了一下用奇数位比较好，因为整体下标首位和末尾都是偶数位，边界值需要额外操作，相对来说不那么方便
func singleNonDuplicate(nums []int) int {
	// 开区间二分
	l, r := -1, len(nums)
	// 这里条件要简单注意一下，因为后续 l 和 r 都只在奇数位上移动，因此中间会有额外的间隔，所以 +2
	for l+2 != r {
		mid := l + (r-l)/2
		// 奇数下标时，判断该数是否是对数的第二个数，是则正常，l 右移
		// 不是则不正常，r 左移
		if mid&1 != 1 {
			if mid-1 >= 0 {
				mid--
			} else {
				mid++
			}
		}
		// 正常 l 右移
		if nums[mid] == nums[mid-1] {
			l = mid
		} else {
			r = mid
		}
	}
	return nums[l+1]
}
