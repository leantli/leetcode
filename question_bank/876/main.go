package main

// https://leetcode.cn/problems/middle-of-the-linked-list/
// 876. 链表的中间结点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 总不能先遍历一遍看看全长再去根据n/2去删中间这个数吧
// 想到快慢指针，令慢指针走一步，快指针走两步
// 此时慢指针走的步数是快指针的1/2，当快指针走到尾部时
// 慢指针就在链表的1/2处
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
