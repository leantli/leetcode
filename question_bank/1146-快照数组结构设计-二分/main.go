package main

// https://leetcode.cn/problems/snapshot-array/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 1146. 快照数组

// 再重新考虑一下，好像是可以和 981 这道题差不多
// map --> index:snap数组
// map --> index: (snap:value)
type SnapshotArray struct {
	version           int
	indexToSnaps      map[int][]int
	indexToSnapMapVal map[int]map[int]int
}

func Constructor(length int) SnapshotArray {
	indexToSnaps := make(map[int][]int)
	indexToSnapMapVal := make(map[int]map[int]int)
	for i := 0; i < length; i++ {
		indexToSnaps[i] = make([]int, 1)
		indexToSnapMapVal[i] = make(map[int]int)
	}
	return SnapshotArray{version: 0, indexToSnaps: indexToSnaps, indexToSnapMapVal: indexToSnapMapVal}
}

func (this *SnapshotArray) Set(index int, val int) {
	if _, ok := this.indexToSnapMapVal[index][this.version]; !ok {
		this.indexToSnaps[index] = append(this.indexToSnaps[index], this.version)
	}
	this.indexToSnapMapVal[index][this.version] = val
}

func (this *SnapshotArray) Snap() int {
	res := this.version
	this.version++
	return res
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	l, r := -1, len(this.indexToSnaps[index])
	for l+1 != r {
		mid := l + (r-l)/2
		// fmt.Printf("mid:%d,index:%d,snap_id:%d,mid->snap_id:%d\n", mid, index, snap_id, this.indexToSnapMapVal[index][mid])
		if this.indexToSnaps[index][mid] <= snap_id {
			l = mid
		} else {
			r = mid
		}
	}
	// l 至少会落在 0 下标指向 snap_id 为 0，否则落在最大的 小于等于 snap_id 的值的下标
	lastSnapId := this.indexToSnaps[index][l]
	return this.indexToSnapMapVal[index][lastSnapId]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

// // map[int][]int --- int-[]int 结构？
// // 甚至就是一个[][]int ?
// // 不过这样的话显然就是简单题了，中等题感觉会超时或者爆内存之类的
// type SnapshotArray struct {
// 	length  int
// 	version int
// 	snapArr [][]int
// }

// func Constructor(length int) SnapshotArray {
// 	snapArr := make([][]int, 0)
// 	snapArr = append(snapArr, make([]int, length))
// 	return SnapshotArray{snapArr: snapArr, version: 0, length: length}
// }

// func (this *SnapshotArray) Set(index int, val int) {
// 	this.snapArr[this.version][index] = val
// }

// func (this *SnapshotArray) Snap() int {
// 	next := make([]int, this.length)
// 	copy(next, this.snapArr[this.version])
// 	this.snapArr = append(this.snapArr, next)
// 	temp := this.version
// 	this.version++
// 	return temp
// }

// func (this *SnapshotArray) Get(index int, snap_id int) int {
// 	return this.snapArr[snap_id][index]
// }
