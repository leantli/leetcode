package main

// https://leetcode.cn/problems/path-sum-ii/
// 113. 路径总和 II

// 在 112. 路径总和 的基础上，要求找到所有满足条件的路径
// bfs 的方法就不太方便使用了？毕竟没有一路记录对应的节点，确定了之后需要返回去查看对应的节点，不过也还好
// 递归的方式先试试
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	curArr := make([]int, 0)
	var findPathSum func(root *TreeNode, targetSum int)
	findPathSum = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		curArr = append(curArr, root.Val)
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			res = append(res, append([]int{}, curArr...))
		}
		findPathSum(root.Left, targetSum-root.Val)
		findPathSum(root.Right, targetSum-root.Val)
		curArr = curArr[:len(curArr)-1]
	}
	findPathSum(root, targetSum)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
