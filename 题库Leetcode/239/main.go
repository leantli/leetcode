package main

// https://leetcode.cn/problems/sliding-window-maximum/
// 239. 滑动窗口最大值

// 二刷
// 最朴素的思路，维系一个 k 大小的滑动窗口，每次移动都排序一次取最大值，时间复杂度显然很高
// 应该是需要空间换时间，借助某种数据结构去存储滑动窗口中的较大的值，从大到小排列，只记录最大值时的话
// 最大值被移出后，又要重新遍历找第二大值，显然是不合理的，所以需要借助一个数据结构，从大到小存滑动窗口中的值
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	// 需要一个单调递减的队列，保证队列中元素数量小于等于 k，并且单调递减
	// 这样能保证在最大值被移出后，还能马上得到第二大的值
	queue := make([]int, 0)
	for i := range nums {
		// 当单调递减队列的队尾小于当前要进滑窗的数，则将队尾去掉，保证队列的单调递减性质
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		// 当滑窗满 k 个元素后，需要排出数字时，查看一下排出的数字，是不是最大值，是则从 queue 中去掉
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		// 当滑窗满 k 个元素后，开始记录每个滑窗的最大值
		if i >= k-1 {
			ans = append(ans, queue[0])
		}
	}
	return ans
}

// // 定长滑动窗口，窗口倒是正常滑动
// // 但是我们需要记录窗口中的最大值，窗口右移可能会导致最大值被弹出，此时我们需要第二大的值
// // 这里我们可以想到，弹出最大值，于是此时数组的第二大值变成了最大值，显然应该是个单调队列
// // 这里我们可以借助单调递减队列去解决该问题
// func maxSlidingWindow(nums []int, k int) []int {
// 	queue := make([]int, 0)
// 	res := make([]int, 0)
// 	// 正常定长滑窗
// 	for r := 0; r < len(nums); r++ {
// 		// 先排出一下窗口要排出的数，如果这个数刚好是最大值，那么从 queue 中排出
// 		// 并且 r 要已经正式成长为 k 长度的滑动窗口后才开始判断是否需要排出
// 		if r >= k && queue[0] == nums[r-k] {
// 			queue = queue[1:]
// 		}
// 		// 维护单调递减队列
// 		for len(queue) > 0 && queue[len(queue)-1] < nums[r] {
// 			queue = queue[:len(queue)-1]
// 		}
// 		queue = append(queue, nums[r])
// 		if r >= k-1 {
// 			res = append(res, queue[0])
// 		}
// 	}
// 	return res
// }

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
