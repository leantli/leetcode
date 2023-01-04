package main

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 167. 两数之和 II - 输入有序数组

// 有序？双指针？时间复杂度是 O(n)
// 感觉没必要再用啥二分了，二分的话感觉至少得固定一次数，时间复杂度是 nlogn
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{-1, -1}
}
