package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 45. 把数组排成最小的数

// 借用 api 简易方法
// func minNumber(nums []int) string {
// 	strs := make([]string, len(nums))
// 	for i := range nums {
// 		strs[i] = strconv.Itoa(nums[i])
// 	}
// 	sort.Slice(strs, func(i, j int) bool {
// 		return strs[i]+strs[j] < strs[j]+strs[i]
// 	})
// 	bs := strings.Builder{}
// 	for i := range strs {
// 		bs.WriteString(strs[i])
// 	}
// 	return bs.String()
// }

// 越小的在越前面，但是 30 又在 3 前面。。
// 如何比较比较好？303 < 330， maybe 两个转成字符串之后再拼接再比较？
// 排序判断规则：若拼接字符串 x+y>y+x ，则 x “大于” y ；反之，若 x+y<y+x ，则 x “小于” y
func minNumber(nums []int) string {
	strNums := make([]string, 0, len(nums))
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}
	// 用于快排比较，组合起来 ab 比 ba 小的返回 true
	compare := func(a, b string) bool {
		num1, _ := strconv.Atoi(a + b)
		num2, _ := strconv.Atoi(b + a)
		return num1 < num2
	}
	var quickSort func(nums []string, l, r int) []string
	quickSort = func(nums []string, l, r int) []string {
		if l >= r {
			return nums
		}
		i, j, pivot := l, r, nums[l]
		for i < j {
			// pivot nums[j] < nums[j] pivot 时继续
			for i < j && compare(pivot, nums[j]) {
				j--
			}
			// nums[i] pivot < pivot nums[i] 时继续
			for i < j && !compare(pivot, nums[i]) {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[l], nums[i] = nums[i], nums[l]
		quickSort(nums, l, i-1)
		quickSort(nums, i+1, r)
		return nums
	}
	strNums = quickSort(strNums, 0, len(nums)-1)
	b := strings.Builder{}
	for _, num := range strNums {
		b.WriteString(num)
	}
	return b.String()
}

// 这里数字处理有问题，得用字符串 比如 30,31 时下面就不方便了
// func minNumber(nums []int) string {
// 	// 用于快排比较，组合起来 ab 比 ba 小的返回 true
// 	compare := func(a, b int) bool {
// 		return a*10+b < b*10+a
// 	}
// 	var quickSort func(nums []int, l, r int) []int
// 	quickSort = func(nums []int, l, r int) []int {
// 		if l >= r {
// 			return nums
// 		}
// 		i, j, pivot := l, r, nums[l]
// 		for i < j {
// 			// pivot nums[j] < nums[j] pivot 时继续
// 			for i < j && compare(pivot, nums[j]) {
// 				j--
// 			}
// 			// nums[i] pivot < pivot nums[i] 时继续
// 			for i < j && !compare(pivot, nums[i]) {
// 				i++
// 			}
// 			nums[i], nums[j] = nums[j], nums[i]
// 		}
// 		nums[l], nums[i] = nums[i], nums[l]
// 		quickSort(nums, l, i-1)
// 		quickSort(nums, i+1, r)
// 		return nums
// 	}
// 	nums = quickSort(nums, 0, len(nums)-1)
// 	b := strings.Builder{}
// 	for _, num := range nums {
// 		b.WriteString(strconv.Itoa(num))
// 	}
// 	return b.String()
// }

// 先回顾一下快排思路
// 取一个基准点 pivot，常取头部数字或者尾部数字，或中间随机一个数，再提到头部
// 然后把比该基准点小的，放在其左边，把比基准点大的放在其右边，每次排序只能确保排好一个值的位置，这也意味着最差时间复杂度是O(N^2)
// 接着在对基准点左右两个区间分治，按上面的方法继续排序，如此到最后

// 这里基于基准点，放小的在左，放大的在右，有两种方式，交换法和填坑法
// 首先设最左为 i，最右为 r，当取 i 为 pivot 时 (下面一行是交换法)
// r 先循环 --，直至遇到比 pivot 小的，此时 l 再开始 ++，直至遇到比 pivot 大的，此时交换两个数，如此往复，最后 l 和 r 相遇的位置再换进 pivot 对应的值
// 首先设最左为 i，最右为 r，当取 i 为 pivot 时 (下面一行是填坑法)
// r 先循环 --，直至遇到比 pivot 小的，此时直接交换 pivot 和 r, l 再开始 ++，直至遇到比 pivot 大的，此时直接交换 pivot 和 r，如此往复，最后 l 和 r 相遇的位置即可继续分治
// 二者效率差异其实不大, 见 benchmark 基准测试 go test -bench='.'

// 交换法
func quickSort(nums []int, l, r int) []int {
	if l >= r {
		return nums
	}
	i, j, pivot := l, r, nums[l]
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[l] = nums[l], nums[i]
	quickSort(nums, l, i-1)
	quickSort(nums, i+1, r)
	return nums
}

// 填坑法
func quickSort2(nums []int, l, r int) []int {
	if l >= r {
		return nums
	}
	i, j, pivot := l, r, nums[l]
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quickSort2(nums, l, i-1)
	quickSort2(nums, i+1, r)
	return nums
}
