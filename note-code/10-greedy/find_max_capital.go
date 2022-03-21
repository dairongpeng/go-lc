package main

import (
	"container/heap"
)

// Item 项目
type Item struct {
	C int
	P int
}

// MinCostQ 项目最小花费。由花费组织的小根堆
type MinCostQ struct {
	Items []*Item
}

func (c MinCostQ) Len() int {
	return len(c.Items)
}

// Less i对应的花费C的值小于j对应的值为true，则从小到大排序，即小根堆
func (c MinCostQ) Less(i, j int) bool {
	return c.Items[i].C < c.Items[j].C
}

func (c MinCostQ) Swap(i, j int) {
	c.Items[i], c.Items[j] = c.Items[j], c.Items[i]
}

func (c *MinCostQ) Push(h interface{}) {
	c.Items = append(c.Items, h.(*Item))
}

func (c *MinCostQ) Pop() (x interface{}) {
	n := len(c.Items)
	x = c.Items[n-1]
	c.Items = c.Items[:n-1]
	return x
}

// MaxProfitQ 项目最大利润，由利润组织的大根堆
type MaxProfitQ struct {
	Items []*Item
}

func (c MaxProfitQ) Len() int {
	return len(c.Items)
}

// Less i对应的利润P的值大于j对应的值为true，则从大到小排序，即大根堆
func (c MaxProfitQ) Less(i, j int) bool {
	return c.Items[i].P > c.Items[j].P
}

func (c MaxProfitQ) Swap(i, j int) {
	c.Items[i], c.Items[j] = c.Items[j], c.Items[i]
}

func (c *MaxProfitQ) Push(h interface{}) {
	c.Items = append(c.Items, h.(*Item))
}

func (c *MaxProfitQ) Pop() (x interface{}) {
	n := len(c.Items)
	x = c.Items[n-1]
	c.Items = c.Items[:n-1]
	return x
}

// findMaximizedCapital 找到项目最大利润。由于Profits和Capital一一对应
// K表示你只能串行的最多K个项目，M表示你的初始资金。
func findMaximizedCapital(K, W int, Profits, Capital []int) int {
	Items := make([]*Item, 0)
	for i := 0; i < len(Profits); i++ {
		im := &Item{
			C: Capital[i],
			P: Profits[i],
		}
		Items = append(Items, im)
	}
	minC := &MinCostQ{
		Items: Items,
	}

	maxQ := &MaxProfitQ{
		Items: Items,
	}

	// 由花费组织的小根堆。初始化
	heap.Init(minC)
	// 由利润组织的大根堆。初始化
	heap.Init(maxQ)

	// 做k轮项目
	for i := 0; i < K; i++ {
		// 小根堆不为空，且堆顶的花费被我当前启动资金cover住
		for len(minC.Items) != 0 && minC.Items[len(minC.Items)-1].C <= W {
			// 小根堆的堆顶扔到大根堆中去
			maxQ.Push(minC.Pop())
		}
		// 大根堆没有可以做的项目，直接返回总钱数
		if len(maxQ.Items) == 0 {
			return W
		}
		// 大根堆不为空，堆顶元素的利润直接加到我们的总钱数上
		// 大根堆弹出堆顶元素
		W += maxQ.Pop().(Item).P
	}
	return W
}
