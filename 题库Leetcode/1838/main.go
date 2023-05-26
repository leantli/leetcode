package main

import (
	"sort"
)

// https://leetcode.cn/problems/frequency-of-the-most-frequent-element/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1838. 最高频元素的频数

// 排序 + 二分 + 前缀和
// 肯定要排序吧，排序之后以便于选取最接近的几个数，去得到最小的趋同时需要的操作数
// 排序时间复杂度 nlogn, 对每个数都将其当做连续区间中最大的值，并去计算它的连续区间的最大长度
// 这个可趋同区间，就由前缀和求 趋同 所需要的 操作数
// 将每个数当做区间中的最大值时，我们只需要从左侧去枚举形成区间，然后去看操作数是不是会超，但枚举显然太麻烦，排序后的区间长度显然和操作数是有单调关系的
// 因此我们可以用二分去选取每个值能与左侧形成的最长趋同长度区间(即最长的区间，并且该区间趋同所需的操作数小于等于 k)
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	sum := make([]int, len(nums)+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	for i := 1; i < len(nums); i++ {
		// 这里注意一下，用的开区间模板，l 和 r 为开区间左右值
		l, r := -1, i+1
		for l+1 != r {
			mid := l + (r-l)/2
			need := nums[i]*(i-mid+1) - (sum[i+1] - sum[mid])
			if need > k {
				l = mid
			} else {
				r = mid
			}
		}
		ans = max(ans, i-r+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// // 第二次做的思路
// // 如果我们要让几个数经过 k 次递增操作后相等，显然得找最相近的几个数，但是数组本身未排序
// // 这里可能要排序一下 nlogn
// // 然后是找最相近的几个数，那么是几个数？这个地方感觉得用滑动窗口，不然取几个数这个问题，没想到其他好解决的思路，dp 感觉也不是很适合？
// // 接下来我们维护一个滑动窗口，只要判断窗口内的数都趋同，所需要的 递增次数是否超过 k，超过则缩左，不超过则扩右
// // 接下来就是如何计算窗口内操作数了
// // 每次右扩时，就是让原窗口内的数，都向新进数趋同，此时计算为 原先窗口总共需要操作的次数 + (新进数 - 原窗口趋同的数(原窗口最后一个数)) * 原窗口长度
// // 每次左缩时，就是少了窗口内最左侧数趋同需要的操作数，此时计算为 原窗口总共需要操作的次数 - (原窗口趋同的数(原窗口最后一个数) - 原窗口最左侧数)
// // 总体时间复杂度还是 nlogn
// func maxFrequency(nums []int, k int) int {
// 	// 先排序，这样方便找最邻近的几个数
// 	sort.Ints(nums)
// 	// 下面基于滑窗去确定要找哪些邻近的数，以及窗口最长长度
// 	l, r := 0, 1
// 	ans := 1
// 	var need int
// 	for r < len(nums) {
// 		need += (nums[r] - nums[r-1]) * (r - l)
// 		for need > k {
// 			need -= nums[r] - nums[l]
// 			l++
// 		}
// 		r++
// 		ans = max(ans, r-l)
// 	}
// 	return ans
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 还得理解一下 二分+前缀和 思路
// func maxFrequency(a []int, k int) (ans int) {
// 	sort.Ints(a)
// 	sum := make([]int, len(a)+1)
// 	for i, v := range a {
// 		sum[i+1] = sum[i] + v
// 	}
// 	for r, v := range a {
// 		l := sort.Search(r, func(l int) bool { return (r-l)*v-sum[r]+sum[l] <= k })
// 		ans = max(ans, r-l+1)
// 	}
// 	return
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 先排序，排序后再按区间去滑动窗口
// // 以区间最右侧数字为目标，去对区间内左侧的值进行递增操作
// // 判断区间内同一值的操作数是否小于等于 k，是则滑动窗口右扩，不是则缩小
// func maxFrequency(nums []int, k int) int {
// 	sort.Slice(nums, func(a, b int) bool { return nums[a] < nums[b] })
// 	res := 1
// 	l, r := 0, 1
// 	var cost int
// 	for r < len(nums) {
// 		// 这里其实要特别关注，对于操作数的计算，如果是每次都对窗口内的所有的数重新计算，显然时间复杂度会更高
// 		// 最好采用以下计算方式，如果直接采取 新进的值*长度-此前操作综总和 的计算方式，则可能超过半数 maxInt
// 		cost += (nums[r] - nums[r-1]) * (r - l)
// 		for cost > k {
// 			cost -= nums[r] - nums[l]
// 			l++
// 		}
// 		r++
// 		res = max(res, r-l)
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
