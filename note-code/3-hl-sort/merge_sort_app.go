package main

func smallSum(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	return sSum(arr, 0, len(arr)-1)
}

// arr[L..R]既要排好序，也要求小和返回
// 所有merge时，产生的小和，累加
// 左 排序   merge
// 右 排序  merge
// arr 整体 merge
func sSum(arr []int, l, r int) int {
	// 只有一个数，不存在右组，小和为0
	if l == r {
		return 0
	}

	mid := l + (r-l)/2
	// 左侧merge的小和+右侧merge的小和+整体左右两侧的小和
	return sSum(arr, l, mid) + sSum(arr, mid+1, r) + sumMerge(arr, l, mid, r)
}

func sumMerge(arr []int, L, M, R int) int {
	// merge过程申请辅助数组，准备copy
	help := make([]int, 0)
	p1 := L
	p2 := M + 1
	res := 0
	// p1未越界且p2未越界
	for p1 <= M && p2 <= R {
		// 当前的数是比右组小的，产生右组当前位置到右组右边界数量个小和，累加到res。否则res加0
		if arr[p1] < arr[p2] {
			help = append(help, arr[p1])
			res += (R - p2 + 1) * arr[p1]
			p1++
		} else {
			help = append(help, arr[p2])
			res += 0
			p2++
		}
	}

	// p2越界的情况
	for p1 <= M {
		help = append(help, arr[p1])
		p1++
	}

	// p1越界的情况
	for p2 <= R {
		help = append(help, arr[p2])
		p2++
	}

	// 把辅助数组help中整体merge后的有序数组，copy回原数组arr中去
	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}
	return res
}
