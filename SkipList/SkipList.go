package main

import (
	"math"
	"math/rand"
)

// 跳跃表的实现



// 跳跃表的最高层数
const MAX_LEVEL  = 16

// 跳跃表节点结构体
type SkipListNode struct {
	val interface{}  // 保存的值
	score int   // 用于排序的分值
	level int  // 层高
	forwards []*SkipListNode  // 每层前进指针
}

// 跳跃表结构体
type SkipList struct {
	head *SkipListNode
	level int
	length int
}

// 新建跳跃表节点
func NewSkipListNode(v interface{}, score, level int) *SkipListNode {
	return &SkipListNode{
		val:      v,
		score:    score,
		level:    level,
		forwards: nil,
	}
}

// 新建跳跃表对象
func NewSkipList() *SkipList {
	// 头结点方便操作
	head := NewSkipListNode(0, math.MaxInt32, MAX_LEVEL)
	return &SkipList{
		head:  head,
		level: 1,
		length:   0,
	}
}

// 获取跳跃表长度
func (this *SkipList)Length() int {
	return this.length
}

// 获取跳跃表层级
func (this *SkipList)Level() int {
	return this.level
}

// 插入节点到跳跃表中
func (this *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	// 查找插入位置
	cur := this.head

	// 记录每层的路径
	update := [MAX_LEVEL]*SkipListNode{}

	i := MAX_LEVEL - 1

	for ; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].val == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	// 随机选取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31() % 7 == 1 {
			level++
		}
	}

	// 新建跳跃表节点
	newNode := NewSkipListNode(v, score, level)

	// 原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	// 更新跳跃表层数
	if level > this.level {
		this.level = level
	}

	this.length++
	return 0

}

// 跳跃表的查找
func (this *SkipList) Find(v interface{}, score int) *SkipListNode  {
	if v == nil || this.length == 0 {
		return nil
	}
	// 查找的节点
	cur := this.head

	for i := this.level-1; i >=0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].val == v {
				return cur.forwards[i]
			}else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}
	return nil
}

// 删除节点
func (this *SkipList) Delete(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找前驱节点
	cur := this.head

	// 记录前驱路径
	update := [MAX_LEVEL]*SkipListNode{}

	for i:= this.level - 1; i>=0; i-- {
		update[i] = this.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].val == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]

	for i:=cur.level-1; i>=0; i-- {
		if update[i] == this.head && cur.forwards[i] == nil {
			this.level = i
		}

		if update[i].forwards[i] == nil {
			update[i].forwards[i] = nil
		}else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	this.length--
	return 0
}

func main() {
	
}
