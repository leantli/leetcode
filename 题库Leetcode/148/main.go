package main

// https://leetcode.cn/problems/sort-list/submissions/
// 148. 排序链表

// 类似归并的思路？
// 通过快慢指针不断拆分成两个链表， 自顶向下拆分最终只剩一个节点
// 接着自底向上归并排序，因为底层的链表都是已排序链表，因此合并两个排序链表思路即可
func sortList(head *ListNode) *ListNode {
	return divide(head)
}

func divide(head *ListNode) *ListNode {
	// 没有节点或只有一个节点就可以直接返回了
	if head == nil || head.Next == nil {
		return head
	}
	// 由于我们需要基于中间节点拆开链表，因此让 fast 先走两步
	// 这样 fast 走出循环时，slow.Next 就是中间节点, 既能基于 slow.Next 进行下一步拆分，又能令 slow.Next = nil，将前半部分拆开
	fast, slow := head.Next.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	r := divide(slow.Next)
	slow.Next = nil
	l := divide(head)
	return merge(l, r)
}

func merge(a, b *ListNode) *ListNode {
	fakeHead := &ListNode{}
	cur := fakeHead
	for a != nil && b != nil {
		if a.Val > b.Val {
			cur.Next = b
			b = b.Next
		} else {
			cur.Next = a
			a = a.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}
	return fakeHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
