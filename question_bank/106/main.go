package main

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 106. 从中序与后序遍历序列构造二叉树

// 我们需要明确一个点，要根据序列构造二叉树，必须要有中序遍历
// 因为基于中序遍历，才能确定哪部分是左子树，哪部分是右子树
// 接下来我们看题目以及给出的这个函数的传参和返回结果，我们可以想到，其实这个传参和返回
// 是可以逐步自上往下递归的，每次递归调用函数，我们都能确定一个节点(根结点)
// 并根据这个根结点的位置，为其左右子树的构建，重新划分左右子树的中序和后序遍历数组
// 首先，我们知道，后序遍历的最后一个节点，就是根结点
// 此时再根据这个根结点，在中序遍历中，我们能确定根结点的左右子树的范围
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 || len(postorder) == 1 {
		return root
	}
	// 接下来，我们考虑构建 root 的左右子树
	// 此时，我们需要根据 root 在中序和后序的位置，确定左右子树各自的中序和后序遍历数组
	// 我们先来确定右子树的中序和后序遍历数组区间
	// 中序自然不必多说，根结点的左侧都是左子树的中序遍历数组，根结点的右侧都是右子树的中序遍历数组
	// 右子树的后序遍历数组是，[根结点在中序遍历数组位置+1的值 在后续遍历数组的位置，后序遍历数组总长-1(因为根结点在最后一个)]
	// 左子树的后序遍历数组是，[:根结点在中序遍历数组位置+1的值 在后续遍历数组的位置)
	// 但这显然太过麻烦，因为我们需要确定这个值---根结点在中序遍历数组位置+1的值
	// 事实上，rootInIndex，root所在的位置的index，即是左子树节点的总数量，我们可以巧妙地运用这个数量去划分
	// 右子树的后序遍历数组是，[rootInIndex，后序遍历数组总长-1(因为根结点在最后一个)]
	// 左子树的后序遍历数组是，[:rootInIndex)
	var rootInIndex int
	for i, v := range inorder {
		if v == root.Val {
			rootInIndex = i
			break
		}
	}
	root.Left = buildTree(inorder[:rootInIndex], postorder[:rootInIndex])
	root.Right = buildTree(inorder[rootInIndex+1:], postorder[rootInIndex:len(postorder)-1])
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
