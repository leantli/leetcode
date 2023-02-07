package main

// https://leetcode.cn/problems/find-latest-group-of-size-m/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1562. 查找大小为 M 的最新分组

// 还有并查集的做法，等以后练并查集题型再补上

// 我们是要求连续1的长度恰好为 m，在第几个步骤，并且该步骤是最后的步骤
// 因此我们可以每次都去判断，我翻了某个0为1，是否会与左右两侧构成新的连续1
// 如果构成，说明左右两侧已经是1了就不会再被翻1，我们只需考虑其他位置，我们也只需要将连续1的左右两侧赋值新的连续1长度
// good solution
func findLatestStep(arr []int, m int) int {
	if m == len(arr) {
		return m
	}
	// 这里要额外多两个空位，防止越界
	lens := make([]int, len(arr)+2)
	ret := -1
	for i, index := range arr {
		// 将 index 翻为 1 后，与左右两侧构建出的连续 1 的新长度
		newLen := lens[index-1] + lens[index+1] + 1
		// 当 两侧连续 1 的长度为 m 时，可以保存一下当前的步骤情况，ret 为 m 的情况已经特殊处理
		if lens[index-1] == m || lens[index+1] == m {
			ret = i
		}
		// 将新的连续1的左右两侧都赋值新的长度
		lens[index-lens[index-1]] = newLen
		lens[index+lens[index+1]] = newLen
	}
	return ret
}

// 其实这道题就像是之前一道求区间的题 436
// 用 map 记录每一个相连的 1 区间, key-value = 起始位置-长度？还是结束位置?
// 如果两个区间中间刚好多了个 1，如何连接两个区间？
// 正向模拟感觉会比较麻烦，区间融合时会需要多些处理，再结合其实我们只要最后步骤长度等于m的情况，因此可以考虑逆向处理
// 逆向处理，直接基于全 1 区间，从后往前遍历 arr，拆分区间，然后遇到 区间长度等于 m 即可返回
// 最后还是超时了
// func findLatestStep(arr []int, m int) int {
// 	if m == len(arr) {
// 		return m
// 	}
// 	groups := make([][]int, 0)
// 	groups = append(groups, []int{1, len(arr)})
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		// 在 groups 中找本次 arr[i] 处在哪个区间
// 		l, r := -1, len(groups)
// 		for l+1 != r {
// 			mid := l + (r-l)/2
// 			if groups[mid][0] <= arr[i] {
// 				l = mid
// 			} else {
// 				r = mid
// 			}
// 		}
// 		// 拆分区间，如果该区间就一个数，则直接删去该区间
// 		if groups[l][0] == groups[l][1] {
// 			groups = append(groups[:l], groups[l+1:]...)
// 		} else {
// 			temp := make([][]int, len(groups[l+1:]))
// 			copy(temp, groups[l+1:])
// 			groups = append(groups[:l+1], []int{arr[i] + 1, groups[l][1]})
// 			groups[l][1] = arr[i] - 1
// 			groups = append(groups, temp...)
// 		}
// 		// 检验本次拆分是否满足条件
// 		for _, group := range groups {
// 			if group[1]-group[0]+1 == m {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }

// // 该数组只包含 1~n， 长度为 n 二进制字符串，默认全为 0
// // 先尝试一下纯模拟，正常来说肯定是会超时的
// func findLatestStep(arr []int, m int) int {
// 	n := len(arr)
// 	str := make([]byte, n)
// 	for i := 0; i < n; i++ {
// 		str[i] = '0'
// 	}
// 	last := -1
// 	for i, num := range arr {
// 		str[num-1] = '1'
// 		for _, slice := range strings.Split(strings.Trim(string(str), "0"), "0") {
// 			if len(slice) == m {
// 				last = i + 1
// 			}
// 		}
// 	}
// 	return last
// }
