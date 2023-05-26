package main

// https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1574. 删除最短的子数组使剩余数组有序

// 找到每一个数的右侧 比它大的最小值，获取其下标
// 如果没有意味着要删除其右侧全部
// 其实好像不行，如果找到了可删除的最短的，但是可能得删一大块，唉不知道怎么描述
// 比如 [1,2,3,10,4,5,4,3,2,1] ，按我上面的描述可能删去10是最短的，但是其后面还有[4,3,2,1]必须删除
// 因此没办法说就是只找到最短的一个，因为我们还要保证全局单调
// 此方法 pass

// // 由于上个思路我们知道，我们必须保证删除某一段后，全局仍然单调
// // 其实我们可以将整个数组分成三个部分，单调不递减部分 | 无规则或必须删除部分 | 单调不递减部分
// // 寻找左右两个各一个点，删去两点区间中的数，仍能保证整个数组单调不递减
// // 时间复杂度为 nlogn
// func findLengthOfShortestSubarray(arr []int) int {
// 	n := len(arr)
// 	// 打算分成三个区间，分别是 [0,i), [i,j) [j,n)
// 	i, j := 1, n-1
// 	// 先把左侧的非递减区间给确定好 [0, i)
// 	for i < n && arr[i] >= arr[i-1] {
// 		i++
// 	}
// 	// 如果 i 一直到底，则说明数组已经是有序的了，直接返回即可
// 	if i == n {
// 		return 0
// 	}
// 	// 再确定右侧的非递减区间 [j, n)
// 	for j >= i && arr[j] >= arr[j-1] {
// 		j--
// 	}
// 	// 下面就是寻找最短的删除长度，根据左侧区间中的每个数，找到右侧区间比它大的最小值，最终计算出要删除的长度
// 	// 因此我们遍历左侧区间的每个数，再基于二分找右侧区间的数，右侧区间的取值范围为 [j,n-1]，下面开区间二分
// 	// 最差情况是左侧区间全都得删掉，即左侧区间没有数比右侧区间的数小，此时删除长度刚好是 j
// 	ans := j
// 	for k := 0; k < i; k++ {
// 		l, r := j-1, n
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if arr[mid] < arr[k] {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// l 落在最后一个 小于 arr[k] 的下标，r 落在第一个大于等于 arr[k] 的下标
// 		ans = min(ans, r-k-1)
// 	}
// 	return ans
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 由于上个方法我们知道，我们必须保证删除某一段后，全局仍然单调
// 其实我们可以将整个数组分成三个部分，单调不递减部分 | 无规则或必须删除部分 | 单调不递减部分
// 寻找左右两个各一个点，删去两点区间中的数，仍能保证整个数组单调不递减
// 时间复杂度为 nlogn
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	// 打算分成三个区间，分别是 [0,i), [i,j) [j,n)
	i, j := 1, n-1
	// 先把左侧的非递减区间给确定好 [0, i)
	for i < n && arr[i] >= arr[i-1] {
		i++
	}
	// 如果 i 一直到底，则说明数组已经是有序的了，直接返回即可
	if i == n {
		return 0
	}
	// 再确定右侧的非递减区间 [j, n)
	for j >= i && arr[j] >= arr[j-1] {
		j--
	}
	// 其实我们可以注意到，当我们区分好三个区间后，左右都是非递减区间，中间是必须删除的区域
	// 此时可以基于 i，j 做一个滑动窗口，窗口初始化时即为中间区域，接下来我们的最终目的是保证两侧区间 [0,i) [j,n) 整体是非递减
	l, r := 0, j
	ans := j
	for l < i && r < n {
		if arr[l] <= arr[r] {
			ans = min(ans, r-l-1)
			l++
		} else {
			r++
		}
	}
	return min(ans, n-i)
}
