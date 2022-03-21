package main

import (
	"sort"
	"strings"
)

// 方法1 暴力法穷举，排列组合略

// LowestStringByGreedy 方法2 贪心法
func LowestStringByGreedy(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	Sort(strs, func(a, b string) int {
		return strings.Compare(a, b)
	})

	res := ""
	for i := 0; i < len(strs); i++ {
		res += strs[i]
	}
	return res
}

type Comparator func(a, b string) int

func Sort(values []string, comparator Comparator) {
	sort.Sort(sortable{values, comparator})
}

type sortable struct {
	values     []string
	comparator Comparator
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
