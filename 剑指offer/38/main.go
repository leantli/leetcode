package main

// https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 38. 字符串的排列

// 填坑递归 + set
func permutation(s string) []string {
	set := make(map[string]struct{})
	bs := []byte(s)
	var dfs func(s []byte, index int, cur []byte)
	dfs = func(s []byte, index int, cur []byte) {
		// 填完坑就直接加进 set
		if index == len(s) {
			set[string(cur)] = struct{}{}
			return
		}
		for i := 0; i < len(bs); i++ {
			if bs[i] != '-' {
				temp := bs[i]
				bs[i] = '-'
				dfs(s, index+1, append(cur, temp))
				bs[i] = temp
			}
		}
	}
	dfs(bs, 0, []byte(""))
	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}
