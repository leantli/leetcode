package main

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 3. 数组中重复的数字

// 还看到一种 鸠占鹊巢 做法，其实和下面的 trick 的思路很相似
func findRepeatNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

// 还有一种做法，看到数组长度为 n， 且数字范围为0～n-1，因此可以合理使用这个 trick
// func findRepeatNumber(nums []int) int {
// 	n := len(nums)
// 	for _, num := range nums {
// 		if nums[num%n] >= n {
// 			return num % n
// 		}
// 		nums[num%n] += n
// 	}
// 	return -1
// }

// 马上想到的常规做法
// func findRepeatNumber(nums []int) int {
// 	m := make(map[int]struct{})
// 	for _, num := range nums {
// 		if _, ok := m[num]; ok {
// 			return num
// 		}
// 		m[num] = struct{}{}
// 	}
// 	return -1
// }
