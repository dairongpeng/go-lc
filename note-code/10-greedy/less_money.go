package main

import (
	"container/heap"
	"fmt"
)

// CutCost Array 例如[10, 20, 30]表示价值为60的金条，需要切割成10 20 30的条段给三个人分
type CutCost struct {
	Array []int
}

func (c CutCost) Len() int {
	return len(c.Array)
}

// Less 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
func (c CutCost) Less(i, j int) bool {
	return c.Array[i] < c.Array[j]
}

func (c CutCost) Swap(i, j int) {
	c.Array[i], c.Array[j] = c.Array[j], c.Array[i]
}

func (c *CutCost) Push(h interface{}) {
	c.Array = append(c.Array, h.(int))
}

func (c *CutCost) Pop() (x interface{}) {
	n := len(c.Array)
	x = c.Array[n-1]
	c.Array = c.Array[:n-1]
	return x
}

// 切金条，贪心解法，建立一个小根堆，把所有数扔进去
func lessMoney(c *CutCost) int {
	fmt.Println("原始slice: ", c.Array)

	heap.Init(c)
	// 通过堆初始化后的arr
	fmt.Println("堆初始化后的slice:", c.Array)

	sum := 0
	cur := 0
	for len(c.Array) > 1 {
		// 每一次弹出两个数，合并成一个数
		// 合成的数一定输非叶子节点
		cur = c.Pop().(int) + c.Pop().(int)
		// 把合成的数累加到sum中去
		sum += cur
		// 把合成的数加入小根堆中
		c.Push(cur)
	}
	return sum
}
