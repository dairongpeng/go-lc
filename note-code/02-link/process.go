package main

import "fmt"

func getMax(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("arr len is zero")
	}
	return process(arr, 0, len(arr)-1), nil
}

func process(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := l + (r-l)/2
	// 左范围最大值
	lm := process(arr, l, mid)
	// 右范围最大值
	rm := process(arr, mid+1, r)
	if lm > rm {
		return lm
	} else {
		return rm
	}
}

func main() {
	arr := []int{1, 4, 2, 6, 992, 4, 2234, 83}
	m, err := getMax(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
