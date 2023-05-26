package main

// https://leetcode.cn/problems/delete-node-in-a-linked-list/
// 237. 删除链表中的节点

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 按刚才的思路跑了一遍，结果击败百分比非常之低，虽然能过但显然还有问题
// 反思一下，有必要把所有后面的值都重新拷贝吗？其实没必要
// 我们只需要将下一个的值复制到当前节点，再删除下一个节点即可
// 并且看到一个很恐怖评论:
// 如何让自己在世界上消失，但又不死？ —— 将自己完全变成另一个人，再杀了那个人就行了
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// func deleteNode(node *ListNode) {
// 	// 不能操作前面的节点，那只能每次都将后一个节点的值
// 	// 赋值到当前节点，并且将最后一个节点归为 nil
// 	for node.Next.Next != nil {
// 		node.Val = node.Next.Val
// 		node = node.Next
// 	}
// 	node.Val = node.Next.Val
// 	node.Next = nil
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
