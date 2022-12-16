package main

// https://leetcode.cn/problems/add-two-numbers/?favorite=2cktkvj
// 2. 两数相加

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 当节点还未遍历完，当还有进位时就要继续相加？

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	cur := fakeHead
	var up int
	for l1 != nil || l2 != nil || up != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + up
		up = sum / 10
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
	}
	return fakeHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
