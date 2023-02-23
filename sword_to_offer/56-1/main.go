package main

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 56-1. 数组中数字出现的次数

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 1. 最基础的解法：如果想不到位运算，map 后两次遍历也能解决，但是要求空间复杂度为 O(1)，因此舍去

// 2. 出现两次的数字，其二进制数是相同的，可以采用异或，遍历后剩余的数即是 A^B

func singleNumbers(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}

	// 但是次数结果是两个数的异或结果
	// 如何将二者区分？
	// 这里我们已经得到二者异或的结果了，异或的结果是由 0 和 1 组成
	// 其中 1 表示 这两个数的二进制数，在该位上是不同的
	// 此时我们可以想到，通过某个相异的位，去对数组做个分组
	// 这里我们取到异或结果最低位的 1，然后再遍历一遍数组，只对该位为 0 的数做操作
	// 此时这个数组就会被分成两个部分，一个部分是该位为 1 的数，另一个部分是该位为 0 的数
	// 两个相异或的数各在一个部分(因为我们是根据他们异或结果来找的相异的位)
	// 此时就得到其中一个数，再根据异或结果得到另一个数

	// flag 得到 xor 异或结果最低的一位 1
	flag := 1
	for xor&flag == 0 {
		flag <<= 1
	}

	var res int
	for _, num := range nums {
		if num&flag == 0 {
			res ^= num
		}
	}
	return []int{res, res ^ xor}
}

// // 二刷
// // 空间复杂度是 O(n) 的话就简单
// // 但是要求 O(1)
// // 考虑是位运算，出现两次的基于异或结果是 0
// // 最终得到的结果的二进制，是我们要的两个数的异或结果
// // 此时我们有这两个数的异或结果
// // 那么我们可以基于这个异或结果去取得我们要的两个数其中之一
// // 比如说异或结果是 00100，此时我们知道，两个数的二进制位在倒数第三位上是 1
// // 两个数的二进制在倒数第三位上是不同的
// // 我们可以基于这个第三位，再一次遍历数组，只有第三位为 1 的再异或
// // 此时 第三位为 1 的，就只有一部分重复的数和我们要的两个数之一
// // 得到了两个数之一，再与此前的异或结果相异或，就能得到另一个数
// func singleNumbers(nums []int) []int {
// 	var twoRes int
// 	for _, num := range nums {
// 		twoRes ^= num
// 	}
// 	// 取异或结果最低的一位 1
// 	flag := 1
// 	for flag&twoRes == 0 {
// 		flag <<= 1
// 	}
// 	var a, b int
// 	for _, num := range nums {
// 		if num&flag == 0 {
// 			a ^= num
// 		} else {
// 			b ^= num
// 		}
// 	}
// 	return []int{a, b}
// }
