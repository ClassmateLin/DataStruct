package main

// 线性查找
func LineSearch(arr[]int, val int) int {
	res := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			res = i
			break
		}
	}
	return res
}


// 二分查找递归实现
func BinarySearchRecursive(arr []int, v, low, hight int)  int {
	if low > hight{
		return -1
	}
	mid := (low+hight) / 2
	if arr[mid] == v {
		return mid
	}else if arr[mid] > v {
		return BinarySearchRecursive(arr, v, low, mid-1)
	}else {
		return BinarySearchRecursive(arr, v, mid+1, hight)
	}
}


// 二分查找非递归实现
func BinaraySearch(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == v {
			return mid
		}else if arr[mid] > v {
			high = mid - 1
		}else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, v int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	low := 0
	high := n-1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] > v {
			high = mid - 1
		}else if arr[mid] < v {
			low = mid + 1
		}else {
			if mid == n - 1 || arr[mid+1] != v {
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}



func main() {
	
}
