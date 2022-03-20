package main

import (
	"fmt"
	"sort"
)

// Comparator
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator func(a, b interface{}) int

// 定义可排序的结构
type sortable struct {
	values []interface{}
	// 该结构携带一个自定义的排序策略
	comparator Comparator
}

// Sort 使用Go原生的排序进行包装，该排序在数据规模大的时候使用快排，数据规模小的时候使用插入排序
func Sort(values []interface{}, comparator Comparator) {
	sort.Sort(sortable{values, comparator})
}

func (s sortable) Len() int {
	return len(s.values)
}

func (s sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s sortable) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}

// IntComparator 是自定义的整形排序策略，可以实现其他自定义排序策略
func IntComparator(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

func main() {
	tests := [][]interface{}{
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{11, 22, -1},
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}
	for _, test := range tests {
		Sort(test, IntComparator)
		fmt.Println(test)
	}
}
