package main

type Heap struct {
	a []int
	n int
	count int
}

func NewHeap(capacity int) *Heap  {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity+1)
	heap.count = 0
	return heap
}

func (heap *Heap) Insert(data int) {
	if heap.count == heap.n { // 堆满
		return
	}
	heap.count++
	heap.a[heap.count] = data // 插入到最后

	// 与父节点比较，节点上浮
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		heap.a[i], heap.a[parent] = heap.a[parent], heap.a[i]
		i = parent
		parent = i / 2
	}
}


// 删除最大元素(堆顶), 节点需要进行下沉操作
func (heap *Heap) removeMax() {
	if heap.count == 0 {
		return
	}
	
	// 最大元素和最后一个元素进行交换
	heap.a[1], heap.a[heap.count] = heap.a[heap.count], heap.a[i]
	heap.count--

	// 节点下沉
	heapifyUpDown(heap.a, heap.count)
}

func heapifyUpDown(a []int, count int)  {
	for i := 1; i <= count/2; {
		maxIdx := i
		if a[i] < a[i*2] {
			maxIdx = i * 2
		}
		if i * 2 + 1 <= count && a[maxIdx] < a[i*2+1] {
			maxIdx = i * 2 + 1
		}
		if maxIdx == i {
			break
		}
		a[i], a[maxIdx] = a[maxIdx], a[i]
		i = maxIdx
	}
}


func main() {
	
}
