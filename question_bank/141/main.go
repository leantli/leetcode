package main

// https://leetcode.cn/problems/linked-list-cycle/
// 141. 环形链表

// 我们知道，在一个环形跑道中，两个速度不一致的人
// 从同一起点出发，总会相遇，而如果是直线跑到就不会相遇
// 那么显然我们可以用两个指针去实现这样的思路
func hasCycle(head *ListNode) bool {
	one, two := head, head
	for two != nil && two.Next != nil {
		two = two.Next.Next
		one = one.Next
		if one == two {
			return true
		}
	}
	return false
}

// // 常规情况下应该是存储每个节点在一个 set 中，遍历这个链表
// // 遍历过程中如果遇到 set 已有的就说明链表中存在环
// // 但是这种情况还额外使用了存储空间，并且其实很大程度上
// // 存的节点只会有第一个重复的节点会被使用到，其他都不会被再利用
// func hasCycle(head *ListNode) bool {
// 	s := make(map[*ListNode]struct{})
// 	for head != nil {
// 		if _, ok := s[head]; ok {
// 			return true
// 		}
// 		s[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return false
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
