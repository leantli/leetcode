package main

// https://leetcode.cn/problems/single-number-iii/
// 260. 只出现一次的数字 III

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 这道题和 剑指56-1 基本一致
// 这里考虑到只能使用常量级别额外空间+O(n)时间复杂度
// 借用 map 的思路肯定是不行的，排序再比较也不行
// 不得不考虑位运算了，因为除了要求的数，其他书都出现两边
// 基于异或，我们能得到我们要的两个数的异或结果
// 再基于异或的结果去区分两个数即可
// 因为异或结果中，是 1 的位，说明两个数在该位上是不同的
// 那么基于这个不同的位，我们能够区分要的两个数
func singleNumber(nums []int) []int {
	var orRes int
	for _, num := range nums {
		orRes ^= num
	}
	// 此时求异或结果最低的一位 1
	flag := 1
	for orRes&flag == 0 {
		flag <<= 1
	}
	// 为什么基于异或结果相异的位能得到其中一个数呢？
	// 基于这个相异的位，我们能够将该位为 0 的分到一个组，将该位为 1 的分到另一个组
	// 显然，这两个组个包含一个我们要求的数，以及其他重复的数
	// 此时我们只需要不断异或其中一个分组的所有数
	// 就能得到该组我们需要的只出现一次的数
	var a int
	for _, num := range nums {
		if flag&num == 0 {
			a ^= num
		}
	}
	return []int{a, a ^ orRes}
}
