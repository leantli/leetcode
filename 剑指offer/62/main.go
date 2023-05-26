package main

// https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 62. 圆圈中最后剩下的数字

// 印象中第一次做这题的时候是大一用 C 写链表的时候哈哈
// 不过当下这道题的 n 和 m 取值范围较大，显然不能常规模拟
// 得考虑一下其数学规律
// 手写一下观察观察规律。。。一番手写后，没发现规律，dbq
// 那怎么才能找到最后一个剩下的数？
// 其实我们只要找到最后一个数，在初始数组中的下标即可
// 逆着去补位
// 我们知道最后的剩下的数，此时在数组中必定是 0 下标
// 逆折往前推 idx = (idx+m)%i (i 为这一轮数组的长度)
// 如此到最后即可得到 最后一个数 在初始数组中的下标
// 而在初始数组中，下标==数
func lastRemaining(n int, m int) int {
	// 逆着推，最后一个数，最后时在数组中下标为 0
	idx := 0
	// i 代表当前数组的长度, 从 2 开始，逐步推到初始长度 n
	for i := 2; i <= n; i++ {
		idx = (idx + m) % i
	}
	return idx
}
