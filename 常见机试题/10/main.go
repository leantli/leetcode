package main

import "fmt"

/**
题目描述:
给定非空字符串s，将该字符串分割成一些子串，使每个子串的ASCII码值的和均为 水仙花数
1、若分割不成功，则返回0
2、若分割成功且分割结果不唯一，则返回-1
3、若分割成功且分割结果唯一，则返回分割后子串的数目
输入描述:
1、输入字符串的最大长度为200
输出描述
根据题目描述中情况，返回相应的结果
补充说明:
水仙花数”是指一个三位数，每位上数字的立方和等于该数字本身，如371是“水仙花数”，因为: 371 = 3^3 +7^3 + 1^3
**/

func main() {
	fmt.Println(work("AXdddF"))
	fmt.Println(work("abc"))
	fmt.Println(work("f3@d5a8"))
}

func work(str string) int {
	var solutions int
	var getRes func(str string, curSum int) int
	getRes = func(str string, curSum int) int {
		var resCnt int
		for i := range str {
			curSum += int(str[i])
			if curSum > 999 {
				break
			}
			// fmt.Println(curSum)
			if curSum > 99 && curSum < 1000 && isFlower(curSum) {
				resCnt++
				if i == len(str)-1 {
					solutions++
					return resCnt
				}
				getRes(str[i+1:], curSum)
				curSum = 0
			}
		}
		return 0
	}
	times := getRes(str, 0)
	if solutions > 1 {
		return -1
	}
	return times
}

func isFlower(num int) bool {
	res := num
	var cnt int
	for num != 0 {
		n := num % 10
		num /= 10
		cnt += n * n * n
	}
	return res == cnt
}
