package main

import (
	"fmt"
)

func main() {
	fmt.Println(work("3A3B", "2A4B"))
	fmt.Println(work("5Y5Z", "5Y5Z"))
	fmt.Println(work("4Y5Z", "9Y"))
}

func work(str1, str2 string) string {
	list1, list2 := make([][2]int, 0), make([][2]int, 0)
	total := getList(str1, &list1)
	getList(str2, &list2)
	index := 0
	len1, len2 := len(list1), len(list2)
	var l1, l2 int // list1和list2当前遍历到的下标
	var r1, r2 int // list1和list2当前遍历到的右边界
	res := 0
	for l1 < len1 && l2 < len2 {
		r1 += list1[l1][1]
		r2 += list2[l2][1]
		list1[l1][1] = 0 // 置0，避免后续重复增加r1
		list2[l2][1] = 0 // 置0，避免后续重复增加r2
		now1 := list1[l1][0]
		now2 := list2[l2][0]
		for index < r1 && index < r2 {
			if now1 != now2 {
				res++
			}
			index++
		}

		if index >= r1 {
			l1++
		}
		if index >= r2 {
			l2++
		}
	}
	return fmt.Sprintln(res, "/", total)
}

func getList(str string, list *[][2]int) int {
	var index, cnt int
	for index < len(str) {
		num := 0
		for index < len(str) && isDigit(str[index]) {
			num *= 10
			num += int(str[index] - '0')
			index++
		}

		c := str[index]
		*list = append(*list, [2]int{int(c), num})
		cnt += num
		index++
	}
	return cnt
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
