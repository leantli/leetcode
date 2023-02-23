package main

import (
	"container/heap"
	"sort"
)

// https://leetcode.cn/problems/sliding-window-median/
// 480. 滑动窗口中位数

// // 更易理解更好写的思路
// // 原先想着始终保持窗口有序
// // 正常情况想的是每次添加删除数字后都要重新排序，想来肯定是过不了
// // 事实上，我们只需要初始化长度k的窗口时排一次序
// // 后面基于二分去定位增删，始终保持窗口有序即可，
// // 时间复杂度为 klogk 远低于常规的 (n-k)klogk
// func medianSlidingWindow(nums []int, k int) []float64 {
// 	windows := make([]int, k)
// 	copy(windows, nums[:k])
// 	sort.Ints(windows)
// 	res := append([]float64{}, getMid(windows))
// 	for i := k; i < len(nums); i++ {
// 		// 先删除
// 		deleteIdx := binarySearch(windows, nums[i-k])
// 		windows = append(windows[:deleteIdx], windows[deleteIdx+1:]...)
// 		// 再插入
// 		addIdx := binarySearch(windows, nums[i]) + 1
// 		temp := append([]int{nums[i]}, windows[addIdx:]...)
// 		windows = append(windows[:addIdx], temp...)
// 		// 中位数加入结果
// 		res = append(res, getMid(windows))
// 	}
// 	return res
// }

