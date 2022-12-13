package main

import "fmt"

// https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 16. 数值的整数次方

// 快速幂
// 3^(11) = 3^(1011)_2 = 3^0 * 3^1 * 3^3 =

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	var res float64 = 1
	for n != 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(myPow(3.1, -3))
}
