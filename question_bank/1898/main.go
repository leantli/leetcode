package main

import "strings"

// https://leetcode.cn/problems/maximum-number-of-removable-characters/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1898. 可移除字符的最大数目

// p 是 s 的子序列
// removable 数组，元素 互不相同，该数组的值是 s 下标的子集
// 找到一个 k (0 <= k <= len(removable))
// 根据 removable 数组的 前 k 个值，从 s 中移除对应下标的字符
// 并且移除后 p 仍是 s 的子序列
// 找到满足以上条件的 最大 k

// 没啥好说的，肯定就是枚举 k，但是这里我们可以采用二分去选取 k, 设为 mid，再去模拟移除操作
// 当仍满足 p 是 s 时，说明 k 可取，并且 k 可能还能继续增大，此时 l = mid
// 不满足时说明 k 不可取，k 只能减小，r = mid

// 这里我们要多些两个函数
// 一个是验证 p 是否是 s 的子序列
// 一个是根据 removable 数组和 k，去移除 s 中对应下标的字符
func maximumRemovals(s string, p string, removable []int) int {
	// k 的取值为 [0,len(removable)]，下面采用开区间
	l, r := -1, len(removable)+1
	for l+1 != r {
		mid := l + (r-l)/2
		if check(remove(removable, mid, s), p) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// 验证 p 是否是 s 的子序列，是则返回 true
// 思路：
// 两个指针，j, i 各自指向 p 和 s 的开头
// 不等时，i 指针移动，相等时 i, j 指针都移动
// 当 j 走到尾则成功，没走到尾则 p 不是 s 的子序列
func check(s, p string) bool {
	if len(p) > len(s) {
		return false
	}
	ss, pp := []byte(s), []byte(p)
	var j int
	for i := 0; i < len(ss); i++ {
		if ss[i] == pp[j] {
			j++
		}
		if j == len(pp) {
			return true
		}
	}
	return j == len(pp)
}

// 根据 removable 数组和 k，去移除 s 中对应下标的字符，并返回最终的 string
// 思路：
// 遍历 s，将字符 append 到新的字符串中，当 下标等于 removable 中的数时略过不 append
// 这里用个 set 存一下 removable 中的下标，效率会高些，或者排序后直接双指针
func remove(removable []int, k int, s string) string {
	set := make(map[int]struct{})
	for i := 0; i < k; i++ {
		set[removable[i]] = struct{}{}
	}
	bs := strings.Builder{}
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if _, ok := set[i]; ok {
			continue
		}
		bs.WriteByte(ss[i])
	}
	return bs.String()
}
