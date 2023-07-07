package main

// https://leetcode.cn/problems/single-number-ii/
// 137. 只出现一次的数字 II

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 这里注意一个问题，这道题和 剑指56-2 基本一样
// 但是这道题还包含负数，而 go 的 int 默认是 64 位
// 这会导致首位的符号位没有起到作用
// 所以注意各个位相加时，取 int32，最后再转成 int 返回
func singleNumber(nums []int) int {
	var bitSum [32]int
	for i := 0; i < 32; i++ {
		// 取 num 每一位是否为 1，是就在该位上 +1
		for _, num := range nums {
			bitSum[i] += (num >> i) & 1
		}
	}
	var res int32
	for i := 0; i < 32; i++ {
		res += int32((bitSum[i] % 3) << i)
	}
	return int(res)
}
