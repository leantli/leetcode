package main

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 32-3 从上到下打印二叉树 III
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		temp := make([]int, 0, length)
		flag := len(res) & 1
		for i := 0; i < length; i++ {
			// 正常地将 queue[i] 的左右子节点入栈
			// 但是对于每一行的数值却不一定按照顺序，如果是偶数行，正常顺序；奇数行则从后往前
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
