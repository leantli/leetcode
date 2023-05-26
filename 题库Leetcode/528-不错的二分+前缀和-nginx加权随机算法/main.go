package main

// https://leetcode.cn/problems/random-pick-with-weight/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 528. 按权重随机选择

// // w[i]最小为1;
// // 数组假设为[3, 1, 2, 4], 前缀和为[3, 4, 6, 10];
// // 1, 2, 3 -> 下标0
// // 4 ->	下标1
// // 5, 6 -> 下标2
// // 7, 8, 9, 10 -> 下标3
// // 即寻找第一个大于等于随机数的下标
// type Solution struct {
// 	sum []int
// }

// func Constructor(w []int) Solution {
// 	sum := make([]int, len(w)+1)
// 	for i, v := range w {
// 		sum[i+1] = sum[i] + v
// 	}
// 	return Solution{sum: sum}
// }

// func (this *Solution) PickIndex() int {
// 	// 随机一个数，其值取值范围为 [0,前缀和最后一个数)
// 	target := rand.Intn(this.sum[len(this.sum)-1])
// 	// 接下来基于二分，寻找第一个大于该随机数的前缀和位置
// 	l, r := -1, len(this.sum)
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if this.sum[mid] > target {
// 			r = mid
// 		} else {
// 			l = mid
// 		}
// 	}
// 	// 最终 r 落在第一个大于 随机数 的下标
// 	// 但是我们要注意，sum 比原 w 数组下标都多 1，因此我们最后要减去 1
// 	return r - 1
// }

// // nginx 加权随机方法
// type Solution struct {
// 	now   []int // 当前权重值，用于根据权重返回随机下标
// 	init  []int // 初始权重值
// 	total int   // 总权重值
// }

// func Constructor(w []int) Solution {
// 	// 计算总权重值
// 	var total int
// 	for i := 0; i < len(w); i++ {
// 		total += w[i]
// 	}
// 	// 先复制个 当前权重 数组，便于后续操作
// 	now := make([]int, len(w))
// 	copy(now, w)
// 	return Solution{now: now, init: w, total: total}
// }

// func (this *Solution) PickIndex() int {
// 	var maxIndex int
// 	// 获取 now 数组中权重最大的下标
// 	for i := 0; i < len(this.now); i++ {
// 		if this.now[i] > this.now[maxIndex] {
// 			maxIndex = i
// 		}
// 	}
// 	fmt.Printf("maxIndex:%d\n", maxIndex)
// 	// 接着将 now 数组中的最大权重值减去总权重值
// 	this.now[maxIndex] -= this.total
// 	fmt.Printf("now:%v\n", this.now)
// 	// 再对 now 数组加上 init 数组(初识数组) 相对应的值
// 	for i := 0; i < len(this.now); i++ {
// 		this.now[i] += this.init[i]
// 	}
// 	fmt.Printf("now:%v\n", this.now)
// 	fmt.Println()
// 	return maxIndex
// }

// // 第一想法是根据权重生成一个数组，再随机，但是这样的话如果有权重 999999999 这类的就会有问题
// // 有个很坏的想法，不知道能不能行，就我也不随机，就顺序返回，但是我按权重顺序返回，试试看 // 没成功
// type Solution struct {
// 	w     []int
// 	ww    []int
// 	index int
// }

// func Constructor(w []int) Solution {
// 	ww := make([]int, len(w))
// 	copy(ww, w)
// 	return Solution{w: w, ww: ww}
// }

// func (this *Solution) PickIndex() int {
// 	if this.index == len(this.w) {
// 		this.index = 0
// 	}
// 	for this.w[this.index] == 0 {
// 		this.w[this.index] = this.ww[this.index]
// 		if this.index < len(this.w) {
// 			this.index++
// 		} else {
// 			this.index = 0
// 		}

// 	}
// 	this.w[this.index]--
// 	temp := this.index
// 	this.index++
// 	return temp
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
