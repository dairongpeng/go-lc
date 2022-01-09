package main

import "fmt"

// 1.两数之和 Easy
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	// 申请一个hash表
	m := map[int]int{}
	// 对给定的数组进行遍历
	for i, v := range nums {
		// 判断target-v在不在hash表中，如果存在，直接返回
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		// 如果不存在，把当前的v存入hash表中，继续寻找
		m[v] = i
	}
	// 遍历完hash表后，仍未找到，返回nil
	return nil
}
