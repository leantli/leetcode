package main

// https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together/
// 1151. 最少交换次数来组合所有的 1

// // 但其实我们可以转换问题，转换完就会发现这道题和之前很多滑窗题差不多
// // 我们肯定是要先遍历一遍确定有多少个 1 的，这里设为 count1
// // 此时，我们可以确定一个定长为 count1 长度的滑动窗口
// // 这个滑动窗口中，我们只需要统计窗口中非 1 的个数，非 1 的个数就是需要交换的次数
// // 接着不断右移这个窗口，每次移动都统计一次非 1 的数量，取最小非 1 数量
// func minSwaps(data []int) int {
// 	var count1 int
// 	for _, num := range data {
// 		if num == 1 {
// 			count1++
// 		}
// 	}
// 	var zeroCnt int
// 	for _, num := range data[:count1] {
// 		if num == 0 {
// 			zeroCnt++
// 		}
// 	}
// 	minZeroCnt := zeroCnt
// 	for i := count1; i < len(data); i++ {
// 		r, l := data[i], data[i-count1]
// 		if l == r {
// 			continue
// 		}
// 		if r == 0 {
// 			zeroCnt++
// 		}
// 		if l == 0 {
// 			zeroCnt--
// 		}
// 		if zeroCnt < minZeroCnt {
// 			minZeroCnt = zeroCnt
// 		}
// 	}
// 	return minZeroCnt
// }

// // 通过交换位置，令所有 1 都相邻，并且返回所有可能性中交换操作最少的次数
// // 暴力模拟 - 每遇到一个 1，就基于这个1去看其后面的count1-1个位置都为1需要多少次交换
// // 这个暴力思路有问题哈哈，显然是不行
// func minSwaps(data []int) int {
// 	var count1 int
// 	for _, num := range data {
// 		if num == 1 {
// 			count1++
// 		}
// 	}
// 	minRes := count1
// 	for i := 0; i <= len(data)-count1; i++ {
// 		if data[i] == 1 {
// 			var cnt int
// 			for j := i; j < count1; j++ {
// 				if data[j] == 0 {
// 					cnt++
// 				}
// 			}
// 			if cnt < minRes {
// 				minRes = cnt
// 			}
// 		}
// 	}
// 	return minRes
// }
