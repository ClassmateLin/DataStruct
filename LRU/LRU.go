package main

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH = 100
)

// LRU节点结构体定义
type LRUNode struct {
	prev *LRUNode  // 前驱节点
	next *LRUNode // 后继节点

	key int // 键
	val int // 值

	hnext *LRUNode // 拉链
}

// LRU结构体定义
type LRUCache struct {
	node []LRUNode // 哈希链表

	head *LRUNode // 头结点
	tail *LRUNode // 尾结点

	capacity int // 容量
	used int // 已使用
}

// 初始化
func Constructor(capacity int) LRUCache {
	return LRUCache{
		node:     make([]LRUNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

// 根据键取值
func (this *LRUCache)  Get(key int) int {
	if this.tail == nil {
		return -1
	}
	if tmp:=this.searchNode(key); tmp != nil {
		this.moveToTail(tmp)
		return tmp.val
	}
	return -1
}

/*
插入数据:
1. 首次插入数据
2. 插入数据不在LRU中
3. 插入数据在LRU中
4. 插入数据不在LRU中， 并且LRU已满
 */
func (this *LRUCache) Put(key, val int)  {
	if tmp := this.searchNode(key); tmp != nil {
		tmp.val = val
		this.moveToTail(tmp)
		return
	}
	this.addNode(key, val)

	if this.used > this.capacity { // 最后需要进行判满操作，淘汰节点数据
		this.delNode()
	}
}


// 查找节点
func (this *LRUCache) searchNode(key int) *LRUNode {
	if this.tail == nil {
		return nil
	}

	tmp := this.node[hash(key)].hnext

	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

// 将指定节点移动至末尾
func (this *LRUCache) moveToTail(node *LRUNode) {
	if this.tail == node {
		return
	}
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	}else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	node.next = nil
	this.tail.next = node
	this.tail.prev = this.tail

	this.tail = node
}

// 淘汰节点数据
func (this *LRUCache) delNode()  {
	if this.head == nil {
		return
	}
	prev := &this.node[hash(this.head.key)]
	tmp := prev.hnext
	for tmp != nil && tmp.key != this.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	this.head = this.head.next
	this.head.prev = nil
	this.used--
}


// 哈希函数
func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}else {
		return (key ^ (key) >> 16) & (LENGTH - 1)
	}
}


// 添加节点
func (this *LRUCache) addNode(key int, val int)  {
	newNode := &LRUNode{
		key:key,
		val:val,
	}

	tmp := &this.node[hash(key)]
	newNode.hnext = tmp.hnext
	this.used++

	if this.tail == nil {
		this.tail, this.head = newNode, newNode
	}
	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}


func main() {

}
