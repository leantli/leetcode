package main

// https://leetcode.cn/problems/implement-stack-using-queues/
// 225. 用队列实现栈

// 队列是先进先出，元素依序进入队列 1，再依序出到 队列2 时，仍然无法实现栈的先进后出
// 因此用栈实现队列的思路，在这里是行不通的
// 这里我们可以这样，A 队列，只留下一个元素，其他元素全部出到 B 队列，此时返回 A 队列中最后一个元素
// 再用一个 curIdx 指明现在哪个队列存着真正的元素即可
type MyStack struct {
	curIdx int // 0 或者 1
	queue  [][]int
}

func Constructor() MyStack {
	q := make([][]int, 2)
	for i := range q {
		q[i] = make([]int, 0)
	}
	return MyStack{0, q}
}

func (this *MyStack) Push(x int) {
	this.queue[this.curIdx] = append(this.queue[this.curIdx], x)
}

func (this *MyStack) Pop() int {
	for len(this.queue[this.curIdx]) != 1 {
		// 弹出队首给另一个队列
		this.queue[this.curIdx^1] = append(this.queue[this.curIdx^1], this.queue[this.curIdx][0])
		this.queue[this.curIdx] = this.queue[this.curIdx][1:]
	}
	num := this.queue[this.curIdx][0]
	this.queue[this.curIdx] = this.queue[this.curIdx][1:]
	this.curIdx ^= 1 // 变更当前存元素的队列 idx
	return num
}

func (this *MyStack) Top() int {
	for len(this.queue[this.curIdx]) != 1 {
		// 弹出队首给另一个队列
		this.queue[this.curIdx^1] = append(this.queue[this.curIdx^1], this.queue[this.curIdx][0])
		this.queue[this.curIdx] = this.queue[this.curIdx][1:]
	}
	num := this.queue[this.curIdx][0]
	this.queue[this.curIdx] = this.queue[this.curIdx][1:]
	this.curIdx ^= 1 // 变更当前存元素的队列 idx
	this.queue[this.curIdx] = append(this.queue[this.curIdx], num)
	return num
}

func (this *MyStack) Empty() bool {
	return len(this.queue[this.curIdx]) == 0
}

// // 再来一点优化，其实两个队列，只有一个是真正返回结果的，另一个只是做数据的备份
// // 其实我们一个队列就可以搞定，将当前长度-1的元素，全部弹出，丢到当前队列尾部即可，此时首部就是之前最后的元素
// // 也实现了先进后出，并且不需要另一个队列和 curIdx 变更等操作
// type MyStack struct {
// 	queue []int
// }

// func Constructor() MyStack {
// 	return MyStack{[]int{}}
// }

// func (this *MyStack) Push(x int) {
// 	this.queue = append(this.queue, x)
// }

// func (this *MyStack) Pop() int {
// 	n := len(this.queue)
// 	for i := 0; i < n-1; i++ {
// 		// 弹出队首到队尾，只弹出 n-1 次，原队尾的数就到了队首
// 		this.queue = append(this.queue, this.queue[0])
// 		this.queue = this.queue[1:]
// 	}
// 	num := this.queue[0]
// 	this.queue = this.queue[1:]
// 	return num
// }

// func (this *MyStack) Top() int {
// 	n := len(this.queue)
// 	for i := 0; i < n-1; i++ {
// 		// 弹出队首到队尾，只弹出 n-1 次，原队尾的数就到了队首
// 		this.queue = append(this.queue, this.queue[0])
// 		this.queue = this.queue[1:]
// 	}
// 	num := this.queue[0]
// 	this.queue = append(this.queue, this.queue[0])
// 	this.queue = this.queue[1:]
// 	return num
// }

// func (this *MyStack) Empty() bool {
// 	return len(this.queue) == 0
// }

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
