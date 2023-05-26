package main

// https://leetcode.cn/problems/longest-palindromic-substring/?favorite=2cktkvj
// 5. 最长回文子串

// 但事实上，这个dp感觉并没有什么优化，也还是算一种暴力？
// 只不过最辣鸡的暴力是枚举子串，判断子串是否为回文串
// 而这个dp本质上也是一种暴力，我们也可以枚举中心，然后从中心向左右扩展
// 记录最长的即可
func longestPalindrome(s string) string {
	var maxL, maxR int
	expand := func(l, r int) {
		for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
		}
		if (r-1)-(l+1) > maxR-maxL {
			maxL = l + 1
			maxR = r - 1
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[maxL : maxR+1]
}

// // 其实我们可以注意到，怎么判断一个字符串是回文：
// // 1. 单个字符
// // 2. 两个相同的字符
// // 3. 长度>2的字符串，首位相同，那么其是否为回文，取决于去掉首尾后的字符串是否是回文
// // 4. 首尾不同的直接不为回文串
// // 抛去前两种case，第三种case是典型的可以使用动态规划解题的特征
// // 因此我们先初始化单个字符成回文串的dp情况
// // 在从字符串长度为2开始主键截取字符串去判断是否为回文串
// func longestPalindrome(s string) string {
// 	dp := make([][]bool, len(s))
// 	// 每个字符肯定都是单独成回文子串的
// 	for i := 0; i < len(s); i++ {
// 		dp[i] = make([]bool, len(s))
// 		dp[i][i] = true
// 	}
// 	var maxL, maxR int
// 	// 从长度为2开始枚举子串长度，因为我们的dp推导显然需要从长度短的情况向上推
// 	for length := 2; length <= len(s); length++ {
// 		for l := 0; l < len(s)-length+1; l++ {
// 			r := l + length - 1
// 			// 对三种case的判断
// 			if s[l] != s[r] {
// 				continue
// 			} else if length < 3 {
// 				dp[l][r] = true
// 			} else {
// 				dp[l][r] = dp[l+1][r-1]
// 			}
// 			if dp[l][r] && r-l > maxR-maxL {
// 				maxL = l
// 				maxR = r
// 			}
// 		}
// 	}
// 	return s[maxL : maxR+1]
// }

// // 暴力是怎么样呢？遍历每一个子串，判断该子串是否为回文串
// // 时间复杂度为O(n*3)，l和r的两轮循环，每一个子串的判断又是一轮
// func longestPalindrome(s string) string {
// 	maxL, maxR := 0, 1
// 	for l := 0; l < len(s); l++ {
// 		for r := l + 1; r <= len(s); r++ {
// 			// fmt.Printf("s:%s, rs:%s\n", s[l:r], rs[l:r])
// 			if isPalindrome(s[l:r]) {
// 				if r-l > maxR-maxL {
// 					maxL = l
// 					maxR = r
// 				}
// 			}
// 		}
// 	}
// 	return s[maxL:maxR]
// }

// // 判断是否为回文串
// func isPalindrome(s string) bool {
// 	bs := []byte(s)
// 	for l, r := 0, len(bs)-1; l < r; {
// 		if bs[l] != bs[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }

// func main() {
// 	tt := "abcdefg"
// 	ttt := "abcdefgh"
// 	t := "abccba"
// 	fmt.Println(isPalindrome(tt))
// 	fmt.Println(isPalindrome(ttt))
// 	fmt.Println(isPalindrome(t))
// }
