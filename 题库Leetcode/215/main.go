package main

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
// 215. 数组中的第K个最大元素

// 其实我们能观察到，partition 会返回基准点的下标，显然，下标+1 指向的值，就是整个数组中第 下标+1 大的数
// 因此，我们无需对整个数组进行排序，我们只需要排序到 k-1 下标的位置的值即可
// 值得注意的是，之前的 partition，我都是传入了新的数组，而不是基于 l 和 r 进行区分边界
// 因此返回的基准点，对 k 和 nums 数组并不通用，所以我们在 partition 时，需要传入 nums 当前分区的左右闭区间边界下标
// 保证输出的基准点的 index 是符合 nums 整个数组的
func findKthLargest(nums []int, k int) int {
	quickSelect(nums, k, 0, len(nums)-1)
	return nums[k-1]
}

func quickSelect(nums []int, k, i, j int) {
	mid := partition(nums, i, j)
	if mid == k-1 {
		return
	}
	if mid > k-1 {
		quickSelect(nums, k, i, mid-1)
	} else {
		quickSelect(nums, k, mid+1, j)
	}
}

// 基于基准点分区，返回基准点的下标
func partition(nums []int, i, j int) int {
	pivot, l, r := nums[i], i, j
	for l < r {
		for l < r && nums[r] < pivot {
			r--
		}
		for l < r && nums[l] >= pivot {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[l], nums[i] = nums[i], nums[l]
	return l
}

// // 快排，找基准点，接着基于基准点将数组区分成三个部分：比基准点大的一个区间，比基准点小的一个区间，基准点自己位于中间
// func findKthLargest(nums []int, k int) int {
// 	quickSort(nums)
// 	return nums[k-1]
// }

// func quickSort(nums []int) {
// 	if len(nums) == 0 {
// 		return
// 	}
// 	index := partition(nums)
// 	quickSort(nums[:index])
// 	quickSort(nums[index+1:])
// }

// // 根据基准值切分左右区间，左边的区间都是大于基准值的，右边的区间都是比基准值小的
// func partition(nums []int) int {
// 	pivot, l, r := nums[0], 0, len(nums)-1
// 	for l < r {
// 		// 注意：这里要先移动右指针，为什么呢？右指针不断地遍历，停下时指向的一定是 >= pivot 的值，如果此时 l==r 跳出循环
// 		// nums[0] 和 l 或者 r 指向的值交换时，仍能够保证我们要的 partition 性质
// 		// 如果是先移动左指针，停下时指向的一定是 < pivot 的，此时如果l==r跳出循环，
// 		// 这个值和 pivot 交换，那么这个比 pivot 小的值就位于了左区间，而左区间存的是要比 pivot 大的，所以这样不符合性质
// 		for l < r && nums[r] < pivot {
// 			r--
// 		}
// 		for l < r && nums[l] >= pivot {
// 			l++
// 		}
// 		nums[l], nums[r] = nums[r], nums[l]
// 	}
// 	nums[0], nums[l] = nums[l], nums[0]
// 	return l
// }

// // 先尝试一下堆排，最大的k个元素，我们使用小根堆，存储最大的k个元素，堆顶是最大的k个元素里面最小的，只要比这个最小的大，就可以了替换掉
// func findKthLargest(nums []int, k int) int {
// 	heapify(nums, k)
// 	if k == len(nums) {
// 		return nums[0]
// 	}
// 	for i := k; i < len(nums); i++ {
// 		if nums[i] > nums[0] {
// 			nums[0] = nums[i]
// 			heapSort(nums, k, 0)
// 		}
// 	}
// 	return nums[0]
// }

// // 小根堆实现
// func heapify(nums []int, n int) {
// 	// 从最后一个非叶子节点开始向前遍历排序
// 	for i := (n - 1) / 2; i >= 0; i-- {
// 		heapSort(nums, n, i)
// 	}
// }

// func heapSort(nums []int, n, i int) {
// 	minest := i
// 	l, r := i*2+1, i*2+2
// 	if l < n && nums[l] < nums[minest] {
// 		minest = l
// 	}
// 	if r < n && nums[r] < nums[minest] {
// 		minest = r
// 	}
// 	if minest == i {
// 		return
// 	}
// 	nums[i], nums[minest] = nums[minest], nums[i]
// 	heapSort(nums, n, minest)
// }
