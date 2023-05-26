package main

// https://leetcode.cn/problems/design-hashset/
// 705. 设计哈希集合

// // 哈希函数
// // 哈希冲突的解决方案---常用链地址法；开放地址法；
// // 扩容，当哈希表中元素过多，我们要考虑扩容
// // 基于数组，数组下标通过hash函数计算后快速得出，存取届基于哈希函数，减少遍历/二分寻找时间
// // 由于简单题，不考虑扩容，直接搞个大点的数组，这道题操作最多1e4次
// // 并且注意，由于是 set，不要 add 重复数
// type MyHashSet struct {
// 	set  [][]int
// 	base int
// }

// func Constructor() MyHashSet {
// 	base := 10000
// 	return MyHashSet{set: make([][]int, base), base: base}
// }

// func (this *MyHashSet) hash(params int) (key int) {
// 	return params % this.base
// }

// func (this *MyHashSet) Add(key int) {
// 	if len(this.set[this.hash(key)]) < 0 {
// 		this.set[this.hash(key)] = make([]int, 0)
// 	}
// 	if this.Contains(key) {
// 		return
// 	}
// 	this.set[this.hash(key)] = append(this.set[this.hash(key)], key)
// }

// func (this *MyHashSet) Remove(key int) {
// 	if len(this.set[this.hash(key)]) == 0 {
// 		return
// 	}
// 	deleteIdx := -1
// 	for i, v := range this.set[this.hash(key)] {
// 		if v == key {
// 			deleteIdx = i
// 			break
// 		}
// 	}
// 	if deleteIdx == -1 {
// 		return
// 	}
// 	this.set[this.hash(key)] = append(this.set[this.hash(key)][:deleteIdx], this.set[this.hash(key)][deleteIdx+1:]...)
// }

// func (this *MyHashSet) Contains(key int) bool {
// 	for _, v := range this.set[this.hash(key)] {
// 		if v == key {
// 			return true
// 		}
// 	}
// 	return false
// }

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
