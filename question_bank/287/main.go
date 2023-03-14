package main

// https://leetcode.cn/problems/find-the-duplicate-number/
// 287. 寻找重复数

// 相似题型：剑指03,56-1;;leetcode-287,26,136,137,260

// 第n次做，这次用快慢指针
// 我们知道，这个数组长度为n+1,内部的数值范围为[1,n]
// 因此我们大可以，从index为0的下标开始，，将下标指向的值作为新的下标进行遍历
// 由于数组中存在重复的数，因此这种遍历方式一定会成环
// 比如 [1,3,4,2,2] --> 1,3,[2,4],[2,4],[2,4],...
// 比如 [3,1,3,4,2] --> [3,4,2],[3,4,2],...
// 结合环形链表2，当快慢指针在环中相遇时，
// 分为成环前道路a,成环后slow走的为b，slow走的道路为k=a+b;fast走的道路为2k=a+2b+c
// 显然c=a，此时快指针回到起始点，走c/a步，慢指针正常走也是c/a步，刚好在成环处相遇
func findDuplicate(nums []int) int {
	fast, slow := nums[nums[0]], nums[0]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}

// // 常规来说，一般是用 map 解决，但空间复杂度不满足题意(要求O(1))
// // 又看到 n+1 个数，数字范围是[1,n]，下标与数对齐(鸠占鹊巢)方法也可以，但是不满足题意(不修改数组nums)
// // 关注到，它圈定了数字范围，[1,n]，并且只有一个重复数，至少重复一次
// // 这里我们可以用二分去猜，谁是重复数，遍历一遍看看小于等于这个重复数的数量，是否是重复数的值
// // 比如正常 1 2 3 4 4，假设重复数是 3，此时小于等于 3 的有三个数，它不是重复数，假设重复数是 4
// // 此时小于等于 4 的有5个数，它显然是重复数
// // 时间复杂度为 nlogn
// func findDuplicate(nums []int) int {
// 	l, r := 0, len(nums)
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		var cnt int
// 		for _, num := range nums {
// 			if num <= mid {
// 				cnt++
// 			}
// 		}
// 		if cnt <= mid {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	// 此时 l 落在导致后面计数错误的数(重复数)的左侧，r 即为重复数
// 	return r
// }

// // 第二次用 二分 刷到这道题
// // 正常情况我们能基于 交换法，或者 取模+下标位置找到重复数，但是题目要求不修改数组
// // 这里我们又注意到 数字 都在 [1,n] 范围内，找一个重复的整数
// // 显然这里其实可以考虑二分？但是我们还需要找到这个二分缩进的规律
// // 有重复的数，这意味着 重复的数 及其前面的数 的总数是异常的，怎么理解这个异常呢？
// // 比如 [1,n] 范围内，假设 n 为 5，重复出现了 3，那么必然是 1,2,3,3,4,5
// // 可以发现，[1,2] 范围内只有两个数，这是正常的,[1,4] 范围内正常应该也只有 4 个数，然而此时却存在 5 个数，这显然是异常的
// // 因此我们可以根据此 特性 去二分缩进
// // 正常则左缩，异常则右缩，最终右指针落在第一个异常的数上，即是重复的数
// func findDuplicate(nums []int) int {
// 	l, r := 0, len(nums)
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		var cnt int
// 		for _, num := range nums {
// 			if num <= mid {
// 				cnt++
// 			}
// 		}
// 		if cnt <= mid {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	return r
// }

// // 还能用什么方法呢？看起来也不能用 摩尔投票，因为出现的次数不一定超过数组长度一半
// // 看了一下题解，发现可以用 二分？！
// // 我是没想到的。。。
// // 整个思路时类似 根据大小判断猜数字
// // l 为 1， r 为 n，因为数字可能出现的范围就是 [1,n]
// // 因此确定上下边界开始二分猜数字
// // 那么如何确定下一次猜数字是往上区间还是下区间呢？
// // 根据本数组中，小于当前数字的数的出现次数
// // 显然，每个数只出现 1 次，则小于当前数的所有数出现的次数相加，也必定小于当前数
// // 比如 3，小于等于 3 的数只有 1、2、3，出现次数为 3，必定是小于等于 3 的
// // 确保下次猜数字往上区间去猜(即往有异常的方向去猜)
// func findDuplicate(nums []int) int {
// 	l, r := 1, len(nums)-1
// 	// l < r 保证退出循环时，l 和 r 重合
// 	for l < r {
// 		mid := l + (r-l)/2
// 		cnt := 0
// 		for _, num := range nums {
// 			if num <= mid {
// 				cnt++
// 			}
// 		}
// 		// 当 cnt 大于 mid 时异常，则左侧区间有重复的，不正常
// 		// 下一次要猜的数字，一定在 [l, mid] 区间中
// 		if cnt > mid {
// 			r = mid
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return l
// }

// // 这里首先是 [1,n] 中只存在一个 重复的 数，并且该数不确定出现几次，可能 2+ 次
// // 因此我们无法对该数直接进行位运算，因为我们连其出现的次数都无法确定
// // 这里我们可以关注到，数组为 n+1, 并且每个数都是在 [1,n] 的，因此我们可以原地操作
// // 遇到一个数，就将该数作为下标，将对应下标的数 +n, 当我们遇到某个数第二次时，发现该下标对应的数已经大于 n 了，因此可以直接返回该重复的数
// // emmm, 做完后发现题目要求不对 num 进行操作，因此该方法不行
// func findDuplicate(nums []int) int {
// 	n := len(nums)
// 	for _, num := range nums {
// 		// 将 num 对应的下标 进行判断操作
// 		index := num % n
// 		if nums[index] <= n {
// 			nums[index] += n
// 		} else {
// 			return num
// 		}
// 	}
// 	return -1
// }
