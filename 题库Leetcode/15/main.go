package main

import "sort"

// https://leetcode.cn/problems/3sum/
// 15. 三数之和

// 返回所有和为0且不重复的三元组
// 注意这次是输出值而不是下标，因此我们可以先排序
// 然后基于双指针进行选取，这样可以不用 set 占用过多空间
// 不过我们也要注意，存储的东西不能重复，这里我们可以用 set
// 也可以在得到 三数和为 0 时，判断 b 和 c 是否可能存在重复取值
// a 不取之前取过的值
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for a := 0; a < len(nums)-2; a++ {
		// 当第一个数大于0，那么三数之和一定无法等于 0 (已排序过，后面两个数只会更大)
		if nums[a] > 0 {
			break
		}
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b, c := a+1, len(nums)-1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == 0 {
				res = append(res, []int{nums[a], nums[b], nums[c]})
				b++
				c--
				for b < c && nums[b] == nums[b-1] {
					b++
				}
				continue
			}
			if sum > 0 {
				c--
			} else {
				b++
			}
		}
	}
	return res
}
