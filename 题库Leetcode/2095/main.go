package main

// https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/
// 2095. 删除链表的中间节点

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
// 慢指针就在链表的1/2处，这里我们不借助前节点，对当前节点进行删除
// copy 下一个节点的值并删除下一个节点
// 但是我们要对传入链表只有 1,2 个结点进行特殊处理
// 因为传入链表长度为[1,2]时，要删除的节点是尾结点
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow.Next == nil {
		head.Next = nil
		return head
	}
	slow.Val = slow.Next.Val
	slow.Next = slow.Next.Next
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
