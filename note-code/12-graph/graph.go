package main

import (
	"fmt"
)

// Node 图中的点元素表示
type Node struct {
	// 点的身份标识
	value int
	// 入度，表示有多少个点连向该点
	in int
	// 出度，表示从该点出发连向别的节点多少
	out int
	// 直接邻居：表示由自己出发，直接指向哪些节点。指向节点的总数等于out
	nexts []*Node
	// 直接下级边：表示由自己出发的边有多少
	edges []*Edge
}

// Edge 图中的边元素表示
type Edge struct {
	// 边的权重信息
	weight int
	// 出发的节点
	from *Node
	// 指向的节点
	to *Node
}

// Graph 图结构
type Graph struct {
	// 点的集合，编号为1的点是什么，用map
	nodes map[int]*Node
	// 边的集合(用hash实现set)
	edges map[*Edge]string
}

// 从node出发，对图进行宽度优先遍历。借助队列
func (node *Node) bfs() {
	if node == nil {
		return
	}
	queue := make([]*Node, 0)
	// 图需要用set结构，因为图相比于二叉树有可能存在环
	// 即有可能存在某个点多次进入队列的情况。使用Set可以防止相同节点重复进入队列
	set := make(map[*Node]string, 0)
	queue = append(queue, node)
	set[node] = ""
	for len(queue) != 0 {
		// 出队列
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.value)
		for _, next := range cur.nexts {
			// 直接邻居，没有进入过Set的进入Set和队列
			// 用set限制队列的元素，防止有环队列一直会加入元素
			if _, ok := set[next]; !ok { // Set中不存在, 则加入队列
				set[next] = ""
				queue = append(queue, next)
			}
		}
	}
}

// 从node出发，对图进行深度优先遍历。借助栈
func (node *Node) dfs() {
	if node == nil {
		return
	}

	stack := make([]*Node, 0)
	// Set的作用和宽度优先遍历类似，保证重复的点不要进栈
	set := make(map[*Node]string, 0)
	// 进栈
	stack = append(stack, node)
	set[node] = ""
	// 打印时机是在进栈的时候
	// 同理该步可以换成其他处理逻辑，表示深度遍历处理某件事情
	fmt.Println(node.value)

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 枚举当前弹出节点的后代
		for _, next := range cur.nexts {
			// 只要某个后代没进入过栈，进栈
			if _, ok := set[next]; !ok {
				// 把该节点的父亲节点重新压回栈中
				stack = append(stack, cur)
				// 再把自己压入栈中
				stack = append(stack, next)
				set[next] = ""
				// 打印当前节点的值
				fmt.Println(next.value)
				// 直接break，此时栈顶是当前next节点，达到深度优先的目的
				break
			}
		}
	}
}

// sortTopology 图的拓扑排序。返回拓扑排序的顺序list
func (graph *Graph) sortTopology() []*Node {
	// key：某一个node, value：该节点剩余的入度
	inMap := make(map[*Node]int)
	// 剩余入度为0的点，才能进这个队列
	zeroInQueue := make([]*Node, 0)
	// 拿到该图中所有的点集
	for _, node := range graph.nodes {
		// 初始化每个点，每个点的入度是原始节点的入度信息
		// 加入inMap
		inMap[node] = node.in
		// 由于是有向无环图，则必定有入度为0的起始点。放入到zeroInQueue
		if node.in == 0 {
			zeroInQueue = append(zeroInQueue, node)
		}
	}

	// 拓扑排序的结果，依次加入result
	result := make([]*Node, 0)

	for len(zeroInQueue) != 0 {
		// 该有向无环图初始入度为0的点，直接弹出放入结果集中
		cur := zeroInQueue[0]
		zeroInQueue = zeroInQueue[1:]
		result = append(result, cur)
		// 该节点的下一层邻居节点，入度减一且加入到入度的map中
		for _, next := range cur.nexts {
			inMap[next] = inMap[next] - 1
			// 如果下一层存在入度变为0的节点，加入到0入度的队列中
			if inMap[next] == 0 {
				zeroInQueue = append(zeroInQueue, next)
			}
		}
	}
	return result
}
