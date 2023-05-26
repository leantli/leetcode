package main

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 34. 二叉树中和为某一值的路径

func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	// 先 dfs 得到所有到达叶子节点的路径，最后判断是否等于给定目标和
	var dfs func(root *TreeNode, target, sum int, tmp []int)
	dfs = func(root *TreeNode, target, sum int, path []int) {
		if root == nil {
			return
		}
		// 当节点的左右都为空时，说明该节点为叶子节点
		if root.Left == nil && root.Right == nil {
			if sum+root.Val == target {
				path = append(path, root.Val)
				// temp := make([]int, len(path))
				// copy(temp, path) 这里有点丑
				res = append(res, append([]int{}, path...))
				path = path[:len(path)-1]
			}
			return
		}
		path = append(path, root.Val)
		dfs(root.Left, target, sum+root.Val, path)
		dfs(root.Right, target, sum+root.Val, path)
		path = path[:len(path)-1]
	}
	dfs(root, target, 0, []int{})
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
