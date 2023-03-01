package main

// https://leetcode.cn/problems/longest-mountain-in-array/
// 845. 数组中的最长山脉

// // 不定长滑窗思路
// // 窗口性质--窗口内满足 山脉数组的性质, 先升后降
// func longestMountain(arr []int) int {
// 	var l, r int
// 	var ins, des bool // 表示是否上升过，是否下降过
// 	var res int
// 	for r < len(arr)-1 {
// 		// 这里维护窗口性质, 这里分类有点麻烦，只能慢慢细细考虑
// 		// 1. 还没降过，此时要升 r++; ins=true
// 		// 2. 升过准备降 r++; des=true
// 		// 3. 降着要升 l = r r++; des=false
// 		// 4. 其他情况对结果无任何作用，直接将 l,r 都置于新的位置 l = r+1; r++; ins=des=false
// 		if !des && arr[r] < arr[r+1] {
// 			ins = true
// 		} else if ins && arr[r] > arr[r+1] {
// 			des = true
// 		} else if des && arr[r] < arr[r+1] {
// 			l = r
// 			des = false
// 		} else {
// 			l = r + 1
// 			ins = false
// 			des = false
// 		}
// 		r++
// 		// 已经下降过，才能开始计算山脉长度
// 		if des && res < r-l+1 {
// 			res = r - l + 1
// 		}
// 	}
// 	return res
// }

// 看到一个更简洁的思路，其实看起来和原先的思路差不多
// 也是寻找从一个山脚到另一边的山脚
// 假设有山脉，那么一定会一直升到顶，再降到底
// 此时计算升了多少，降了多少，有升有降才是山脉，再计算山脉长度
// 最后再看看是否遇到重复值，重复值一定不会成山脉，直接略过
func longestMountain(arr []int) int {
	var res int
	r := 1
	for r < len(arr) {
		var ins, des int
		// 假设遇到山脉，一直升到山脉顶，再一次降到山脉底
		// 并记录升降的长度
		for r < len(arr) && arr[r-1] < arr[r] {
			r++
			ins++
		}
		for r < len(arr) && arr[r-1] > arr[r] {
			r++
			des++
		}
		// 有升有降才是山脉，再计算山脉长度
		if ins > 0 && des > 0 && res < ins+des+1 {
			res = ins + des + 1
		}
		for r < len(arr) && arr[r-1] == arr[r] {
			r++
		}
	}
	return res
}
