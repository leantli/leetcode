package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 32-2 从上到下打印二叉树 2

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			temp = append(temp, queue[i].Val)
		}
		queue = queue[length:]
		res = append(res, temp)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
