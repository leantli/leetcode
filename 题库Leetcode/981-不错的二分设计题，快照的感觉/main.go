package main

// https://leetcode.cn/problems/time-based-key-value-store/?envType=study-plan&id=binary-search-basic&plan=binary-search&plan_progress=bv79h7h
// 981. 基于时间的键值存储

// get 某个 key 时，会根据 timestamp 去取value，只取最大的 小于等于 timestamp 的值，若不存在，则返回空字符串
// 显然，key 对应的 timestamp 需要是一个数组，方便进行二分
// key : timestamp数组，并且该数组需要为升序
// 每个 timestamp 和 value 是一个键值对
// 因此我们下面需要两个 map，一个 map 存储 key-time数组，一个 map 存储 key-(time-value)
type TimeMap struct {
	keyToTimestamp         map[string][]int          // key-timestamp
	keyToTimestampMapValue map[string]map[int]string // key-(timestamp-value)
}

func Constructor() TimeMap {
	return TimeMap{
		keyToTimestamp:         make(map[string][]int),
		keyToTimestampMapValue: make(map[string]map[int]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keyToTimestamp[key] = append(this.keyToTimestamp[key], timestamp)
	if _, ok := this.keyToTimestampMapValue[key]; !ok {
		this.keyToTimestampMapValue[key] = make(map[int]string)
	}
	this.keyToTimestampMapValue[key][timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	l, r := -1, len(this.keyToTimestamp[key])
	for l+1 != r {
		mid := l + (r-l)/2
		if this.keyToTimestamp[key][mid] <= timestamp {
			l = mid
		} else {
			r = mid
		}
	}
	// 此时 l 会落在最大的小于等于 timestamp 的时间戳下标上
	// 当不存在小于等于 timestamp 时返回空字符串
	if l == -1 {
		return ""
	}
	return this.keyToTimestampMapValue[key][this.keyToTimestamp[key][l]]
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
