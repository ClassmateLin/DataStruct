package main

import "fmt"

type ArrayStack struct {
	data []interface{}
	top int
}

func NewArrayStack() *ArrayStack  {
	return &ArrayStack{
		data:make([]interface{}, 0, 32),
		top: -1,
	}
}

// 判断栈是否为空
func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

// 入栈操作
func (this *ArrayStack) Push(v interface{})  {
	if this.top < 0 {
		this.top = 0
	}else {
		this.top += 1
	}
	// 栈空间已满，添加空间
	if this.top > len(this.data) - 1 {
		this.data = append(this.data, v)
	}else { // 直接入栈
		this.data[this.top] = v
	}
}

// 出栈操作
func (this *ArrayStack) Pop() interface{}  {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.top]
	this.top--
	return v
}

// 取栈顶元素
func (this *ArrayStack) Top() interface{}  {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.top]
}

// 清空栈
func (this *ArrayStack) Flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("Empty stack!")
	}else {
		for i := this.top; i>=0; i++ {
			fmt.Println(this.data[i])
		}
	}
}

func main() {
	
}
