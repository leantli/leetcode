package main

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
// 438. 找到字符串中所有字母异位词

// 定长滑动窗口？窗口的长度固定为 p 字符串的长度
// 窗口内满足条件即可将窗口的左边界 l 加入结果数组中
// 如何判断 窗口中字符串是 p 的异位词？
// 只要窗口内的字符以及字符出现的次数相同，就说明满足条件
// 这里显然需要两个 map，分别记录 窗口内存在的字符:该字符出现的次数 和 p 中存在的字符:该字符出现的次数
// 并且这里还有一个小技巧，纯粹比较两个 map 其实效率还是会低一些
// 可以直接记录窗口中，目前有多少满足条件的字符，以及 完全满足窗口为p异位词需要多少个字符
func findAnagrams(s string, p string) []int {
	n, m := len(s), len(p)
	if m > n {
		return []int{}
	}
	// 先记录 p 中有多少个字符，以及各个字符出现的次数
	pMap := make([]int, 128)
	var need int
	for _, c := range []byte(p) {
		if pMap[c] == 0 {
			need++
		}
		pMap[c]++
	}
	// 初始化窗口
	var satisfy int
	bs := []byte(s)
	windowsMap := make([]int, 128)
	res := make([]int, 0)
	for _, c := range bs[:m] {
		windowsMap[c]++
		// 如果该字符属于满足 窗口子串 为 p 异位词的字符(即该字符是 p 中出现的字符)
		// 则判断该字符在窗口内出现的次数是否达标
		if pMap[c] != 0 && pMap[c] == windowsMap[c] {
			satisfy++
		}
	}
	if satisfy == need {
		res = append(res, 0)
	}
	for r := m; r < n; r++ {
		l := r - m
		windowsMap[bs[l]]--
		// 这里要注意：本来是满足的，该字符被舍弃后，数量不达标，satisfy 再减 1
		if pMap[bs[l]] != 0 && windowsMap[bs[l]]+1 == pMap[bs[l]] {
			satisfy--
		}
		windowsMap[bs[r]]++
		if pMap[bs[r]] != 0 && pMap[bs[r]] == windowsMap[bs[r]] {
			satisfy++
		}
		if satisfy == need {
			res = append(res, l+1)
		}
	}
	return res
}
