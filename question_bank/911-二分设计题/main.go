package main

// https://leetcode.cn/problems/online-election/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 911. 在线选举

// 只用记录每个时刻的优胜者即可
// 只需要两个东西，一个是时刻数组，用来二分定位最近时间，题目就有给
// 一个用来根据定位到的时刻返回优胜者即可
type TopVotedCandidate struct {
	times       []int
	TimeMapKing map[int]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	// 1. 初始化
	// count数组 -- 计数桶
	count := make([]int, len(persons))
	// 各个时刻对应的票王
	timeMapKing := make(map[int]int)
	// 初始化一个边界值
	timeMapKing[times[0]] = persons[0]
	count[persons[0]]++
	// 2. 计算各个时刻的票王
	for i := 1; i < len(persons); i++ {
		count[persons[i]]++
		// 只有当本次被投票的人的票数比上一时刻的高，此人才可成为票王
		if count[persons[i]] >= count[timeMapKing[times[i-1]]] {
			timeMapKing[times[i]] = persons[i]
		} else {
			// 否则票王就还是上一时刻的票王
			timeMapKing[times[i]] = timeMapKing[times[i-1]]
		}
	}
	return TopVotedCandidate{
		times:       times,
		TimeMapKing: timeMapKing,
	}
}

// t 显然不能只用拿来用，基于 t 用二分定位到可选取的时刻
func (this *TopVotedCandidate) Q(t int) int {
	l, r := -1, len(this.times)
	for l+1 != r {
		mid := l + (r-l)/2
		if this.times[mid] <= t {
			l = mid
		} else {
			r = mid
		}
	}
	// 最终 l 所在的时刻位置，即为最后一个小于等于 t 的时刻
	return this.TimeMapKing[this.times[l]]
}
