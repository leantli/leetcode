package main

// https://leetcode.cn/problems/happy-number/
// 202. 快乐数

// 并且显然，在知道成环后，完全可以用快慢指针解决，减少了set空间的存储
func isHappy(n int) bool {
	var getNextN func(n int) int
	getNextN = func(n int) int {
		var sum int
		for n != 0 {
			cur := n % 10
			sum += cur * cur
			n /= 10
		}
		return sum
	}
	fast, slow := getNextN(n), n
	for fast != 1 {
		if slow == fast {
			return false
		}
		slow = getNextN(slow)
		fast = getNextN(getNextN(fast))
	}
	return true
}

// // 这道题，知道原理就很简单，但是正常情况下感觉很难想到
// // ----无限循环的过程中，出现的数居然会成环!!!
// func isHappy(n int) bool {
// 	set := make(map[int]struct{})
// 	var sum int
// 	for sum != 1 {
// 		sum = 0
// 		for n != 0 {
// 			cur := n % 10
// 			sum += cur * cur
// 			n /= 10
// 		}
// 		if _, ok := set[sum]; ok {
// 			return false
// 		}
// 		set[sum] = struct{}{}
// 		n = sum
// 	}
// 	return true
// }
