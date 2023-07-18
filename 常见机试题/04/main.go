package main

import (
	"fmt"
	"strings"
)

/**
题目描述:
[敏感字段加密]给定一个由多个命令字组成的命令字符串:
1、字符串长度小于等于127字节，只包含大小写字母，数字，下划线和偶数个双引号:
2、命令字之间以一个或多个下划线 进行分割;
3、可以通过两个双引号""来标识包含下划线 的命令字或空命令字(仅包含两个双引号的命令字)，双引号不会在命令字内部出现:
请对指定索引的敏感字段进行加密，替换为******(6个*)，并删除命令字前后多余的下划线 。如果无法找到指定索引的命令字，输出字符串QERROR。
输入描述:
输入为两行，第一行为命令字索引K (从0开始)，第二行为命令字符串S.
输出描述:
输出处理后的命令字符串，如果无法找到指定索引的命令字，输出字符串ERROR
示例1
输入
password_a12345678_timeout_100
输出
password_******_timeout_100
示例2
输入
aaa_password_"a12_45678"_timeout_100_""_
输出
aaa_password_******_timeout_100
**/

func main() {
	fmt.Println(work("password_a12345678_timeout_100", 1))
	fmt.Println(work(`aaa_password_"a12_45678"_timeout_100_""_`, 2))
}

func work(s string, idx int) string {
	s = strings.Trim(s, "_")
	n := len(s)
	var l, r int
	strs := make([]string, 0)
	for r < n {
		for r < n && s[r] == '_' {
			r++
		}
		l = r
		for r < n && s[r] != '_' {
			r++
		}
		if s[l] == '"' && s[r-1] != '"' {
			r++
			for r < n && s[r] != '_' {
				r++
			}
		}
		if l == r && l == n {
			break
		}
		strs = append(strs, s[l:r])
	}
	if len(strs)-1 < idx {
		return "ERROR"
	}
	strs[idx] = "******"
	return strings.Join(strs, "_")
}
