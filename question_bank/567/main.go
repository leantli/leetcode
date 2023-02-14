package main

// https://leetcode.cn/problems/permutation-in-string/
// 567. 字符串的排列

// 主要是判断 s 的子串中，是否要包含 p 的异位词
// 这里可以活用 go 的一些特性，我们不用 slice，而是直接用数组，这样数组直接可以比较
// 不过这种数组之间的比较相当于是 全量地判断两个字符串是否为异位词，效率比原先的思路低
// 因此我们在这种数组思路上用上原先的考虑，就是考虑窗口内的 有效字符情况(指满足p异位词的字符情况)
func checkInclusion(p string, s string) bool {
	n, m := len(p), len(s)
	if n > m {
		return false
	}
	var windowsCount, pCount [26]int
	bs := []byte(s)
	// 定长滑窗初始化，定长为 p 的长度，我们只需要检验 s 中长度为 n 的子串即可
	// 同时记录 p 字符串中出现字符及其出现次数
	// 还要顺带记录 窗口内和p有几个字符(同个字符只计一次)不同
	var diff int
	for i, c := range []byte(p) {
		windowsCount[bs[i]-'a']++
		tempChar := c - 'a'
		if pCount[tempChar] == 0 {
			diff++
		}
		pCount[tempChar]++
	}
	// 初始化 diff 的值，判断当前窗口还需要满足几个字符才能成为 p 的异位词
	for i, pcnt := range pCount {
		if pcnt == 0 {
			continue
		}
		if pcnt <= windowsCount[i] {
			diff--
		}
	}
	if diff == 0 {
		return true
	}
	for r := n; r < m; r++ {
		lChar, rChar := bs[r-n]-'a', bs[r]-'a'
		windowsCount[lChar]--
		if pCount[lChar] != 0 && windowsCount[lChar]+1 == pCount[lChar] {
			diff++
		}
		windowsCount[rChar]++
		if pCount[rChar] != 0 && windowsCount[rChar] == pCount[rChar] {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

// // 主要是判断 s 的子串中，是否要包含 p 的异位词
// // 这里可以活用 go 的一些特性，我们不用 slice，而是直接用数组，这样数组直接可以比较
// // 不过这种数组之间的比较相当于是 全量地判断两个字符串是否为异位词，效率比原先的思路低
// func checkInclusion(p string, s string) bool {
// 	n, m := len(p), len(s)
// 	if n > m {
// 		return false
// 	}
// 	var windowsCount, pCount [26]int
// 	bs := []byte(s)
// 	// 定长滑窗初始化，定长为 p 的长度，我们只需要检验 s 中长度为 n 的子串即可
// 	// 同时记录 p 字符串中出现字符及其出现次数
// 	for i, c := range []byte(p) {
// 		windowsCount[bs[i]-'a']++
// 		pCount[c-'a']++
// 	}
// 	if windowsCount == pCount {
// 		return true
// 	}
// 	for r := n; r < m; r++ {
// 		windowsCount[bs[r-n]-'a']--
// 		windowsCount[bs[r]-'a']++
// 		if windowsCount == pCount {
// 			return true
// 		}
// 	}
// 	return false
// }

// // 这道题和 438 找异位词差不多，定长滑动窗口
// // 主要优化点在于子串的判断上，用 满足的字符数量 来判断会更高效(相比每次都比较两个字符串是否为异位词)
// func checkInclusion(p string, s string) bool {
// 	n, m := len(s), len(p)
// 	if m > n {
// 		return false
// 	}
// 	// 先记录 p 中有多少个字符，以及各个字符出现的次数
// 	pMap := make([]int, 128)
// 	var need int
// 	for _, c := range []byte(p) {
// 		if pMap[c] == 0 {
// 			need++
// 		}
// 		pMap[c]++
// 	}
// 	// 初始化窗口
// 	var satisfy int
// 	bs := []byte(s)
// 	windowsMap := make([]int, 128)
// 	for _, c := range bs[:m] {
// 		windowsMap[c]++
// 		// 如果该字符属于满足 窗口子串 为 p 异位词的字符(即该字符是 p 中出现的字符)
// 		// 则判断该字符在窗口内出现的次数是否达标
// 		if pMap[c] != 0 && pMap[c] == windowsMap[c] {
// 			satisfy++
// 		}
// 	}
// 	if satisfy == need {
// 		return true
// 	}
// 	for r := m; r < n; r++ {
// 		l := r - m
// 		windowsMap[bs[l]]--
// 		// 这里要注意：本来是满足的，该字符被舍弃后，数量不达标，satisfy 再减 1
// 		if pMap[bs[l]] != 0 && windowsMap[bs[l]]+1 == pMap[bs[l]] {
// 			satisfy--
// 		}
// 		windowsMap[bs[r]]++
// 		if pMap[bs[r]] != 0 && pMap[bs[r]] == windowsMap[bs[r]] {
// 			satisfy++
// 		}
// 		if satisfy == need {
// 			return true
// 		}
// 	}
// 	return false
// }
