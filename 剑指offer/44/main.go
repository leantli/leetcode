package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 44. 数字序列中某一位的数字

// 0 1-9  10-99  100-999 1000-9999
// 0  9    90     900       9000 90000
// 0 1-9  10-189 190-189+900*3
// 硬找规律，其实也不是规律了，就是找几位数的起始数和起始 n，然后算差距

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	digit, start := 1, 1
	count := 9
	startNum := 1
	for n-count*digit > 0 {
		n -= count * digit
		start += count * digit
		digit++
		count *= 10
		startNum *= 10
	}
	fmt.Printf("start:%d, startNum:%d, n:%d, digit:%d\n", start, startNum, n, digit)
	num := startNum + ((n - 1) / digit)
	// 最后这里也可以迭代一下，这里偷个懒，直接字符串操作了
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}

func main() {
	fmt.Println(findNthDigit(10))
	fmt.Println(findNthDigit(11))
	fmt.Println(findNthDigit(190))
	fmt.Println(findNthDigit(191))
	fmt.Println(findNthDigit(192))
}
