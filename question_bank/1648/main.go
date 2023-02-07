package main

// https://leetcode.cn/problems/sell-diminishing-valued-colored-balls/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1648. 销售价值减少的颜色球

// 显然，每次取都取最大的取 orders 次
// 这样下来，结果往往会是使得 inventory 数组中的数趋同或相近
// 题目就可以转换为，令 inventory 中的数趋于某个数的操作次数小于等于 orders
// 在此基础上，去计算 数组中每个数与当前均值的等差之和 + 补齐 orders 次数的 max 操作
func maxProfit(inventory []int, orders int) int {
	mod := int(1e9 + 7)
	l, r := 0, getMaxOfArr(inventory)+1
	for l+1 != r {
		mid := l + (r-l)/2
		// 需要的操作次数少，说明还可以操作更多次，操作更多次的话，mid 作为 avg 可以再小些
		if getOrdersBaseOnAvg(inventory, mid) <= orders {
			r = mid
		} else {
			l = mid
		}
	}
	// 最终 r 为均值，先计算数组中每个数与均值的等差和
	// 注意, r 为均值时，指的是数组中每个数当前都为 r
	// 因此我们等差求和时，起始值是 r+1，终点值是 原值, (r+1+num)*(num-(r+1)+1)
	var res int
	for _, num := range inventory {
		if num > r {
			res = (res + (num-r)*(num+r+1)/2) % mod
		}
	}
	need := getOrdersBaseOnAvg(inventory, r)
	if need == orders {
		return res
	}
	for need < orders {
		res = (res + r) % mod
		need++
	}
	return res
}

// 从数组中取最大值
func getMaxOfArr(arr []int) int {
	var max int
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

// 根据均值，判断数组中每个数趋于均值需要的操作次数
func getOrdersBaseOnAvg(inventory []int, avg int) int {
	var cnt int
	for _, num := range inventory {
		if num > avg {
			cnt += num - avg
		}
	}
	return cnt
}
