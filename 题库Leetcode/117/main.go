package main

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
// 117. 填充每个节点的下一个右侧节点指针 II

// 看了大佬的题解，思路是一样，但是实现上多了不少小处理，更简洁明了
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	first := root
	for first != nil {
		// 这里大佬使用了一个伪头结点，起到了两个作用
		// 1. 原先我需要判断pre是否为空，空的话另做处理
		// 现在 pre 一定不为空，节省了多余的判断，并且也能正常串联原先的节点
		// 2. 有了伪头结点，其next节点就是下一层的第一个节点，节省掉了我原先麻烦的寻找下一层首节点的操作
		// 除了这个伪头结点之外，大佬还减少了连接两个相近节点的判断代码，两个 if 解决，比我原先的判断少了
		nextLevelHead := &Node{Val: 0}
		pre := nextLevelHead
		for cur := first; cur != nil; cur = cur.Next {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = cur.Left
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = cur.Right
			}
		}
		first = nextLevelHead.Next
	}
	return root
}

// // 和 116 的区别在于，这次给的树并不一定是完全二叉树
// // 可能存在左子树的左节点和右子树的右节点相连通的情况
// // 这种情况，我们纵览刚才 116 的三种解法，感觉递归不是很方便
// // 先基于 next 指针尝试一下
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return root
// 	}
// 	first := root
// 	for first != nil {
// 		// 这里要连接下一层的两个相近的节点
// 		var pre *Node
// 		for cur := first; cur != nil; cur = cur.Next {
// 			if cur.Left == nil && cur.Right == nil {
// 				continue
// 			}
// 			if cur.Left != nil && cur.Right != nil {
// 				if pre != nil {
// 					pre.Next = cur.Left
// 				}
// 				cur.Left.Next = cur.Right
// 				pre = cur.Right
// 				continue
// 			}
// 			if pre == nil {
// 				if cur.Left != nil {
// 					pre = cur.Left
// 				} else {
// 					pre = cur.Right
// 				}
// 				continue
// 			}
// 			if cur.Left != nil {
// 				pre.Next = cur.Left
// 				pre = cur.Left
// 			} else {
// 				pre.Next = cur.Right
// 				pre = cur.Right
// 			}
// 		}
// 		// 找到每层第一个
// 		for first != nil && first.Left == nil && first.Right == nil {
// 			first = first.Next
// 		}
// 		if first == nil {
// 			break
// 		}
// 		if first.Left != nil {
// 			first = first.Left
// 		} else {
// 			first = first.Right
// 		}
// 	}
// 	return root
// }

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
