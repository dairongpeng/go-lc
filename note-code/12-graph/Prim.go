package main

// primMST prim算法实现图的最小生成树
func (graph *Graph) primMST() map[*Edge]string {
	// 哪些点被解锁出来了
	nodeSet := make(map[*Node]string, 0)
	// 边的小根堆
	edgesHeap := make(Edges, 0)
	// 已经考虑过的边，不要重复考虑
	edgeSet := make(map[*Edge]string, 0)
	// 依次挑选的的边在resultSet里
	resultSet := make(map[*Edge]string, 0)
	// 随便挑了一个点,进入循环处理完后直接break

	for _, node := range graph.nodes {
		// node 是开始点
		if _, ok := nodeSet[node]; !ok {
			// 开始节点保留
			nodeSet[node] = ""
			// 开始节点的所有邻居节点全部放到小根堆
			// 即由一个点，解锁所有相连的边
			for _, edge := range node.edges {
				if _, ok := edgeSet[edge]; !ok {
					edgeSet[edge] = ""
					// 加入小根堆
					edgesHeap.Push(edge)
				}
			}

			for len(edgesHeap) != 0 {
				// 弹出解锁的边中，最小的边
				edge := edgesHeap.Pop().(*Edge)
				// 可能的一个新的点,from已经被考虑了，只需要看to
				toNode := edge.to
				// 不含有的时候，就是新的点
				if _, ok := nodeSet[toNode]; !ok {
					nodeSet[toNode] = ""
					resultSet[edge] = ""
					for _, nextEdge := range toNode.edges {
						// 没加过的，放入小根堆
						if _, ok := edgeSet[nextEdge]; !ok {
							edgeSet[nextEdge] = ""
							edgesHeap.Push(edge)
						}
					}
				}
			}
		}
	}
	// 直接break意味着我们不用考虑森林的情况
	// 如果不加break我们可以兼容多个无向图的森林的生成树
	// break;
	return resultSet
}

// Edges 边的集合。实现小根堆
type Edges []*Edge

func (es Edges) Less(i, j int) bool {
	return es[i].weight <= es[j].weight
}

func (es Edges) Len() int {
	return len(es)
}

func (es Edges) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es *Edges) Push(v interface{}) {
	*es = append(*es, v.(*Edge))
}

func (es *Edges) Pop() (x interface{}) {
	n := len(*es)
	x = (*es)[n-1]
	*es = (*es)[:n-1]
	return x
}
