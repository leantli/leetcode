package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/submissions/
// 83. 删除排序链表中的重复元素

// 本身链表是有序的，要删除重复的元素，首先要保留当前遍历到的数值的前一个节点
// 判断当前节点与前一个节点是否重复，重复则 pre.Next = cur.Next，跳过当前节点，并且更新 cur = cur.Next
// 如果没遇到重复的，就更新 pre 节点为当前节点，并且 cur = cur.Next
// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	h := head
// 	pre := head
// 	cur := head.Next
// 	for cur != nil {
// 		if pre.Val == cur.Val {
// 			pre.Next = cur.Next
// 		} else {
// 			pre = cur
// 		}
// 		cur = cur.Next
// 	}
// 	return h
// }

// 或者直接基于 pre 进行判断
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	for pre != nil {
		if pre.Next != nil && pre.Val == pre.Next.Val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
