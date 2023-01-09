package main

// https://leetcode.cn/problems/find-the-student-that-will-replace-the-chalk/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 1894. 找到需要补充粉笔的学生编号

// 还是模拟，但是可以先取得所有学生回答一次用的粉笔数量
// 再取模，只判断一轮
func chalkReplacer(chalk []int, k int) int {
	var sum int
	for _, num := range chalk {
		sum += num
	}
	k %= sum
	var i int
	for k >= chalk[i] {
		k -= chalk[i]
		i++
		i %= len(chalk)
	}
	return i
}

// // 模拟？结果遇到 [1] 100000000000 就超时了
// func chalkReplacer(chalk []int, k int) int {
// 	var i int
// 	for k >= chalk[i] {
// 		k -= chalk[i]
// 		i++
// 		i %= len(chalk)
// 	}
// 	return i
// }
