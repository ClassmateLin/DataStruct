package main

import "fmt"

// 比较函数
func compare(a, b int) bool {
	return  a > b
}

// 冒泡排序, comp传入一个闭包, 作为是否交换值的条件判断
func BubbleSort(arr []int, comp func(a, b int) bool) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 0; i < n; i++ {

		flag := false // 提前退出标志位

		for j := 0; j < n - i - 1; j++ {
			if comp(arr[j], arr[j+1])  {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag { // 如果没有数据交换，可以提前退出操作
			break
		}
	}
	return arr
}

func main() {
	arr := []int{1,5,3,2,5,7,6,113,44,-1}
	arr = BubbleSort(arr, compare)
	fmt.Println(arr)
}
