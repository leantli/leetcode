package main

// https://leetcode.cn/problems/top-k-frequent-elements/
// 347. 前 K 个高频元素

type pair struct {
	num, frequency int
}

func topKFrequent(nums []int, k int) []int {
	// 使用 map 统计每个数出现的频率
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	// 基于统计的频率生成 pair 数组
	pairs := make([]pair, 0, len(m))
	for num, fre := range m {
		pairs = append(pairs, pair{num, fre})
	}
	// fmt.Println(pairs)
	pairs = quickSelect(0, len(pairs)-1, k, pairs)
	// fmt.Println(pairs)
	res := make([]int, k)
	for i, pair := range pairs[:k] {
		res[i] = pair.num
	}
	return res
}

// 返回排出频率前 k 高的数组
func quickSelect(l, r, k int, nums []pair) []pair {
	mid := quickSort(l, r, nums)
	if mid+1 == k {
		return nums
	}
	if mid+1 > k {
		return quickSelect(l, mid-1, k, nums)
	}
	return quickSelect(mid+1, r, k, nums)
}

// 传入要排序的数组，传入排序的左右边界起始点，返回基准点排序后的下标位置
// 升序排序，右边是小于 pivot，左边大于等于 pivot
func quickSort(l, r int, nums []pair) int {
	pivot, i, j := nums[l].frequency, l, r
	for i < j {
		for j > i && nums[j].frequency < pivot {
			j--
		}
		for j > i && nums[i].frequency >= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[i] = nums[i], nums[l]
	// 返回最终基准点的位置
	return i
}
