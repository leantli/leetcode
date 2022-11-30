package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 32-3 从上到下打印二叉树 III
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		var temp []int
		flag := len(res) % 2
		for i := 0; i < length; i++ {
			if flag == 0 {
				temp = append(temp, queue[i].Val)
			} else {
				temp = append(temp, queue[length-i-1].Val)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
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
