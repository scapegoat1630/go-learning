package heap

import (
	"container/heap"
	"sort"
)

// HeapSort 堆排序
func HeapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一个个从堆顶取出元素
	for i := n - 1; i >= 0; i-- {
		// 移动当前根到末尾
		arr[i], arr[0] = arr[0], arr[i]

		// 调用 max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

// heapify 将以 i 为根的子树调整为最大堆
func heapify(arr []int, n int, i int) {
	largest := i // 初始化最大为根
	l := 2*i + 1 // 左子节点
	r := 2*i + 2 // 右子节点

	// 如果左子节点大于根
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	// 如果右子节点大于当前的最大值
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// 如果最大值不是根
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换

		// 递归地堆化受影响的子树
		heapify(arr, n, largest)
	}
}

func findKthLargest(nums []int, k int) int {
	return 0
}

func quickSort(nums []int, l, r, k int) int {
	return 0
}

// Hp 默认实现是小顶堆，即第一个元素为最小值
type Hp struct {
	sort.IntSlice
}

func (h *Hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *Hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type MedianFinder struct {
	// 大顶堆
	QueMin Hp
	// 小顶堆
	QueMax Hp
}

func Constructor() MedianFinder {
	return MedianFinder{
		QueMin: Hp{},
		QueMax: Hp{},
	}
}
func (m *MedianFinder) AddNum(num int) {
	if m.QueMin.Len() == 0 || num <= -m.QueMin.IntSlice[0] {
		heap.Push(&m.QueMin, -num)
		if m.QueMin.Len() > m.QueMax.Len()+1 {
			heap.Push(&m.QueMax, -heap.Pop(&m.QueMin).(int))
		}
	} else {
		heap.Push(&m.QueMax, num)
		if m.QueMax.Len() > m.QueMin.Len() {
			heap.Push(&m.QueMin, -heap.Pop(&m.QueMax).(int))
		}
	}
}

func (m *MedianFinder) GetMedian() float64 {
	if m.QueMin.Len() > m.QueMax.Len() {
		return float64(-m.QueMin.IntSlice[0])
	} else {
		return float64(m.QueMax.IntSlice[0]-m.QueMin.IntSlice[0]) / 2
	}
}
