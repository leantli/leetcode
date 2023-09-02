package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
// 82. 删除排序链表中的重复元素 II

// 直接删除重复的全部元素，因此我们需要使用 pre 存放上一个不重复元素的节点
// 判断 pre.Next 和 pre.Next.Next 是否重复，是则记录重复值为 v，并且不断判断 pre.Next.Val == v
// 是则 pre.Next 跳过该节点，直接 pre.Next = pre.Next.Next，但此时 pre 仍然不变，还是上一个不重复元素
// 继续 for pre.Next != nil && pre.Next.Val == v，将所有重复的节点都跳过
// 不重复再移动 pre，pre = pre.Next
func deleteDuplicates(head *ListNode) *ListNode {
	fakeHead := &ListNode{Val: -1, Next: head}
	pre := fakeHead
	for pre.Next != nil && pre.Next.Next != nil {
		if pre.Next.Val == pre.Next.Next.Val {
			v := pre.Next.Val
			for pre.Next != nil && pre.Next.Val == v {
				pre.Next = pre.Next.Next
			}
		} else {
			pre = pre.Next
		}
	}
	return fakeHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
