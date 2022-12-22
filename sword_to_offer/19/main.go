package main

// https://leetcode.cn/problems/zheng-ze-biao-da-shi-pi-pei-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 19. 正则表达式匹配

// 看到字符串匹配，马上想到双指针 and dp，先考虑考虑 dp
// (DP) 是解决此类字符串匹配问题的通用方法（关键！牢记！）
// 1. 状态定义
// dp[i][j] 表示 s 的前 i 个字符和 p 的前 j 个字符能否匹配
// 2. 状态转义
//   1) 当 p[j] 是正常字母时:
//       dp[i][j] = dp[i-1][j-1] (if s[i] == p[j]) or dp[i][j] = false (if s[i] != p[j])
//   2) 当 p[j] == '.' 时:
//       p[j] 必定能和 s[i] 匹配成功，因此 dp[i][j] = dp[i-1][j-1]
//
//   上面的匹配和状态转移方程都很简单也容易理解，主要是下面的，需要一定的推导
//
//   3) 当 p[j] == '*' 时： ( 最关键的
//       3.1) 当 s[i] != p[j-1] 时，则可以不虚考虑 p[j] + p[j-1]，直接回到 p[j-2] --> dp[i][j] = dp[i][j-2]，
//       不匹配的情况还是比较好考虑的，但是匹配的状态转移方程就不好推导了
//       3.2) 当 s[i] == p[j-1] 时，dp[i][j] = dp[i-1][j-2]
//            当 s[i-1] == p[j-1] 时, dp[i][j] = dp[i-2][j-2]
//            .... 当 s[i-k] == p[j-1] 时，dp[i][j] = dp[i-k-1][j-2]
//       此时可以得到
//       dp[i][j] = dp[i][j-2] || ({dp[i-1][j-2]} && s[i]==p[j-1]) {||
//                  (dp[i-2][j-2] & s[i-1]==p[j-1]) || .... || (dp[i-k-1][j-2] & s[i-k][j-1])}
//       但是这个匹配的方程式情况太多了，我们不可能一个个都去或操作吧？
//       这里真的难想到。。。。
//       这里我们要考虑到如下
//       dp[i-1][j] = { dp[i-1][j-2] || (dp[i-2][j-2] && s[i-1] == p[j-1]) ||
//                    .... || (dp[i-k-1][j-2] && s[i-k] == p[j-1]) }
//       可以观察到 dp[i-1][j] 是基本可以等同于 dp[i][j] 后面一系列的或操作的(看上面方程式中的大括号{}内容)
//       二者之间的差距，仅在于 s[i] == p[j-1]，因此将 dp[i][j] = 一系列或操作简化为如下方程:
//       dp[i][j] = dp[i][j-2] || ( dp[i-1][j] && s[i] == p[j-1] )
//
// 此时得到全部的状态转移方程
//              |--- dp[i-1][j-1]              , s[i]==p[j]
//              |
//  dp[i][j] ---|--- dp[i][j-2]                , p[j] == '*' && s[i]!=p[j-1]
//              |
//              |--- dp[i-1][j] || dp[i][j-2]  , p[j] == '*' && s[i]==p[j-1]
//  以上简单忽略了 p[i]='.' 的处理，因为必定能匹配上，自己处理一下即可

func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	// 初始化
	dp := make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}
	dp[0][0] = true
	// 这里要注意，由于下面状态转移时，i=0会导致数组越界，因此我们在这里做一个边界处理
	// 只需要考虑 p[j-1] == '*' --> dp[0][j] = dp[0][j-2] 即可，因为 此时 i=0, p 无论为啥都不能匹配上
	for j := 1; j <= pLen; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// 状态转移
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if match(s[i-1], p[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if match(s[i-1], p[j-2]) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[sLen][pLen]
}

// 判断是否匹配
func match(s, p byte) bool {
	if p == '.' {
		return true
	}
	return s == p
}
