package main

// https://leetcode.cn/problems/daily-temperatures/
// 739. 每日温度

// 常规做法就是 n^2 遍历

// 这里我们尽量让时间复杂度变低，先说说题目目的
// 找到每一个数右侧第一个比该数大的数
// 显然我们能够使用数据结构--栈，从栈底到栈顶单调递减，
// 当遍历到的数值大于栈顶的数，则将栈顶弹出，将当前数丢进去
// 弹出时，我们就知道，当前的数是栈顶数右侧第一个比它大的值
// 接着，由于题目要求是求二者的idx距离，因此栈中存的是下标，而不是具体的值
// 栈中存放的是数值的 idx，也便于后续 res[idx]定位
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			preIdx := stack[len(stack)-1]
			res[preIdx] = i - preIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
