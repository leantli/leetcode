package main

import "fmt"

/**
题目描述:
小明从糖果盒中随意抓一把糖果，每次小明会取出一半的糖果分给同学们。
当糖果不能平均分配时，小明可以选择从糖果盒中(假设盒中糖果足够)取出一个糖果或放回一个糖果。
小明最少需要多少次 (取出、放回和平均分配均记一次) ，能将手中糖果分至只剩一颗。
输入描述:
抓取的糖果数 (<10000000000)
15
输出描述:
最少分至一颗糖果的次数:
5
补充说明:
解释: (1)15+1=16;(2)16/2=8;(3)8/2=4;(4)4/2=2:(5)2/2=1;
**/

// dfs 虽然好，但是这道题也很像 剑指14-1，我们可以考虑 dp
// dp[i] 定义为，当有 i 个糖果时，至少要分 dp[i] 次才能让手上剩 1 个
func main() {
	var candies int
	fmt.Scanln(&candies)
	dp := make([]int, candies+1)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= candies; i++ {
		if (i & 1) == 0 {
			// 当 i 刚好是 2 的倍数，此时 dp[i] = dp[i/2] + 1，不需要考虑其他太多
			dp[i] = dp[i/2] + 1
		} else {
			// 当 i 不是 2 的倍数时，向下取 i/2，向上取 (i+1)/2，每次都要进行两次操作
			dp[i] = min(dp[i/2]+2, dp[(i+1)/2]+2)
		}
		fmt.Printf("dp[%d]=%d\n", i, dp[i])
	}
	fmt.Println(dp[candies])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // dfs 递归穷举, 最终获取最少的次数
// var leastCnt int = math.MaxInt

// func minimumSteps(candies, cnt int) {
// 	if candies == 1 {
// 		leastCnt = min(leastCnt, cnt)
// 		return
// 	}
// 	if candies%2 == 0 {
// 		minimumSteps(candies/2, cnt+1)
// 	} else {
// 		minimumSteps(candies+1, cnt+1)
// 		minimumSteps(candies-1, cnt+1)
// 	}
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	var candies int
// 	fmt.Scanln(&candies)
// 	minimumSteps(candies, 0)
// 	fmt.Println(leastCnt)
// }
