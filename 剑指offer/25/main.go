package main

// https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 25. 合并两个排序的链表

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	cur := fakeHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return fakeHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
