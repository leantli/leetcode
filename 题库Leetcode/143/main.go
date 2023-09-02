package main

// https://leetcode.cn/problems/reorder-list/description/
// 143. 重排链表

// 没想到特别节省空间时间的方法，感觉还是得借助数组存储对应的下标和节点
// 存储后根据下标重建链表
func reorderList(head *ListNode) {
	lists := make([]*ListNode, 0)
	for head != nil {
		lists = append(lists, head)
		head = head.Next
	}
	l, r := 0, len(lists)-1
	for l < r {
		lists[l].Next = lists[r]
		l++
		if l == r {
			break
		}
		lists[r].Next = lists[l]
		r--
	}
	lists[l].Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
