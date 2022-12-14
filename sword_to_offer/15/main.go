package main

// https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 15. 二进制中1的个数

// &1 取最后一位 1，>>1 移动
func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}
