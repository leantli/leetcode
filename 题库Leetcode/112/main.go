package main

// https://leetcode.cn/problems/path-sum/
// 112. 路径总和

// 根结点到叶子节点路径上的值综合刚好等于 target 则返回 true
// 显然这里用不了自底向上，因为我们需要确保每个叶子节点的值都被取到，都被跑到
// 而自底向上会舍弃一半的可能值
// 显然自顶向下+前序，到叶子节点时，判断是否值为 targetSum，路径中不判断
// 这类相加取值判断是否等于 target，都可以用 target-取到的值是否等于 0 替代
// 减少掉一个 curSum 的中间态，并且在本题中可以很巧妙地结合递归实现代码简洁且思路清晰

// 官解思路：假定从根节点到当前节点的值之和为 val，
// 我们可以将这个大问题转化为一个小问题：是否存在从当前节点的子节点到叶子的路径，满足其路径和为 sum - val
// 不难发现这满足递归的性质，若当前节点就是叶子节点，那么我们直接判断 sum 是否等于 val 即可
// 因为路径和已经确定，就是当前节点的值，我们只需要判断该路径和是否满足条件
// 若当前节点不是叶子节点，我们只需要递归地询问它的子节点是否能满足条件即可
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	// 求其左右节点的路径和是否等于 targetSum - root.Val
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// // 根结点到叶子节点路径上的值综合刚好等于 target 则返回 true
// // 显然这里用不了自底向上，因为我们需要确保每个叶子节点的值都被取到，都被跑到
// // 而自底向上会舍弃一半的可能值
// // 显然自顶向下+前序，到叶子节点时，判断是否值为 targetSum，路径中不判断
// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	var findTargetSum func(root *TreeNode, sum int)
// 	var res bool
// 	findTargetSum = func(root *TreeNode, sum int) {
// 		if root == nil || res {
// 			return
// 		}
// 		sum += root.Val
// 		if root.Left == nil && root.Right == nil && sum == targetSum {
// 			res = true
// 		}
// 		findTargetSum(root.Left, sum)
// 		findTargetSum(root.Right, sum)
// 	}
//     findTargetSum(root, 0)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
