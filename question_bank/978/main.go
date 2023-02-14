package main

// https://leetcode.cn/problems/longest-turbulent-subarray/
// 978. 最长湍流子数组

// 更清晰更简洁的官方滑窗解法
// 直接比较三个值去移动 r
func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	ans := 1
	left, right := 0, 0
	for right < n-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 不定长滑动窗口 返回最大的长度
// // 窗口性质--窗口内是湍流数组
// // 则窗口内应满足 [0]>[1]<[2]>[3]<[4]>[5]....
// // 或 [0]<[1]>[2]<[3]>[4]<[5]....
// // 这里可以假设，如果窗口 r 右移时，新进的数不满足湍流规则
// // 则 l 直接右移至当前 r -1 的位置，然后判断当前 r-1 和 r 两个数的大小关系
// // 由此大小关系决定下一个新进数是否满足规则
// func maxTurbulenceSize(arr []int) int {
// 	n := len(arr)
// 	if n <= 1 {
// 		return n
// 	}
// 	var l int
// 	maxLen := 1
// 	var nextShouldLT bool // 下一个关系是否是否应为<，true为是，否为>
// 	for r := 1; r < n; r++ {
// 		// 保证窗口始终满足性质

// 		// 当窗口不满足湍流规则时，l 右移至 r-1 位置，并重置 nextShouldLT
// 		// 当二者相等时的特殊处理
// 		if arr[r-1] == arr[r] {
// 			for r < n && arr[r-1] == arr[r] {
// 				r++
// 			}
// 			if r == n {
// 				return maxLen
// 			}
// 			l = r - 1
// 			if arr[l] > arr[r] {
// 				nextShouldLT = true
// 			} else {
// 				nextShouldLT = false
// 			}
// 		} else if nextShouldLT && arr[r-1] > arr[r] { // 下一个应该为 小于 关系，但此时 r-1 与 r 为大于等于关系
// 			l = r - 1
// 			nextShouldLT = true
// 		} else if !nextShouldLT && arr[r-1] < arr[r] { // 下一个应为大于关系，此时 r-1 与 r 为小于等于关系
// 			l = r - 1
// 			nextShouldLT = false
// 		} else { // 其余都是满足湍流规则的，此时需要正常变更 nextShouldLT
// 			nextShouldLT = !nextShouldLT
// 		}
// 		maxLen = max(maxLen, r-l+1)
// 	}
// 	return maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
