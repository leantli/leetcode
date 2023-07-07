package main

// https://leetcode.cn/problems/path-sum-iii/
// 437. 路径总和 III

// 二刷
// 二叉树的前缀和？先算算
func pathSum(root *TreeNode, targetSum int) int {
	// 需要一个 map 存储当前已经计算到的节点的路径和，map 存放的 key-value 指 路径和-出现次数
	// 由于 路径方向必须是向下的，我们可以采用 前序遍历，保证晚计算的都在之前计算的节点的下面
	// 此时我们要计算是否满足 targetSum，需要 curSum - map 中所有其他已有的值 == target，符合则 res+= map中对应的value
	// 不过这样需要每次都遍历全部 map，并不好，我们可以 curSum - target = map 中满足条件的 key，基于此无需遍历
	var res int
	m := map[int]int{0: 1}
	var getSum func(root *TreeNode, curSum int)
	getSum = func(root *TreeNode, curSum int) {
		if root == nil {
			return
		}
		curSum += root.Val
		res += m[curSum-targetSum]
		m[curSum]++
		getSum(root.Left, curSum)
		getSum(root.Right, curSum)
		m[curSum]--
	}
	getSum(root, 0)
	return res
}

// // 在 112. 路径总和 我们求过根结点到叶子节点路径和
// // 但这次并没有要求从根结点到叶子节点，即路径上任意连通的节点和都可以，只要满足条件即可计算
// // 那我们显然可以在 112 的基础上，减少掉对叶子节点的判断，只需sum相等即可
// // 并且 112 我们固定是从根结点开始的，但本次显然从其他节点开始也是可以的，因此两个 dfs
// // 一个是 dfs 每个节点作为根结点，另一个 dfs 就是 112 题的解法，计算该节点到叶子节点的路径上有多少满足条件的情况
// func pathSum(root *TreeNode, targetSum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res := getNodeSum(root, targetSum)
// 	// 这里去不断递归，选取路径的不同起始节点
// 	res += pathSum(root.Left, targetSum)
// 	res += pathSum(root.Right, targetSum)
// 	return res
// }

// // 获取这个节点到叶子节点，这一路有多少路径是满足条件的
// // 这里路径的结尾都是叶子节点，因此我们需要另一个函数，去确保起始节点的不同
// func getNodeSum(node *TreeNode, targetSum int) int {
// 	if node == nil {
// 		return 0
// 	}
// 	l := getNodeSum(node.Left, targetSum-node.Val)
// 	r := getNodeSum(node.Right, targetSum-node.Val)
// 	if node.Val == targetSum {
// 		return l + r + 1
// 	}
// 	return l + r
// }

// 但是显然，这个嵌套 dfs，在一定程度上是有不少重复计算的
// 也会一下子想到二叉树的前缀和的感觉？
// bfs 计算各个和，并存入一个数组中，然后按照二叉树数组的性质去相减取结果
// 但是实现起来挺麻烦的, 还是继续考虑递归解决
// 先整个前序遍历把各个节点的前缀和算出来再说
// 常规来说，算出前序和之后，我们就需要将其与其之前的前序和逐个相减
// 求中间路径的节点之和，再判断这些节点之和是否等于 targetSum
// 但这显然太过麻烦，我们可以用 map 存储达到这个节点的路径的前序和
// 正常情况下，我们都是遍历 map，用当前总和减去各个前序和，判断是否能得到 targetSum
// 但是这样一来我们需要遍历整个 map，较为麻烦，这里我们可以直接用当前总和 - targetSum
// 然后直接查看 map 中是否有这个和的值，有的话，就说明有这个路径和的路径都是可取的
// 显然，这里的 map key-value<-->前序和-出现次数
// func pathSum(root *TreeNode, targetSum int) int {
// 	var res int
// 	preSumMapfrequency := map[int]int{0: 1}
// 	var getPreSum func(root *TreeNode, curSum int)
// 	getPreSum = func(root *TreeNode, curSum int) {
// 		if root == nil {
// 			return
// 		}
// 		curSum += root.Val
// 		res += preSumMapfrequency[curSum-targetSum]
// 		preSumMapfrequency[curSum]++
// 		getPreSum(root.Left, curSum)
// 		getPreSum(root.Right, curSum)
// 		preSumMapfrequency[curSum]--
// 	}
// 	getPreSum(root, 0)
// 	return res
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
