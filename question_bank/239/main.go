package main

// https://leetcode.cn/problems/sliding-window-maximum/
// 239. 滑动窗口最大值

// 定长滑动窗口，窗口倒是正常滑动
// 但是我们需要记录窗口中的最大值，窗口右移可能会导致最大值被弹出，此时我们需要第二大的值
// 这里我们可以想到，弹出最大值，于是此时数组的第二大值变成了最大值，显然应该是个单调队列
// 这里我们可以借助单调递减队列去解决该问题
func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	res := make([]int, 0)
	// 正常定长滑窗
	for r := 0; r < len(nums); r++ {
		// 先排出一下窗口要排出的数，如果这个数刚好是最大值，那么从 queue 中排出
		// 并且 r 要已经正式成长为 k 长度的滑动窗口后才开始判断是否需要排出
		if r >= k && queue[0] == nums[r-k] {
			queue = queue[1:]
		}
		// 维护单调递减队列
		for len(queue) > 0 && queue[len(queue)-1] < nums[r] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[r])
		if r >= k-1 {
			res = append(res, queue[0])
		}
	}
	return res
}

// // 初始写法，后面再整合
// // 定长滑动窗口，窗口倒是正常滑动
// // 但是我们需要记录窗口中的最大值，窗口右移可能会导致最大值被弹出，此时我们需要第二大的值
// // 这里我们可以想到，弹出最大值，于是此时数组的第二大值变成了最大值，显然应该是个单调队列
// // 这里我们可以借助单调非递增队列去解决该问题
// func maxSlidingWindow(nums []int, k int) []int {
// 	queue := make([]int, 0)
// 	// 初始化窗口和 queue
// 	for _, num := range nums[:k] {
// 		if len(queue) > 0 && num > queue[len(queue)-1] {
// 			var j int
// 			for ; j < len(queue); j++ {
// 				if queue[j] < num {
// 					queue = queue[:j]
// 					break
// 				}
// 			}
// 		}
// 		queue = append(queue, num)
// 	}
// 	res := make([]int, 0)
// 	res = append(res, queue[0])
// 	// 正常滑窗
// 	for r := k; r < len(nums); r++ {
// 		// 先排出一下窗口要排出的数，如果这个数刚好是最大值，那么从 queue 中排出
// 		if queue[0] == nums[r-k] {
// 			queue = queue[1:]
// 		}
// 		// 维护单调递减队列
// 		for len(queue) > 0 && queue[len(queue)-1] < nums[r] {
// 			queue = queue[:len(queue)-1]
// 		}
// 		queue = append(queue, nums[r])
// 		res = append(res, queue[0])
// 	}
// 	return res
// }
