package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Left, Right *Node
	Data interface{}
}

/**
通过传入的切片来进行创建二叉树
 */
func (node *Node) CreateFromSlice(data []interface{}) *Node {

	if node == nil {
		return nil
	}
	var bin *Node
	if len(data) < 2 {
		return nil
	}
	bin = new(Node)
	bin.Data = data[0]  // 设置当前结点的值
	bin.Left = bin.CreateFromSlice(data[1:]) // 递归创建左子树
	bin.Right = bin.CreateFromSlice(data[2:]) // 递归创建右子树
	return bin
}

/**
通过从键盘输入的方式来创建二叉树
 */
func (node *Node) CreateFromInput() *Node {
	reader := bufio.NewReader(os.Stdin) //读取输入的内容
	fmt.Print("请输入结点数字(输入-1结束):")
	in, err := reader.ReadString('\n')
	fmt.Println()
	if err != nil {
		fmt.Println("获取输入出错, 程序退出!")
		os.Exit(1)
	}
	in = strings.Replace(in, "\n", "", 1) //　输入in带有\n, 无法转数字,先去除
	num, err := strconv.Atoi(in)  // 转换为数字类型

	if err != nil  {
		fmt.Println("输入非法, 程序退出!", err.Error())
		os.Exit(1)
	}

	if num == -1 {
		 return nil
	}

	var bin *Node
	bin = new(Node)
	bin.Data = num
	fmt.Printf("添加结点: {%d} 左子树:\n", bin.Data)
	bin.Left = bin.CreateFromInput()
	fmt.Printf("添加结点: {%d} 右子树:\n", bin.Data)
	bin.Right = bin.CreateFromInput()
	return bin
}

/**
先序遍历:  根->左->右
先打印当前结点值,
再递归调用左子树的先序遍历方法
接着递归调用右子树的先序遍历方法
 */
func (node *Node) PrevSort() {
	if node == nil {
		return
	}

	fmt.Print(node.Data, " ")
	node.Left.PrevSort()
	node.Right.PrevSort()
}

/**
中序遍历: 左->根->右
 */
func (node *Node) MidSort() {
	if node == nil {
		return
	}
	node.Left.MidSort() // 左子树递归调用此方法
	fmt.Print(node.Data, " ") // 中, 打印结点
	node.Right.MidSort() // 右子树递归调用此方法
}

/**
后序遍历: 左->右->根
 */
func (node *Node) PostSort() {
	if node == nil {
		return
	}
	node.Right.PostSort()
	node.Right.PostSort()
	fmt.Print(node.Data, " ")
}

func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	leftHeight := node.Left.Height() // 左子树递归进入
	rightHeight := node.Right.Height() // 右子树递归进入

	if leftHeight > rightHeight { // 左右子树进行比较, 累加递归返回值
		leftHeight++
		return leftHeight
	}else {
		leftHeight++
		return leftHeight
	}
}

/**
Go不支持函数默认值和Static且不想使用全局遍历,　手动传入Num为0!
 */
func (node *Node) LeafNum(num int)  {
	if node == nil { // 递归出口
		return
	}

	if node.Left == nil && node.Right == nil {
		num++
	}

	node.Left.LeafNum(num)
	node.Right.LeafNum(num)
}



func main()  {
	//nums := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12}
	node := new(Node)
	fmt.Println("创建根节点: ")
	node = node.CreateFromInput()
	fmt.Printf("先序遍历: \n\n")
	node.PrevSort()
	fmt.Println(node.Height())
}