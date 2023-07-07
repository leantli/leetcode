package main

// https://leetcode.cn/problems/copy-list-with-random-pointer/
// 138. 复制带随机指针的链表

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// 遍历两遍，并将原结点和新结点都存入一个 map 作为映射
func copyRandomList(head *Node) *Node {
	fakeHead := &Node{}
	curNew := fakeHead
	m := make(map[*Node]*Node)
	for head != nil {
		curNew.Next = &Node{
			Val:    head.Val,
			Random: head.Random,
		}
		curNew = curNew.Next
		m[head] = curNew
		head = head.Next
	}
	curNew = fakeHead.Next
	for curNew != nil {
		curNew.Random = m[curNew.Random]
		curNew = curNew.Next
	}
	return fakeHead.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
