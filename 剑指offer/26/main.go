package main

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 26. 树的子结构
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 印象中是有非常简化的写法，不过还是一步步来
// 先考虑如何判断是否是子结构，再遍历，遍历时再进行判断
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// 题目显示规定 B 为空，则不为 A 的子结构
	if B == nil {
		return false
	}
	return search(A, B)
}

// 比较判断是否为子结构
func compare(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		// A [1,3] B[1] 这种
		return true
	}
	if A == nil {
		// 此时 B 还有， A 已无节点，则失败
		return false
	}
	// 此时比较 A 和 B 的值，顺便 compare 二者的左右子节点
	return A.Val == B.Val && compare(A.Left, B.Left) && compare(A.Right, B.Right)
}

// 遍历所有节点并比较
func search(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		// B 不为空，此时 A 为空，直接返回
		return false
	}
	if compare(A, B) {
		// 当 A 和 B 值相等时，进行 compare，成了就返回 true
		// 写完发现这个 比较 其实可以和下面的 return || 合起来, 上面的 A == nil 可以和最开始的  B == nil 合起来
		return true
	}
	// 不相等则继续搜
	return search(A.Left, B) || search(A.Right, B)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 简化后
// func isSubStructure(A *TreeNode, B *TreeNode) bool {
// 	if A == nil || B == nil {
// 		return false
// 	}
// 	// 比较，失败则继续搜左右子节点
// 	return compare(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
// }

// func compare(A *TreeNode, B *TreeNode) bool {
// 	if B == nil {
// 		return true
// 	}
// 	if A == nil {
// 		return false
// 	}
// 	return A.Val == B.Val && compare(A.Left, B.Left) && compare(A.Right, B.Right)
// }
