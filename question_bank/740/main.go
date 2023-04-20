package main

// https://leetcode.cn/problems/delete-and-earn/
// 740. 删除并获得点数

// 当然我们也可以不用排序，而是遍历寻找最大值
func deleteAndEarn(nums []int) int {
	var maxVal int
	m := make([]int, 1e4+1)
	for _, num := range nums {
		m[num] += num
		if num > maxVal {
			maxVal = num
		}
	}
	// 相邻的不取，怎么保证取最大？
	// 定义 dp[i] 为前 i 个数能取到的最大点数
	// dp[i] = max(dp[i-1], dp[i-2]+nums[i]) // 要么i-1位置取过了，要么取i-2位置和当前值
	// 再关注到 i 只由 i-1 和 i-2 决定，因此我们可以用两个局部变量替代整个 dp 数组的声明
	ppre, pre := m[0], max(m[0], m[1])
	for i := 2; i <= maxVal; i++ {
		ppre, pre = pre, max(ppre+m[i], pre)
	}
	return pre
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 取任意nums[i]后，需要删掉该nums[i]和所有nums[i]-1、nums[i]+1
// // 其实相当于，取全部nums[i]，并且放弃所有nums[i]+-1
// // 这里就要判断 到底是取全部nums[i]点数更大，还是nums[i]+-1点数更大
// // map存 {nums[i], sumOfNumsI}，接下来就是打家劫舍了，相当于相邻的不能取，并且要取最大点数
// // 但是 map 的key-value不一定有序，因此我们可以考虑直接用数组存 nums[i]--该nums[i]的总和
// // 然后遍历这个数组做打家劫舍
// func deleteAndEarn(nums []int) int {
// 	sort.Ints(nums)
// 	maxVal := nums[len(nums)-1]
// 	m := make([]int, maxVal+1)
// 	for _, num := range nums {
// 		m[num] += num
// 	}
// 	// 相邻的不取，怎么保证取最大？
// 	// 定义 dp[i] 为前 i 个数能取到的最大点数
// 	// dp[i] = max(dp[i-1], dp[i-2]+nums[i]) // 要么i-1位置取过了，要么取i-2位置和当前值
// 	// 再关注到 i 只由 i-1 和 i-2 决定，因此我们可以用两个局部变量替代整个 dp 数组的声明
// 	ppre, pre := m[0], max(m[0], m[1])
// 	for i := 2; i <= maxVal; i++ {
// 		ppre, pre = pre, max(ppre+m[i], pre)
// 	}
// 	return pre
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
