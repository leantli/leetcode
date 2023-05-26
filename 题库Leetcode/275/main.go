package main

// https://leetcode.cn/problems/h-index-ii/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 275. H 指数 II

// citations 升序
// citations[i] 表示研究者的第 i 篇论文被引用的次数
// 求最大 h 指数，有 h 篇论文被引用了 >=h 次，其他的引用次数 < h
// 这里显然是要用二分的，数组排序+对数时间复杂度
// 显然常规是枚举 h，但这里我们可以二分 h
// 以 [0,1,3,5,6] 3 为例
// 假设 h=1， >= 1 引用的论文有 4 篇，此时 h 可增大，l = mid
// 假设 h=4, >= 4 引用的论文有 2 篇，此时 h 得减小, r = mid
// 不确定这个缩进是否正确，先尝试一下
// 没啥问题，但是这样显然时间复杂度还是太高了，关键在于 check，时间复杂度为 n，导致总体 nlogn
// func hIndex(citations []int) int {
// 	// h 的取值范围为 [0,n]？下面采用开区间二分
// 	l, r := -1, len(citations)+1
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if check(citations, mid) {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	return l
// }

// // 当 h 需要增大时，返回 true
// func check(citations []int, h int) bool {
// 	var cnt int
// 	for _, cicitation := range citations {
// 		if cicitation >= h {
// 			cnt++
// 		}
// 	}
// 	return cnt >= h
// }

// 但显然我们没有利用本题的「数组有序」的特性
// 根据对 H 指数 定义，如果 citations 升序，在最大的符合条件的分割点 x 的右边（包含分割点），必然满足 citations[i] >= x，
// 我们应当对其进行计数，对于分割点的左边，必然不满足 citations[i] >= x，无需进行计数。
// 因此，我们可以利用 分割点右边论文的个数与分割点 citations[x] 的大小关系进行二分
// 假设存在真实分割点下标 x，其值大小为 citations[x]，分割点右边的数值个数为 n - x，根据 H 指数 的定义，必然有 citations[x] >= n - x 关系：
// 在分割点 x 的右边：citations[i] 非严格单调递增，而论文的个数严格单调递减，仍然满足 citations[i] >= n - i 关系；
// 在分割点 x 的左边：citations[i] 非严格单调递增，论文的个数严格单调递增，x 作为真实分割点，因此必然不满足 citations[i] >= n - i 关系
// 即现在并不是是之前的二分找 h，而是找分割点
// 找索引 i，并且满足条件 n-i <= h，此时找最小索引 i，能得到最大 h
// 即满足上述条件就缩 r，最终 r 落在满足条件的最小 i 上，再 n-r 得到 h
// 但是这个转换感觉挺难想到的。。。
func hIndex(citations []int) int {
	// h 的取值范围为 [0,n]？下面采用开区间二分
	n := len(citations)
	l, r := -1, n
	for l+1 != r {
		mid := l + (r-l)/2
		if citations[mid] >= n-mid {
			r = mid
		} else {
			l = mid
		}
	}
	return n - r
}
