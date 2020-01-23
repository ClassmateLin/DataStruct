package main

import "fmt"

/*
单链表的实现
 */

type SinListNode struct {
	next *SinListNode
	val int
}

type SinLinkedList struct {
	head *SinListNode
	len uint
}


// 创建结点
func NewListNode(v interface{}) *SinListNode {
	return &SinListNode{nil, v}
}

// 获取链表下一个节点
func (this *SinListNode) GetNext() *SinListNode  {
	return this.next
}

// 获取节点值
func (this *SinListNode) GetValue() int  {
	return this.val
}

// 创建链表
func NewLinkedList() *SinLinkedList  {
	return &SinLinkedList{NewListNode(0), 0}
}


// 在指定节点后面插入节点
func (this *SinLinkedList) InsertAffter(p *SinListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	// 创建一个新节点
	newNode, oldNext := NewListNode(v), p.next
	// 新节点插入到指定节点后，原节点的下一节点插入到新节点的下一节点
	p.next, newNode.next = newNode, oldNext
	this.len++
	return true
}


// 在某个节点前面插入节点
func (this *SinLinkedList) InsertBefore(p *SinListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}

	cur := this.head.next // 记录当前节点
	pre := this.head // 记录前缀节点

	// 查找p的前缀节点
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode  // 前缀节点指向新节点
	newNode.next = cur // 新节点指向原来的节点P
	this.len++
	return true
}

// 在链表头部插入节点
func (this *SinLinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAffter(this.head, v)
}

// 在链表尾部插入节点
func (this *SinLinkedList) InsertToTail(v interface{}) bool  {
	cur := this.head

	// 查找尾节点
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAffter(cur, v)
}


// 通过索引查找节点
func (this *SinLinkedList) FindByIndex(idx uint) *SinListNode {
	if idx >= this.len {
		return nil
	}
	cur := this.head.next

	// 移动节点值idx位置
	var i uint = 0
	for ; i < idx; i++ {
		cur = cur.next
	}

	return cur
}

func (this *SinLinkedList) DeleteNode(p *SinListNode) bool {
	if p == nil {
		return false
	}
	cur := this.head.next
	pre := this.head

	// 查找p的前驱节点
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 找不到p节点
	if cur == nil {
		return false
	}
	// p的前驱节点指向p的后继节点
	pre.next = p.next
	p = nil // 释放p节点
	this.len--
	return true
}

// 打印链表
func (this *SinLinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}


// 单链表翻转,时间复杂度O(n)
func (this *SinLinkedList)Reverse()  {
	if this.head == nil || this.head.next == nil || this.head.next.next == nil {
		return
	}
	var pre *SinListNode = nil

	cur := this.head.next

	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre, cur = cur, tmp
	}
	this.head.next = pre
}


// 快慢指针判断链表是否有环
func (this *SinLinkedList) HasCycle() bool {
	if this.head != nil {
		slow := this.head
		fast := this.head

		if fast != nil && fast.next != nil{
			slow = slow.next // 慢指针一次移动一个节点
			fast = fast.next.next  // 快指针一次移动两个节点
			if slow == fast { // 指针相遇证明有环
				return true
			}
		}
	}
	return false
}


// 合并两个有序单链表
func MergeSortedList(l1, l2 *SinLinkedList) *SinLinkedList {
	// 链表1为空直接返回链表2
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}

	// 链表2为空直接返回链表1
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}
	l := &SinLinkedList{head:&SinListNode{}}
	cur := l.head
	// 使用双指针思想进行合并
	curl1 := l1.head.next
	curl2 := l2.head.next

	for curl1 != nil && curl2 != nil {

		if curl1.GetValue() > curl2.GetValue() {
			cur.next = curl2
			curl2 = curl2.next
		}else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}
	if curl1 != nil {
		cur.next = curl1
	}else if curl2!= nil {
		cur.next = curl2
	}
	return l
}

// 删除倒数第N个节点
func (this *SinLinkedList) DeleteBottomN(n int)   {
	if n <= 0 || this.head == nil || this.head.next == nil {
		return
	}
	fast := this.head

	// 查找第N个节点
	for i := 1; i <= n && fast!= nil; i++ {
		fast = fast.next
	}

	if fast == nil {
		return
	}

	slow := this.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

// 获取链表的中间节点
func (this *SinLinkedList) FindMiddleNode() *SinListNode  {
	if nil == this.head || this.head.next == nil {
		return nil
	}
	// 链表只有一个元素则直接返回该节点
	if this.head.next.next == nil {
		return this.head.next
	}

	slow, fast := this.head, this.head

	// 使用快慢指针, 快指针移动两个节点，慢指针移动一个节点，当快指针遍历结束,那么慢指针指向的节点为中间节点
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}