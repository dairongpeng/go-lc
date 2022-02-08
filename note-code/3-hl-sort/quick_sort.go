package main

// swap 交换数组中的两个位置的数
func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// partition 对数组进行partition处理
func partition(arr []int, L, R int) int {
	if L > R {
		return -1
	}
	if L == R {
		return L
	}
	// 选定左边界的左边一个位置作为小于区域的起点
	lessEqual := L - 1
	index := L
	// 每次搞定一个位置
	for index < R {
		if arr[index] <= arr[R] {
			lessEqual++
			swap(arr, index, lessEqual)
		}
		index++
	}
	lessEqual++
	swap(arr, lessEqual, R)
	return lessEqual
}

//  arr[L...R] 玩荷兰国旗问题的划分，以arr[R]做划分值
//  小于arr[R]放左侧  等于arr[R]放中间  大于arr[R]放右边
//  返回中间区域的左右边界
func netherlandsFlag(arr []int, L, R int) []int {
	// 不存在荷兰国旗问题
	if L > R {
		return []int{-1, -1}
	}

	// 已经都是等于区域，由于用R做划分返回R位置
	if L == R {
		return []int{L, R}
	}

	// < 区 右边界
	less := L - 1
	// > 区 左边界
	more := R
	index := L
	for index < more {
		// 当前值等于右边界，不做处理，index++
		if arr[index] == arr[R] {
			index++
		} else if arr[index] < arr[R] { // 小于交换当前值和左边界的值
			less++
			swap(arr, index, less)
			index++
		} else { // 大于右边界的值
			more--
			swap(arr, index, more)
		}
	}
	// 比较完之后，把R位置的数，调整到等于区域的右边，至此大于区域才是真正意义上的大于区域
	swap(arr, more, R)
	return []int{less + 1, more}
}

func QuickSort1(arr []int) {
	if len(arr) < 2 {
		return
	}

	sortByPartition(arr, 0, len(arr)-1)
}

func sortByPartition(arr []int, L int, R int) {
	if L >= R {
		return
	}

	// L到R上进行partition 标记位为arr[R] 数组被分成  [   <=arr[R]   arr[R]    >arr[R]  ]，M为partition之后标记位处在的位置
	M := partition(arr, L, R)
	sortByPartition(arr, L, M-1)
	sortByPartition(arr, M+1, R)
}

func QuickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	sortByNetherlandsFlag(arr, 0, len(arr)-1)
}

func sortByNetherlandsFlag(arr []int, L int, R int) {
	if L >= R {
		return
	}

	// 每次partition返回等于区域的范围,荷兰国旗问题
	equalArea := netherlandsFlag(arr, L, R)
	// 对等于区域左边的小于区域递归，partition
	sortByNetherlandsFlag(arr, L, equalArea[0]-1)
	// 对等于区域右边的大于区域递归，partition
	sortByNetherlandsFlag(arr, equalArea[1]+1, R)
}
