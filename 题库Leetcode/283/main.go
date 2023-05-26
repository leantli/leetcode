package main

// https://leetcode.cn/problems/move-zeroes/
// 283. 移动零

// 对数组原地修改，并且按题目的要求，就相当于把0往后放即可
// 原地修改显然需要用两个指针
// 一个指针指向“新”数组当前填补到的位置，另一个指针遍历全数组以填充“新”数组
func moveZeroes(nums []int) {
	var new int
	// 先按序将非0的数往前补
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[new] = nums[i]
			new++
		}
	}
	// 最后把数组后面的全部补为0
	for new < len(nums) {
		nums[new] = 0
		new++
	}
}
