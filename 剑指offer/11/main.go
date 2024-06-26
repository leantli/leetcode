package main

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 11. 旋转数组的最小数字

// 这道二分的判断条件想了好一会。。
// 还存在重复元素，还得线性-- 去重。。
// 感觉不算 easy 题
func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + (r-l)/2
		// 去除首尾的重复元素，保证两个单调区间的单调性
		if numbers[mid] == numbers[r] {
			r--
			continue
		}
		if numbers[mid] < numbers[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return numbers[r]
}