// // 返回最后一个 小于等于 target 的数的下标
// func binarySearch(nums []int, target int) int {
// 	l, r := -1, len(nums)
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if nums[mid] <= target {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	return l
// }

// func getMid(nums []int) float64 {
// 	n := len(nums)
// 	if n&1 == 1 {
// 		return float64(nums[n/2])
// 	}
// 	return (float64(nums[n/2]) + float64(nums[n/2-1])) / 2.0
// }

// 参考官方的对顶堆实现 + map 延迟删除
// 实现默认是小根堆，这里我们可以通过取负值存取，间接用成大根堆
// 或者自己实现一个 slice，主要点在于 less 函数的实现，但是需要实现两个堆
type Heap struct {
	sort.IntSlice
	size int
}

func (h *Heap) Push(v interface{}) {
	h.size++
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *Heap) Pop() interface{} {
	v := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return v
}

func (h *Heap) pop() int {
	h.size--
	return heap.Pop(h).(int)
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var res []float64
	maxHp, minHp := &Heap{}, &Heap{} // 大根堆放较小的一批数，小根堆放较大的一批数
	deleted := make(map[int]int)
	// 清理堆顶已被删除的数
	prune := func(hp *Heap) {
		// 这里注意用 len 而不是 size，size 是指堆中实际有用数的数量
		// 并且注意这里用 heap.Pop() 而不是 pop()，因为这里删除的size变化之前已经操作了
		for hp.Len() > 0 {
			num := hp.IntSlice[0]
			if hp == maxHp {
				num = -num
			}
			if deleted[num] > 0 {
				deleted[num]--
				heap.Pop(hp)
			} else {
				break
			}
		}
	}
	// 平衡大小根堆，这里保证大根堆的数量永远比小根堆多0或1
	ensureBalance := func() {
		for maxHp.size-minHp.size > 1 {
			heap.Push(minHp, -maxHp.pop())
			// 移除了 maxHp 堆顶，那么再检查一下其新的堆顶
			prune(maxHp)
		}
		for maxHp.size < minHp.size {
			heap.Push(maxHp, -minHp.pop())
			prune(minHp)
		}
	}
	// 优先放大根堆，注意 Heap 实现默认是小根堆，因此大根堆我们取反
	insert := func(num int) {
		if maxHp.Len() == 0 || num <= -maxHp.IntSlice[0] {
			heap.Push(maxHp, -num)
		} else {
			heap.Push(minHp, num)
		}
		ensureBalance()
	}
	// 删除这里注意，我们是虚删除，所以需要手动 size--
	// 等待 prune 的时候直接调用 Pop() 而不是 pop()，注意区分两个 pop
	delete := func(num int) {
		deleted[num]++
		if num <= -maxHp.IntSlice[0] {
			maxHp.size--
			if num == -maxHp.IntSlice[0] {
				prune(maxHp)
			}
		} else {
			minHp.size--
			if num == minHp.IntSlice[0] {
				prune(minHp)
			}
		}
		ensureBalance()
	}
	for i := 0; i < len(nums); i++ {
		insert(nums[i])
		if i >= k {
			delete(nums[i-k])
		}
		// 开始给结果添加滑窗的中位数
		if i >= k-1 {
			if k&1 == 1 {
				res = append(res, float64(-maxHp.IntSlice[0]))
			} else {
				mid := (float64(-maxHp.IntSlice[0]) + float64(minHp.IntSlice[0])) / 2.0
				res = append(res, mid)
			}
		}
	}
	return res
}

// func main() {
// 	s := &Heap{}
// 	fmt.Println(s)
// 	heap.Push(s, 111)
// 	heap.Push(s, 99)
// 	heap.Push(s, 88)
// 	heap.Push(s, 44)
// 	heap.Push(s, 66)
// 	heap.Push(s, 77)
// 	heap.Push(s, 11111111111)
// 	heap.Push(s, 1)
// 	fmt.Println(s)

// 	ss := sort.IntSlice{}
// 	fmt.Println(ss)
// 	ss = append(ss, 111)
// 	ss = append(ss, 99)
// 	ss = append(ss, 88)
// 	ss = append(ss, 44)
// 	ss = append(ss, 66)
// 	ss = append(ss, 77)
// 	ss = append(ss, 11111111111111)
// 	ss = append(ss, 1)
// 	fmt.Println(ss)
// }

// 增加--重平衡--删除--看堆顶是否需要递归删除--重平衡
// // 我的思路：
// // 295. 数据流的中位数的小进阶题
// // 两个堆，一个小顶堆一个大顶堆（对顶堆），但是这里主要还有个问题
// // 滑动窗口排出一个数时，堆也要移除这个数，但是堆一般只支持移除堆顶
// // 这里得考虑一下怎么删除，比如说用map记录一下已经删除的数
// // 当取中位数的时候，弹出的堆顶是这个数，就移除这个数，重新维护堆，map中的记录--
// // 这里还需要注意一个问题，就是两个堆中目前的大小，不能直接 len(heap) 取
// // 因为堆中是存在未删除的数，所以我们还得额外再声明两个变量去计算堆正确的大小
// // 这里大部分情况都能过，但是在删除上还存在一点问题，比如此时要删除滑窗排出的数
// // 而大根堆和小根堆堆顶都是这个数，此时要删除谁？这里有些问题
// // 思路没问题，但是细节上没处理好
// func medianSlidingWindow(nums []int, k int) []float64 {
// 	res := make([]float64, 0)    // 用于记录最终输出的各个窗口的中位值
// 	deleted := make(map[int]int) // 用于记录需要被删除的数及其被删除的次数
// 	maxHeap := make([]int, 0)    // 大根堆存较小的一批数字
// 	minHeap := make([]int, 0)    // 小根堆存较大的一批数
// 	var maxCnt, minCnt int       // 大小根堆目前实际存储数量
// 	// 定长滑窗
// 	for r := 0; r < len(nums); r++ {
// 		// 要保证两个条件：1. 小根堆中数量只大于大根堆数量0~1； 2. 求中位数时涉及到的数不能将要被删除
// 		// 即 正常 nums[r] 入堆，直接进小根堆，然后重新构建，此时看两个堆堆顶是否都未被删除

// 		// 1. 正常的入堆(这里要注意对顶堆性质---保证较大的都在小根堆，较小的都在大根堆)
// 		if minCnt == 0 || nums[r] > minHeap[0] {
// 			minHeap = append(minHeap, nums[r])
// 			buildMinHeap(minHeap)
// 			minCnt++
// 			if minCnt-maxCnt > 1 {
// 				maxHeap = append(maxHeap, minHeap[0])
// 				maxCnt++
// 				minHeap = minHeap[1:]
// 				minCnt--
// 				buildMaxHeap(maxHeap)
// 				buildMinHeap(minHeap)
// 			}
// 		} else {
// 			maxHeap = append(maxHeap, nums[r])
// 			buildMaxHeap(maxHeap)
// 			maxCnt++
// 			if maxCnt > minCnt {
// 				minHeap = append(minHeap, maxHeap[0])
// 				minCnt++
// 				maxHeap = maxHeap[1:]
// 				maxCnt--
// 				buildMaxHeap(maxHeap)
// 				buildMinHeap(minHeap)
// 			}
// 		}

// 		fmt.Printf("正常入堆后---r:%d, minheap:%v, maxheap:%v, deleted:%v,maxcnt:%d,mincnt:%d\n", r, minHeap, maxHeap, deleted, maxCnt, minCnt)

// 		// 2. 看看堆顶是否需要删除
// 		// map 中表示该数需要被删除，并且确认要被删除的数在哪个堆，给对应的堆减少数量
// 		if r >= k {
// 			deleted[nums[r-k]]++
// 			if nums[r-k] <= maxHeap[0] {
// 				maxCnt--
// 			} else {
// 				minCnt--
// 			}
// 		}
// 		// 此时得到两个未被删除过的堆，如果他们的堆顶不在要被删除的名单内，就可以了正常使用
// 		for maxCnt > 0 && deleted[maxHeap[0]] > 0 {
// 			deleted[maxHeap[0]]--
// 			maxHeap = maxHeap[1:]
// 			buildMaxHeap(maxHeap)
// 		}
// 		for minCnt > 0 && deleted[minHeap[0]] > 0 {
// 			deleted[minHeap[0]]--
// 			minHeap = minHeap[1:]
// 			buildMinHeap(minHeap)
// 		}

// 		fmt.Printf("检查是否删除后---r:%d, minheap:%v, maxheap:%v, deleted:%v,maxcnt:%d,mincnt:%d\n", r, minHeap, maxHeap, deleted, maxCnt, minCnt)

// 		// 3. 平衡两个堆中的数量
// 		for minCnt-maxCnt > 1 {
// 			maxHeap = append(maxHeap, minHeap[0])
// 			maxCnt++
// 			minHeap = minHeap[1:]
// 			minCnt--
// 			buildMaxHeap(maxHeap)
// 			buildMinHeap(minHeap)
// 		}
// 		for minCnt < maxCnt {
// 			minHeap = append(minHeap, maxHeap[0])
// 			minCnt++
// 			maxHeap = maxHeap[1:]
// 			maxCnt--
// 			buildMaxHeap(maxHeap)
// 			buildMinHeap(minHeap)
// 		}

// 		fmt.Printf("平衡后---r:%d, minheap:%v, maxheap:%v, deleted:%v,maxcnt:%d,mincnt:%d\n", r, minHeap, maxHeap, deleted, maxCnt, minCnt)

// 		// 4. 开始给 res 添加窗口的中位值了
// 		if r >= k-1 {
// 			// 确认中位数
// 			var mid float64
// 			if k&1 == 0 {
// 				mid = float64(minHeap[0]+maxHeap[0]) / 2.0
// 			} else {
// 				mid = float64(minHeap[0])
// 			}

// 			fmt.Printf("确认中位数---r:%d, minheap[0]:%d, maxheap[0]:%d, mid:%f\n", r, minHeap[0], maxHeap[0], mid)
// 			res = append(res, mid)
// 		}
// 	}
// 	return res
// }

// // 要用对顶堆的话，先写两个堆化函数+构建堆函数

// // 大根堆构建函数
// func buildMaxHeap(nums []int) {
// 	n := len(nums)
// 	// 最后一个非叶子节点是 n/2-1 的下标
// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapifyMax(nums, i)
// 	}
// }

// // 小根堆构建函数
// func buildMinHeap(nums []int) {
// 	n := len(nums)
// 	// 最后一个非叶子节点是 n/2-1 的下标
// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapifyMin(nums, i)
// 	}
// }

// // 大根堆的维护函数
// func heapifyMax(nums []int, i int) {
// 	// k 为当前的堆顶，我们要不断地去置换堆顶及其下面的节点
// 	n := len(nums)
// 	left, right, maximum := 2*i+1, 2*i+2, i
// 	if left < n && nums[left] > nums[maximum] {
// 		maximum = left
// 	}
// 	if right < n && nums[right] > nums[maximum] {
// 		maximum = right
// 	}
// 	if maximum == i {
// 		return
// 	}
// 	nums[i], nums[maximum] = nums[maximum], nums[i]
// 	heapifyMax(nums, maximum)
// }

// // 小根堆的维护函数
// func heapifyMin(nums []int, i int) {
// 	// k 为当前的堆顶，我们要不断地去置换堆顶及其下面的节点
// 	n := len(nums)
// 	left, right, minumum := 2*i+1, 2*i+2, i
// 	// 找到最小值的位置
// 	if left < n && nums[left] < nums[minumum] {
// 		minumum = left
// 	}
// 	if right < n && nums[right] < nums[minumum] {
// 		minumum = right
// 	}
// 	// 如果最小值就是 i，直接返回了
// 	if minumum == i {
// 		return
// 	}
// 	// 否则就替换并且继续维护下面的节点
// 	nums[i], nums[minumum] = nums[minumum], nums[i]
// 	heapifyMin(nums, minumum)
// }
