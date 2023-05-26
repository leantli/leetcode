package main

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// 1456. 定长子串中元音的最大数目

// 定长 子串 元音 最大数目
// 定长区间滑动窗口
func maxVowels(s string, k int) int {
	// 初始化元音字母的 map，便于查验
	vowels := make([]int, 128)
	vowels['a'], vowels['e'], vowels['i'], vowels['o'], vowels['u'] = 1, 1, 1, 1, 1
	bs := []byte(s)
	var count, maxCount int
	for _, c := range bs[:k] {
		count += vowels[c]
	}
	maxCount = count
	for i := k; i < len(bs); i++ {
		count += vowels[bs[i]] - vowels[bs[i-k]]
		maxCount = max(maxCount, count)
	}
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
