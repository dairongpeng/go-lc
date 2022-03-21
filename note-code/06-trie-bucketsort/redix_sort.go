package main

import "fmt"

func RadixSort(nums []int) []int {
	numberBit := howManyBit(maximum(nums))
	// 循环的次数
	// 定义一个rec 二维切片 rec[i][x] 用来接受尾数是 i的数字

	for i := 0; i < numberBit; i++ {
		rec := make([][]int, 10)

		for _, num := range nums {
			rec[(num/pow10(i))%10] = append(rec[(num/pow10(i))%10], num)
		}
		// flatten the rec slice to the one dimension slice
		numsCopy := make([]int, 0)
		for j := 0; j < 10; j++ {
			numsCopy = append(numsCopy, rec[j]...)
		}
		// refresh nums，使得他变为 经过一次基数排序之后的数组
		nums = numsCopy
	}
	return nums
}

func pow10(num int) int {
	res := 1
	base := 10
	for num != 0 {
		if num&1 == 1 {
			num -= 1
			res *= base
		}
		num >>= 1
		base *= base
	}
	return res
}

func maximum(list []int) int {
	max := 0
	for _, i2 := range list {
		if i2 > max {
			max = i2
		}
	}
	return max
}

func howManyBit(number int) int {
	count := 0
	for number != 0 {
		number = number / 10
		count += 1
	}
	return count
}

func main() {
	var theArray = []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 233, 144}
	fmt.Print("排序前")
	fmt.Println(theArray)
	fmt.Print("排序后")
	fmt.Println(RadixSort(theArray))
}
