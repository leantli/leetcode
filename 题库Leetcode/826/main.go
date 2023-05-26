package main

import "sort"

// https://leetcode.cn/problems/most-profit-assigning-work/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 826. 安排工作以达到最大收益

// n 人 m 事
// 每个工人只能安排一个事
// 对每个工人都找到他们能接受的价值最高的工作
// 这里还得注意，并不是难度越高，价值就越高
// 因此我们需要先确定 该工人 能接受的难度范围
// 再在该范围内去寻找价值最高的
// 对每个工人去二分找其能接受难度最高的范围，其实就已经是 nlogn 时间复杂度了
// 因此这里我们同时间复杂度内，大胆采用排序
// 此时难度已经是升序的了，我们再额外起一个，当前难度下，最大的价值是多少
// 这样就可以减少找到难度范围后，再去找最高价值
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	// 先对难度做个排序
	sorts := make([][]int, n)
	for i := 0; i < n; i++ {
		sorts[i] = append(sorts[i], difficulty[i])
		sorts[i] = append(sorts[i], profit[i])
	}
	sort.Slice(sorts, func(i, j int) bool { return sorts[i][0] < sorts[j][0] })
	// 计算难度升序下，对应的最高价值s
	profit[0] = sorts[0][1]
	for i := 1; i < n; i++ {
		profit[i] = max(profit[i-1], sorts[i][1])
	}
	var finalProfit int
	for _, w := range worker {
		// 确定难度的边界，[0,n-1]，下面开区间二分
		l, r := -1, n
		for l+1 != r {
			mid := l + (r-l)/2
			// 工作难度比工人能力低时，工作难度还可以再高，因此右移 l
			if w >= sorts[mid][0] {
				l = mid
			} else {
				r = mid
			}
		}
		// 当 l == -1 时，该工人啥都做不了
		if l == -1 {
			continue
		}
		// 最后 l 落在 最后一个 工作难度 <= 工人能力的 下标
		finalProfit += profit[l]
	}
	return finalProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
