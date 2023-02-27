package main

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 160. 相交链表

// 但是我们仔细观察，会发现如果两个链表相交，那么相交后的链表长度一定相等，设为 c
// 此时 A 链表相交前的长度为 a, B 链表相交前的长度为 b
// 两个节点分别从 A,B 头部开始遍历，遍历结束后，再从另一个链表头开始遍历
// 最终他们会在相交处相遇，因为两个指针都共同走了 a+b+c 的长度
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

// // 这里最常规仍然是使用 set 存储已出现过的节点
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	set := make(map[*ListNode]struct{})
// 	for headA != nil {
// 		set[headA] = struct{}{}
// 		headA = headA.Next
// 	}
// 	for headB != nil {
// 		if _, ok := set[headB]; ok {
// 			return headB
// 		}
// 		headB = headB.Next
// 	}
// 	return nil
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
