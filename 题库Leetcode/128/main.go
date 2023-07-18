package main

// https://leetcode.cn/problems/longest-consecutive-sequence/
// 128. 最长连续序列

func longestConsecutive(nums []int) int {
	// 记录某一个数是否出现过，若出现过，以该数为结尾的最长连续序列有多长
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	var res int
	// 遍历 m 中出现过的数
	for num := range m {
		// 如果这个数之前没出现过，就以这个数作为起始点开始向后尝试遍历
		if !m[num-1] {
			curLen := 1
			// 后面有连续的，就++
			for m[num+1] {
				curLen++
				num++
			}
			if curLen > res {
				res = curLen
			}
		}
	}
	return res
}
