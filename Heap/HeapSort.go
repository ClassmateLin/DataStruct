package main

import "fmt"

func buidHeap(a []int, n int) {

	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}

}

// 排序索引从1开始的n个元素, 索引为0的元素不参与排序
func sort(a []int, n int) {
	buidHeap(a, n)

	k := n
	for k >= 1 { // 堆顶移动到最后, 节点进行下沉操作
		a[1], a[k] = a[k], a[1]
		heapifyUpToDown(a, 1, k-1)
		k--
	}
}

// 节点下城
func heapifyUpToDown(a []int, top int, count int) {

	for i := top; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		a[i], a[maxIndex] = a[maxIndex], a[i]
		i = maxIndex
	}

}


func main() {
	// 二叉堆排序索引从1开始的n个元素, 数组中的0不参与排序
	arr := []int{0, 2,4,6,8,1, 11}
	sort(arr, len(arr)-1)
	fmt.Print(arr[1:])
}
