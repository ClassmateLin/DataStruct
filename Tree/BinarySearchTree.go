package main

// Binary Search Tree
type BST struct {
	// Data interface{}
	Val int
	Left *BST
	Right *BST
}

// BST插入操作递归操作
func (node *BST) Insert(val int) {
	if val < node.Val {
		if node.Left != nil {
			node.Left.Insert(val)
		}else {
			node.Left = &BST{Val:val}
		}
	}else {
		if node.Right != nil {
			node.Right.Insert(val)
		}else {
			node.Right = &BST{Val:val}
		}
	}
}



// BST查找操作递归实现
func (node *BST) SearchRecursion(key int) bool {
	if node == nil {
		return false
	}else if key == node.Val { // 找到结果
		return true
	}else if key > node.Val { // 查找的值大于当前节点在右子树找
		return node.Right.SearchRecursion(key)
	}else if key < node.Val { // 查找的值比当前节点小，在左子树上查找
		return node.Left.SearchRecursion(key)
	}
	return false
}

// BST查找操作迭代实现
func (node *BST) SearchIteration(key int) bool  {
	tree := node
	for {
		if tree == nil{
			return false
		}else if tree.Val == key {
			return true
		}else if tree.Val > key{
			tree = tree.Left
		}else {
			tree = tree.Right
		}
	}
}

// BST最小值
func (node *BST) Min() int {
	tree := node
	for {
		if tree.Left != nil {
			tree = tree.Left
		}else {
			return tree.Val
		}
	}
}


// BST最大值
func (node *BST) Max() int {
	tree := node
	for {
		if tree.Right != nil {
			tree = tree.Right
		}else {
			return tree.Val
		}
	}
}

// BST删除操作
func (node *BST) Remove(key int) *BST {
	if node == nil {
		return nil
	}
	// 向左找
	if key < node.Val {
		return node.Left.Remove(key)
	}

	if key > node.Val {
		return node.Right.Remove(key)
	}

	// 如果该节点没左右子树直接删除
	if node.Left == nil && node.Right == nil {
		node = nil
		return node
	}

	// 有右子树直接用右子树覆盖该节点
	if node.Left == nil {
		node = node.Right
		return node
	}

	// 有左子树直接用左子树覆盖该节点
	if node.Right != nil {
		node = node.Left
		return node
	}
	// 左右子树均存在， 找到右子树的最左节点替换当前节点
	mostLeftMode := node.Right
	for {
		if mostLeftMode != nil && mostLeftMode.Left != nil {
			mostLeftMode = mostLeftMode.Left
		}else {
			break
		}
	}
	node.Val = mostLeftMode.Val

	// 删除右子树的最左节点
	node.Right = node.Right.Remove(node.Val)
	return node
}


func main() {

}
