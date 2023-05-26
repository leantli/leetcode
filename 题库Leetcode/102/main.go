package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal/
// 102. 二叉树的层序遍历

// 层序遍历，单靠递归左右肯定是都不行的，因为会直接深入到某一边的叶子节点
// 因此我们显然要想其他方法
// 一层一层的遍历，那我们显然可以借助数据结构存储每一个节点的左右子节点
// 并且保证每一层都从左到右逐个遍历
// 比如先存入根结点的左右子节点，数组？
// 然后从数组首位取，如此把数组替换成队列？每遍历到一个节点，就将其存入队列
// 用到时弹出，存时先左后右，这样出队列时，也是按得从左到右
// 但是这样显然没法对节点分层
// 这里可以每到下一层是先取一个长度len，保证每一层都只遍历对应层数的节点数
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		raw := make([]int, 0)
		for i := 0; i < size; i++ {
			raw = append(raw, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		res = append(res, raw)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
