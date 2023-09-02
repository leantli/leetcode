package main

// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
// 25. K 个一组翻转链表

// k 个一组进行翻转，显然我们首先需要定位到原先 head 和 tail，然后取得 head.pre 和 tail.next
// 保证我们基于 head 和 tail 翻转后，翻转后的链表的新 head 和 tail，能够成功和原先的 head.pre 和 tail.next 连接起来
// 所以先写一个基于 head 和 tail 翻转的函数，返回新的 head 和 tail
func reverseKGroup(head *ListNode, k int) *ListNode {
	fakeHead := &ListNode{Next: head}
	pre := fakeHead
	tail := fakeHead
	for tail != nil {
		// 先找到 k 个节点的尾部
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return fakeHead.Next
			}
		}
		// 记录下 tail 的下一个节点，便于翻转后，新的 tail 连接到 after 节点
		after := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = after
		pre = tail
		head = after
	}
	return fakeHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := head
	cur := head.Next
	for pre != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return tail, head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
