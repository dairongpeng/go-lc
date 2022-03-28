package main

import "math"

// dijkstra算法-图的最短路径算法
// 给定一个图的节点，返回这个节点到图的其他点的最短距离
// 某个点不在map中记录，则from到该点位正无穷
func dijkstra(from *Node) map[*Node]int {
	// 从from出发到所有点的最小距离表
	distanceMap := make(map[*Node]int, 0)
	// from到from距离为0
	distanceMap[from] = 0
	// 已经求过距离的节点，存在selectedNodes中，不会再被选中记录
	selectedNodesSet := make(map[*Node]string)
	// from 0 得到没选择过的点的最小距离
	minNode := getMinDistanceAndUnselectedNode(distanceMap, selectedNodesSet)
	// 得到minNode之后
	for minNode != nil {
		// 把minNode对应的距离取出,此时minNode就是桥连点
		distance := distanceMap[minNode]
		// 把minNode上所有的邻边拿出来
		// 这里就是要拿到例如A到C和A到桥连点B再到C哪个距离小的距离
		for _, edge := range minNode.edges {
			// 某条边对应的下一跳节点toNode
			toNode := edge.to
			// 如果关于from的distencMap中没有去toNode的记录，表示正无穷，直接添加该条
			if _, ok := distanceMap[toNode]; !ok {
				// from到minNode的距离加上个minNode到当前to节点的边距离
				distanceMap[toNode] = distance + edge.weight
			} else { // 如果有，看该距离是否更小，更小就更新
				minDistance := int(math.Min(float64(distanceMap[toNode]), float64(distance+edge.weight)))
				distanceMap[edge.to] = minDistance
			}
		}
		// 锁上minNode，表示from通过minNode到其他节点的最小值已经找到
		// minNode将不再使用
		selectedNodesSet[minNode] = ""
		// 再在没有选择的节点中挑选MinNode当成from的桥接点
		minNode = getMinDistanceAndUnselectedNode(distanceMap, selectedNodesSet)
	}
	// 最终distanceMap全部更新，返回
	return distanceMap
}

// getMinDistanceAndUnselectedNode 得到没选择过的点的最小距离
func getMinDistanceAndUnselectedNode(distanceMap map[*Node]int, selectedNodesSet map[*Node]string) *Node {
	var minNode *Node = nil
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		// 没有被选择过，且距离最小
		if _, ok := selectedNodesSet[node]; !ok && distance < minDistance {
			minNode = node
			minDistance = distance
		}
	}
	return minNode
}

// 我们可以借助小根堆来替代之前的distanceMap。达到优化算法的目的
// 原因是之前我们要遍历hash表选出最小距离，现在直接是堆顶元素
// 但是我们找到通过桥节点更小的距离后，需要临时更该堆结构中元素数据
// 所以系统提供的堆我们需要改写。略
