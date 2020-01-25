package main

import "fmt"

func SelectionSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 0; i < n; i++ {
		minIdx := i  // 记录最小值的索引

		// 查找最小值的索引
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// 当前值与最小值进行交换
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}


func main() {
	arr := []int{22,44,11,13,67, -1,9,3}
	arr = SelectionSort(arr)
	fmt.Println(arr)
}
