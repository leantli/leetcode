package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 46. 把数字翻译成字符串

// 极简的切分数字，输出有多少种分类
// dp[i] = { dp[i-1], dp[i-2] + dp[i-1] }
// dp[i] = dp[i-1] 时，i与i-1不在 [0,26)
// dp[i] = dp[i-2] + dp[i-1] 时，i与i-1在 [0,26)
// 还是类似跳台阶，但是多了条件限制，需要判断

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	s := strconv.Itoa(num)
	n := len(s)
	// 这里值得注意一下，dp[i] 对应的是 s[i-1] 的位置
	dp := make([]int, n+1, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		temp := s[i-2 : i]
		if temp >= "10" && temp < "26" {
			dp[i] = dp[i-1] + dp[i-2]
			continue
		}
		dp[i] = dp[i-1]
	}
	return dp[n]
}

func main() {
	fmt.Println(translateNum(12258))
}
