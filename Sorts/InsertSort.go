package main

import "fmt"

// 插入排序
func InsertSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		// 查找需要插入的位置并移动数据
		for ; j>=0; j-- { // 往左边找一个比当前value大的元素
			if arr[j] > value {
				arr[j+1] = arr[j]
			}else { // 找到了
				break
			}
		}
		arr[j+1] = value // 将元素插入找到位置中
	}
	return arr
}

func main() {
	arr := []int{1, 3, 2, 4, 7, 5,9}
	arr = InsertSort(arr)
	fmt.Println(arr)
}
