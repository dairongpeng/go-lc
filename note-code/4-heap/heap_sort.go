package main

// HeapSort 堆排序额外空间复杂度O(1)
func HeapSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 原始版本, 调整arr满足大根堆结构。O(N*logN)
	//for i := 0; i < len(arr); i++ { // O(N)
	//	heapInsert(arr, i) // O(logN)
	//}

	// 优化版本：heapInsert改为heapify。从末尾开始看是否需要heapify=》O(N)复杂度。
	// 但是这只是优化了原有都是构建堆（O(NlogN)），最终的堆排序仍然是O(NlogN)。比原始版本降低了常数项
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	// 实例化一个大根堆,此时arr已经是调整后满足大根堆结构的arr
	mh := MaxHeap{
		heap:     arr,
		limit:    len(arr),
		heapSize: len(arr),
	}

	mh.heapSize--
	swap(arr, 0, mh.heapSize)
	// O(N*logN)
	for mh.heapSize > 0 { // O(N)
		heapify(arr, 0, mh.heapSize) // O(logN)
		mh.heapSize--
		swap(arr, 0, mh.heapSize) // O(1)
	}

}
