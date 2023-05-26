package main

// https://leetcode.cn/problems/contains-duplicate-ii/
// 219. 存在重复元素 II

// 再看到这里其实 i j 之间的差值小于等于 k ，其实相当于一个定长滑动窗口？
// 窗口只要一直没有重复数，最终就能返回 false
// 窗口中有重复数就返回 true
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 形成窗口后，先舍弃最左边的再正常添加窗口右侧的数
		if i > k {
			m[nums[i-k-1]]--
		}
		// 添加右侧数
		m[nums[i]]++
		// 判断是否重复
		if m[nums[i]] > 1 {
			return true
		}
	}
	return false
}

// // 记录 num 上次出现的位置，时间复杂度为 O(n)
// // 返回是否存在 i 和 j，使得 nums[i] == nums[j] 并且 abs(i-j) <= k
// // 这里显然不能排序，否则打乱 i j
// // i,j只有正整数，因此我们需要一个 map 记录每个数字上次出现的位置
// // 只需记录其上次出现的位置即可，更前的相等数，下标之间的绝对差值一定更大(下标单调递增)
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	m := make(map[int]int)
// 	for i, num := range nums {
// 		if j, ok := m[num]; ok {
// 			if i-j <= k {
// 				return true
// 			}
// 		}
// 		m[num] = i
// 	}
// 	return false
// }
