package main

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合

var mapper = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	var dfs func(idx int, cur string)
	dfs = func(idx int, cur string) {
		if idx == len(digits) {
			res = append(res, cur)
			return
		}
		for _, letter := range mapper[string(digits[idx])] {
			dfs(idx+1, cur+letter)
		}
	}
	dfs(0, "")
	return res
}
