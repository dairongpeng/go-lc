package main

import "fmt"

// 有序数组二分查找
func exist(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	var L = 0
	var R = len(arr) - 1
	var mid = 0
	for L < R {
		// 防止整数越界
		mid = L + (R-L)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] < target { // mid位置已经比较过，L置为mid+1
			L = mid + 1
		} else if arr[mid] > target { // mid位置已经比较过，R置为mid-1
			R = mid - 1
		}
	}
	return arr[L] == target
}

// 在一个有序数组上，找到大于等于value的最左位置
func nearestLeftIndex(arr []int, value int) int {
	L := 0
	R := len(arr) - 1
	index := -1
	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] >= value {
			index = mid
			R = mid - 1
		} else { // 大于等于时
			L = mid + 1
		}
	}
	return index
}

// 在一个有序数组上，找到小于等于value的最右位置
func nearestRightIndex(arr []int, value int) int {
	L := 0
	R := len(arr) - 1
	index := -1
	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] <= value {
			index = mid
			L = mid + 1
		} else {
			L = mid - 1
		}
	}
	return index
}

// IsoOr 异或交换数据
func IsoOr() {
	a := 3
	b := 4
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a)
	fmt.Println(b)

	arr := []int{3, 1, 100}
	IsoOrSwap(arr, 0, 3)
	fmt.Println(arr)

	// i和j指向同一块内存，这种位运算交换变量的方法就不可行了。
	IsoOrSwap(arr, 0, 0)
}

func IsoOrSwap(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
