package main

// https://leetcode.cn/problems/can-place-flowers/
// 605. 种花问题

// 什么情况才能种花？i-1,i,i+1都为0时，i才能种花
// 并且还要注意，当i为首末时，只需要他们旁边的数为0即可
// 这里我们可以给原数组前后都价格0，方便统一判断
func canPlaceFlowers(flowerbed []int, n int) bool {
	temp := append([]int{0}, flowerbed...)
	temp = append(temp, 0)
	var cnt int
	for i := 1; i < len(temp)-1; i++ {
		if temp[i] == 1 {
			continue
		}
		if temp[i-1] == 0 && temp[i] == 0 && temp[i+1] == 0 {
			cnt++
			temp[i] = 1
		}
	}
	return cnt >= n
}
