package main

// https://leetcode.cn/problems/design-hashmap/
// 706. 设计哈希映射

// 哈希函数
// 哈希冲突的解决方案---常用链地址法；开放地址法；
// 扩容，当哈希表中元素过多，我们要考虑扩容
// 基于数组，数组下标通过hash函数计算后快速得出，存取届基于哈希函数，减少遍历/二分寻找时间
// 由于简单题，不考虑扩容，直接搞个大点的数组，这道题操作最多1e4次
type MyHashMap struct {
	base int
	m    []*Entry
}

type Entry struct {
	key  int
	val  int
	next *Entry
}

func Constructor() MyHashMap {
	base := 10000
	return MyHashMap{base: base, m: make([]*Entry, base)}
}

func (this *MyHashMap) hash(params int) (key int) {
	return params % this.base
}

func (this *MyHashMap) Put(key int, value int) {
	node := &Entry{key: key, val: value}
	mIndex := this.hash(key)
	if this.m[mIndex] == nil {
		this.m[mIndex] = node
		return
	}
	cur := this.m[mIndex]
	for cur.next != nil {
		if cur.key == key {
			cur.val = value
			return
		}
		cur = cur.next
	}
	if cur.key == key {
		cur.val = value
		return
	}
	cur.next = node
}

func (this *MyHashMap) Get(key int) int {
	mIndex := this.hash(key)
	if this.m[mIndex] == nil {
		return -1
	}
	cur := this.m[mIndex]
	for cur != nil {
		if cur.key == key {
			return cur.val
		}
		cur = cur.next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	mIndex := this.hash(key)
	if this.Get(key) == -1 {
		return
	}
	cur := this.m[mIndex]
	if cur.key == key {
		this.m[mIndex] = cur.next
		return
	}
	for cur.next != nil && cur.next.key != key {
		cur = cur.next
	}
	cur.next = cur.next.next
}
