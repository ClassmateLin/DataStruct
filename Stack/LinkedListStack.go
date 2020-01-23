package main

import "fmt"

/*
基于链表实现的栈
 */



type Node struct {
	Next *Node
	Val interface{}
}

type LinkedListStack struct {
	TopNode *Node // 栈顶结点
}

// 创建栈
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

// 判断栈是否为空栈
func (this *LinkedListStack) IsEmpty() bool  {
	return this.TopNode == nil
}

// 出栈操作
func (this *LinkedListStack) Pop() interface{}  {
	if this.IsEmpty(){
		return nil
	}
	v := this.TopNode.Val
	this.TopNode = this.TopNode.Next
	return v
}

// 返回栈顶元素
func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty(){
		return nil
	}
	return this.TopNode.Val
}


// 清空栈
func (this *LinkedListStack) Flush() {
	this.TopNode = nil
}

// 栈的遍历
func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("Empty stack!")
	}else {
		cur := this.TopNode
		for cur != nil {
			fmt.Println(cur.Val)
			cur = cur.Next
		}
	}
}

func main() {
	
}
