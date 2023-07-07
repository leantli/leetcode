package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 20. 表示数值的字符串

// 有不少判断
// 我感觉用正则比较好
// 其次就是自己手动写各种条件判断 或 状态机写法
// 1. 预处理前后空格
// 2. 判断是否存在非法字符
// 3. 判断是小数还是整数
// 4. 小数判断处理
// 5. e/E后处理

// 正则少量语法
// \s 匹配所有空白符，包括换行 \S 匹配除换行外的所有空白符
// * 表示前面的子表达时出现 0 次或多次 -> zo* 能匹配 "z" 以及 "zoo"。* 等价于 {0,}
// + 表示前面的子表达时出现 1 次或多次 -> zo* 能匹配 "z" 以及 "zoo"。* 等价于 {0,}
// ? 表示前面的子表达式出现 0 次或 1 次 -> do(es)? 可以匹配 "do" 、 "does"、 "doxy" 中的 "do" 。? 等价于 {0,1}
// [..] 表示匹配中括号中的所有字符，[^..] 表示除了中括号中的字符，其他字符都匹配
// \d 匹配数字，等价于[0-9] \D 匹配非数字，等价于[^0-9]
// | 表示左右两个表达式之间的一个选择

// 先 ^ 和 $ 标识一下开头结尾
// 再基于\s把空格等匹配掉 [\s]*
// 再把开头可能存在的+-匹配掉 [+-]?
// 把e后这部分匹配一下 ([eE][+-]?\d+) eE后一定要跟着个整数，因此是 \d+
// 最难的中间的，同时满足整数和小数，实在不行就 | 并上一堆子表达式
// (()|()|()|())
// 当然，整合各种匹配条件后，两个子表达式就够了
// (\d*\.\d+) 匹配小数的情况，包括 1.1, .1 等情况
// (\d+\.?\d*) 匹配整数和小数的部分情况，包括 111, 1., 1.1 等情况
// var matchRE = regexp.MustCompile(`^[\s]*[+-]?((\d*\.\d+)|(\d+\.?\d*))([eE][+-]?\d+)?[\s]*$`)

// func isNumber(s string) bool {
// 	return matchRE.MatchString(s)
// }

func main() {
	fmt.Println(isNumber("  +100e+5.6    "))
	fmt.Println(isNumber("  +100e+5    "))
	fmt.Println(isNumber("  +100E+5.6    "))
	fmt.Println(isNumber("  -100e-5    "))
	fmt.Println(isNumber("  +-100e-5    "))
	fmt.Println(isNumber("  +.5    "))
	fmt.Println(isNumber("  +5.5.5    "))
}

// 常规
// 空格 符号 小数/整数(小数中 '.' 点前可以有符号和数字，点后至少有一位数字) e/E 符号 整数 空格
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	var dotFlag, expFlag, numFlag bool // 是否遇到过小数点、E/e、数字
	for i := range s {
		switch {
		case isDigit(s[i]):
			numFlag = true
			// 如果是 + -，要么符号是在首位，要么前一位是 e/E
		case (s[i] == '+' || s[i] == '-') && (i == 0 || s[i-1] == 'E' || s[i-1] == 'e'):
			continue
			// 遇到 e/E，前面要遇到过数字，没出现过 e/E；后面也要有数字，因此将 numFlag 重新置为 false
		case (s[i] == 'E' || s[i] == 'e') && (!expFlag && numFlag):
			expFlag = true
			numFlag = false
			// 遇到 . 前面不能出现过 exp，也不能重复出现小数点
		case (s[i] == '.') && (!expFlag && !dotFlag):
			dotFlag = true
		default:
			return false
		}
	}
	return numFlag
}

func isDigit(b byte) bool {
	switch b {
	case '0':
		return true
	case '1':
		return true
	case '2':
		return true
	case '3':
		return true
	case '4':
		return true
	case '5':
		return true
	case '6':
		return true
	case '7':
		return true
	case '8':
		return true
	case '9':
		return true
	}
	return false
}
