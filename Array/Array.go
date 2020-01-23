package main

import (
	"errors"
	"fmt"
)

type Array struct {
	data []int
	len uint
}


// 为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data: make([]int, capacity, capacity),
		len:  0,
	}
}

// 返回数组长度
func (this *Array) Len() uint {
	return this.len
}


// 判断数组是否越界
func (this *Array) isIndexOutOfRange(idx uint) bool {
	if idx >= uint(cap(this.data)) {
		return true
	}
	return false
}

// 通过索引访问数组元素
func (this *Array) Find(idx uint) (int, error)  {
	if this.isIndexOutOfRange(idx) {
		return 0, errors.New("Out of index range!")
	}
	return this.data[idx], nil
}

// 插入数值到索引idx上
func (this *Array) Insert(idx uint, v int) error {
	// 数组已满
	if this.Len() == uint(cap(this.data)) {
		return errors.New("Full array!")
	}

	// 索引越界
	if idx != this.Len() && this.isIndexOutOfRange(idx) {
		return errors.New("Out of index range!")
	}

	// 将idx开始的元素全部向尾部移动一个单位
	for i := this.Len(); i > idx; i-- {
		this.data[i] = this.data[i-1]
	}
	// 插入到idx位置上
	this.data[idx] = v
	this.len++
	return nil
}

// 插入到数组末尾
func (this *Array) InsertToTail(v int) error  {
	return this.Insert(this.Len(), v)
}

// 删除指定索引的元素,并返回其值
func (this *Array) Delete(idx uint) (int, error)  {
	if this.isIndexOutOfRange(idx) {
		return 0, errors.New("Out of index range!")
	}
	v := this.data[idx]

	// 将idx后的元素向前移动一个单位, 删除idx上的元素
	for i := idx; i < this.Len()-1; i ++ {
		this.data[i] = this.data[i+1]
	}
	this.len--
	return v, nil
}


// 打印数组的元素
func (this *Array) Print() {
	var format string

	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}


func main() {
	arr:= NewArray(10)
	for i:=0; i<10; i++ {
		_ = arr.Insert(uint(i), i * i)
	}
	arr.Print()
	arr.Delete(1)
	arr.Print()
	arr.Delete(2)
	arr.Print()
	_ = arr.Insert(5, 10)
	arr.Print()
	val, _ := arr.Find(1)
	fmt.Println(val)
}
