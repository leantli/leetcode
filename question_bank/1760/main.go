package main

// https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1760. 袋子里最少数目的球

// 操作至多 maxOperations 次，操作次数可 小于等于 maxOperations
// 题目要求 为 "最小化所有袋子中单个袋子球数目的最大值" 即 "尽量让所有袋子的球的数目均匀"
// 因此容易考虑为 为了减小最大值，在分袋子时自然会选择球最多的袋子，而且要尽可能平均分才能使分开后的最大值更小
// 在思考过后会发现，在只有“操作至多 maxOperations 次”这么一个条件的限制下，要解决“如何分袋子能使袋子的最大值最小？”
// 这个问题并没有十分有效的算法，因为我们很难把握每次分袋子时把袋子平均分成几份才是最好的，除非我们在分每一次袋子时都枚举一遍所有可能性，但这样时间复杂度将十分恐怖
// 与其在每一次分袋子时都枚举一遍分成的份数，不如直接枚举 maxOperations 次操作后的开销是否可能是 y
// 所以我们考虑人为增加一个条件，即“每个袋子至多有y个球”
// 此时问题便转化成了“给定 maxOperations 次操作次数，能否可以使得单个袋子里球数目的最大值不超过 y”
// 如何计算单个袋子数目最大值不超过 y ？ 比如 这个袋子目前球数量为 5-8，y 为 4，此时只需操作 1 次即可
// 因此我们可以理解 对于第 i 个袋子，其中有nums[i] 个球，那么需要的操作次数为 (nums[i]-1)/y 次
// 将所有袋子达到最大值不超过 y 所需要的操作次数累加，得到操作次数 count
// 当操作次数 count < maxOperations 时，说明还能进行额外的拆分动作，y 开销还能更小，缩小 r
// 当操作次数 count > maxOperations 时，说明不能进行这么多的拆分动作，y 开销得更大，扩大 l
// 当操作次数 count == maxOperations，说明 y 值到底了，y 就是 maxOperations 限制下，开销值最小化的值
func minimumSize(nums []int, maxOperations int) int {
	var max int
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	// 确定 l/r 的取值范围为 [1,max]，下面采用开区间二分
	l, r := 0, max+1
	for l+1 != r {
		// 这个 mid 就是上面思路的 y (每个袋子数量最大为 y )
		// 我们需要找到最后一个 count <= maxOperations 的情况
		mid := l + (r-l)/2
		var count int
		for _, num := range nums {
			count += (num - 1) / mid
		}
		if count <= maxOperations {
			r = mid
		} else {
			l = mid
		}
	}
	// 结束二分后，r 落在最后一个 count <= maxOperations 的值上
	return r
}
