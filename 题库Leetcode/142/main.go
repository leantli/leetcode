package main

// https://leetcode.cn/problems/linked-list-cycle-ii/
// 142. 环形链表 II

// 但是以前做过这道题，这道题还是可以用快慢指针解决
// 当有环的时候，快慢指针一定会相遇，无环时快指针会先到链表尾
// 但是有环时怎么找到环的第一个节点？
// 假设慢指针走了 k 步，那么快指针就走了 2k 步
// 并且我们将他们相遇时走的路程分成三个端
// 入环前的路程为 a，入环后慢指针走的路程为 b，此时 a+b=k
// 快指针走过的路程为 a+b+b+x=2k，此时 2a+2b = a+2b+x
// 显然 x = a，而这个 x 的路程，就是相遇点到环的第一个节点的路程
// 让快指针回到初始节点，并且令其开始每次走一步
// 快慢指针再相遇的点，就是入环的第一个节点，因为 x=a
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast != slow {
			continue
		}
		fast = head
		for fast != slow {
			fast = fast.Next
			slow = slow.Next
		}
		return fast
	}
	return nil
}

// // 返回开始入环的第一个节点，如果是快慢，其实不一定会在环的第一个节点就相遇
// // 显然，还是 set 最容易想到以及写出
// func detectCycle(head *ListNode) *ListNode {
// 	set := make(map[*ListNode]struct{})
// 	for head != nil {
// 		if _, ok := set[head]; ok {
// 			return head
// 		}
// 		set[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return nil
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
