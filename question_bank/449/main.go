package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/serialize-and-deserialize-bst/
// 449. 序列化和反序列化二叉搜索树

// 虽然最后做完 297，其实会发现，297的通用解法，时间空间复杂度都是O(n)，即便不利用
// 二叉搜索树的性质也是完全可以的

// 刚才的写法我们显然没很好地利用二叉搜索树的性质，仅仅是序列化时只取前序遍历
// 虽然减少了一次中序遍历的时间，但在反序列化时还没能很好利用二叉搜索树的性质
// 这里序列化部分和之前一样，只取前序遍历
type Codec struct{}

func Constructor() Codec { return Codec{} }

// 这里我们对二叉搜索树进行一个前序遍历，得到一个数组，并将其转换为字符串
func (this *Codec) serialize(root *TreeNode) string {
	var preorderTravel func(root *TreeNode)
	preorderPath := make([]string, 0)
	preorderTravel = func(root *TreeNode) {
		if root == nil {
			return
		}
		preorderPath = append(preorderPath, strconv.Itoa(root.Val))
		preorderTravel(root.Left)
		preorderTravel(root.Right)
	}
	preorderTravel(root)
	return strings.Join(preorderPath, ",")
}

// 这次我们努力利用二叉搜索树的性质，前序遍历数组的头部一定是根结点
// 第一种容易想到的做法：此时我们能够根据根结点的值，将其余的值分成左子树和右子树(基于二分找到分界下标)
// 并继续递归，将划分好的区间传入 rebuild 函数(传入参数为一个数组)
// 第二种比较难想到的处理：我们知道，前序遍历顺序是[中 左 右]
// 其中左子树区间就是根结点左子树的前序遍历情况，且我们能知道，左子节点的值，一定是有范围的<-->(-1,root.Val)<-->(-1来自题目划定的范围[0,10000])
// 如果不在这个范围内，说明不存在左子节点甚至可以说该根结点没有左子树
// 因此我们可以划定范围，判断当前根结点的值是否在范围内，是的话就满足节点条件，不满足就说明该节点应该在另一个位置
// 而这个位置的递归，我们先左后右，按照前序遍历的顺序进行重建
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || data == "" {
		return nil
	}
	prePath := strings.Split(data, ",")
	var rebuild func(low, high int) *TreeNode
	rebuild = func(low, high int) *TreeNode {
		if len(prePath) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(prePath[0])
		// 当 val 不在范围内，说明 val 属于其他位置，而不属于当前子树
		if val < low || val > high {
			return nil
		}
		prePath = prePath[1:]
		root := &TreeNode{Val: val}
		// 这里要先左后右，按照前序遍历的顺序
		root.Left = rebuild(low, val)
		root.Right = rebuild(val, high)
		return root
	}
	return rebuild(-1, 10001)
}

// // 常规来说，我们都是直接用 json, xml 等形式对数据做序列化和反序列化
// // 但是显然，这种形式在一定程度上额外存储了树的结构字段，有一定的浪费
// // 事实上，我们往往只需要根据(前序/后序)+中序遍历即可重构一棵儿二叉树
// // 而二叉搜索树，在一定程度上还能额外减少序列化后二叉树结构的数据大小
// // 因为二叉搜索树只需要序列化其 前序/后序 遍历数组即可，无需中序
// // 反序列化时，对前后序排序后即可得到中序遍历的数组
// type Codec struct{}

// func Constructor() Codec { return Codec{} }

// // 这里我们对二叉搜索树进行一个前序遍历，得到一个数组，并将其转换为字符串
// func (this *Codec) serialize(root *TreeNode) string {
// 	var preorderTravel func(root *TreeNode)
// 	preorderPath := make([]string, 0)
// 	preorderTravel = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		preorderPath = append(preorderPath, strconv.Itoa(root.Val))
// 		preorderTravel(root.Left)
// 		preorderTravel(root.Right)
// 	}
// 	preorderTravel(root)
// 	return strings.Join(preorderPath, ",")
// }
// // 常规写法：解析完前序遍历数组后，先排序取得中序遍历数组，再根据前中序数组重建二叉树
// // 但显然，还不够好，没很好地利用上二叉搜索树的性质
// func (this *Codec) deserialize(data string) *TreeNode {
// 	if len(data) == 0 || data == "" {
// 		return nil
// 	}
// 	// 解析字符串，获取前序遍历数组
// 	preStringPath := strings.Split(data, ",")
// 	prePath := make([]int, len(preStringPath))
// 	for i, v := range preStringPath {
// 		num, err := strconv.Atoi(v)
// 		if err != nil {
// 			log.Panicf("v:%s -- Atoi failed\n", v)
// 			return nil
// 		}
// 		prePath[i] = num
// 	}
// 	// 获取中序遍历数组
// 	inPath := make([]int, len(prePath))
// 	copy(inPath, prePath)
// 	sort.Ints(inPath)
// 	return this.rebuildTree(prePath, inPath)
// }

// // 根据前序遍历和中序遍历重构二叉树
// // 首先前序遍历的首个节点必定是根结点
// func (this *Codec) rebuildTree(pre, in []int) *TreeNode {
// 	if len(pre) == 0 {
// 		return nil
// 	}
// 	root := &TreeNode{Val: pre[0]}
// 	// 找到 root 在中序遍历中的下标，并且下标==root左子树节点的数量
// 	var rootInIndex int
// 	for i, v := range in {
// 		if v == root.Val {
// 			rootInIndex = i
// 		}
// 	}
// 	root.Left = this.rebuildTree(pre[1:rootInIndex+1], in[:rootInIndex])
// 	root.Right = this.rebuildTree(pre[rootInIndex+1:], in[rootInIndex+1:])
// 	return root
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
