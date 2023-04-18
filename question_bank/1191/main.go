package main

// https://leetcode.cn/problems/k-concatenation-maximum-sum/
// 1191. K 次串联后最大子数组之和

// 另一思路---这个太巧了，我感觉我是想不到的，写的也很巧
// 最大子数组存在于单个数组内，此时最大和即为单个数组内的最大子数组和
// 最大子数组横跨两个数组，此时最大和即为单个数组内的最大前缀和加上最大后缀和
// 最大子数组横跨多个数组，此时要求单个数组总和大于0，此时最大和即为k-2个数组和加上最大前缀和加上最大后缀和
func kConcatenationMaxSum(arr []int, k int) int {
	var s, mxPre, miPre, mxSub int
	for _, x := range arr {
		s += x
		mxPre = max(mxPre, s)       // 最大前缀和
		miPre = min(miPre, s)       // 最小前缀和
		mxSub = max(mxSub, s-miPre) // 最大子数组和
	}
	const mod = 1e9 + 7
	ans := mxSub
	if k == 1 {
		return ans % mod
	}
	// 求最大后缀和
	mxSuf := s - miPre
	// 判断最大子数组和是之前求出的还是最大前后缀和相加
	ans = max(ans, mxSuf+mxPre)
	if s > 0 {
		ans = max(ans, mxSuf+(k-2)*s+mxPre)
	}
	return ans % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // 环形数组失败的原因主要在于同一数组中，同个数不会被重复取
// // 而实际k次重复后的数组，显然是可以重复取数的
// // 那么这次，我们可以直接重复一次数组，求最大子数组和
// // 这里给出两个例子 [30 -30 55 66 -121 -10] k=3，此时我们会发现，最大子数组和为 121，并且其无法与后面的数组连接起来
// // [300 -301 55 66 -122 10] k=3，此时我们会找到最大子数组和为 [10 300 -301 .... 10 300]
// // 我们发现，此时最大子数组是跨越单个原数组的，如果 k=2，我们会发现找到的最大子数组和是 [10 300]
// // 但是此时 k=3，那么能不能取到下一个跨界的 [10 300]呢？三个数组中，存在两个[10 300]
// // 此时我们发现，如果要到达下一个最大子数组，中间需要跨越 原数组除了[10 300]的其他数
// // [300 -301 55 66 -122 10 300 -301 55 66 -122 10 300 -301 55 66 -122 10]
// // 即 10 300 到下一个 10 300，需要跨越 -301 55 66 -122，此时我们发现，这四个要跨越的数，加上最大子数组[10 300]
// // 其实就是原数组，并且这个原数组本身和大于10，因此才可以继续正常衔接，否则会像例 1 ，无法衔接起来，只能取121，无论后面有多长
// // 因此我们可以得到规律，基于 k=2 的数组求出最大子数组和，接着根据原数组和，判断是否大于0
// // 若小于等于0，则最终结果就是已取到的最大子数组和
// // 若大于0，则说明该k=2的最大子数组，后面还能继续衔接下一个最大子数组，衔接后增长的值为原数组和
// // 即为 k=2时最大子数组和+(k-2)*原数组和
// func kConcatenationMaxSum(arr []int, k int) int {
// 	n, mod := len(arr), int(1e9+7)
// 	var dp, biggest, sum int
// 	for i := 0; i < min(2, k)*n; i++ {
// 		idx := i % n
// 		dp = max(dp+arr[idx], arr[idx])
// 		biggest = max(dp, biggest)
// 		if i < n {
// 			sum += arr[i]
// 		}
// 	}
// 	if sum > 0 && k > 2 {
// 		// // 这里因为go的int类型足够大，因此可以直接相乘后取模
// 		// k -= 2
// 		// for k > 0 {
// 		// 	biggest = (biggest + sum) % mod
// 		// 	k--
// 		// }
// 		biggest = (biggest + sum*(k-2)) % mod
// 	}
// 	return max(0, biggest) % mod
// }

// func max(arr ...int) int {
// 	res := arr[0]
// 	for _, num := range arr {
// 		if res < num {
// 			res = num
// 		}
// 	}
// 	return res
// }

// func min(arr ...int) int {
// 	res := arr[0]
// 	for _, num := range arr {
// 		if res > num {
// 			res = num
// 		}
// 	}
// 	return res
// }

// // 最简单的思路应该是先复现重复k次的数组，再求最大子数组和
// // 但是很可能会有测试用例超内存，一次你我们得考虑，是否有其他可利用的点
// // 除非最大子数组和是整个数组，否则重复k次只需要考虑首尾相连的情况即可
// // 因此这里我们可以直接对整个数组求总和，如果最大子数组和相等，则可以直接*k%mod
// // 否则就是正常取环形数组的最大子数组和
// // 并且，注意这道题的子数组长度可以是0，那么最大子数组和最小都是0，无需考虑负数
// // 写完部分测试用例过了，但本质上是不行的，因为参考环形数组求最大子数组和时
// // 假设有该数组重复一次,分别是a和b，环形数组中，a中取过的数，b中无法再取
// // 因此该思路错误
// func kConcatenationMaxSum(arr []int, k int) int {
// 	biggest, smallest, sum := arr[0], arr[0], arr[0]
// 	big, small := arr[0], arr[0]
// 	mod := int(1e9 + 7)
// 	for i := 1; i < len(arr); i++ {
// 		sum += arr[i]
// 		big = max(big+arr[i], arr[i])
// 		small = min(small+arr[i], arr[i])
// 		biggest = max(biggest, big)
// 		smallest = min(smallest, small)
// 	}
// 	if biggest < 0 {
// 		return 0
// 	}
// 	if biggest == sum {
// 		return biggest * k % mod
// 	}
// 	return max(biggest, sum-smallest, 0)
//
// }

// func max(arr ...int) int {
// 	res := arr[0]
// 	for _, num := range arr {
// 		if res < num {
// 			res = num
// 		}
// 	}
// 	return res
// }

// func min(arr ...int) int {
// 	res := arr[0]
// 	for _, num := range arr {
// 		if res > num {
// 			res = num
// 		}
// 	}
// 	return res
// }
