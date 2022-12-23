package main

// https://leetcode.cn/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 43. 1～n 整数中 1 出现的次数

// 先观察一下，看看有没有规律
// 1,10,11,...,19,21,31,...,91,100,101,...,199
// emm，虽然没观察到规律，但是我们是否可以考虑只关注 1 出现的情况
// 即统计每一位上 1 出现的次数，如个位上 1 出现的次数，十位上 1 出现的 次数
// 但是如何考虑呢？
// 以 2103 为例子
// 设个位为 1，2103 会出现多少次个位为 1 ？ --> 简单考虑一下就是 0 ~ 210 次，即 210 - 0 + 1 = 211 次，更简单可以拿 12 举个栗子，个位上出现 1 的情况为 01 和 11
// 设十位为 1，2103 会出现多少次十位为 1 ？ --> 此时十位未满 1，因此我们可以从前面取 1 满上，换算考虑为 2019 --> 此时千百位为 00~20 -> 20-0+1 = 21; 个位为 0~9, 因此为 21 * 10 = 210
// 设百位为 1，2103 会出现多少次百位为 1 ？ --> 此时百位为 1，但显然高位为 2 时的低位(个位+十位)未满，因此我们不能粗暴地直接取 0~2
// 因此我们需要做拆分，先取(0~1) * (0~99)，2*100 = 200，再去取高位为 2 情况下的低位 (00~03)=4 -> 200 + 4 = 204
// 设千位为 1，2103 会出现多少次千位为 1 ?  --> 此时千位为 2，其高位显然为 0， 低位为 0~999，总计 0*1000 + 1000 = 1000
// 总计为 211 + 210 + 204 + 1000 = 3524
func countDigitOne(n int) int {
	var res int
	high, cur, low := n/10, n%10, 0
	digit := 1
	// 当高低位还能继续计算就不停止
	for high != 0 || cur != 0 {
		switch cur {
		case 0:
			res += high * digit
		case 1:
			res += high*digit + low + 1
		default:
			res += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return res
}

// // 先暴力试试，不过应该会超时
// func countDigitOne(n int) int {
// 	countOne := func(n int) int {
// 		var res int
// 		for n != 0 {
// 			if n%10 == 1 {
// 				res++
// 			}
// 			n /= 10
// 		}
// 		return res
// 	}
// 	var res int
// 	for i := 1; i <= n; i++ {
// 		res += countOne(i)
// 	}
// 	return res
// }
