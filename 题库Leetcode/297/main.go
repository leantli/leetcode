package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
// 297. 二叉树的序列化与反序列化

// 常规来说应该是序列化时一次前序一次中序
// 反序列化时基于前中序重构二叉树
// 但是这样太麻烦了，其实我们可以依靠一次带 nil 的前序遍历，依赖nil提供的额外信息
// 只依赖前/后序遍历即可完成二叉树的重建
type Codec struct{}

func Constructor() Codec { return Codec{} }

// 序列化，这里注意，虽然同样是前序遍历，但这次我们要把 nil 也填入
// 这样叫做序列化，nil 值也能提供信息，如果仅仅只是前序遍历，就还需要中序遍历才能重建二叉树
func (this *Codec) serialize(root *TreeNode) string {
	prePath := make([]string, 0)
	var preTravel func(root *TreeNode)
	preTravel = func(root *TreeNode) {
		if root == nil {
			prePath = append(prePath, "X")
			return
		}
		prePath = append(prePath, strconv.Itoa(root.Val))
		preTravel(root.Left)
		preTravel(root.Right)
	}
	preTravel(root)
	return strings.Join(prePath, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	prePath := strings.Split(data, ",")
	var rebuild func() *TreeNode
	rebuild = func() *TreeNode {
		if prePath[0] == "X" {
			prePath = prePath[1:]
			return nil
		}
		val, _ := strconv.Atoi(prePath[0])
		prePath = prePath[1:]
		root := &TreeNode{Val: val}
		root.Left = rebuild()
		root.Right = rebuild()
		return root
	}
	return rebuild()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
