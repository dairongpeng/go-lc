package main

import (
	"errors"
)

type Heap interface {
	IsEmpty() bool
	IsFull() bool
	Push(value int) error
	Pop() int
}

func assertListImplementation() {
	var _ Heap = (*MaxHeap)(nil)
}

type MaxHeap struct {
	// 大根堆的底层数组结构
	heap []int
	// 分配给堆的空间限制
	limit int
	// 表示目前这个堆收集了多少个数，即堆大小。也表示添加的下一个数应该放在哪个位置
	heapSize int
}

// NewMaxHeap 初始化一个大根堆结构
func NewMaxHeap(limit int) *MaxHeap {
	maxHeap := &MaxHeap{
		heap:     make([]int, 0),
		limit:    limit,
		heapSize: 0,
	}
	return maxHeap
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *MaxHeap) IsFull() bool {
	return h.heapSize == h.limit
}

func (h *MaxHeap) Push(value int) error {
	if h.heapSize == h.limit {
		return errors.New("heap is full")
	}

	h.heap[h.heapSize] = value
	// heapSize的位置保存当前value
	heapInsert(h.heap, h.heapSize)
	h.heapSize++
	return nil
}

// Pop 返回堆中的最大值，并且在大根堆中，把最大值删掉。弹出后依然保持大根堆的结构
func (h *MaxHeap) Pop() int {
	tmp := h.heap[0]
	h.heapSize--
	swap(h.heap, 0, h.heapSize)
	heapify(h.heap, 0, h.heapSize)
	return tmp
}

// 往堆上添加数，需要从当前位置找父节点比较
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// 从index位置，不断的与左右孩子比较，下沉。下沉终止条件为：1. 左右孩子都不大于当前值 2. 没有左右孩子了
func heapify(arr []int, index int, heapSize int) {
	// 左孩子的位置
	left := index*2 + 1
	// 左孩子越界，右孩子一定越界。退出循环的条件是：2. 没有左右孩子了
	for left < heapSize {
		var largestIdx int
		rigth := left + 1
		// 存在右孩子，且右孩子的值比左孩子大，选择右孩子的位置
		if rigth < heapSize && arr[rigth] > arr[left] {
			largestIdx = rigth
		} else {
			largestIdx = left
		}

		// 1. 左右孩子的最大值都不大于当前值，终止寻找。无需继续下沉
		if arr[largestIdx] <= arr[index] {
			break
		}
		// 左右孩子的最大值大于当前值
		swap(arr, largestIdx, index)
		// 当前位置移动到交换后的位置，继续寻找
		index = largestIdx
		// 移动后左孩子理论上的位置，下一次循环判断越界情况
		left = index*2 + 1
	}
}

// swap 交换数组中的两个位置的数
func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
