package main

import "fmt"

// 循环队列结构定义
type CircularQueue struct {
	q []interface{}
	capacity int
	head int
	tail int
}

// 新建队列
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{
		q: make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}


// 队列判空
func (this *CircularQueue) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}


// 队列判满
func (this *CircularQueue) IsFull() bool {
	if this.head == (this.tail + 1) % this.capacity {
		return true
	}
	return false
}

// 入队操作
func (this *CircularQueue) EnQueue(v interface{}) bool {
	if this.IsFull() {
		return false
	}
	this.q[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	return true
}

// 出队操作
func (this *CircularQueue) DeQueue() interface{}  {
	if this.IsEmpty() {
		return false
	}
	v := this.q[this.head]
	this.head = (this.head+1) % this.capacity
	return v
}

// 队列的遍历
func (this *CircularQueue) String() string {
	if this.IsEmpty() {
		return "empty queue!"
	}
	result := "head"
	var i = this.head
	for {
		result += fmt.Sprintf("<-%+v", this.q[i])
		i = (i + 1) % this.capacity
		if i == this.tail {
			break
		}
	}
	result += "<-tail"
	return result
}