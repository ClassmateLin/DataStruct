package main

import "fmt"

func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int)  {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}


func partition(arr []int, start, end int) int  {
	pivot := arr[end] // 选取最后一位当基准数据
	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func main()  {
	arr := []int{1,4,6,2,3,9,0}
	QuickSort(arr)
	fmt.Println(arr)
}