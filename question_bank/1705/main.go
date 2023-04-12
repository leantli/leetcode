package main

import "container/heap"

// https://leetcode.cn/problems/maximum-number-of-eaten-apples/
// 1705. 吃苹果的最大数目

// 和 1353 很相似，但这道题还多了一个变量，即苹果的数量
// 但是没有关系，在堆中的，配合苹果的最后能吃时间组成一个数量即可
// 接下来按照贪心策略---优先吃最快要腐烂的
// 根据最快要腐烂的时间做小顶堆排序即可
// 能吃的吃正常加上，不能吃的就从堆中剔除
// 并且注意，在n天后仍能继续吃苹果
// 因此我们在二阶段还可以额外判断，苹果的腐烂时间短还是剩下的苹果数少，取小值增加天数和吃的苹果数
func eatenApples(apples []int, days []int) int {
	h := hp{}
	var cnt, curDay int
	// 1. 模拟苹果树还在长苹果的阶段
	for i := 0; i < len(apples); i++ {
		if apples[i] > 0 && days[i] > 0 {
			// 苹果能吃的区间[i,i+days[i]-1]
			h.push(pair{end: i + days[i] - 1, left: apples[i]})
		}
		// 剔除已经腐烂的苹果
		for h.Len() > 0 && h[0].end < curDay {
			h.pop()
		}
		// 当堆中还有苹果，就吃
		if h.Len() > 0 {
			h[0].left--
			if h[0].left == 0 {
				h.pop()
			}
			cnt++
		}
		curDay++
	}
	// 2. 苹果树不长苹果后还能继续吃苹果，只要堆中还有苹果即可
	for h.Len() > 0 {
		// 剔除已经腐烂的苹果
		for h.Len() > 0 && h[0].end < curDay {
			h.pop()
		}
		if h.Len() == 0 {
			break
		}
		// 同一天将要过期，我们取数量和(过期时间-现在时间+1)的最小值
		// 因为只能吃这一定数量的苹果
		// cueDay是已经吃过了的？s
		p := h.pop()
		count := min(p.end-curDay+1, p.left)
		cnt += count
		curDay += count
	}
	return cnt
}

// 实现一个小顶堆
type pair struct{ end, left int } // 苹果不腐烂的最后一天时间，剩下的苹果

type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(x pair)        { heap.Push(h, x) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
