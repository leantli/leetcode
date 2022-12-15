package main

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
// 56-2. 数组中数字出现的次数 II

// 把所有数的二进制位相加，最后 % 3，剩余的数就是只出现一次的数
//
//	  1 1 1
//	  1 1 1
//	  1 1 1
//	  1 0 1
//	= 4 3 4
//	% 3 3 3
//	= 1 0 1
func singleNumber(nums []int) int {
	temp := make([]int, 32, 32)
	for _, num := range nums {
		addBinarty(temp, num)
	}
	var res int
	for i := 0; i < 32; i++ {
		res += ((temp[i] % 3) << i)
	}
	return res
}

func addBinarty(a []int, b int) {
	for i := 0; i < 32; i++ {
		a[i] += (b >> i) & 1
	}
}
