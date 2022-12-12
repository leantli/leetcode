package main

// https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 40. 最小的k个数

// 堆排的优化思路 20 ms，全部堆排后取前 k (32 ms)
// 只维护一个大小为 k 的大根堆，然后进行遍历，当遇到小于 大根堆 堆顶的数时，再弹出堆顶，重新维护堆
// 这种和全部堆排再取前 k 相比，效率应该更高些
func getLeastNumbers(arr []int, k int) []int {
	if k == len(arr) {
		return arr
	}
	getHeap(arr[:k], k)
	for _, num := range arr[k:] {
		if num < arr[0] {
			arr[0] = num
			heapify(arr[:k], k, 0)
		}
	}
	return arr[:k]
}

// 堆排的另一种操作，先写个维护堆性质的函数, i 为待维护的局部堆堆顶下标
func heapify(heap []int, len, i int) {
	largest, left, right := i, i*2+1, i*2+2
	if left < len && heap[largest] < heap[left] {
		largest = left
	}
	if right < len && heap[largest] < heap[right] {
		largest = right
	}
	if largest == i {
		return
	}
	heap[largest], heap[i] = heap[i], heap[largest]
	heapify(heap, len, largest)
}

func getHeap(arr []int, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		heapify(arr, len, i)
	}
}

// func getLeastNumbers(arr []int, k int) []int {
// 	if len(arr) == 0 || len(arr) == k {
// 		return arr
// 	}
// 	arr = mergeSort(arr)
// 	return arr[:k]
// }

// // 最后再来个归并好了
// func mergeSort(nums []int) []int {
// 	// 考虑递归截止条件，什么时候停止继续切分
// 	if len(nums) <= 1 {
// 		return nums
// 	}
// 	// 先分后合
// 	mid := len(nums) / 2
// 	left, right := mergeSort(nums[:mid]), mergeSort(nums[mid:])
// 	// 最后合
// 	return merge(left, right)
// }

// // 合并两个数组
// func merge(left, right []int) []int {
// 	l, r := 0, 0
// 	res := make([]int, 0, len(left)+len(right))
// 	for l < len(left) && r < len(right) {
// 		if left[l] < right[r] {
// 			res = append(res, left[l])
// 			l++
// 		} else {
// 			res = append(res, right[r])
// 			r++
// 		}
// 	}
// 	res = append(res, left[l:]...)
// 	res = append(res, right[r:]...)
// 	return res
// }

// // 快排+剪枝 == 快速选择，接下来再看看堆排序咋搞
// // 先写个常规的堆排

// // 先写个大根堆的维护函数, heap 是堆的数组表现形式，len 是该堆的长度，i 是待维护堆(也可以是堆中的局部堆)的堆顶的下标
// func heapify(heap []int, len, i int) {
// 	// 假设当前堆顶 i 是最大值，我们用 lagest 记录它的下标
// 	// i 的左右子节点在数组中是 i*2+1 和 i*2+2
// 	largest, left, right := i, i*2+1, i*2+2
// 	// 将 largest 定位到 i, left, right 三个节点中最大的下标
// 	if left < len && heap[largest] < heap[left] {
// 		largest = left
// 	}
// 	if right < len && heap[largest] < heap[right] {
// 		largest = right
// 	}
// 	// 将最大的节点换到 i 的位置，即堆顶
// 	heap[i], heap[largest] = heap[largest], heap[i]
// 	// 如果堆顶已经是最大的了，就无需再看下面局部堆的情况了
// 	// 此时还要递归地维护下面的局部堆, 此时 largest 指向的位置已经是换过的了
// 	if largest == i {
// 		return
// 	}
// 	heapify(heap, len, largest)
// }

// // 堆排序入口
// func heapSort(heap []int, len int) {
// 	// 建堆，从堆的最后一个非叶子节点开始
// 	for i := len/2 - 1; i >= 0; i-- {
// 		heapify(heap, len, i)
// 	}
// 	// 建成大根堆后，利用大根堆的性质来排序
// 	// 将数组分成两部分，前面的部分仍是大根堆，后面的部分是已排序数组
// 	// 从数组尾部开始放堆顶, 再重新维护前面的大根堆，得到新的堆顶，再放在已排序数组的首位
// 	for i := len - 1; i >= 0; i-- {
// 		// 堆顶放在已排序数组的最前端
// 		heap[0], heap[i] = heap[i], heap[0]
// 		// 再重新维护大根堆的堆顶(0)
// 		heapify(heap, i, 0)
// 	}
// }

// // 非常经典的一道题，一看到就能想到快速搜索(快排+只排最小 k 个)
// // 但是不管怎么样，先写个快排吧
// func getLeastNumbers(arr []int, k int) []int {
// 	if len(arr) == 0 || len(arr) == k {
// 		return arr
// 	}
// 	return quickSearch(arr, 0, len(arr)-1, k)
// }

// // 我们知道 pivot 基准点，就是一次排序的中值，排序后的 i 也就是这个中值在整个数组中的下标
// // 当这个下标为 k 时，[0,k)其实就是最小的 k 个数了，
// // 尽管还未对其进行更细化的排序，但是他们已经是确定的最小的 k 个数，直接返回即可
// // 因此我们把正常快排做一个拆分，我们并不直接在一个函数中递归快排，而是做一个拆分
// // 我们把是否继续递归放在另一个函数，把排序逻辑抽出来，每次排序都返回 pivot 对应的下标
// // 是否递归就看这个下标是否已经到 k 了，还大于 k 的话，就继续递归
// func quickSort(nums []int, l, r int) int {
// 	i, j, pivot := l, r, nums[l]
// 	for i < j {
// 		for i < j && nums[j] > pivot {
// 			j--
// 		}
// 		for i < j && nums[i] <= pivot {
// 			i++
// 		}
// 		nums[i], nums[j] = nums[j], nums[i]
// 	}
// 	nums[l], nums[i] = nums[i], nums[l]
// 	return i
// }

// func quickSearch(nums []int, l, r, k int) []int {
// 	mid := quickSort(nums, l, r)
// 	if mid == k {
// 		return nums[:k]
// 	}
// 	// 再根据 mid 目前排序到的位置，去判断接下来是要往左排还是往右排
// 	if mid < k {
// 		return quickSearch(nums, mid+1, r, k)
// 	}
// 	return quickSearch(nums, l, mid-1, k)
// }
