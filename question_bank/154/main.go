package main

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=c8d11zm
// 154. 寻找旋转排序数组中的最小值 II

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		// 去除首尾的重复元素，保证两个单调区间的单调性
		if nums[mid] == nums[r] {
			r--
			continue
		}
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[r]
}

/*
   与旋转数组I非常类似,唯一的不同就是有重复元素
   主体思路仍是二分法:
   nums总是可分为:nums1与nums2两个子数组,并且两个数组都是递增的,且nums1[i]>=nums2[j]
   mid的索引取地板除,最终要根据nums[mid]与nums[right]进行范围缩小
   最终要找的是nums2的起始索引对应的值
   1.nums[mid]>nums[right],如[5,6,7,8,4],那么nums2起点必定严格位于mid(不含)右边,left=mid+1
   2.nums[mid]<nums[right],如[5,1,2,3,4],那么nums2起点必定位于mid(含)左边,right=mid
   3.nums[mid]==nums[right],如[5,1,2,2,2],那么right--继续进入下一轮循环
       3.1 这个操作不会越界:right>left>=0
       3.2 这个操作会一直尝试收缩右边界间接缩小mid直至[mid,right]跳出相等范围序列
       而被收缩的原本的nums[right]必定不是唯一的最小值,换句话说有效范围依然合法
       反证:nums[right]是唯一最小值,left<right && mid<right(地板除)
           因此不可能出现nums[mid]==nums[right]与题设矛盾
       有以下几种情况:[3,4,2,2,2,2];[3,2,2,2,2,2];[2,2,2,2,2,2]
       都可以正确求出最小值索引
   最后left==right直接返回nums[left]即可
*/
