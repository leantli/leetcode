package main

// https://leetcode.cn/problems/find-k-closest-elements/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 658. 找到 K 个最接近的元素

// 二刷，在双指针题再次遇到
// 首先都是排序好的，我们显然可以用二分找到最接近 x 的两个数
// 接着确定这两个数哪个是最接近的，l,r 都定位在其下标
// 此时二者都判断左右两边的数，哪个更接近，逐步扩散，每次只扩散一个(一边)
// 最终扩散 k 次, 时间复杂度为 logn+k
// 并且注意，绝对差值相等时，优先取较小的数( l 侧)
func findClosestElements(arr []int, k int, x int) []int {
	// 开区间二分
	l, r := -1, len(arr)
	for l+1 != r {
		mid := l + (r-l)/2
		if arr[mid] <= x {
			l = mid
		} else {
			r = mid
		}
	}
	// 最后 l 和 r 落在最接近 x 的两个数上
	// 此时找到最接近 x 的第一个数，k 一定大于 0，所以至少会有一个数
	if l == -1 {
		return arr[:k]
	} else if r == len(arr) {
		return arr[len(arr)-k:]
	}
	if abs(arr[l], x) <= abs(arr[r], x) {
		r = l
	} else {
		l = r
	}
	// 此时 l，r 都在同一个数上，l 考虑向左移动，r 考虑向右移动
	for r-l+1 < k {
		// 如果 l，r 指针都到边界了，就直接移动另一边的指针
		if l-1 < 0 {
			r++
			continue
		}
		if r+1 >= len(arr) {
			l--
			continue
		}
		if abs(arr[l-1], x) <= abs(arr[r+1], x) {
			l--
		} else {
			r++
		}
	}
	// 结束循环时，l，r 直接的数量就是 k 个最接近 x 的数
	return arr[l : r+1]
}
func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

// arr 升序
// 找到 k 个数，其满足在该数组中，与 x 的差值最小
// 基于二分，找到 < x 和 >= x 的值，分别标为 l 和 r
// 再去向左和向右扩散，并不断比较 l 和 r 指向的值 与 x 的差值
// 时间复杂度为 logn + k + klogk
// func findClosestElements(arr []int, k int, x int) []int {
// 	res := make([]int, 0, k)
// 	l, r := -1, len(arr)
// 	// logn
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if arr[mid] < x {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	// k
// 	// 此时 l 位于最后一个 < x 的位置，r 位于第一个 >= x 的位置
// 	for len(res) != k {
// 		if l >= 0 && r < len(arr) {
// 			if abs(arr[l], x) <= abs(arr[r], x) {
// 				res = append(res, arr[l])
// 				l--
// 			} else {
// 				res = append(res, arr[r])
// 				r++
// 			}
// 		} else if l >= 0 {
// 			res = append(res, arr[l])
// 			l--
// 		} else {
// 			res = append(res, arr[r])
// 			r++
// 		}
// 	}
// 	// klogk
// 	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
// 	return res
// }

// // 这里再考虑一下，其实没必要用 res 去返回？我们只需要确认了左右边界，直接从原数组截取即可，免去最后多余的排序，也节省了空间
// // 此时时间复杂度为 logn + k
// func findClosestElements(arr []int, k int, x int) []int {
// 	l, r := -1, len(arr)
// 	// logn
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if arr[mid] <= x {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	// k
// 	// 此时 l 位于最后一个 < x 的位置，r 位于第一个 >= x 的位置
// 	for ; k > 0; k-- {
// 		if l >= 0 && r < len(arr) {
// 			if abs(arr[l], x) <= abs(arr[r], x) {
// 				l--
// 			} else {
// 				r++
// 			}
// 		} else if l >= 0 {
// 			l--
// 		} else {
// 			r++
// 		}
// 	}
// 	return arr[l+1 : r]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func abs(a, b int) int {
// 	temp := a - b
// 	if temp < 0 {
// 		return -temp
// 	}
// 	return temp
// }
