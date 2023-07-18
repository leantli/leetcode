package main

import (
	"fmt"
	"sort"
)

/**
题目描述:
在一条笔直的公路上安装了N个路灯，从位置0开始安装，路灯之间间距固定为100米。每个路灯都有自己的照明半径，请计算第一个路灯和最后一个路灯之间，无法照明的区间的长度和
输入描述:
第一行为一个数N，表示路灯个数，1<=N<=100000第二行为N个空格分隔的数，表示路径的照明半径，1<=照明半径<=100000*100
输出描述:
第一个路灯和最后一个路灯之间，无法照明的区间的长度和
**/

// 读取每个路灯的照明半径，分别设置其左右照明区间
// 对区间左边界排序，然后贪心？--类似于合并区间
func main() {
	// 区间输入和初始化
	var n int
	fmt.Scan(&n)
	interval := make([][]int, n)
	var temp int
	fmt.Scan(&temp)
	interval[0] = []int{0, temp}
	for i := 1; i < n; i++ {
		var half int
		fmt.Scan(&half)
		mid := i * 100
		interval[i] = []int{mid - half, mid + half}
	}
	// 逻辑处理
	var res int // 未覆盖到的长度
	// 左区间小的排在前面
	sort.Slice(interval, func(i, j int) bool { return interval[i][0] < interval[j][0] })
	fmt.Println(interval)
	last := interval[0][1]
	for i := 1; i < n; i++ {
		if interval[i][0] > last {
			res += interval[i][0] - last
		}
		last = interval[i][1]
	}
	fmt.Println(res)
}
