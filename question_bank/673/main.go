package main

// https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
// 673. 最长递增子序列的个数

// // 二刷
// // 由于做过T300，显然，我们能够基于常规dp，计算最长长度时，有多少个
// // 不清楚的先补做一下T300，更方便理解
// // dp[i] 为 nums[i] 结尾的子序列的最长长度
// // 因此，我们先正常写出LIS的常规dp，然后再考虑怎么获取LIS的个数
// // 知道怎么求LIS后，我们注意，求LIS个数时
// // 我们不能每求完一个dp[i]后，再去判断是否长度为最长，然后count++
// // 这样的话会忽略掉同一个dp[i]下，不同的子序列(而这些子序列的长度刚好都达到了最长的情况)
// // 这里假设我们知道dp[i]的长度是最长的，那我们怎么求其LIS的个数？
// // 遍历dp[j]时，当dp[j]+1时等同于最长时，我们可以计数
// // 但是计数多少呢？我们显然不能仅仅是count+1，因为组成dp[j]的LIS很可能也存在多种情况
// // 显然，这里构成了重复子问题，我们需要额外一个数组记录 以nums[i]结尾的dp[i]最大长度下，LIS子序列的个数 -- cnt数组
// // dp[i] = max{dp[i], dp[j](i前的所有nums[j])}
// // cnt[i] = 当dp[j]+1==dp[i]时，cnt[i]+=cnt[j]；；；当dp[j]+1>dp[i]时，重置dp[i]以及cnt[i]---dp[i]=dp[j]+1,cnt[i]=cnt[j]
// // count += cnt[i],当dp[i]==maxLen时；；；dp[i]>maxLen时则直接重置count=cnt[i]
// func findNumberOfLIS(nums []int) int {
// 	dp := make([]int, len(nums))
// 	cnt := make([]int, len(nums))
// 	for i := range nums {
// 		dp[i] = 1
// 		cnt[i] = 1
// 	}
// 	count := 1
// 	maxLen := 1
// 	for i := 1; i < len(nums); i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				if dp[j]+1 > dp[i] {
// 					dp[i] = dp[j] + 1
// 					cnt[i] = cnt[j]
// 				} else if dp[j]+1 == dp[i] {
// 					cnt[i] += cnt[j]
// 				}
// 			}
// 		}
// 		if dp[i] == maxLen {
// 			count += cnt[i]
// 		} else if dp[i] > maxLen {
// 			maxLen = dp[i]
// 			count = cnt[i]
// 		}
// 	}
// 	return count
// }

// 但显然，300的贪心+二分，在这里也是可以使用的
// 之前tail[i]表示i+1长度结尾的最小数，事实上，每次更新时，都代表着更新的数能构成长度为i+1的子序列
// 那么我们这次用一个二维数组,tail[i]表示i+1长度结尾的数组
// 这个时候再用cnt[i][j]表示，长度i+1情况下，j下标指向的数结尾的子序列，有多少个数
// 而显然 cnt[i][j] = cnt[i-1][k]累加，k<j(这里特指k,j下标指向的数字，而非下标)
// 有思路，但是实现上有些问题，题解参考 https://leetcode.cn/problems/number-of-longest-increasing-subsequence/solution/yi-bu-yi-bu-tui-dao-chu-zui-you-jie-fa-2-zui-chang/
func findNumberOfLIS(nums []int) int {
	tail, cnt := [][]int{}, [][]int{}
	for _, v := range nums {
		// 1. 找到 新进数 要插入的 tail 的位置
		// i 指 tail 数组中的 i，表示 i+1 长度结尾的数字们
		i := findInsertIndex(tail, v)
		// 2. 计算新进数的子序列情况
		count := 1 // 默认count至少为1，即自身
		// 当长度大于 1 时，需要根据tail[i-1]的cnt情况计算新入数的cnt
		if i > 0 {
			// 找到 新进数 不能衔接的部分的最后一个数，tail[i-1] 数组是一个非递增数组，新进数肯定能衔接 tail[i-1] 位置最后的数
			// 但是前面的数可不一定，因此我们需要找到不能衔接的最后一个数，方便计算前缀和，这个前缀和表示的是对应 cnt[i-1][k+1~~~len-1] 这部分(能衔接的)
			k := findPreNumIndex(tail[i-1], v)
			// k 为 -1 时，说明 tail[i-1] 中的所有数，都比新进数小，因此 count 为上一长度的总前缀和
			if k == -1 {
				count = cnt[i-1][len(cnt[i-1])-1]
			} else {
				count = cnt[i-1][len(cnt[i-1])-1] - cnt[i-1][k]
			}
		}
		// 3. 准备插入，当插入的位置为末尾时，需要新建一维数组，因此特殊处理
		if i == len(tail) {
			// 此时说明 v 是新出现的最大值，将其添加到 tail 尾部，新建tail和cnt的第二层数组
			tail = append(tail, []int{v})
			cnt = append(cnt, []int{count})
		} else {
			// 否则就在对应的位置插入 v，并且也插入对应的前缀和
			tail[i] = append(tail[i], v)
			cnt[i] = append(cnt[i], cnt[i][len(cnt[i])-1]+count)
		}
	}
	c := cnt[len(cnt)-1]
	return c[len(c)-1]
}

// 找到 target 将要插入的 tail 的位置
func findInsertIndex(tail [][]int, target int) int {
	l, r := -1, len(tail)
	for l+1 != r {
		mid := l + (r-l)/2
		if tail[mid][len(tail[mid])-1] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

// 找到最后一个大于等于 target 的位置(l)，用于排除这部分之前的前缀和
func findPreNumIndex(tail []int, target int) int {
	l, r := -1, len(tail)
	for l+1 != r {
		mid := l + (r-l)/2
		if tail[mid] >= target {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// // 常规 dp 的基础上，再额外加上计数
// // 值得注意的是 计数 防重防漏
// // 可能会漏掉的情况，同一个 nums[i] 结尾有多种情况形成最长子序列，然而只计算一次
// // 这里计数也要基于 dp 去计数，cnt[i] 表示以 nums[i] 为结尾的最长子序列的个数
// func findNumberOfLIS(nums []int) int {
// 	n := len(nums)
// 	// dp[i] 表示，以 nums[i] 为结尾的子序列最长的长度
// 	// cnt[i] 表示，以 nums[i] 为结尾的最长子序列的个数
// 	// 初始化，默认至少为 1
// 	dp := make([]int, n)
// 	cnt := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		dp[i] = 1
// 		cnt[i] = 1
// 	}
// 	// 初始化，dp[0] 最大长度为 1，此时最长子序列的个数为 1
// 	maxLen, count := 1, 1
// 	for i := 1; i < n; i++ {
// 		for j := 0; j < i; j++ {
// 			// i 前面任意一个 j 元素比 i 小，则可以考虑在 以j为末尾的子序列上补个 i，此时判断长度是否需要变更
// 			if nums[j] < nums[i] {
// 				if dp[j]+1 > dp[i] {
// 					dp[i] = dp[j] + 1
// 					cnt[i] = cnt[j] // 当前以 nums[i] 为结尾的最长子序列的个数的重置计数
// 				} else if dp[j]+1 == dp[i] {
// 					cnt[i] += cnt[j]
// 				}
// 			}
// 		}
// 		if dp[i] == maxLen {
// 			count += cnt[i]
// 		} else if dp[i] > maxLen {
// 			maxLen = dp[i]
// 			count = cnt[i] // 总体计数的重置计数
// 		}
// 	}
// 	return count
// }
