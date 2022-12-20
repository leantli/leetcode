package main

import (
	"fmt"
	"regexp"
)

// https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 20. 表示数值的字符串

// 有不少判断
// 我感觉用正则比较好
// 其次就是自己手动写各种条件判断
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
var matchRE = regexp.MustCompile(`^[\s]*[+-]?((\d*\.\d+)|(\d+\.?\d*))([eE][+-]?\d+)?[\s]*$`)

func isNumber(s string) bool {
	return matchRE.MatchString(s)
}

func main() {
	fmt.Println(isNumber("  +100e+5.6    "))
	fmt.Println(isNumber("  +100e+5    "))
	fmt.Println(isNumber("  +100E+5.6    "))
	fmt.Println(isNumber("  -100e-5    "))
	fmt.Println(isNumber("  +-100e-5    "))
	fmt.Println(isNumber("  +.5    "))
	fmt.Println(isNumber("  +5.5.5    "))
}
