package main

// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 复杂链表的复制

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// 官方解法 1 回溯+哈希备忘录
// 新建一个 NewNode 后就存入 map 中
// 对 next/random 到的节点就继续 copy
// 已存在 newNode 就直接返回新的节点
// 由于有 next 存在，必定会遍历完全部的节点
var nodeMap = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	nodeMap = make(map[*Node]*Node)
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if newNode, ok := nodeMap[node]; ok {
		return newNode
	}
	newNode := &Node{Val: node.Val}
	nodeMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

// // 暴力一点，遍历两遍，并将原结点和新结点都存入一个 map 作为映射
// func copyRandomList(head *Node) *Node {
// 	fakeHead := &Node{}
// 	curNew := fakeHead
// 	m := make(map[*Node]*Node)
// 	for head != nil {
// 		curNew.Next = &Node{
// 			Val:    head.Val,
// 			Random: head.Random,
// 		}
// 		curNew = curNew.Next
// 		m[head] = curNew
// 		head = head.Next
// 	}
// 	curNew = fakeHead.Next
// 	for curNew != nil {
// 		curNew.Random = m[curNew.Random]
// 		curNew = curNew.Next
// 	}
// 	return fakeHead.Next
// }

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
