package main

// https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips/
// 995. K 连续位的最小翻转次数

// 暂时只想到了模拟，并且不能够确定，这种常规模拟是否真的能得到最小的翻转次数
// 逐个遍历，遇到0就翻转，最后遍历一遍看看还有没有0存在，感觉其实这样并不一定能得到最小翻转次数
// 其实这样还真的可以得到最小翻转次数，因为后面区间的翻转并不会影响前面的元素
// 只是这样确实会超时，因为我们真实得进行了翻转操作，实际上效率过低，里面存在很多重复无用的翻转操作

// 其实我们要注意一个情况：A[i] 翻转偶数次的结果是 A[i]；翻转奇数次的结果是 A[i] ^ 1
// 这是一道滑窗练习题，怎么把这道题和滑窗联系上呢？
// 当需要翻转时，会影响到其本身以及之后的k-1个数，这里是否有点定长滑窗的感觉？
// 窗口右移时，将窗口最左的数也排出，因为排出的数即便翻转了也不会影响到此时新加入的数
// 此时我们只需要看窗口内有多少个会引起翻转的数，并结合 奇/偶数次的翻转规律
// 判断当前数是否会引起翻转，是的话则 计数 并 将其下标加入窗口，不会的话则略过
func minKBitFlips(nums []int, k int) int {
	// 这个 cases 存储 [i-k+1,i-1] 会引起翻转的数的下标
	// 也就是说，这个窗口最多只需要维持 k-1 的长度
	// 在确认好 i 位置是否需要翻转后，
	// 再移除 i-k+1 位置的数和加入 i 位置的数(如果需要)
	cases := make([]int, 0)
	var res int
	for i := 0; i < len(nums); i++ {
		// 先移除窗口中不会影响到当前下标的最左边的数
		// len(cases) 表示当前窗口中会导致当前位置数被翻转的次数
		if len(cases) != 0 && cases[0] == i-k {
			cases = cases[1:]
		}
		// 如果 i 被翻转的次数为奇数，则变动 nums[i]
		// 翻转次数为偶数次则无需变动
		if len(cases)&1 == 1 {
			nums[i] ^= 1
		}
		// 此时判断当前数是否会引起翻转，是则翻转并加入窗口中
		if nums[i] == 0 && i > len(nums)-k {
			return -1
		} else if nums[i] == 0 {
			nums[i] = 1
			cases = append(cases, i)
			res++
		}
	}
	return res
}

// // 后面区间的翻转，不会影响前面的元素。因此可以使用贪心策略，从左到右遍历，遇到每个 0 都把 它以及后面的  K−1 个元素进行翻转
// // 因此我们可以单纯模拟遍历翻转(超时，但前面的大部分用例都没问题)
// func minKBitFlips(nums []int, k int) int {
// 	var res int
// 	n := len(nums)
// 	for i, num := range nums[:n-k] {
// 		if num == 0 {
// 			j := i
// 			for j < i+k && j < n {
// 				nums[j] ^= 1
// 				j++
// 			}
// 			res++
// 		}
// 	}
// 	for _, num := range nums {
// 		if num == 0 {
// 			return -1
// 		}
// 	}
// 	return res
// }
