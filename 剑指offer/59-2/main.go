package main

// https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 59-2. 队列的最大值

// 一道单调队列的题
// max 数组为单调队列，单调递增，参考 59-1 思路
type MaxQueue struct {
	queue []int
	max   []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		max:   make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.max) > 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	if res == this.max[0] {
		this.max = this.max[1:]
	}
	return res
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
