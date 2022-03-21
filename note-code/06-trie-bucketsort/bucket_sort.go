package main

import "math"

// BucketSort 计数排序
func BucketSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	max := math.MinInt
	for i := 0; i < len(arr); i++ {
		max = int(math.Max(float64(max), float64(arr[i])))
	}

	bucket := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	k := 0
	for i := 0; i < len(bucket); i++ {
		bucket[i]--
		for bucket[i] > 0 {
			arr[k] = i
			k++
		}
	}
}
