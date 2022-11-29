package main

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 53-2 0～n-1中缺失的数字

// map 和遍历当然可以，但这两种都太低效了，考虑一下其他的方法
// 还是二分吧，找到第一个 index != nums[index] 的值

// 印象中之前做过一道题，一个数组中所有数字都是重复，只有一个数字只出现过一次，可以用位运算
// 这里也可以转换一下思路
// 从题目可以看出多出来的值必定是 len(nums)
// index 和 nums[index] 都是相等的情况，二者异或得 0
// 此时遍历一遍进行异或，就会多出来一个单独的 index, 它得不到异或，就是结果
// 其他的 index 和 num 都会异或得 0
func missingNumber(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		res ^= i ^ num
	}
	return res
}

// 二分
// func missingNumber(nums []int) int {
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		m := l + (r-l)/2
// 		if nums[m] == m {
// 			l = m + 1
// 		} else {
// 			r = m - 1
// 		}
// 	}
// 	return l
// }

// 常规，秒解但面试肯定不行
// func missingNumber(nums []int) int {
// 	for i, num := range nums {
// 		if i != num {
// 			return i
// 		}
// 	}
// 	return -1
// }
