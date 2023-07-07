package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/xu-lie-hua-er-cha-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 37. 序列化二叉树 == 297. 二叉树的序列化与反序列化 https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
// 层序遍历？

// 常规来说应该是序列化时给出一次前序一次中序或者层序遍历
// 反序列化时基于前中序重构二叉树
// 但是这样太麻烦了，其实我们可以依靠一次带 nil 的前序遍历，依赖nil提供的额外信息
// 只依赖前/后序遍历即可完成二叉树的重建

// 序列化，这里注意，虽然同样是前序遍历，但这次我们要把 nil 也填入
// 这样叫做序列化，nil 值也能提供信息，如果仅仅只是前序遍历，就还需要中序遍历才能重建二叉树
func serialize(root *TreeNode) string {
	strs := make([]string, 0)
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			strs = append(strs, "X")
			return
		}
		strs = append(strs, strconv.Itoa(root.Val))
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return strings.Join(strs, ",")
}

func deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	var rebuild func() *TreeNode
	rebuild = func() *TreeNode {
		if vals[0] == "X" {
			vals = vals[1:]
			return nil
		}
		v, _ := strconv.Atoi(vals[0])
		vals = vals[1:]
		root := &TreeNode{Val: v}
		root.Left = rebuild()
		root.Right = rebuild()
		return root
	}
	return rebuild()
}

// func Serialize(root *TreeNode) string {
// 	if root == nil {
// 		return "[]"
// 	}
// 	bs := strings.Builder{}
// 	// 层序遍历
// 	queue := make([]*TreeNode, 0)
// 	queue = append(queue, root)
// 	for len(queue) != 0 {
// 		node := queue[0]
// 		if node == nil {
// 			bs.WriteString("null,")
// 		} else {
// 			queue = append(queue, node.Left)
// 			queue = append(queue, node.Right)
// 			bs.WriteString(strconv.Itoa(node.Val) + ",")
// 		}
// 		queue = queue[1:]
// 	}
// 	s := []byte(bs.String())
// 	s = s[:len(s)-1]
// 	return "[" + string(s) + "]"
// }

// // 再一次层序遍历
// func Deserialize(s string) *TreeNode {
// 	if s == "[]" {
// 		return nil
// 	}
// 	bs := []byte(s)
// 	s = string(bs[1 : len(bs)-1])
// 	ss := strings.Split(s, ",")
// 	// [100,99,98,1,null,null,2,null,null,null,null]
// 	queue := make([]*TreeNode, 0)
// 	num, err := strconv.Atoi(ss[0])
// 	if err != nil {
// 		return nil
// 	}
// 	root := &TreeNode{Val: num}
// 	queue = append(queue, root)
// 	i := 1
// 	for len(queue) != 0 {
// 		node := queue[0]
// 		if ss[i] != "null" {
// 			num, err := strconv.Atoi(ss[i])
// 			if err != nil {
// 				return nil
// 			}
// 			node.Left = &TreeNode{Val: num}
// 			queue = append(queue, node.Left)
// 		}
// 		i++
// 		if ss[i] != "null" {
// 			num, err := strconv.Atoi(ss[i])
// 			if err != nil {
// 				return nil
// 			}
// 			node.Right = &TreeNode{Val: num}
// 			queue = append(queue, node.Right)
// 		}
// 		i++
// 		queue = queue[1:]
// 	}
// 	return root
// }

// // 用 map 构造存一下(后来判断发现其实不可，还是再一次层序)，或者再一次层序遍历也ok
// func Deserialize(s string) *TreeNode {
// 	if s == "[]" {
// 		return nil
// 	}
// 	bs := []byte(s)
// 	s = string(bs[1 : len(bs)-1])
// 	ss := strings.Split(s, ",")
// 	// [100,99,98,1,null,null,2,null,null,null,null]
// 	m := make(map[int]*TreeNode, len(ss))
// 	for i, s := range ss {
// 		if s == "null" {
// 			continue
// 		}
// 		num, err := strconv.Atoi(s)
// 		if err != nil {
// 			return nil
// 		}
// 		m[i] = &TreeNode{
// 			Val: num,
// 		}
// 	}
// 	for i := 0; i < len(ss); i++ {
// 		if m[i] != nil {
// 			m[i].Left = m[2*i+1]
// 			m[i].Right = m[2*i+2]
// 		}
// 	}
// 	return m[0]
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

// func main() {
// 	root := TreeNode{
// 		Val: 100,
// 		Left: &TreeNode{
// 			Val: 99,
// 			Left: &TreeNode{
// 				Val: 1,
// 			},
// 		},
// 		Right: &TreeNode{
// 			Val: 98,
// 			Right: &TreeNode{
// 				Val: 2,
// 			},
// 		},
// 	}
// 	fmt.Println(root)
// 	fmt.Println(Serialize(&root))
// 	fmt.Println(Serialize(Deserialize(Serialize(&root))))
// }
