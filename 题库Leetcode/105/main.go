package main

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 105. 从前序与中序遍历序列构造二叉树

// 根据前序和中序遍历构造二叉树，重构二叉树必须有两个序列且包含中序
// 因为中序才能根据根结点的位置，确定哪部分区间是左子树和右子树的
// 并且我们能很清楚的想到，根据前序或者后序，我们能找到当前两个序列的根结点
// 但是要继续得到左右子节点，我们得根据这个根结点去拆分前序和中序遍历的区间
// 得到左右子树各自的前序和中序遍历数组，才能再继续得到下一个节点
// 再根据这个函数的传参和返回结果，我们应该想到，这是一个可以自上而下的递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	var rootInIndex int
	// 这里注意 root 在中序遍历中的下标位置，还等同于左子树的节点数量
	// 这个性质，能够帮助我们在前后序遍历中，方便地划分左右子树的区间
	for i, v := range inorder {
		if v == root.Val {
			rootInIndex = i
			break
		}
	}
	root.Left = buildTree(preorder[1:1+rootInIndex], inorder[:rootInIndex])
	root.Right = buildTree(preorder[1+rootInIndex:], inorder[rootInIndex+1:])
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
