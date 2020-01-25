package main

import "fmt"

type ArrayQueue struct {
	q []interface{}
	capacity int // 队列容量
	head int // 队头指针
	tail int // 队尾指针
}

// 创建队列
func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// 入队操作
func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.tail == this.capacity { // 队列已满
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}


// 出队操作
func (this *ArrayQueue) DeQueue() interface{}  {
	if this.head == this.tail { // 队列已空
		return false
	}
	v := this.q[this.head]
	this.head++
	return v
}

// 队列不为空返回所有元素，否则返回空
func (this *ArrayQueue) String() string  {
	if this.head == this.tail {
		return  "empty queue!"
	}
	result := "head"
	for i := this.head; i <= this.tail - 1; i++ {
		result += fmt.Sprintf("<-%+v", this.q[i])
	}
	result += "<-tail"
	return result
}


// 判断队列是否为空
func (this *ArrayQueue) Empty() bool {
	if this.head == this.tail {
		return true
	}else {
		return false
	}
}

// 返回队列长度
func (this *ArrayQueue) Len() int {
	return (this.tail - this.head + this.capacity) % this.capacity
}

// 返回队头元素
func (this *ArrayQueue) GetHead() interface{} {
	if this.Empty() {
		return nil
	}
	return this.q[this.head]
}


// 判断队列是否已满
func (this *ArrayQueue) Full() bool {
	if this.tail == this.head {
		return true
	}
	return false
}


// 清空队列
func (this *ArrayQueue)  Clear() {
	this.head = 0
	this.tail = 0
}
