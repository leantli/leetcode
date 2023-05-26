package main

// https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 66. 构建乘积数组

// 此时时间复杂度为 O(2n), 空间复杂度为O(2n)
// 空间复杂度还能优化到 O(1)
func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return a
	}
	res := make([]int, n)
	res[0] = 1
	// 直接将左乘积放在 res 中
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * a[i-1]
	}
	// 右乘积标识
	temp := 1
	for i := n - 2; i >= 0; i-- {
		temp *= a[i+1]
		// 此时 res[i] 表示的是 a[i] 的左乘积
		// temp 为 a[i] 的右乘积
		res[i] *= temp
	}
	return res
}

// 这样的一个数组，第一个感觉的方法就是先求出总乘积再对单个值做除法
// 但显然题目要求了不能使用除法
// 那不使用该方法还能咋搞？我们先想想最普通的做法
// 遍历每个数，根据其下标，遍历整个数组相乘，只跳过当前下标
// 此时我们需要发现，每个数的乘积，都是其左右数组的乘积相乘
// 那么我们可以采用数组，去存放两边数组的乘积，减少重复计算，空间换时间

// func constructArr(a []int) []int {
// 	n := len(a)
//     if n == 0 {
//         return a
//     }
// 	left, right := make([]int, n), make([]int, n)
// 	left[0], right[n-1] = 1, 1
// 	for i := 1; i < n; i++ {
// 		left[i] = left[i-1] * a[i-1]
// 		right[n-i-1] = right[n-i] * a[n-i]
// 	}
// 	res := make([]int, n)
// 	for i := range a {
// 		res[i] = left[i] * right[i]
// 	}
// 	return res
// }
