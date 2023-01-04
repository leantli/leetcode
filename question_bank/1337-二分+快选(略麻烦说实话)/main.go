package main

import "sort"

// https://leetcode.cn/problems/the-k-weakest-rows-in-a-matrix/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 1337. 矩阵中战斗力最弱的 K 行

// 最普通的模拟就是逐行遍历，找到最后一个军人出现的下标，然后再比较排序
// 这里可以优化的点在于逐行遍历寻找，太过麻烦，可以二分寻找最后一个军人出现的下标
// 比较排序可以采用快速选择，只排出最弱的 k 行，而不需要全部排序
// 原先的时间复杂度是 n^2 + nlogn, 优化后是 nlogn+n(快速选择的算法时间复杂度约为O(n))
// 这里的二分采用万用模板，因为要求分界
// 但是这里还有一个问题，就是最后排序时，同战斗力时，还需要考虑行号，但是我们采用的排序算法一般都是不稳定算法
// 所以还要考虑如何保证同战斗力下，行号低的在前面，需要对排序方式进行一定的变更
// 我们本来只是基于战斗力排序，此时可以考虑这样 战斗力*100+行号 作为排序号
// 因为行数最多为 100，因此这里乘上 100，可以避免对大行数对 战斗力 的影响
// 可以保证同战斗力下，行号低的在前面，又不会因为行号过大导致战斗力出错
// 懒得写快速选择版本
// func kWeakestRows(mat [][]int, k int) []int {
// 	n := len(mat[0])
// 	for i := 0; i < len(mat); i++ {
// 		l, r := -1, n
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if mat[i][mid] == 1 {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// 结束时 l 位于最后一个 1，r 位于最开始的 0
// 		// 将该行原本的 下标 和 最后一个军人下标 加在该行最后，便于排序和操作
// 		mat[i] = append(mat[i], i)
// 		mat[i] = append(mat[i], l*100+i) // 最后一个数本来是最后一个军人下标(战斗力)，但是我们需要 战斗力*100+行号，保证同战斗力下，行号低的在前面
// 	}
// 	// 此时每行的 n 下标为该行排序前的下标位置， n+1 下标为该行最后一个军人下标
// 	// 我们根据最后一个军人的下标进行排序
// 	sort.Slice(mat, func(i, j int) bool {
// 		return mat[i][n+1] < mat[j][n+1]
// 	})
// 	ans := make([]int, 0, k)
// 	for i := 0; i < k; i++ {
// 		ans = append(ans, mat[i][n])
// 	}
// 	return ans
// }

// 这里再考虑一下写个快速选择？
// 需要先写个快排，并且该快排会返回本次排序确定的下标，先写一维的，后面再改成二维的，循序渐进
// func quickSort(nums []int, l, r int) int {
// 	left, right, pivot := l, r, nums[l]
// 	for left < right {
// 		for left < right && nums[right] > pivot {
// 			right--
// 		}
// 		for left < right && nums[left] < pivot {
// 			left++
// 		}
// 		nums[left], nums[right] = nums[right], nums[left]
// 	}
// 	nums[left], nums[l] = nums[l], nums[left]
// 	return left
// }

// // 快速选择，根据 当前排到的下标位置，去选择下一次快排的位置
// func quickSearch(nums []int, l, r, k int) []int {
// 	index := quickSort(nums, l, r)
// 	if index == k {
// 		return nums[:k]
// 	}
// 	if index < k {
// 		return quickSearch(nums, index+1, r, k)
// 	}
// 	return quickSearch(nums, l, index-1, k)
// }

// 最普通的模拟就是逐行遍历，找到最后一个军人出现的下标，然后再比较排序
// 这里可以优化的点在于逐行遍历寻找，太过麻烦，可以二分寻找最后一个军人出现的下标
// 比较排序可以采用快速选择，只排出最弱的 k 行，而不需要全部排序
// 原先的时间复杂度是 n^2 + nlogn, 优化后是 nlogn+n(快速选择的算法时间复杂度约为O(n))
// 这里的二分采用万用模板，因为要求分界
// 但是这里还有一个问题，就是最后排序时，同战斗力时，还需要考虑行号，但是我们采用的排序算法一般都是不稳定算法
// 所以还要考虑如何保证同战斗力下，行号低的在前面，需要对排序方式进行一定的变更
// 我们本来只是基于战斗力排序，此时可以考虑这样 战斗力*100+行号 作为排序号
// 因为行数最多为 100，因此这里乘上 100，可以避免对大行数对 战斗力 的影响
// 可以保证同战斗力下，行号低的在前面，又不会因为行号过大导致战斗力出错
// 好了，上面写出一维数组版本的快速选择，这里根据题目改写一下，根据每行的最后一个战斗力值进行排序
func quickSort(nums [][]int, l, r int) int {
	n := len(nums[0])
	left, right, pivot := l, r, nums[l][n-1]
	for left < right {
		for left < right && nums[right][n-1] > pivot {
			right--
		}
		for left < right && nums[left][n-1] <= pivot {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[l] = nums[l], nums[left]
	return left
}

// 快速选择，根据 当前排到的下标位置，去选择下一次快排的位置
func quickSearch(nums [][]int, l, r, k int) [][]int {
	index := quickSort(nums, l, r)
	// 事实上，这里 快选 一般是 index == k，实际上 index+1 == k 即可
	// 但是在做最小k和最大k时，存在k==0的情况，因此最好是index=k
	// 而此处，k 必定大于等于 0，且 k == len(nums) 时会越界，因此这里最好是 index+1 == k
	if index+1 == k {
		return nums[:k]
	}
	if index < k {
		return quickSearch(nums, index+1, r, k)
	}
	return quickSearch(nums, l, index-1, k)
}

// 快速选择 + 二分版本
func kWeakestRows(mat [][]int, k int) []int {
	n := len(mat[0])
	for i := 0; i < len(mat); i++ {
		l, r := -1, n
		for l+1 != r {
			mid := l + (r-l)/2
			if mat[i][mid] == 1 {
				l = mid
			} else {
				r = mid
			}
		}
		// 结束时 l 位于最后一个 1，r 位于最开始的 0
		// 将该行原本的 下标 和 最后一个军人下标 加在该行最后，便于排序和操作
		mat[i] = append(mat[i], i)
		mat[i] = append(mat[i], l*100+i) // 最后一个数本来是最后一个军人下标(战斗力)，但是我们需要 战斗力*100+行号，保证同战斗力下，行号低的在前面
	}
	// 此时每行的 n 下标为该行排序前的下标位置， n+1 下标为该行最后一个军人下标
	// 我们根据最后一个军人的下标先进行快速选择，得到最小的 k 组，再做升序排序
	res := quickSearch(mat, 0, len(mat)-1, k)
	sort.Slice(res, func(i, j int) bool {
		return res[i][n+1] < res[j][n+1]
	})
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, res[i][n])
	}
	return ans
}
