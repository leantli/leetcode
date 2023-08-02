package main

// https://leetcode.cn/problems/trapping-rain-water/
// 42. 接雨水

// // 二刷
// // 显然，我们就对每一列的左右各自的最大高度，取最小，减去当前的高度
// // min(maxL,maxR)-curHeight，值大于 0，即为当前位置能接的雨水
// // 此时我们就知道，只要某个位置的左右有比它高的，就能接水
// // 左右都比之前的高，显然可以利用单调递减栈，从栈底到栈顶是单调递减的
// // 此时我们遇到一个新的位置，高度比栈顶的高，则可以弹出栈顶，新位置的高度和此时新的栈顶，就构成了弹出的位置的左右高
// // 此时再 (min(maxL,maxR)-curHeight) * (curIdx - 新栈顶的下标位置 -1) ，即为弹出位置可以接的一部分
// func trap(height []int) int {
// 	stack := make([]int, 0)
// 	var res int
// 	for i, rightHeight := range height {
// 		for len(stack) > 0 && rightHeight > height[stack[len(stack)-1]] {
// 			midHeight := height[stack[len(stack)-1]]
// 			stack = stack[:len(stack)-1]
// 			if len(stack) > 0 {
// 				leftIdx := stack[len(stack)-1]
// 				res += (min(rightHeight, height[leftIdx]) - midHeight) * (i - leftIdx - 1)
// 			}
// 		}
// 		stack = append(stack, i)
// 	}
// 	return res
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// // 我们知道，要求下雨能接多少雨水，最基础的思路的是
// // 对每一列求其左右的最高高度，取两侧最高的较低者，减去当前的高度，则为当前列能容纳的水
// // 这里我们可以先两次遍历，分别求出每一列左侧的最高值和每一列右侧的最高值，最后一次遍历直接计算接雨水的量即可，此时 O(3n)
// // 也可以采取单调栈的方式，这里显然需要求出某一列小，两侧大的情况，使用单调递减栈，从栈底到栈顶，单调递减
// // 此时遍历到下标 r，若 height[r] > stack.top()，弹出栈顶下标，获取 midHeight，此时新的栈顶一定比 midHeight 高
// // 并且当前 height[r] 也比 midHeight 高，则能够计算接的雨水, (min(height[r],height(stack.top()))-midHeight) * (r-stack.top()-1)
// func trap(height []int) int {
// 	stack := make([]int, 0)
// 	var res int
// 	for i, rightHeight := range height {
// 		for len(stack) > 0 && rightHeight > height[stack[len(stack)-1]] {
// 			midHeight := height[stack[len(stack)-1]]
// 			stack = stack[:len(stack)-1]
// 			if len(stack) > 0 {
// 				res += (min(rightHeight, height[stack[len(stack)-1]]) - midHeight) * (i - stack[len(stack)-1] - 1)
// 			}
// 		}
// 		stack = append(stack, i)
// 	}
// 	return res
// }
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// // 我们关注到 逐列取值 的优化方案--动态规划，先求出每个列的左右最高列，其实只需要关注左右两个最高列值即可
// // 这里我们可以采用双指针优化
// // 双指针，左右各自指向最高的位置，哪边更低哪边向中间移动，保证移动后的位置，如果更低，左右都有最高值
// // 移动时，如果遇到更高的数，则更新对应指针侧的最高值，否则该侧最高值减去当前指针的高度，确认当前位置能装多少水
// // 因为当前位置比其所在侧的最高高度要低，其次另一侧一定比当前侧的最高高度高，因此该位置的水深就是该侧最高值减去当前位置的高度
// func trap(height []int) int {
// 	l, r := 0, len(height)-1
// 	topL, topR := height[l], height[r]
// 	var res int
// 	for l < r {
// 		if height[l] < height[r] {
// 			l++
// 			if height[l] < topL {
// 				res += topL - height[l]
// 			} else {
// 				topL = height[l]
// 			}
// 		} else {
// 			r--
// 			if height[r] < topR {
// 				res += topR - height[r]
// 			} else {
// 				topR = height[r]
// 			}
// 		}
// 	}
// 	return res
// }
// // 其实我们只要关注每一列左右两侧最高的列值，就可以了求出当前列能接的积水
// // 因此我们直接从左右两侧开始度量两侧的最高值，哪边低就先移动哪边
// // 移动后，如果新的位置比该侧初始最高值更高，则更新最高值，无需计算此处的积水，因为该侧之前的最高值比它低，无法接水
// // 如果新位置比该侧最高值低，显然可以接水，接的水就是该侧的最高值减去当前位置的高度
// func trap(height []int) int {
//     n := len(height)
//     l, r := 0, n-1
//     maxL, maxR := height[l], height[r]
//     var res int
//     for l < r {
//         if height[l] < height[r] {
//             l++
//             if height[l] > maxL {
//                 maxL = height[l]
//             } else {
//                 res += maxL - height[l]
//             }
//         } else {
//             r--
//             if height[r] > maxR {
//                 maxR = height[r]
//             } else {
//                 res += maxR - height[r]
//             }
//         }
//     }
//     return res
// }

