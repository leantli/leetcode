package main

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 19. 删除链表的倒数第 N 个结点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 但是现在我还知道了快慢指针，不错子，毕竟递归嵌套调函数的话
// 链表过长可能会爆栈(当然，刷题正常来说不会遇到这种用例)
// 但是两个指针遍历链表就不会有这种可能
// 不过值得注意的是，这里是删除倒数第 n 个结点，而不是返回倒数第 n 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for range make([]int, n+1) {
		if fast == nil {
			return head.Next
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// // 以前我都是这样做的，一个递归，从末尾开始计数是第几个
// // 然后到倒数第 n + 1 时，直接使其 next 指向下下个节点，略过第 n 个结点
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	var findNthFromEnd func(head *ListNode) int
// 	findNthFromEnd = func(head *ListNode) int {
// 		if head == nil {
// 			return 0
// 		}
// 		curCount := findNthFromEnd(head.Next) + 1
// 		if curCount == n+1 {
// 			head.Next = head.Next.Next
// 		}
// 		return curCount
// 	}
// 	curCount := findNthFromEnd(head)
// 	if curCount == n {
// 		return head.Next
// 	}
// 	return head
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
