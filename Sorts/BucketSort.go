package main

import "fmt"

// 获取待排序数组中最大值
func getMax(a []int)int {
	max := a[0]

	for i:=1; i<len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}


// 桶排序
func BucketSort(source []int)  {
	if len(source) <= 1 {
		return
	}
	array := make([]int, getMax(source)+1) // 创建一个桶
	for i:=0; i<len(source);i++{ // 桶计各个桶中元素出现次数
		array[source[i]]++
	}
	c := make([]int, 0)
	for i:=0; i < len(array); i++ {
		for array[i] != 0 { // 待排序数组可能有多个相同元素， 这里要根据统计值进行添加
			c = append(c,i)
			array[i]--
		}
	}
	copy(source, c)
}

func main() {
	arr := []int{2,1, 1,1,4,3,2,3,4}
	BucketSort(arr)
	fmt.Println(arr)
}
