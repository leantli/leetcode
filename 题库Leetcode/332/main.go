package main

import (
	"sort"
)

// https://leetcode.cn/problems/reconstruct-itinerary/
// 332.重新安排行程

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	res := make([]string, 0)
	// 将每个 ticket 对应的目的地，都放入出发点的 value 列表中
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	// 对每个出发点对应的 目的地列表进行排序，便于 dfs 操作
	for key := range m {
		sort.Strings(m[key])
	}
	// 开始 dfs
	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}
	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

// // 所有机票只能使用一次，存在多种，按字典序返回最小行程组合
// func findItinerary(tickets [][]string) []string {
// 	res := make([][]string, 0)
// 	used := make([]bool, len(tickets))
// 	// 所有机票都从 JFK 出发
// 	cur := []string{"JFK"}
// 	var dfs func()
// 	dfs = func() {
// 		// 所有票都用完了
// 		if len(cur) == len(tickets)+1 {
// 			// 先不进行字典序排序--等会再写
// 			res = append(res, append([]string{}, cur...))
// 			return
// 		}
// 		for i, ticket := range tickets {
// 			if !used[i] && ticket[0] == cur[len(cur)-1] {
// 				cur = append(cur, ticket[1])
// 				used[i] = true
// 				dfs()
// 				used[i] = false
// 				cur = cur[:len(cur)-1]
// 			}
// 		}
// 	}
// 	dfs()
// 	// OOM，无法通过先获取全部路程再排序的方式进行比较
// 	sort.Slice(res, func(i, j int) bool {
// 		for idx := range res[i] {
// 			if res[i][idx] == res[j][idx] {
// 				continue
// 			}
// 			return res[i][idx] < res[j][idx]
// 		}
// 		return res[i][0] <= res[j][0]
// 	})
// 	return res[0]
// }
