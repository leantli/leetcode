package main

type CQueue struct {
	stack1 []int
	stack2 []int
}

// 每次要进东西，就进 stack1
// 每次要出东西，就从 stack2 拿，没有的话就把 stack1 的东西都倒进 Stack2

func Constructor() CQueue {
	return CQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		this.stack1 = make([]int, 0)
	}
	if len(this.stack2) == 0 {
		return -1
	}
	res := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return res
}
