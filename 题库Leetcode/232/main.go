package main

// https://leetcode.cn/problems/implement-queue-using-stacks/submissions/
// 232. 用栈实现队列

// 栈是先进后出的，队列是先进先出的，用两个栈实现队列，A 栈进 ABC,弹出顺序将会逆序为 CBA，此时若要先进先出，则将 A 栈元素全部依序弹出到 B 栈，此时再一次逆序就变成了先进先出，CBA在B栈中弹出为 ABC
// 实现思路：
// 有 push 时，都加入 in 中，需要 pop 时，从 out 栈出，如果 out 栈没元素了，再将 in 栈中所有元素丢到 out 栈
// 最后弹出 out 栈的栈顶，此时满足先进先出；
type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

// in 出栈并将其加入 out 栈
func (this *MyQueue) inToOut() {
	// 题目表示调用 pop 时，一定不为空，所以这里我们无需加额外判断
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		this.inToOut()
	}
	res := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		this.inToOut()
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
