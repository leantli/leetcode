package main

import (
	"fmt"
	"unsafe"
)

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 58-2 左旋转字符

// 空间复杂度O(1) 旋转类型题目要记住这个操作，常见
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 三次反转，空间复杂度为 O(1)
	// 反转前 n 个数字，注意下标 n-1
	reverseLToR(0, n-1, b)
	reverseLToR(n, len(b)-1, b)
	reverseLToR(0, len(b)-1, b)
	return string(b)
}

func reverseLToR(l, r int, b []byte) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}

// 马上想到，但还最好再考虑原地修改
// func reverseLeftWords(s string, n int) string {
// 	return s[n:] + s[:n]
// }

func main() {
	fmt.Println(reverseLeftWords("1234567", 3))
	fmt.Println(string(String2Bytes("1234")))
}

// 突然又想到 string 和 []byte 无拷贝互转
func String2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
