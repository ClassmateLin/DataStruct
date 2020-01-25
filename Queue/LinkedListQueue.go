package main

import "fmt"

// 队列结点
type QueNode struct {
	Val interface{}
	Next *QueNode
}


// 链式队列
type LinkedListQueue struct {
	Head *QueNode
	Tail *QueNode
	Length int
}

// 新建队列
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}


// 入队
func (this LinkedListQueue) EnQueue(v interface{}) {
	node := &QueNode{v, nil}

	if this.Tail == nil {
		this.Tail = node
		this.Head = node
	}else {
		this.Tail.Next = node
		this.Tail = node // 队尾始终指向队尾元素
	}
	this.Length++
}

// 出队
func (this *LinkedListQueue) Dequeue() interface{} {
	if this.Head == nil {
		return nil
	}
	v := this.Head.Val
	this.Head = this.Head.Next
	this.Length--
	return v
}

// 判断队列是否为空
func (this *LinkedListQueue) Empty() bool {
	return this.Tail == nil
}

// 返回队列元素
func (this *LinkedListQueue) String() string {
	if this.Head == nil {
		return "Empty queue！"
	}
	result := "head<-"
	for cur := this.Head; cur != nil; cur = cur.Next {
		result += fmt.Sprintf("<-%+v", cur.Val)
	}
	result += "<-tail"
	return result
}