package main

// https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 22. 链表中倒数第k个节点

// 后面又想起最基础的快慢指针做法
func getKthFromEnd(head *ListNode, k int) *ListNode {
	l, r := head, head
	// 让 r 先走 k 步
	for i := 0; i < k; i++ {
		r = r.Next
	}
	for r != nil {
		l = l.Next
		r = r.Next
	}
	return l
}

// 第一想法是递归处理一下，从尾部开始往前数即可
// 这里 count 为全局遍历，每次都初始化一下，不然 leetcode 复用会导致报错
// var count int

// func getKthFromEnd(head *ListNode, k int) *ListNode {
// 	count = 0
// 	return get(head, k)
// }

// func get(head *ListNode, k int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	res := getKthFromEnd(head.Next, k)
// 	count++
// 	if count == k {
// 		res = head
// 		return res
// 	}
// 	if count > k {
// 		return res
// 	}
// 	return head
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
