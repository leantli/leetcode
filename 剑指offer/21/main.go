package main

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 21. 调整数组顺序使奇数位于偶数前面

// 马上想到两次遍历取奇偶，一个数组 append，不过这样应该不是最优解，应该有一次遍历解决的操作？

// func exchange(nums []int) []int {
// 	res := make([]int, 0, len(nums))
// 	for _, num := range nums {
// 		if (num & 1) == 1 {
// 			res = append(res, num)
// 		}
// 	}
// 	for _, num := range nums {
// 		if (num & 1) == 0 {
// 			res = append(res, num)
// 		}
// 	}
// 	return res
// }

// // 第一次优化，一次遍历，但用两个数组
// func exchange(nums []int) []int {
// 	res := make([]int, 0, len(nums))
// 	// 第一次优化，用两个数组
// 	another := make([]int, 0)
// 	for _, num := range nums {
// 		if (num & 1) == 1 {
// 			res = append(res, num)
// 		} else {
// 			another = append(another, num)
// 		}
// 	}
// 	res = append(res, another...)
// 	return res
// }

// // 第二次优化，一个数组+一次遍历
// func exchange(nums []int) []int {
// 	n := len(nums)
// 	res := make([]int, n, n)
// 	l, r := 0, n-1
// 	for _, num := range nums {
// 		if (num & 1) == 1 {
// 			res[l] = num
// 			l++
// 		} else {
// 			res[r] = num
// 			r--
// 		}
// 	}
// 	return res
// }

// 双指针+原地修改解
func exchange(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		// 是奇数则左指针正常向右遍历，直到遇见偶数
		if l < r && (nums[l]&1) == 1 {
			l++
		}
		// 是偶数则右指针正常向左遍历，直到遇见奇数
		if l < r && (nums[r]&1) != 1 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
