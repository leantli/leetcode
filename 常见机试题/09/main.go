package main

import (
	"fmt"
)

/**
题目描述:
给定一个以顺序储存结构存储整数值的完全二叉树序列(最多1000个整数》，请找出此完全二叉树的所有非叶子节点部分，然后采用后序遍历方式将此部分树 (不包含叶子) 输出。
1、只有一个节点的树，此节点认定为根节点 (非叶子)
2、此完全二叉树并非满二叉树，可能存在倒数第二层出现叶子或者无右叶子的情况
**/

func main() {
	temp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	afterOrder(temp, 0, len(temp)/2-1)
}

func afterOrder(nums []int, i, k int) {
	if i > k {
		return
	}
	afterOrder(nums, i*2+1, k)
	afterOrder(nums, i*2+2, k)
	fmt.Printf("%d ", nums[i])
}
