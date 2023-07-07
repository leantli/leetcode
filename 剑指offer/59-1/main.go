package main

// https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 59-1. 滑动窗口的最大值

// 怎么想到用单调递减队列？
// 我们需要维护窗口内的最大值，总不能每一次移动，都排序获取最大值吧？最大值要是被移除了，我们又要马上获取窗口内的第二大值对吧
// 第二大值也被移除了，我们就需要马上获取到第三大的值，显然，最好是基于滑动窗口内的值，使用单调递减队列
// 入时，判断 queue 中最后的数，如果比当前的数小，则将queue最后的数去除掉，保证当前的数，能够成功队列中第 k 大的值
// 出时，每一次移动，判断队列长度是否等于 k，是的话，就判断窗口的最大值是不是nums[i-k]的值，是则移除()
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 0, n-k+1)
	// 单调递减队列
	queue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i >= k && queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		if i >= k-1 {
			res = append(res, queue[0])
		}
	}
	return res
}

// // 看了题解，发现使用单调队列，但是怎么想到要用单调队列的呢？
// // 由于我们需要求出的是滑动窗口的最大值，如果当前的滑动窗口中有两个下标 i 和 j，其中 i 在 j 的左侧（i < j)
// // 并且 i 对应的元素不大于 j 对应的元素( nums[i] < nums[j] )，那么会发生什么呢
// // 当滑动窗口向右移动时，只要 i 还在窗口中，那么 j 一定也还在窗口中，这是 i 在 j 的左侧所保证的
// // 因此，由于 nums[j] 的存在，nums[i] 一定不会是滑动窗口中的最大值了，我们可以将 nums[i] 永久地移除。
// // 因此我们可以使用一个队列存储所有还没有被移除的值
// // 这个单调队列有这么两个性质：
// // 1. 队列长度小于等于k
// // 2. 队列内从头到尾单调递减
// // 此时如何保证这个队列的性质？
// // 新进一个值，将队列中比该值小的都剔除掉，新进值位于队尾，保证队列中的单调递减，便于队列前面大的数值被移除后能及时取到新的最大值
// // 新进一个值，此时队列长度大于等于 k，便判断当前的最大值是否是即将要排出的值，是则去掉，不是则无需去掉(因为后进的最大值会将前面进去的值全部排除)
// func maxSlidingWindow(nums []int, k int) []int {
// 	queue := []int{}
// 	res := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		// 剔除掉比新进值小的数，再将新进值添加到队尾
// 		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
// 			queue = queue[:len(queue)-1]
// 		}
// 		queue = append(queue, nums[i])
// 		// 当队列长度大于等于 k 时，则判断最大值是否时要排出的值，是则去掉，不是则无需去掉(因为后进的最大值会将前面进去的值全部排除)
// 		if i >= k && queue[0] == nums[i-k] {
// 			queue = queue[1:]
// 		}
// 		// 当窗口填充完毕后则可以开始添加结果
// 		if i >= k-1 {
// 			res = append(res, queue[0])
// 		}
// 	}
// 	return res
// }

// 重写时思路
// 暴力 nk 的时间复杂度
// 窗口就是堆！堆的维护 O(logn) 堆排才是 nlogn，则该算法的整体时间复杂度是 nlogk
// 虽然可以用堆，但是感觉比较麻烦，考虑一下可否单调队列
// 这里我们得观察到，晚进的最大值，一定比前面的值大，也就是说，比新进的值小的数可以直接从窗口排掉了，因为他们不可能成为窗口的最大值
// 也就是说，对于窗口，我们可以只维护可能成为最大值的数，按照下标从小到大，但窗口内的值要从大到小
// 当移动时，判断窗口的最大值，也就是下标最小的值，是否被排除即可
// 这就是单调队列
// func maxSlidingWindow(nums []int, k int) []int {
//     res := make([]int, 0)
//     queue := make([]int, 0)
//     for i := 0; i < len(nums); i++ {
//         // 不断排除比新进值小的数，因为他们在后续的移动中不可能再成为窗口的最大值
//         for len(queue) != 0 && queue[len(queue)-1] < nums[i] {
//             queue = queue[:len(queue)-1]
//         }
//         queue = append(queue, nums[i])
//         // 判断窗口的移动是否导致最大值变更了
//         if i > k-1 && queue[0] == nums[i-k] {
//             queue = queue[1:]
//         }
//         // 当有窗口后，就可以给结果数组添加窗口最大值了
//         if i >= k-1 {
//             res = append(res, queue[0])
//         }
//     }
//     return res
// }

// // 本想双指针+维系最大值，但仔细再考虑一下其实不行，因为窗口移出的可能是最大值
// // 移出的如果不是最大值，则将最大值和新进值比较即可
// // 但是如果移出的是最大值，此时不得重新遍历 k 计算最大值？ 时间复杂度为 O((n-k+1)*k) 近似 O(nk)
// // 有没有什么办法可以不用重新遍历?
// // 除了滑动窗口之外，再维系一个最大值数组？或者滑动窗口本身就是一个大根堆？如此一来，堆的维护是 O(logk)，整体时间复杂度近似 O(nlogk)
// // 那先实现个 O(nk) 吧
// // 此时跑测试用例会过不了，超时
// func maxSlidingWindow(nums []int, k int) []int {
// 	res := make([]int, 0)
// 	// 窗口初始化
// 	m := math.MinInt
// 	i := 0
// 	window := make([]int, 0, k+1)
// 	for i < k {
// 		m = max(m, nums[i])
// 		window = append(window, nums[i])
// 		i++
// 	}
// 	res = append(res, m)
// 	// 滑动窗口
// 	for ; i < len(nums); i++ {
// 		// 不管怎么样先加新的进来
// 		window = append(window, nums[i])
// 		temp := window[0]
// 		// 移出左侧第一个
// 		window = window[1:]
// 		// 移出的不是最大值时，只需要比较新进的值和最大值即可
// 		if m != temp {
// 			m = max(nums[i], m)
// 			res = append(res, m)
// 			continue
// 		}
// 		// 移除了最大，则重选最大
// 		m = math.MinInt
// 		for _, num := range window {
// 			m = max(m, num)
// 		}
// 		res = append(res, m)
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 暴力的简写版，主要在于窗口的写法
// func maxSlidingWindow(nums []int, k int) []int {
//     res := make([]int, 0)
//     for i := 0; i < len(nums)-k+1; i++ {
//         res = append(res, max(nums[i:i+k]))
//     }
//     return res
// }

// func max(nums []int) int {
//     res := math.MinInt32
//     for _, num := range nums {
//         if num > res {
//             res = num
//         }
//     }
//     return res
// }
