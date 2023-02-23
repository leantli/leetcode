package main

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
// 56-2. 数组中数字出现的次数 II

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 把所有数的二进制位相加，最后 % 3，剩余的数就是只出现一次的数
//
//	  1 1 1
//	  1 1 1
//	  1 1 1
//	  1 0 1
//	= 4 3 4
//	% 3 3 3
//	= 1 0 1
// func singleNumber(nums []int) int {
// 	temp := make([]int, 32, 32)
// 	for _, num := range nums {
// 		addBinarty(temp, num)
// 	}
// 	var res int
// 	for i := 0; i < 32; i++ {
// 		res += ((temp[i] % 3) << i)
// 	}
// 	return res
// }

// func addBinarty(a []int, b int) {
// 	for i := 0; i < 32; i++ {
// 		a[i] += (b >> i) & 1
// 	}
// }

// 二刷
// 将每个数的二进制各个位置都加起来，最后模3，此时得到一个二进制数
// 这个二进制数的十进制就是我们要的数
func singleNumber(nums []int) int {
	var bitsSum [32]int
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bitsSum[i] += (num >> i) & 1
		}
	}
	var res int
	for i := 0; i < 32; i++ {
		res += ((bitsSum[i] % 3) << i)
	}
	return res
}
