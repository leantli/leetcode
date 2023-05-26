package main

// https://leetcode.cn/problems/avoid-flood-in-the-city/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1488. 避免洪水泛滥

// // 我的想法是维持一个湖泊布尔数组，i 代表第 i 个湖泊，true 表示该湖泊已满
// // 每次都以遇到 0 开始分隔，比如 [1,2] [0,0] [2,1] 进行分隔，0 0 转为可以抽干的次数
// // 遇见[0,0]后的[2,1]时，我们需要判断 2和1 湖泊是否有水，有则需要动用能力去抽干，否则会满
// // 这里还要额外注意题目没写清楚的，到最后，所有没必要动用的抽水能力，都要抽湖泊 1
// // 还有一个地方要注意，当后面重复之时，我们动用的抽水能力，必须在该湖泊上次注水之后，否则该抽水能力就白给了
// // 因此我们利用抽水能力时，需要保证该抽水时间大于该湖泊上次注水的时间
// // 总结，注意两个点
// // 1. 确定抽水次数够不够
// // 2. 抽水的时间尽量靠前，并且保证在上次下雨和本次下雨时间区间之间
// func avoidFlood(rains []int) []int {
// 	lake := make(map[int]bool)
// 	last := make(map[int]int) // 湖泊上次注水的时间
// 	// 当 index==len(zero) 时，则说明没有可以抽水的机会了
// 	zero := make([]int, 0) // zero 数组记录可以抽水的下标
// 	res := make([]int, len(rains))
// 	for i, v := range rains {
// 		// 1. 晴天的处理
// 		if v == 0 {
// 			zero = append(zero, i)
// 			continue
// 		}
// 		// 2. 下雨的处理
// 		res[i] = -1
// 		// 2.1. 湖泊本身没水就注水的处理
// 		if !lake[v] {
// 			lake[v] = true
// 			last[v] = i
// 			continue
// 		}
// 		// 2.2. 后面都是湖泊已经有水的处理
// 		if len(zero) < 1 {
// 			return []int{}
// 		}
// 		// 用可以抽水的日子，把今天的这个湖给抽了
// 		// 注意，可以抽水日子的区间是，(上次注水时间，本次注水时间)
// 		// 我们要找到第一个大于上次注水时间的课抽水时间
// 		l, r := -1, len(zero)
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if last[v] < zero[mid] {
// 				r = mid
// 			} else {
// 				l = mid
// 			}
// 		}
// 		if r == len(zero) {
// 			return []int{}
// 		}
// 		last[v] = i
// 		res[zero[r]] = v
// 		zero = append(zero[:r], zero[r+1:]...)
// 	}
// 	for _, v := range zero {
// 		res[v] = 1
// 	}
// 	return res
// }

// 其实上面的代码再仔细斟酌后会发现，没有必要去专门去创建一个湖泊数组去关注没有注水的湖泊
// 总结，注意两个点
// 1. 确定抽水次数够不够
// 2. 抽水的时间尽量靠前，并且保证在上次下雨和本次下雨时间区间之间
func avoidFlood(rains []int) []int {
	lake := make(map[int]int) // 已经注水的湖泊下标 : 注水的时间
	zero := make([]int, 0)    // zero 数组记录可以抽水的下标
	res := make([]int, len(rains))
	for i, v := range rains {
		// 1. 晴天的处理
		if v == 0 {
			zero = append(zero, i)
			continue
		}
		// 2. 下雨的处理
		res[i] = -1
		// 2.1. 湖泊本身没水就注水的处理
		if _, ok := lake[v]; !ok {
			lake[v] = i
			continue
		}
		// 2.2. 后面都是湖泊已经有水的处理
		// 用可以抽水的日子，把今天的这个湖给抽了
		// 注意，可以抽水日子的区间是，(上次注水时间，本次注水时间)
		// 我们要找到第一个大于上次注水时间的课抽水时间
		l, r := -1, len(zero)
		for l+1 != r {
			mid := l + (r-l)/2
			if lake[v] < zero[mid] {
				r = mid
			} else {
				l = mid
			}
		}
		if r == len(zero) {
			return []int{}
		}
		lake[v] = i
		res[zero[r]] = v
		zero = append(zero[:r], zero[r+1:]...)
	}
	for _, v := range zero {
		res[v] = 1
	}
	return res
}
