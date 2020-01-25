package main

import (
	"fmt"
	"math"
)

func CountingSort(arr []int, n int)  {
	if n <= 1 {
		return
	}

	var max int = math.MinInt32

	// 求出数组中的最大值
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}

	c := make([]int, max+1)

	for i:= range arr {
		c[arr[i]]++
	}

	for i:=1;i<=max;i++ {
		c[i] += c[i-1]
	}

	r := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		idx := c[arr[i]] - 1
		r[idx] = arr[i]
		c[arr[i]]--
	}
	copy(arr, r)
}


func main() {
	arr := []int{1,3,2,4,6,8,4,6}
	CountingSort(arr, len(arr))
	fmt.Println(arr)
}
