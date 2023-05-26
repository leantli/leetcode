package main

// https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 数组中出现次数超过一半的数字

// 再重新审一下题目细节，数字出现次数超过一半
// 其他数字出现的次数都不足一半
// 也就是说，我们可以用出现次数去互相抵消，最后就会只剩下出现次数超过一半的数
// 怎么抵消呢？我们需要一个 count 计数，遇到相同的数就 count++
// 遇到不同的数时，count--，如果 count==0，则换新的数并且 count++
// 如此一来，出现次数最多的数，其 count 一定能抵消其他数的出现次数
func majorityElement(nums []int) int {
	var count, res int
	for _, num := range nums {
		if count == 0 {
			count++
			res = num
		} else {
			if res != num {
				count--
			} else {
				count++
			}
		}
	}
	return res
}

// 简单题应该直接 map 解决？先写一下，不过我印象中是有更好的解法
// 时间复杂度应该已经是最佳的了？主要是空间复杂度
// 排序后再取中值，实际上时间复杂度是O(nlogn)，比 map 还差

// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	var maxNum int
// 	for _, num := range nums {
// 		m[num]++
// 		// 这里再判断一下，如果此时出现次数已经大于 len/2 了
// 		// 说明已经确认是该数了，可以提前终止了
// 		if m[num] >= len(nums)/2 {
// 			return num
// 		}
// 		if m[num] > m[maxNum] {
// 			maxNum = num
// 		}
// 	}
// 	return maxNum
// }
