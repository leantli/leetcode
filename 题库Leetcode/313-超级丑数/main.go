package main

import "math"

// https://leetcode.cn/problems/super-ugly-number/
// 313. 超级丑数

// 发现一个事，这样两轮遍历比我一开始的一轮遍历同时获取最小+最小下标数组效率更高
// 可能是数组的创建删除等操作反而使效率变低
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	// 各个素数目前 dp 指针的位置 index
	// ids 本身的下标等同于 primes 的下标，但其指向的值为各素数的 dp 中指针的位置
	ids := make([]int, m)
	for i := 0; i < m; i++ {
		ids[i] = 1
	}
	// 各个素数当前的乘积，res 的下标等同于 primes 的下标
	res := make([]int, m)
	for i, prime := range primes {
		res[i] = dp[ids[i]] * prime
	}
	for i := 2; i <= n; i++ {
		min := math.MaxInt
		for _, num := range res {
			if num < min {
				min = num
			}
		}
		dp[i] = min
		for i, num := range res {
			if num == min {
				ids[i]++
				res[i] = primes[i] * dp[ids[i]]
			}
		}
	}
	return dp[n]
}

// // 按照此前的方法，每求一个新的丑数，需要用数组保存每次的计算，遍历数组取最小，再一次遍历移动指针
// // 很可能会超时，尝试一下 // 居然没超时 并且击败了43.93%....
// func nthSuperUglyNumber(n int, primes []int) int {
// 	dp := make([]int, n+1)
// 	dp[1] = 1
// 	m := len(primes)
// 	// 各个素数目前 dp 指针的位置 index
// 	// ids 本身的下标等同于 primes 的下标，但其指向的值为各素数的 dp 中指针的位置
// 	ids := make([]int, m)
// 	for i := 0; i < m; i++ {
// 		ids[i] = 1
// 	}
// 	// 各个素数当前的乘积，res 的下标等同于 primes 的下标
// 	res := make([]int, m)
// 	for i, prime := range primes {
// 		res[i] = dp[ids[i]] * prime
// 	}
// 	for i := 2; i <= n; i++ {
// 		// 遍历 res 数组，获取最小的数及其下标
// 		minRes, indexs := getMinAndIndexOfArr(res)
// 		dp[i] = minRes
// 		for _, index := range indexs {
// 			ids[index]++
// 			res[index] = dp[ids[index]] * primes[index]
// 		}
// 	}
// 	return dp[n]
// }

// // 获取数组最小值及其下标
// func getMinAndIndexOfArr(arr []int) (int, []int) {
// 	min := math.MaxInt
// 	index := make([]int, 0)
// 	for i, num := range arr {
// 		if num <= min {
// 			if num < min {
// 				min = num
// 				index = make([]int, 0)
// 			}
// 			index = append(index, i)
// 		}
// 	}
// 	return min, index
// }
