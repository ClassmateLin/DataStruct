package main

import "fmt"


// 归并排序
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	// 分治法将数组拆分
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end) // 合并数组
}

// 合并有序数组
func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0

	// 双指针思想进行数组合并
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}

func main()  {
	arr := []int{1,3,5,2,4,6,7}
	MergeSort(arr)
	fmt.Println(arr)
}