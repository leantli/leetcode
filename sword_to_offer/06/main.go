package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 先递归到最末尾的节点，再逐步 append 数组

func reversePrint(head *ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	arr := reversePrint(head.Next)
	arr = append(arr, head.Val)
	return arr
}

type ListNode struct {
	Val  int
	Next *ListNode
}
