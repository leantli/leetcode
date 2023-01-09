package main

// https://leetcode.cn/problems/find-k-closest-elements/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 658. 找到 K 个最接近的元素

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

// 这里再考虑一下，其实没必要用 res 去返回？我们只需要确认了左右边界，直接从原数组截取即可，免去最后多余的排序，也节省了空间
// 此时时间复杂度为 logn + k
func findClosestElements(arr []int, k int, x int) []int {
	l, r := -1, len(arr)
	// logn
	for l+1 != r {
		mid := l + (r-l)/2
		if arr[mid] <= x {
			l = mid
		} else {
			r = mid
		}
	}
	// k
	// 此时 l 位于最后一个 < x 的位置，r 位于第一个 >= x 的位置
	for ; k > 0; k-- {
		if l >= 0 && r < len(arr) {
			if abs(arr[l], x) <= abs(arr[r], x) {
				l--
			} else {
				r++
			}
		} else if l >= 0 {
			l--
		} else {
			r++
		}
	}
	return arr[l+1 : r]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	temp := a - b
	if temp < 0 {
		return -temp
	}
	return temp
}
