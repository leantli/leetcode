package main

import (
	"sort"
)

// https://leetcode.cn/problems/reconstruct-itinerary/
// 332.重新安排行程

type pair struct {
	target  string
	visited bool
}

func findItinerary(tickets [][]string) []string {
	// 1. 先将出发地和对应的能到达的目的地都存入 map 中，并且对目的地数组进行升序排序，便于取字典序最小的路线
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string][]pair)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make([]pair, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], pair{ticket[1], false})
	}
	for k, _ := range targets {
		sort.Slice(targets[k], func(i, j int) bool { return targets[k][i].target < targets[k][j].target })
	}
	// 2. 接着开始 dfs 枚举，这里我们要使用票的话，很不方便通过 used 去判断，因此封装成 pair 结构体，通过 visited 去判断是否使用过机票
	result := []string{"JFK"}
	var backtracking func() bool
	backtracking = func() bool {
		// 当游历过的地点数量 == 机票数量+1 时，说明 dfs 完成
		if len(result) == len(tickets)+1 {
			return true
		}
		// 取出起飞航班对应的目的地
		for _, pair := range targets[result[len(result)-1]] {
			if pair.visited == false {
				result = append(result, pair.target)
				pair.visited = true
				if backtracking() {
					return true
				}
				result = result[:len(result)-1]
				pair.visited = false
			}
		}
		return false
	}
	backtracking()
	return result
}

// // 官解 -- 一笔画问题--欧拉图或半欧拉图
// // 1. 由于题目中说必然存在一条有效路径(至少是半欧拉图)，所以算法不需要回溯（既加入到结果集里的元素不需要删除）
// // 2. 整个图最多存在一个死胡同(出度和入度相差1），且这个死胡同一定是最后一个访问到的，否则无法完成一笔画。
// // 3. DFS的调用其实是一个拆边的过程（既每次递归调用删除一条边，所有子递归都返回后，再将当前节点加入结果集保证了结果集的逆序输出），一定是递归到这个死胡同（没有子递归可以调用）后递归函数开始返回。所以死胡同是第一个加入结果集的元素。
// // 4. 最后逆序的输出即可。
// func findItinerary(tickets [][]string) []string {
// 	m := make(map[string][]string)
// 	res := make([]string, 0)
// 	// 将每个 ticket 对应的目的地，都放入出发点的 value 列表中
// 	for _, ticket := range tickets {
// 		src, dst := ticket[0], ticket[1]
// 		m[src] = append(m[src], dst)
// 	}
// 	// 对每个出发点对应的 目的地列表进行升序排序，便于 dfs 操作同时保证多条路线下获取字典序最小的
// 	for key := range m {
// 		sort.Strings(m[key])
// 	}
// 	// 开始 dfs
// 	var dfs func(curr string)
// 	dfs = func(curr string) {
// 		for {
// 			if dsts, ok := m[curr]; !ok || len(dsts) == 0 {
// 				break
// 			}
// 			dst := m[curr][0]
// 			m[curr] = m[curr][1:]
// 			dfs(dst)
// 		}
// 		// 因为是深搜到最后，从叶子节点最后回到根结点，因此 append 中的地址是逆序的，虽然 dfs("JFK") 最先，但是 "JFK" 在 res 数组中最后，因此最后要逆序一下
// 		res = append(res, curr)
// 	}
// 	dfs("JFK")
// 	for i := 0; i < len(res)/2; i++ {
// 		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
// 	}
// 	return res
// }

// // 修改 res append 方式，减少最后的逆序操作
// func findItinerary(tickets [][]string) []string {
// 	m := make(map[string][]string)
// 	res := make([]string, 0)
// 	// 将每个 ticket 对应的目的地，都放入出发点的 value 列表中
// 	for _, ticket := range tickets {
// 		src, dst := ticket[0], ticket[1]
// 		m[src] = append(m[src], dst)
// 	}
// 	// 对每个出发点对应的 目的地列表进行排序，便于 dfs 操作
// 	for key := range m {
// 		sort.Strings(m[key])
// 	}
// 	// 开始 dfs
// 	var dfs func(curr string)
// 	dfs = func(curr string) {
// 		for {
// 			if dsts, ok := m[curr]; !ok || len(dsts) == 0 {
// 				break
// 			}
// 			dst := m[curr][0]
// 			m[curr] = m[curr][1:]
// 			dfs(dst)
// 		}
// 		res = append([]string{curr}, res...)
// 	}
// 	dfs("JFK")
// 	return res
// }

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
