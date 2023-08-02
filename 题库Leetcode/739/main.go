package main

// https://leetcode.cn/problems/daily-temperatures/
// 739. 每日温度

// 二刷
// 找到 temp[i] 右侧第一个比 temp[i] 大的温度，记录其下标作为 answer[i]
// 最暴力的方法很明显，n^2 的遍历，实现很简单，但很可能在某些用例过不了，时间复杂度过高
// 一次遍历，记录左侧的温度，当遇到比左侧温度大的日子时，弹出对应的下标为 i，则 answer[i] = 当前的日子下标 - i
func dailyTemperatures(temperatures []int) []int {
	// 存放过去的日子的下标，因为我们要取右侧气温比左侧大的日子
	// 因此当栈顶的温度比当前小，显然就可以弹出，此时我们可以发现这个栈是个单调递减栈，从栈底到栈顶是单调递减的
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))
	for i := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			preIdx := stack[len(stack)-1]
			ans[preIdx] = i - preIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// // 这里我们尽量让时间复杂度变低，先说说题目目的
// // 找到每一个数右侧第一个比该数大的数
// // 显然我们能够使用数据结构--栈，从栈底到栈顶单调递减，
// // 当遍历到的数值大于栈顶的数，则将栈顶弹出，将当前数丢进去
// // 弹出时，我们就知道，当前的数是栈顶数右侧第一个比它大的值
// // 接着，由于题目要求是求二者的idx距离，因此栈中存的是下标，而不是具体的值
// // 栈中存放的是数值的 idx，也便于后续 res[idx]定位
// func dailyTemperatures(temperatures []int) []int {
// 	stack := make([]int, 0)
// 	res := make([]int, len(temperatures))
// 	for i, temperature := range temperatures {
// 		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
// 			preIdx := stack[len(stack)-1]
// 			res[preIdx] = i - preIdx
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, i)
// 	}
// 	return res
// }
