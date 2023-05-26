package main

// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 7. 重建二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 依靠前序遍历和中序遍历的特性构建二叉树
// 我们知道前序遍历，是先 root，再 left 和 right
//        中序遍历，是先 left，再 root，最后 right
// 根据这个性质，我们首先能够确认清楚 root 和 left
// 并且后面的 right 也可以作为新的 root 递进判断

// 再说构建，从根节点开始？pre 的第一个肯定就是根节点。
// 我们需要确切知道 root 的左子树的整个区间，便于定位 right 开始的下标位置
// 由于题目不包含重复数字，我们可以采用 map 存放 root 在 pre 中的下标位置，遍历中序，找到 root，此时下标为 inorder_root_index
// 此时下标减去 pre 中 root 的下标，可得到 左子树的节点的区间 [inorder_root_index-diff, inorder_root_index)
// 此时 pre 中 root 下标 + 左子树节点数量 + 1，就是右子树的 root 节点，再重复以上过程，就可以得到全部的区间
// 没搞出来，这里其实弄岔了一个点，主要是创建节点和节点之间的连接，没有很好地通过递归完成
// 可以直接这样，0下标一定为本次构建的 root 节点，每次 buildTree 就把左右子树的区间传入，这样保证左右子节点都是左右子树的 root

// 简易理解
// 先根据 pre[0] 建好 root，再在 inorder 的基础上，确定左子树节点的数量，确定左子树的区间
// 将对应区间的 preorder 和 inorder 传入，构建局部树,并且返回局部树的 root 作为当前 root 的左子节点
// 右子树也是同样的思路
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	inorderRootIndex := findInorderIndex(preorder[0], inorder)
	root.Left = buildTree(preorder[1:inorderRootIndex+1], inorder[:inorderRootIndex])
	root.Right = buildTree(preorder[inorderRootIndex+1:], inorder[inorderRootIndex+1:])
	return root
}

func findInorderIndex(root int, inorder []int) int {
	for i, num := range inorder {
		if num == root {
			return i
		}
	}
	return -1
}

// java 解法，不像 go 切片可以便捷的复用一定长度的切片，因此只能前序和中序数组都要传，避免重复拷贝，并且传入左右子树对应的前序和中序边界。。。
// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	inHit := make(map[int]int, len(inorder))
// 	for i, num := range inorder {
// 		inHit[num] = i
// 	}
// 	var build func(preorder, inorder []int, pl, pr, il, ir int) *TreeNode
// 	build = func(preorder, inorder []int, pl, pr, il, ir int) *TreeNode {
// 		if pl > pr {
// 			return nil
// 		}
// 		root := &TreeNode{
// 			Val: preorder[pl],
// 		}
// 		// 获取 root 在中序数组中的下标
// 		inRootIndex := inHit[preorder[pl]]
// 		// 计算左子树节点的数量
// 		leftSize := inRootIndex - il
// 		root.Left = build(preorder, inorder, pl+1, pl+leftSize, il, inRootIndex-1)
// 		root.Right = build(preorder, inorder, pl+leftSize+1, pr, inRootIndex+1, pr)
// 		return root
// 	}
// 	n := len(preorder)
// 	return build(preorder, inorder, 0, n-1, 0, n-1)
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
