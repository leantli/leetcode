package main

// https://leetcode.cn/problems/peak-index-in-a-mountain-array/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 852. 山脉数组的峰顶索引

// 这道题只有一个峰顶的样子，无准确值可用，使用模糊值模板
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		// 当是正常递增时
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
