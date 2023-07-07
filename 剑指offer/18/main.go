package main

// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 18. 删除链表的节点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 重新试试递归
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}

// // 先试了一下递归，感觉递归不如直接迭代好处理，还是直接迭代完事，最后特殊处理一下头节点就ok
// func deleteNode(head *ListNode, val int) *ListNode {
// 	if head.Val == val {
// 		return head.Next
// 	}
// 	l, r := head, head
// 	for r != nil {
// 		r = r.Next
// 		if r.Val == val {
// 			l.Next = r.Next
// 			break
// 		}
// 		l = l.Next
// 	}
// 	return head
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
