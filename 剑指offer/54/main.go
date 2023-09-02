package main

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 54. 二叉搜索树的第k大节点

// 二刷
// 二叉搜索树的第 k 大节点，显然我们可以利用二叉搜索树的特性
// 我们可以采用中序遍历，遍历优先级：右子结点、父节点、左子节点
// 如此依序遍历到的第 k 个节点，就是我们需要的第 k 大节点
func kthLargest(root *TreeNode, k int) int {
	var cnt, res int
	var afterorder func(root *TreeNode)
	afterorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		afterorder(root.Right)
		cnt++
		if cnt == k {
			res = root.Val
		}
		afterorder(root.Left)
	}
	afterorder(root)
	return res
}

// 二叉搜索树，性质右大左小
// 采用中序遍历，递归到的第 k 个结点就是第 k 大
// func kthLargest(root *TreeNode, k int) int {
// 	var res int
// 	var dfs func(root *TreeNode)
// 	dfs = func(root *TreeNode) {
// 		if root == nil || k < 0 {
// 			return
// 		}
// 		dfs(root.Right)
// 		fmt.Printf("%d, k=%d", root.Val, k)
// 		k--
// 		if k == 0 {
// 			res = root.Val
// 			return
// 		}
// 		dfs(root.Left)
// 	}
// 	dfs(root)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
