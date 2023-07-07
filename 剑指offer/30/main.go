package main

import "math"

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
// 剑指 Offer 30. 包含min函数的栈

// 首先需要一个 slice 作为栈
// 其次还需要一个 slice，只保存栈中最小的元素
// 每次入栈时，min slice 需要判断一个入栈的值是否小于 min slice 上一个 index 的值

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0, 20000),
		min:   []int{math.MaxInt},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	temp := this.min[len(this.min)-1]
	if temp < x {
		this.min = append(this.min, temp)
	} else {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
