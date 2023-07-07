package main

import "fmt"

// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 10.1 斐波那契数列

// 没想到官方还有一种解法 将 f(n) 用矩阵来解，发现可以使用矩阵快速幂，时间复杂度为 log(n)
func fib(n int) int {
	if n < 2 {
		return n
	}
	// 这个 {{1,1},{1,0}} 也是很关键的
	// {{1,1},{1,0}} * {{F(n)},{F(n-1)}} = {{F(n)+F(n-1)},{F(n)}} = {{F(n+1)},{F(n)}}
	// 以此类推
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

type matrix [2][2]int

const mod int = 1e9 + 7

// 矩阵相乘
// [1 2 4]   [1 2]     [1*1+2*3+4*0    1*2+2*2+4*5]
// ------- * [3 2]  =  ---最终矩阵的行等于前者的行，列等于后者的列
// [2 0 3]   [0 5]     [2*1+0*3+3*0    2*2+0*2+3*5]
// 同时注意，前者的列数等同于后者的行数
// len(a) = 2; len(a[0]) = 3; len(b) = 2;
func multipy(a, b matrix) matrix {
	var res matrix
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			for k := 0; k < len(b); k++ {
				res[i][k] = (res[i][k] + a[i][j]*b[j][k]) % mod
			}
		}
	}
	return res
}

// 矩阵快速幂
func pow(x matrix, n int) matrix {
	// 根据矩阵的性质，在常规快速幂中等同于 res = 1
	res := matrix{{1, 0}, {0, 1}}
	for n != 0 {
		if (n & 1) != 0 {
			res = multipy(x, res)
		}
		n >>= 1
		x = multipy(x, x)
	}
	return res
}

// 复习一下快速幂
// 常规模拟幂
// func pow(x, n int) int {
// 	// 需要一个 res 保证 n=0 时返回 1
// 	res := 1
// 	for n != 0 {
// 		res *= x
// 		n--
// 	}
// 	return res
// }

// 快速幂 本质 --> x^(13) = x^(1101)_2 = x^(2^3) * x^(2^2) * x(2^0) = x^8 * x^4 * x^1
// 本来要反复乘 13 次，现在乘 3(二进制位为1的位数)+4(二进制的位数)
// func pow(x, n int) int {
// 	res := 1
// 	for n != 0 {
// 		if (n & 1) != 0 {
// 			res *= x
// 		}
// 		n >>= 1
// 		x *= x
// 	}
// 	return res
// }

// 常规
// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	res, pre, post := 0, 0, 1
// 	for i := 1; i < n; i++ {
// 		res = (pre + post) % (1e9 + 7)
// 		pre = post
// 		post = res
// 	}
// 	return res
// }

// 先用 map 整一遍
// func fib(n int) int {
// 	m := make(map[int]int)
// 	m[0], m[1] = 0, 1
// 	for i := 2; i <= n; i++ {
// 		m[i] = (m[i-1] + m[i-2]) % (1e9 + 7)
// 	}
// 	return m[n]
// }

func main() {
	a := matrix{{1, 2}, {3, 2}}
	b := matrix{{2, 2}, {3, 1}}
	fmt.Println(multipy(a, b))
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(fib(i))
	// }
}
