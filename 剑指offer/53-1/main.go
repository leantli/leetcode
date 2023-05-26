package main

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 53-1 在排序数组中查找数字 I

// 看到后马上就是两种思路: 1. 单向遍历 2. 二分
// 单向遍历可能会超时，直接不写

func search(nums []int, target int) int {
	index := locateTargetIndex(nums, target)
	if index == -1 {
		return 0
	}
	count := 1
	for i := index - 1; i >= 0; i-- {
		if nums[i] == target {
			count++
		} else {
			break
		}
	}
	for i := index + 1; i < len(nums); i++ {
		if nums[i] == target {
			count++
		} else {
			break
		}
	}
	return count
}

func locateTargetIndex(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		}
		if nums[mid] > target {
			r = mid - 1
		}
	}
	return -1
}

// 大佬的
// 1. 找到最后一个< target的数字下标，记为 L
// 2. 找到最后一个<= target 的数字下标，记为 R
// 3. R-L
// func search(nums []int, t int) int {
// 	N := len(nums)
// 	bisearch := func(cond func(int, int) bool) int {
// 		l, r := -1, N
// 		for l+1 != r {
// 			m := (l + r) >> 1
// 			if cond(m, t) {
// 				l = m
// 			} else {
// 				r = m
// 			}
// 		}
// 		return l
// 	}
// 	r := bisearch(func(m, t int) bool { return nums[m] <= t })
// 	l := bisearch(func(m, t int) bool { return nums[m] < t })
// 	return r - l
// }
