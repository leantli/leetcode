package main

// https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 52. 两个链表的第一个公共节点

// 经典双指针题目，如果不记得具体思路，也可以考虑 map 解决

// 两个指针分别指向 A 和 B，各自先遍历完自身的链表，再从另一条链表的头部开始遍历，直到两节点相等则为重合处
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for headA != headB {
		if headA != nil {
			headA = headA.Next
		} else {
			headA = B
		}
		if headB != nil {
			headB = headB.Next
		} else {
			headB = A
		}
	}
	return headA
}

type ListNode struct {
	Val  int
	Next *ListNode
}
