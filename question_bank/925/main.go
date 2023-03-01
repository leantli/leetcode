package main

// https://leetcode.cn/problems/long-pressed-name/
// 925. 长按键入

// 两个指针分别指向两个字符串首部，分别是 i 和 j
// 当 i j 指向的字母相等时，直接右移
// 不相等时，判断 j 指向的字母是否与 i-1 相等，是的话，则右移 j，但 i 不右移
func isLongPressedName(name string, typed string) bool {
	n, m := len(name), len(typed)
	// 当输入长度少于名字长度直接返回 false
	if m < n {
		return false
	}
	var i, j int
	for j < m {
		// 相等时直接双指针右移
		if i < n && typed[j] == name[i] {
			j++
			i++
			continue
		}
		// 不相等时判断 j 是否和 i-1 相等，是的话 j 右移
		if i > 0 && typed[j] == name[i-1] {
			j++
			continue
		}
		// 两次相等判断都没过，那指定是错了
		return false
	}
	// 出循环后 j == m, 再看看 i 是否 ==n
	return i == n
}
