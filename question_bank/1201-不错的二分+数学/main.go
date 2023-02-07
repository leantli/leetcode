package main

import "math"

// https://leetcode.cn/problems/ugly-number-iii/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1201. 丑数 III

// 这题有想到用二分和求比 mid 小的丑数有几个，但是忘记了容斥原理，更忘记了 gcd/lcm .....
// 关于容斥原理,gcd,lcm解析 https://leetcode.cn/problems/ugly-number-iii/solution/-by-lcfgrn-mau6/

// 换个思路，二分枚举数，然后看该数是第几个丑数
func nthUglyNumber(n int, a int, b int, c int) int {
	ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
	abc := lcm(ab, c)
	minVal := min(a, b, c)
	// 确定二分的左右边界
	var l, r int = minVal - 1, minVal*n + 1
	for l+1 != r {
		mid := l + (r-l)/2
		// 求比 mid 小的丑数有几个(容斥原理)
		count := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if count < n {
			l = mid
		} else {
			r = mid
		}
	}
	// 此时 r 落在count刚好>=n的位置
	return r
}

// 求两个数的最大公约数
func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}
	return y
}

// 求两个数的最小公倍数
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

// min 求任意数群中的最小数
func min(args ...int) int {
	min := math.MaxInt
	for _, a := range args {
		if a < min {
			min = a
		}
	}
	return min
}

// // 注意这里的丑数并不是和之前的一样要求质因数，只要求能被整除即可
// // 先模拟一下吧，感觉很可能会超时，dp数组也没法那么长
// func nthUglyNumber(n int, a int, b int, c int) int {
// 	var res int
// 	ai, bi, ci := 1, 1, 1
// 	for i := 1; i <= n; i++ {
// 		va, vb, vc := ai*a, bi*b, ci*c
// 		res = min(min(va, vb), vc)
// 		if res == va {
// 			ai++
// 		}
// 		if res == vb {
// 			bi++
// 		}
// 		if res == vc {
// 			ci++
// 		}
// 	}
// 	return res
// }
