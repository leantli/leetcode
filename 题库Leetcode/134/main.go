package main

// https://leetcode.cn/problems/gas-station/
// 134. 加油站

// 贪心思路，每一次要走到下一个加油站，其实就是当前总油量减去当前的cost，是否大于等于 0
// 如果总油量小于 0 了，其实也意味着，这之前的所有加油站做起始点，总和的油量减去cost，其实都是无法继续到达下一个点的
// 因为这个地方的 cost 值肯定是比较大的，才能导致前面的总油量减去后，直接低于 0
// 所以下一个起始点，必定是该位置 + 1，再继续重新考虑是否能走完全部
// 还有一个点：总油量大于总 cost，才能走完全部
func canCompleteCircuit(gas []int, cost []int) int {
	// curGas 每到一个加油站，剩余的油量
	// res 起始加油点的位置
	// sum 油量和消耗的最终值
	var curGas, res, sum int
	for i, g := range gas {
		curRest := g - cost[i]
		curGas += curRest
		sum += curRest
		if curGas < 0 {
			res = i + 1
			curGas = 0
		}
	}
	if sum < 0 {
		return -1
	}
	return res
}

// // 暴力模拟，leetcode 过不了
// func canCompleteCircuit(gas, cost []int) int {
// 	// 从每个加油站出发
// 	n := len(gas)
// 	for i := 0; i < n; i++ {
// 		var curGas int
// 		// 这个循环是接下来要走的，可能走不成
// 		for j := i; j < n+i; j++ {
// 			idx := j % n
// 			curGas += gas[idx]
// 			curCost := cost[idx]
// 			curGas -= curCost
// 			if curGas < 0 {
// 				break
// 			}
// 		}
// 		if curGas >= 0 {
// 			return i
// 		}
// 	}
// 	return -1
// }
