package main

// https://leetcode.cn/problems/longest-consecutive-sequence/
// 128. 最长连续序列

// 二刷，我们可以通过 map 存储所有的数
// 再重新遍历一遍 map 中存在的数，只要他们的 num+1 也存在于 map 中，则以它为起点的 cnt++，同时更新 num，继续查看新的 num+1 是否在 map 中
// 以此类推，但这种方法可能存在一个问题，比如 1 2 3 4，我们只需要以 1 为起点判断一次即可，不需要每个数都重复进行判断，浪费资源
// 因此我们可以判断 if !map[num-1]，如果该数是数组中的起始数，则进行 cnt 计算，否则不进行，因为会有其他起始数做 cnt 计算，避免重复
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	var res int
	// 这里遍历 map 中的数而不是遍历 nums 中的数，在一定程度上去重
	for num := range m {
		// 如果该数有其他前置数，则不进行 cnt 计算
		// 我们只从某个起始数开始计算长度，避免不必要的计算浪费
		if m[num-1] {
			continue
		}
		cnt := 1
		for m[num+1] {
			cnt++
			num++
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}

// func longestConsecutive(nums []int) int {
// 	// 记录某一个数是否出现过，若出现过，以该数为结尾的最长连续序列有多长
// 	m := make(map[int]bool)
// 	for _, num := range nums {
// 		m[num] = true
// 	}
// 	var res int
// 	// 遍历 m 中出现过的数
// 	for num := range m {
// 		// 如果这个数之前没出现过，就以这个数作为起始点开始向后尝试遍历
// 		if !m[num-1] {
// 			curLen := 1
// 			// 后面有连续的，就++
// 			for m[num+1] {
// 				curLen++
// 				num++
// 			}
// 			if curLen > res {
// 				res = curLen
// 			}
// 		}
// 	}
// 	return res
// }
