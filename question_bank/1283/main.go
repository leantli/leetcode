package main

// https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1283. 使结果不超过阈值的最小除数

// 整数数组， 正整数
// 选择一个正整数作为除数，数组每个数都 除以 它，对结果求和
// 找到一个 最小的 除数，满足 求和结果 <= 阈值
// 由于题目规定了 nums.length <= threshold <= 10^6
// 因此我们能确定，除数的最大取值是数组中最大的值？
// 枚举范围为[1, max(nums)]，但我们可以采用二分
// mid 为本次采用的除数
// 当 结果和 < 阈值时，说明结果和可以再大些，除数可以继续变小, r=mid
// 当 结果和 > 阈值时，说明结果和得小点，让除数变大一些, l=mid
// 这里规定了数组的取值，其实都是正整数
// 无语了，写完之后发现 相除结果向上 取整，没注意这个，简单改一下
func smallestDivisor(nums []int, threshold int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	l, r := 0, max+1
	for l+1 != r {
		mid := l + (r-l)/2
		var sum int
		for _, num := range nums {
			sum += (num-1)/mid + 1
		}
		if sum <= threshold {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