// 逐列求积水的优化 --- 动态规划，记录每个位置左右两侧的最高高度值即可
// 在逐列求积水的代码中，对每个位置的左右两侧都求其最高高度，显然重复计算过多
// 这里我们可以先将每个位置的左右最高高度一次性求出，后续一次遍历即可，时间复杂度为 O(3n)
// func trap(height []int) int {
//     n := len(height)
//     maxLeft := make([]int, n)
//     maxRight := make([]int, n)
//     for i := 1; i < n; i++ {
//         maxLeft[i] = max(maxLeft[i-1], height[i-1])
//     }
//     for i := n-2; i >= 0; i-- {
//         maxRight[i] = max(maxRight[i+1], height[i+1])
//     }
//     var res int
//     for i := 1; i < n-1; i++ {
//         cur := min(maxLeft[i], maxRight[i]) - height[i]
//         if cur > 0 {
//             res += cur
//         }
//     }
//     return res
// }
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// // 暴力？一行一行往上遍历，同一层时只要当前位置左右有柱子，则可以增加
// // 存在超时用例
// func trap(height []int) int {
//     var max int
//     for i := range height {
//         if max < height[i] {
//             max = height[i]
//         }
//     }
//     var res int
//     for i := 0; i < max; i++ {
//         var cur int
//         // l 记录上次遇到的最后一个柱子的位置
//         l := -1
//         for j := range height {
//             if height[j] > i {
//                 if l != -1 {
//                     cur += j-l-1
//                 }
//                 l=j
//             }
//         }
//         res += cur
//     }
//     return res
// }

// // 正常情况下要计算两个水坑之间的容积
// // 我们得先找到水坑的两个高，保证这两个高中间的水不会流走
// // 但是这并不方便，比如 [6,0,3,0,6]，此时我们会先找到 6 和 3
// // 接着就直接计算了二者之间的容器，但实际上，6 可以和后面的 6 构成更大的容器
// // 所以，我们必须同时找到两边的最高值, 那么如何找到两边的最高值呢？
// // 每到一列，我们就向左向右去找比它高的，如果都找到了，那么就用较矮的高减去其本身的高度即可
// // 但是这样的话，每一列都要找一遍，相当于 n*n 的时间复杂度
// // 这里有个小技巧思路：
// // 我们可以用两个指针分别指向首尾
// // 共同维护两边的最高值，哪边的最高值较低，那么这一边的指针就向中间步进
// // 并判断是否是更高的高度，更高则更新这边的最高值，更低则计算此处能装多少水
// // 因为我们是移动最高值较低的一边，因此另一边的高度是足够支持此处容纳水的
// // 并且如果是同一边有更高的值，也不会影响计算，每一边的高度都只会越来越高
// // 因此遇到低的，就一定能直接计算这一步的容水情况
// func trap(height []int) int {
// 	l, r := 0, len(height)-1
// 	lh, rh := height[l], height[r]
// 	var cap int
// 	for l < r {
// 		if lh < rh {
// 			l++
// 			if lh > height[l] {
// 				cap += lh - height[l]
// 			} else {
// 				lh = height[l]
// 			}
// 		} else {
// 			r--
// 			if rh > height[r] {
// 				cap += rh - height[r]
// 			} else {
// 				rh = height[r]
// 			}
// 		}
// 	}
// 	return cap
// }
