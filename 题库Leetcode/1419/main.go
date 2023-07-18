package main

// https://leetcode.cn/problems/minimum-number-of-frogs-croaking/
// 1419. 数青蛙

func minNumberOfFrogs(croakOfFrogs string) int {
	var c, r, o, a, k int
	var res int
	for i := range croakOfFrogs {
		switch croakOfFrogs[i] {
		case 'c':
			c++
		case 'r':
			r++
		case 'o':
			o++
		case 'a':
			a++
		case 'k':
			k++
		}
		// 一定要符合顺序
		if c < r || r < o || o < a || a < k {
			return -1
		}
		// 只有 c 还没和 k 匹配完的，才能进入至少多少只青蛙的计算
		if res < c-k {
			res = c - k
		}
	}
	// 最后一定要由若干有效的 croak 组成
	if c == r && r == o && o == a && a == k {
		return res
	}
	return -1
}
