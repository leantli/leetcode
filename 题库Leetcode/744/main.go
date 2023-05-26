package main

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/?envType=study-plan&id=binary-search-beginner&plan=binary-search&plan_progress=cnhyx51
// 744. 寻找比目标字母大的最小字母

// 字符数组 非递减顺序 有字符target
// 要求返回 大于 target 的最小字符，不存在则返回 第一个字符
// 尝试一下万金油模板
// func nextGreatestLetter(letters []byte, target byte) byte {
// 	l, r := -1, len(letters)
// 	for l+1 != r {
// 		mid := l + (r-l)/2
// 		if letters[mid] <= target {
// 			l = mid
// 		} else {
// 			r = mid
// 		}
// 	}
// 	if r == len(letters) {
// 		return letters[0]
// 	}
// 	return letters[r]
// }

// 再尝试一下模糊模板
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	// 此时需要特殊处理最后一个字符的大小判断，否则后续二分时，很难判断最后一个字符
	if letters[r] <= target {
		return letters[0]
	}
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l]
}
