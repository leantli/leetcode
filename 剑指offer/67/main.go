package main

import "math"

// https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 67. 把字符串转换成整数

// emm，先随便试试
// 试试就逝世，吐了
// 看了自己以前的题解，主要就是不要在一个大的 for 里面去判读各种情况
// 将整个表达式分成三个部分，空格/符号/数值，分别进行处理最佳
func strToInt(str string) int {
	len := len(str)
	if len == 0 {
		return 0
	}
	index := 0
	// 先对空格做处理，不要在一个大的 for 循环中处理各种情况
	for str[index] == ' ' {
		index++
		if index == len {
			return 0
		}
	}
	// 对符号做处理
	sign := 1
	if str[index] == '-' {
		sign = -1
	}
	if str[index] == '-' || str[index] == '+' {
		index++
	}
	// 对数值做处理，res 用于返回结果, bigJudge 判断是否在[INT_MIN,  INT_MAX]区间内
	// go int 在 64 位机上默认是 64 位，因此不用担心溢出
	// 担心溢出的可以 math.MaxInt32/10 在此阶段进行判断
	res, bigJudge := 0, math.MaxInt32
	for i := index; i < len; i++ {
		if str[i] < '0' || bs[i] > '9' {
			break
		}
		res = res*10 + int(str[i]-'0')
		if res > bigJudge {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	return res * sign
}
