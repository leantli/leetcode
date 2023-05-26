package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 替换空格

const replaceToText = "%20"

func replaceSpace(s string) string {
	var sb strings.Builder
	for _, v := range []byte(s) {
		if v == ' ' {
			sb.WriteString(replaceToText)
			continue
		}
		sb.WriteByte(v)
	}
	return sb.String()
}

func main() {
	fmt.Println(replaceSpace(" We are happy ."))
}
