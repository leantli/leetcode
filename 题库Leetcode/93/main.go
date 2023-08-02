package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/restore-ip-addresses/
// 93. 复原 IP 地址

// 枚举分割数字成4个部分，判断是否为有效 IP，是则加入ip段slice中，当组成四个ip端则加入最终结果并 return
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	cur := make([]string, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		// 当 cur 长度为 4，并且 idx 其实已经到 len(s) 最后了，才能考虑加入最终结果
		if len(cur) == 4 && idx == len(s) {
			res = append(res, strings.Join(cur, "."))
			return
		}
		for i := idx; i < len(s); i++ {
			// 剪枝，存在前导0的数排除，直接跳过，只有 0 能尝试加入
			if s[idx] == '0' && idx != i {
				break
			}
			str := s[idx : i+1]
			num, err := strconv.Atoi(str)
			if err != nil {
				break
			}
			if num >= 0 && num <= 255 {
				cur = append(cur, str)
				dfs(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}
