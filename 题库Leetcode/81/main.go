package main

// https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 81. 搜索旋转排序数组 II

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	l, r := 0, n-1
	// 我们判断mid位于哪个区间，是基于 nums[r] 的
	// 因此先将首尾重复的值消掉，否则会影响两个区间的单调性以及影响首位判断
	for l < r && nums[0] == nums[r] {
		r--
	}
	fakeN := r
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		// 先确定 mid 在什么位置
		if nums[mid] <= nums[fakeN] {
			// 再确定 target是否也在 mid 所在的位置
			// 是的话再看看 mid 怎么缩，只能 mid < target 时右移
			if target <= nums[fakeN] && nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target > nums[fakeN] && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
