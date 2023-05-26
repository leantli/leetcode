package main

import "fmt"

// https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 61. 扑克牌中的顺子

// 简单题，就直接用 API 了吧....
// 算了，当巩固一些快排写法好了
// func isStraight(nums []int) bool {
// 	quickSort(nums, 0, len(nums)-1)
// 	// 用来计数 0 的个数，可充当任意数的次数
// 	var flag int
// 	for i := 0; i < 4; i++ {
// 		// 一些特殊处理
// 		// 0 肯定直接跳过并且增加任意数的次数
// 		if nums[i] == 0 {
// 			flag++
// 			continue
// 		}
// 		// 有两个相等的数就肯定不是顺子了，除了 0
// 		if nums[i+1] == nums[i] {
// 			return false
// 		}
// 		// 正常情况
// 		if nums[i+1]-1 == nums[i] {
// 			continue
// 		}
// 		// 任意数补不满
// 		flag -= nums[i+1] - nums[i] - 1
// 		if flag < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// 草，看了大佬的，感觉自己的想法还是弱了
func isStraight(nums []int) bool {
	var joker int
	quickSort(nums, 0, len(nums)-1)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	// 此时 nums[joker] 指向除 0 之外最小的数，只要满足 max - min > 5，此时就是顺子！
	return (nums[4] - nums[joker]) > 5
}

// 交换法快排
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j, pivot := l, r, nums[l]
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[l] = nums[l], nums[i]
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
}

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 6}))
}
