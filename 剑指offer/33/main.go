package main

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 33. 二叉搜索树的后序遍历序列

// 二叉搜索树，右大左小
// 后序遍历，左->右->中
// 左 < 右， 右 > 中
// 不会连续地大于下一个数？（X）

// 这里还要再观察一下！后序遍历，最后一个数，一定是根节点
// 此时，从前往后遍历，遇到第一个大于根节点的下标，到根节点前一个下标，都是右子树节点区间，其他的都是左子树区间；
// 两个子树区间的最后一个数也分别是根节点的左右子节点
func verifyPostorder(postorder []int) bool {
	// 当 postorder 小于等于 2 时，我们认为该序列一定满足后序遍历，无需再拆分
	if len(postorder) <= 2 {
		return true
	}
	rootIndex := len(postorder) - 1
	// 这里 rightBeginIndex 设为 rootIndex，不要设为 0 或其他，否则没找到比 root 大的节点时分界会出现问题
	var rightBeginIndex int = rootIndex
	for i, num := range postorder {
		if num > postorder[rootIndex] {
			rightBeginIndex = i
			break
		}
	}
	// 此时我们已经区分好左右子树区间了，可以通过其他操作来判断是否符合题目要求
	// 右子树节点一定都比根节点大
	for _, num := range postorder[rightBeginIndex:rootIndex] {
		if num < postorder[rootIndex] {
			return false
		}
	}
	// 递归地去区分左右子树，以及判断
	return verifyPostorder(postorder[:rightBeginIndex]) && verifyPostorder(postorder[rightBeginIndex:rootIndex])
}
