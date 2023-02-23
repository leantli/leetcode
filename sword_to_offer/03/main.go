package main

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 3. 数组中重复的数字

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 二刷
// 还有一种，我们注意到这里长度为 n,所有数字都在 0~n-1 范围内
// 0~n-1不就是数组下标吗，那这里我们可以遍历数组中的数
// 这个数是多少，就把这个数作为下标，和对应下标的数交换位置
// 并且值得注意的是，我们不能直接就遍历一遍去置换
// 每个下标都要 for 循环直到 i == nums[i]，避免遗漏
func findRepeatNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}

// // 常规做法的话肯定就是 map,然后遇到已有的就返回这个数
// func findRepeatNumber(nums []int) int {
//     m := make(map[int]bool)
//     for _, num := range nums {
//         if m[num] {
//             return num
//         }
//         m[num] = true
//     }
//     return -1
// }

// // 还看到一种 鸠占鹊巢 做法，其实和下面的 trick 的思路很相似
// func findRepeatNumber(nums []int) int {
// 	for i := range nums {
// 		for nums[i] != i {
// 			if nums[i] == nums[nums[i]] {
// 				return nums[i]
// 			}
// 			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
// 		}
// 	}
// 	return -1
// }

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
