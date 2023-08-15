package main

// https://leetcode.cn/problems/house-robber-iii/
// 337. 打家劫舍 III

// 二刷
// 假设我们遍历到某一个节点，一定是先看，要不要取其左右子节点，如果左右子节点的和更大，那不选取当前节点
// 而且我们需要返回某一节点取或不取各自的最大金额，否则我们无法判断，本节点是否应该取
// 那么此时 定义 dp[i][j], i 表示第 i 个节点，j 只有 0 和 1，分别表示第 i 个节点取或不取
// 则此时 dp[i][0] = max(dp[i+1][1] + dp[i+2][1] + node[i], max(dp[i+1][0], dp[i+1][1]) + max(dp[i+2][0], dp[i+2][1]))
// 取 i 的话，i+1 和 i+2 必定是不取的；不取 i 的话，i+1 和 i+2 可以都取，也可以取其中一个，另一个不取，单纯看金额大小
// 此时注意到，dp[i] 仍然只由 dp[i+1] 和 dp[i+2] 两个子节点的结果决定，因此我们可以通过 dfs
// 分别获取到 某一节点 的左右子节点取与不取的最大金额，再进行判断
func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (yes, no int)
	dfs = func(root *TreeNode) (yes, no int) {
		if root == nil {
			return 0, 0
		}
		leftY, leftN := dfs(root.Left)
		rightY, rightN := dfs(root.Right)
		return root.Val + leftN + rightN, max(leftY, leftN) + max(rightN, rightY)
	}
	rootY, rootN := dfs(root)
	return max(rootY, rootN)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 为什么不不能像 198 一维数值这样的状态转移方程{} dp[i] = max(dp[i-1], dp[i-2]+nums[i]) }，
// // 而是要额外定义第 i 家偷/不偷的状态呢？
// // 首先 198 是一维数组情况，操作 i, i-1, i-2 其实比较方便，因此直接一套状态转移方程即可
// // 而 337 中 i 为父节点，i-1 为两个子节点， i-2 可能是四个孙子节点，一来数量上不太好直接 max，二来不像一维数组取值方便
// // 因而通过定义 f[i] 和 g[i]，分别为第i家偷/不偷进行操作和做转移方程，会更简单
// // 此时我们发现，有了 f 和 g 两个 map 存储不同的取节点情况
// // 我们已经无需在同一函数中额外多处理孙子节点的情况，并且此时已经形成了一个基本的状态转移过程
// // 并且我们发现，每个节点的取与不取，都只取决于其左右子节点的取与不取
// // 完全不需要通过 map 去过多地存储父子节点之外的情况，只需要存储当前节点左右子节点的情况即可
// // 并且我们注意到，我们是基于后序遍历考虑情况，先处理了左右子节点，最后才处理父节点，
// // 因此我们可以获取左右子节点的情况, 直接将函数的返回参数改为该节点取/不取时的最大可盗窃金额
// // 4ms
// func rob(root *TreeNode) int {
// 	var dfs func(node *TreeNode) (used, unused int)
// 	dfs = func(node *TreeNode) (used int, unused int) {
// 		if node == nil {
// 			return 0, 0
// 		}
// 		lu, luu := dfs(node.Left)
// 		ru, ruu := dfs(node.Right)
// 		used = node.Val + luu + ruu
// 		unused = max(lu, luu) + max(ru, ruu)
// 		return used, unused
// 	}
// 	used, unsued := dfs(root)
// 	return max(used, unsued)
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 接着，基于刚才的 map 存储，我们会发现，其实我们只需要关注父节点和两个子节点取不取
// // 之前考虑的三层结构，其实也只是为了便于区分父节点和两个子节点取不取的状态
// // 那么此时我们可以设 f(o) 表示 选择 o 节点情况下，最大偷窃金额为多少
// // 设 g(o) 表示 不选择 o 节点情况下，最大偷窃金额为多少
// // 则此时 当 o 被选中时，o 的左右子节点都不能被选中，因此 f(o) = g(o.l) + g(o.r),
// // 而当 o 不被选中时，o 的左右子节点可以被选中，也可以不被选中
// // 因此 g(o) = max(f(o.l), g(o.l)) + max(f(o.r), g(o.r))
// // 此时我们在递归中无需显性计算孙子节点以区分左右子节点的取或不取，在递归中就会被考虑到
// // 8ms
// func rob(root *TreeNode) int {
// 	f, g := make(map[*TreeNode]int), make(map[*TreeNode]int)
// 	var dfs func(node *TreeNode)
// 	dfs = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		dfs(node.Left)
// 		dfs(node.Right)
// 		f[node] = g[node.Left] + g[node.Right] + node.Val
// 		g[node] = max(f[node.Left], g[node.Left]) + max(f[node.Right], g[node.Right])
// 	}
// 	dfs(root)
// 	return max(f[root], g[root])
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 那么我们能够基于刚才的思路，加上一个 map，做记忆化搜索，减少重复计算
// // 12ms
// func rob(root *TreeNode) int {
// 	var dfs func(node *TreeNode) int
// 	m := make(map[*TreeNode]int)
// 	dfs = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		if v, ok := m[node]; ok {
// 			return v
// 		}
// 		money := node.Val
// 		if node.Left != nil {
// 			money += dfs(node.Left.Left) + dfs(node.Left.Right)
// 		}
// 		if node.Right != nil {
// 			money += dfs(node.Right.Left) + dfs(node.Right.Right)
// 		}
// 		res := max(money, dfs(node.Left)+dfs(node.Right))
// 		m[node] = res
// 		return res
// 	}
// 	return dfs(root)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 这里涉及到节点的三层情况
// // 选了父节点，就不能选两个子节点，选了两个子节点，就不能选父节点和四个孙子节点
// // ps: 最后会超时，显然我们发现，这里面会出现非常多的重复计算，一个父节点会调用到孙子节点和子节点的dfs
// func rob(root *TreeNode) int {
// 	var dfs func(node *TreeNode) int
// // dfs 获取该节点的最大偷窃金额
// 	dfs = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		moneyOfGrandson := node.Val
// 		if node.Left != nil {
// 			moneyOfGrandson += dfs(node.Left.Left) + dfs(node.Left.Right)
// 		}
// 		if node.Right != nil {
// 			moneyOfGrandson += dfs(node.Right.Left) + dfs(node.Right.Right)
// 		}
// 		return max(moneyOfGrandson, dfs(node.Left)+dfs(node.Right))
// 	}
// 	return dfs(root)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
