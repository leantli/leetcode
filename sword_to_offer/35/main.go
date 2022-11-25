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

// 复制一个复杂链表，该链表的节点中存在一个 randowm 指针指向任意节点或 nil
// 这里的难点就在于 random 指针可能指向后面的节点，此时我们还没复制到后面的节点，这个地方如何处理
// 瞎写出来的做法 0ms，但是内存只击败了 24.37%

// func copyRandomList(head *Node) *Node {
// 	var bucketOriginal = make(map[*Node]int)
// 	var bucketNew = make(map[int]*Node)
// 	var index int
// 	for head != nil {
// 		bucketOriginal[head] = index
// 		node := &Node{
// 			Val: head.Val,
// 		}
// 		bucketNew[index] = node
// 		index++
// 		head = head.Next
// 	}
// 	for node, i := range bucketOriginal {
// 		bucketNew[i].Next = bucketNew[i+1]
// 		if node.Random == nil {
// 			bucketNew[i].Random = nil
// 			continue
// 		}
// 		bucketNew[i].Random = bucketNew[bucketOriginal[node.Random]]
// 	}
// 	return bucketNew[0]
// }

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
