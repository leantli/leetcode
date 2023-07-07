package main

// https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=131d40r
// 65. 不用加减乘除做加法

// 肯定是位运算，用位运算做加法
// 比如 3+5=8
// 0011
// 0101
// 1000
// 这里可以拆分成三步
// 1. 计算本位: 不管啥进位，直接把两个数相加
//     0011
//     0101
// 得到 0110

// 2. 计算进位:只算进位
//     0011
//     0101
// 得到 0010

// 3. 将进位和本位相加
//     0110
//     0010
// 得到 1000

// 1. 本位相加可以通过 a^b，当ab为 11 时因为进位，本位会变为0
// 当ab为10/01时，本位仍为1；当ab为00时，本位仍为0，因此通过异或最 ok
// 2. 是否有进位可以通过 a&b 判断，有则 <<1，表示已进位
// 3. 但上面第三步是进位与本位相加，我们不能直接相加，因此我们可以通过递归，
// 将上面计算得到的 本位 作为 a，进位 作为 b，让他们两个继续相加，即通过 异或 进行本位加
// 直到进位为 0 时，此时已无相加必要，不会再变更了

// 递归，貌似更好理解？
func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}

// 别人的迭代解答
// func add(a, b int) int {
// 	for b != 0 {
// 		carry := uint(a&b) << 1
// 		a ^= b
// 		b = int(carry)
// 	}
// 	return a
// }
