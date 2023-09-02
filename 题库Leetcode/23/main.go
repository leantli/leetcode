package main

import "container/heap"

// https://leetcode.cn/problems/merge-k-sorted-lists/description/
// 23. 合并 K 个升序链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 将所有链表的首节点，加入到小根堆中，每次弹出堆顶时，将堆顶的节点合并到合并链表中
// 再将堆顶的下一个节点加入到堆中
func mergeKLists(lists []*ListNode) *ListNode {
	mh := minHeap{nodes: make([]*ListNode, 0)}
	for _, node := range lists {
		if node != nil {
			mh.push(node)
		}
	}
	fakeHead := &ListNode{}
	cur := fakeHead
	for mh.Len() > 0 {
		node := mh.pop()
		cur.Next = node
		cur = node
		if node.Next != nil {
			mh.push(node.Next)
		}
		node.Next = nil
	}
	return fakeHead.Next
}

type minHeap struct {
	nodes []*ListNode
}

func (h *minHeap) Len() int {
	return len(h.nodes)
}
func (h *minHeap) Less(i, j int) bool {
	return h.nodes[i].Val < h.nodes[j].Val
}
func (h *minHeap) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}
func (h *minHeap) Push(x interface{}) {
	h.nodes = append(h.nodes, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	v := h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	return v
}
func (h *minHeap) push(x *ListNode) {
	heap.Push(h, x)
}
func (h *minHeap) pop() *ListNode {
	return heap.Pop(h).(*ListNode)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
