package main

// https://leetcode.cn/problems/contains-duplicate-iii/
// 220. 存在重复元素 III

// 看了官解--最强思路--滑窗+桶排序--感觉正常想不到吧
// 1. 了解如何分桶，为什么要怎么分桶？
// 我们对数组中每个元素都除以(t+1)，此时得到一个值，将其作为 index
// 这样做的意思是什么？即，这些元素，除以(t+1)后，都是同一个 index，
// 而这个 index 对应的这些元素，一定满足，差的绝对值，<= t
// 但这里我们注意到，以上求元素是否在同一个 index，只限于 元素 >= 0 的情况
// 负数的情况显然要额外处理？为什么？比如 [-4,-3,-2,-1,0,1,2,3] t=3
// num >= 0 的情况下，[0,1,2,3] 都在 index = 0 的桶中，它们也显然都属于同一个桶，并且差的绝对值都 <= t
// 但是我们再一看，-3,-2,-1，除以 t+1，得到的 index 也是 0，但是 -3 和 3 显然就满足我们要的条件(3-(-3)并不<=t)
// 因此对于负数我们需要额外处理，比如 num < 0 时, index = x/(t+1)-1，但是我们注意到，此时 index 为 -1 的桶
// 只装载着 [-3,-2,-1]，这显然不对吧，-4 也应该进去的，因此对于负数， index = (x+1)/(t+1)-1
// 当然，我们这道题实际情况下的话，一个桶一般最多只会有一个数，因为如果有第二个数到这个桶，就可以直接返回 true 了
// 此时我们知道了这些数，应该怎样分进各个桶中
// 2. 为什么要分桶？
// 这样分桶后，非常方便我们判断两个数的差的绝对值是否<=t；比如当前数为 i, index=(i)/(t+1)
// 我们确保，与 i 差的绝对值<=t的数，一定存在同一个桶或相邻的左右两桶中，因为一个桶，差值就是 t+1
// 如果 i 刚好是中间桶的边缘，那么边缘依靠的的另一个桶，是可能存在满足条件的数的
// 3. 如何利用这个桶？
// 遍历，保持窗口长度 k+1，每进一个数，就判断该数所在桶是否已有其他数，没有则判断其两边的桶，是否有数，有数则计算是否满足绝对差值<=t
// 当然，我们在给元素分桶时，也要保持一个 k+1 长度的窗口，窗口右移时，要将最左侧排出的数，从桶中移除
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 也可以 滑动窗口+红黑树
// 每次要进新的数字 num，就在红黑树窗口中搜索两个数
// 这两个数分别是 大于等于 num 的第一个数 和 小于等于 num 的最后一个数
// 即 [:num] 范围中最接近 num 的数 和 [num:] 范围中最接近 num 的数
// 此时再比较 num 和 这两个数的绝对差值即可，总体时间复杂度比之前的暴力求绝对值小
// java 可以用 treeset, go 得自己实现(狗头)
// 如果要这种思路的话建议改 java 写....
// func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

// }

// // 滑动窗口+暴力求绝对值(通过)
// // 返回是否存在 i,j
// // 使得 abs( nums[i] - nums[j] ) <= t
// // abs( i-j ) <= k
// // 这里显然 abs(i-j)<=k 相当于一个定长为 k+1 的滑动窗口，窗口内之间的距离一定满足相减后绝对值小于等于 k
// // 也用不到前缀和这种求子数组和/差的，也不方便排序
// // 否则的话每次窗口新进来一个数，就要和内部的数做一个 abs 差值比较
// // 也能通过，但是时间复杂度显然太高了
// func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
// 	k++
// 	var l int
// 	for r := 1; r < len(nums); r++ {
// 		// 此时要开始舍弃最左边的数了
// 		if r >= k {
// 			l++
// 		}
// 		for i := l; i < r; i++ {
// 			if abs(nums[i], nums[r]) <= t {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func abs(a, b int) int {
// 	temp := a - b
// 	if temp < 0 {
// 		return -temp
// 	}
// 	return temp
// }
