package main

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
// 116. 填充每个节点的下一个右侧节点指针

// 结果看了解答发现还有更巧妙的递归解法
// 这里就是优先不断地去递归连接不在同一棵树的右节点和左节点
// 但是感觉不太好想到这个递归思路
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		l := root.Left
		r := root.Right
		for l != nil {
			l.Next = r
			l = l.Right
			r = r.Left
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root
}

// // 显然这道题还有更巧妙的解法，那么除了基于递归去连接两个节点
// // 我们还要注意到这个next节点，我们可以在上一层有next的情况下，node.r.next = node.next.l
// // 如此一来，递归的重复操作就被避免，并且也能够一次遍历后就将全部的节点连通
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	// 从每一层的第一个节点开始向下一层的第一个节点步进
// 	for first := root; first.Left != nil; first = first.Left {
// 		// 从每一层的第一个节点开始向同一层的右侧节点步进
// 		for cur := first; cur != nil; cur = cur.Next {
// 			cur.Left.Next = cur.Right
// 			if cur.Next != nil {
// 				cur.Right.Next = cur.Next.Left
// 			}
// 		}
// 	}
// 	return root
// }

// 涉及到左右子树的节点间的处理，显然是需要额外一个函数做自上而下的递归
// 这个递归函数需要两个参数，即两个节点，函数内部就做两件事，一件事是串联这两个节点
// 另一件事就是递归这两个节点应该要连接起来的子节点
// 这个虽然好理解，但显然后面的递归是存在重复的操作，这个解法不够好
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	var connect func(l, r *Node)
// 	connect = func(l, r *Node) {
// 		if l == nil || r == nil {
// 			return
// 		}
// 		l.Next = r
// 		connect(l.Left, l.Right)
// 		connect(l.Right, r.Left)
// 		connect(r.Left, r.Right)
// 	}
// 	connect(root.Left, root.Right)
// 	return root
// }

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
