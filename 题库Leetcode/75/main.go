package main

// https://leetcode.cn/problems/sort-colors/
// 75. 颜色分类

// 巧妙的做法，效率也不如一开始我想的，但是真的很巧妙
func sortColors(nums []int) {
	p0, p1 := -1, -1
	for p2 := 0; p2 < len(nums); p2++ {
		tmp := nums[p2]
		if tmp <= 2 {
			nums[p2] = 2
		}
		if tmp <= 1 {
			p1++
			nums[p1] = 1
		}
		if tmp == 0 {
			p0++
			nums[p0] = 0
		}
	}
}

// // 双指针做法，不过效率不如下面的
// // 保证这么一个性质，l 左边的元素都是 0， r 右边的元素都为 2
// // i 指针遍历原数组，遇到 0 和 l 交换，遇到 2 和 r 交换
// // 但是 i 与 r 交换时，不能直接就继续 i++，否则会跳过原来 r 指向的元素，这会导致漏查
// func sortColors(nums []int) {
// 	l, r := 0, len(nums)-1
// 	for i := 0; i <= r; i++ {
// 		if nums[i] == 0 {
// 			nums[l], nums[i] = nums[i], nums[l]
// 			l++
// 		} else if nums[i] == 2 {
// 			nums[r], nums[i] = nums[i], nums[r]
// 			r--
// 			i--
// 		}
// 	}
// }

// // 原地对数组排序，000111222
// // 常数空间的一趟扫描算法，看起来好像要用 sort，但其实根本不用
// // 用三个变量计数各个球出现的次数，然后按顺序填充数组即可
// func sortColors(nums []int) {
// 	var rc, wc, bc int
// 	for _, num := range nums {
// 		if num == 0 {
// 			rc++
// 		} else if num == 1 {
// 			wc++
// 		} else {
// 			bc++
// 		}
// 	}
// 	var index int
// 	for rc > 0 {
// 		nums[index] = 0
// 		index++
// 		rc--
// 	}
// 	for wc > 0 {
// 		nums[index] = 1
// 		index++
// 		wc--
// 	}
// 	for bc > 0 {
// 		nums[index] = 2
// 		index++
// 		bc--
// 	}
// }
