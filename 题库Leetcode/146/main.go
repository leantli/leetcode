package main

// https://leetcode.cn/problems/lru-cache/description/
// 146. LRU 缓存

// 首先，存在或插入的话，要 O(1) 复杂度返回, 显然需要 Map
// 插入的话，超过 cap 需要逐出最久未使用的节点, 而最近是否使用过，显然我们需要使用非 map 的结构去维护存储
// 比如说链表，最近使用过(不过是 get 还是 insert 还是 update)，则取出放到链表尾部
// 此时最久未使用的在链表首部，删除时只需要删除头结点后一个节点即可
type LRUCache struct {
	m        map[int]*LinkListNode
	head     *LinkListNode
	tail     *LinkListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &LinkListNode{}
	tail := &LinkListNode{}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkListNode), head, tail, capacity}
}

func (this *LRUCache) Get(key int) int {
	// 搜得到就返回，否则返回 -1
	if node, ok := this.m[key]; ok {
		// 更新链表的最近使用情况
		this.del(node)
		this.add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 如果 key 存在，不管满不满都能进行操作
	if node, ok := this.m[key]; ok {
		node.val = value
		// 更新链表的最近使用情况
		this.del(node)
		this.add(node)
		return
	}
	// 不存在则需要插入，判断 capacity 满了没
	if len(this.m) == this.capacity {
		// 已经满了，需要做清除操作
		this.del(this.head.next)
	}
	this.add(&LinkListNode{key, value, nil, nil})
}

type LinkListNode struct {
	key, val int
	next     *LinkListNode
	pre      *LinkListNode
}

// 不管是更新还是get，节点都是插入到尾部
func (this *LRUCache) add(node *LinkListNode) {
	// 将节点加入到链表尾部并且更新 tail 指针指向新的尾部
	node.pre = this.tail.pre
	this.tail.pre.next = node
	node.next = this.tail
	this.tail.pre = node
	this.m[node.key] = node
}

// 删除一个节点
func (this *LRUCache) del(node *LinkListNode) {
	// 先将节点从链表中删除，连接该节点的前后节点
	node.pre.next = node.next
	node.next.pre = node.pre
	delete(this.m, node.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
