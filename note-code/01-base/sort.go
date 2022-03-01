package main

import "fmt"

func main() {
	arr := []int{4, 3, 4, 2, 2, 2, 4, 1, 1, 1, 3, 3, 1, 1, 1, 4, 2, 2}
	printNum(arr)
}

func printNum(arr []int) {
	flag := 0
	// 经过循环处理，flag等于这两个不相等的且出现奇数次的异或结果；a ^ b
	for i := 0; i < len(arr); i++ {
		flag = flag ^ arr[i]
	}

	// 由于a != b 所以flag不为0。则flag的二进制位上一定存在1，选最后一位的1
	// 选取办法是，用flag 与 自身取反加1的结果做与运算
	rightOne := flag & ((^flag) + 1)
	onlyOne := 0
	// 经过这层循环的筛选，onlyOne等于a或者b其中的一个
	for j := 0; j < len(arr); j++ {
		if arr[j]&rightOne != 0 {
			onlyOne = onlyOne ^ arr[j]
		}
	}
	result1 := onlyOne
	result2 := flag ^ onlyOne
	// result1和result2就是数组中不相等的且为奇数的两个未知数a、b
	fmt.Println(result1)
	fmt.Println(result2)
}

// 选择排序
func selectionSort(arr []int) {
	// base case
	if len(arr) == 0 || len(arr) < 2 {
		return
	}

	// 在i到n-1的位置依次处理, i从0开始
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		// 根据第一层循环，在i~n-1的位置上依次寻找该位置的最小值，放到当前的i位置
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// 冒泡排序
func bubbleSort(arr []int) {
	if len(arr) == 0 || len(arr) < 2 {
		return
	}

	// 外循环从最大位置开始处理，形成冒泡的效果
	for i := len(arr) - 1; i > 0; i-- {
		// 内循环的一轮次，搞定外循环的一个位置。
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				swap(arr, i, j)
			}
		}
	}
}

// 选择排序
func insertionSort(arr []int) {
	if len(arr) == 0 || len(arr) < 2 {
		return
	}

	// 类比打扑克牌
	for i := 1; i < len(arr); i++ {
		// 每一轮内循环，与前面的元素来一轮比较，达到的效果是最小元素经过一轮内循环总能放到0位置
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
