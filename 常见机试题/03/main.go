package main

import "fmt"

/**
有一个长度为N的01数组(数组中的每个数字要么是0，要么是1)每个数字编号为1~N。定义了一个翻转操作:选取数组中一个连续的区间[L，R]，将其数字0变为数字1，而数字1变为数字0。对数组进行进行一次翻转操作后，数组中数字1的数量最多是多少。
输入 (第一行N表示数组的长度，第二行N个数字，表示数组中对应编号的具体数字:0或1)3
101
输出 (一次翻转操作后，数组中数字1的数量最多是多少)
3
**/

// 贪心解决
// 1. 先遍历查看当前 1 的数量
// 2. 一次遍历，统计翻转后多出的 1 的数量，遇到 0，就翻转后多出的 +1
// 如果翻转后多处的 1 的数量为负，直接置为 0，表示遇到下一个重新开始翻转
func maxOnesAfterFlip1(nums []int) int {
	maxCount := 0
	count := 0
	flipCount := 0

	// 统计原始数组中的1的数量
	for _, num := range nums {
		if num == 1 {
			count++
		}
	}

	// 遍历数组，计算每个位置作为翻转的起点时，翻转后1的数量
	for _, num := range nums {
		// 如果当前位置是0，则翻转后1的数量加1
		if num == 0 {
			flipCount++
		} else {
			// 如果当前位置是1，则翻转后1的数量减1
			flipCount--
		}

		// 更新最大的1的数量
		if flipCount > maxCount {
			maxCount = flipCount
		}

		// 翻转后1的数量不能小于0
		if flipCount < 0 {
			flipCount = 0
		}
	}

	// 返回翻转后的最大1的数量
	return count + maxCount
}

func main() {
	var N int
	fmt.Scan(&N)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	maxOnes1 := maxOnesAfterFlip1(nums)
	maxOnes3 := maxOnesAfterFlip3(nums)
	fmt.Println(maxOnes1)
	fmt.Println(maxOnes3)
}

// 暴力解决
func maxOnesAfterFlip3(nums []int) int {
	maxCount := 0

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			flip(nums, i, j)
			count := countOnes(nums)
			maxCount = max(maxCount, count)
			flip(nums, i, j) // 恢复原始数组
		}
	}
	return maxCount
}

func flip(nums []int, start, end int) {
	for i := start; i <= end; i++ {
		if nums[i] == 0 {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}
}

func countOnes(nums []int) int {
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
