package main

import (
	"sort"
)

// https://leetcode.cn/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1498. 满足条件的子序列数目

// 子序列和顺序无关，其次我们肯定是需要对每个元素都去找其另一个适配的元素，
// 保证其分别为序列中的最小和最大元素，并且和小于等于 target，如果用二分，就是 nlogn，再加个排序
// 滑窗可能也行，但也还是需要先排个序 nlogn

// // 先试试滑窗？双指针，执行效率相对来说应该是比二分高的
// // 初始思路，但会导致越界
// func numSubseq(nums []int, target int) int {
// 	sort.Ints(nums)
// 	mod := (1e9 + 7)
// 	var ans int
// 	// 之前做排序后的两数之和我们知道，我们可以双指针分别从左右两端向中间递进
// 	// 求得最小最大元素和，小于等于 target
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		if nums[l]+nums[r] <= target {
// 			// 这里我们确定好了最小值和最大值，比如说 [0,1,2,3]，我们确定最小值是0，最大值是3
// 			// 我们可以计算该最小值最大值情况下，满足条件的子序列有 2^(3) 种，除了最小值固定外，另外3个值都可取可不取
// 			ans = (ans + 1<<(r-l)) % int(mod)
// 			l++
// 		} else {
// 			r--
// 		}
// 	}
// 	return ans % int(mod)
// }

// 但上面显然存在一定越界问题，数组长度最长为 1e5，2^(1e5) 在幂运算过程中就需要不断取模了
// 这里有两种解决方法，一种是用一个集合去初始化和存储幂运算的结果，另一种就是自己写个快速幂，并在过程中不断取模
func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	mod := (1e9 + 7)
	var ans int
	// 之前做排序后的两数之和我们知道，我们可以双指针分别从左右两端向中间递进
	// 求得最小最大元素和，小于等于 target
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l]+nums[r] <= target {
			// 这里我们确定好了最小值和最大值，比如说 [0,1,2,3]，我们确定最小值是0，最大值是3
			// 我们可以计算该最小值最大值情况下，满足条件的子序列有 2^(3) 种，除了最小值固定外，另外3个值都可取可不取
			ans = (ans + pow(2, r-l)) % int(mod)
			l++
		} else {
			r--
		}
	}
	return ans % int(mod)
}

// 快速幂，比如3^11,本来要3*3乘上11次
// 3^(1011)_2 = 3^(2^3)*3^(2^1)*3^(2^0) = ...(81=(9*9)=(3*3)*(3*3))*(9=3*3)*(3)
// 4次3的乘+3次结果乘
func pow(x, n int) int {
	res := 1
	for n != 0 {
		if n&1 == 1 {
			res = (res * x) % (1e9 + 7)
		}
		n >>= 1
		x = (x * x) % (1e9 + 7)
	}
	return res
}
