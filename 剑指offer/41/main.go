package main

// https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 41. 数据流中的中位数

// 每次进来一个数就快排一次？然后只取最中间的？可以是可以，但是感觉效率应该不行
// 想到一个，要取中位数，将整个数组看成一个已排序数组，前半部分化成一个大根堆，后半部分化成一个小根堆
// 如此每次都可以更快速地取出中位数，排序时也无需对整个数组，而是只对半部分数组进行排序
// 但是这里对 新进数 去哪个堆的分配策略，我还不够确定，因为 为了求中位数，两个堆存放的值数量需要相等，可能还存在堆顶去另一个堆的情况

// 这个叫对顶堆，一般用于求中位数，维护一个大顶堆和一个小顶堆，并且无论何时，都将新进数丢进小顶堆
// 当小顶堆容量大于大顶堆容量2后，再弹出小顶堆堆顶到大顶堆，以保证，奇数时，可以直接取小顶堆堆顶
// 当然，这个也可以反着来

// 做完可做 480

type MedianFinder struct {
	heapMax []int // 大根堆，存放小的半部分数组
	heapMin []int // 小根堆，存放大的半部分数组
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		heapMax: make([]int, 0),
		heapMin: make([]int, 0),
	}
}

// 440ms 先判断处理一下，再看看情况
func (this *MedianFinder) AddNum(num int) {
	// 当小根堆为空的时候优先进小根堆，以及比小根堆堆顶大的值也直接进小根堆
	// 如果此时长度已经比大根堆的长度大了 2+，则把原先的堆顶给大根堆，因为新进的值必定比堆顶大
	if len(this.heapMin) == 0 || num >= this.heapMin[0] {
		this.heapMin = append(this.heapMin, num)
		if len(this.heapMin) > len(this.heapMax)+1 {
			this.heapMax = append(this.heapMax, this.heapMin[0])
			buildHeapMax(this.heapMax)
			this.heapMin = this.heapMin[1:]
			buildHeapMin(this.heapMin)
		}
		return
	}
	// 大根堆这边的逻辑和小根堆也差不多，但是其长度限制为 小于等于 小根堆的长度，保证中位数返回更方便
	// 并且 <= this.heapMin[0] 的 num 进入 heapMax 后还得先堆化一遍
	this.heapMax = append(this.heapMax, num)
	buildHeapMax(this.heapMax)
	if len(this.heapMax) > len(this.heapMin) {
		this.heapMin = append(this.heapMin, this.heapMax[0])
		buildHeapMin(this.heapMin)
		this.heapMax = this.heapMax[1:]
		buildHeapMax(this.heapMax)
	}
}

// // 600ms 不管那么多有的没的判断，不管怎么样都先去另一个堆排个序，再丢去别的堆
// func (this *MedianFinder) AddNum(num int) {
// 	// 只看 len 长度，当两堆长度相等的时候，先丢进大根堆排个序，再把大根堆堆顶丢进小根堆
// 	if len(this.heapMin) == len(this.heapMax) {
// 		this.heapMax = append(this.heapMax, num)
// 		buildHeapMax(this.heapMax)
// 		this.heapMin = append(this.heapMin, this.heapMax[0])
// 		buildHeapMin(this.heapMin)
// 		this.heapMax = this.heapMax[1:]
// 		buildHeapMax(this.heapMax)
// 		return
// 	}
// 	// 不等时肯定是大根堆长度小，先丢进小根堆排个序，再把小根堆堆顶丢进大根堆
// 	this.heapMin = append(this.heapMin, num)
// 	buildHeapMin(this.heapMin)
// 	this.heapMax = append(this.heapMax, this.heapMin[0])
// 	buildHeapMax(this.heapMax)
// 	this.heapMin = this.heapMin[1:]
// 	buildHeapMin(this.heapMin)
// }

func (this *MedianFinder) FindMedian() float64 {
	// fmt.Printf("min:%v ; max:%v\n", this.heapMin, this.heapMax)
	if len(this.heapMin) > len(this.heapMax) {
		return float64(this.heapMin[0])
	}
	return float64(this.heapMax[0]+this.heapMin[0]) / 2.0
}

func buildHeapMax(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyMax(nums, i)
	}
}

func buildHeapMin(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyMin(nums, i)
	}
}

// 先写个 大根堆 堆化函数, i 是指待堆化的局部堆的堆顶下标
func heapifyMax(nums []int, i int) {
	n := len(nums)
	// 选出堆顶的最大值
	largest, left, right := i, i*2+1, i*2+2
	if left < n && nums[largest] < nums[left] {
		largest = left
	}
	if right < n && nums[largest] < nums[right] {
		largest = right
	}
	// 如果 i 下标本身就已经是正确的堆顶，直接返回即可
	if largest == i {
		return
	}
	// 替换位置并递归判断
	nums[largest], nums[i] = nums[i], nums[largest]
	heapifyMax(nums, largest)
}

// 先写个 小根堆 堆化函数, i 是指待堆化的局部堆的堆顶下标
func heapifyMin(nums []int, i int) {
	n := len(nums)
	// 选出堆顶的最大值
	minest, left, right := i, i*2+1, i*2+2
	if left < n && nums[minest] > nums[left] {
		minest = left
	}
	if right < n && nums[minest] > nums[right] {
		minest = right
	}
	// 如果 i 下标本身就已经是正确的堆顶，直接返回即可
	if minest == i {
		return
	}
	// 替换位置并递归判断
	nums[minest], nums[i] = nums[i], nums[minest]
	heapifyMin(nums, minest)
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
