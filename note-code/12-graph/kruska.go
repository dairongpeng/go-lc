package main

// kruskaMST 克鲁斯卡尔最小生成树算法。返回set
func kruskaMST(graph *Graph) map[*Edge]string {
	values := make([]int, 0)
	for k := range graph.nodes {
		values = append(values, k)
	}
	// 初始化一个并查集结构
	unitionSet := InitUnionSet(values)
	edgesHeap := make(Edges, 0)
	// 边按照权值从小到大排序，加入到堆
	for edge := range graph.edges {
		edgesHeap.Push(edge)
	}

	resultSet := make(map[*Edge]string)

	// 堆不为空，弹出小根堆的堆顶
	for len(edgesHeap) != 0 {
		// 假设M条边，O(logM)
		edge := edgesHeap.Pop().(*Edge)
		// 如果该边的左右两侧不在同一个集合中
		if !unitionSet.IsSameSet(edge.from.value, edge.to.value) {
			// 要这条边
			resultSet[edge] = ""
			// 联合from和to
			unitionSet.Union(edge.from.value, edge.to.value)
		}
	}
	return resultSet
}

// UNode 并查集结构中的节点类型
type UNode struct {
	V int
}

type UnionSet struct {
	// 记录样本到样本代表点的关系。值到代表该值的Node节点的关系映射
	Nodes map[int]*UNode
	// 记录某节点到根祖宗节点的关系。
	// 比如b指向a，c指向a，d指向a，a指向自身
	// map中保存的a->a b->a c->a d->a
	RootFatherMap map[*UNode]*UNode
	// 只有当前点，他是代表点，会在sizeMap中记录该代表点的连通个数
	SizeMap map[*UNode]int
}

// InitUnionSet 初始化一个并查集结构
func InitUnionSet(values []int) *UnionSet {
	us := &UnionSet{}
	nodes := make(map[int]*UNode, 0)
	fatherMap := make(map[*UNode]*UNode, 0)
	sizeMap := make(map[*UNode]int, 0)
	for _, v := range values {
		node := &UNode{V: v}
		nodes[v] = node
		fatherMap[node] = node
		sizeMap[node] = 1
	}

	us.Nodes = nodes
	us.RootFatherMap = fatherMap
	us.SizeMap = sizeMap
	return us
}

// FindFather 在并查集结构中找一个节点的父亲根节点
// 从点cur开始，一直往上找，找到不能再往上的代表点，返回
// 通过把路径上所有节点指向最上方的代表节点，目的是把findFather优化成O(1)的
func (set *UnionSet) FindFather(cur *UNode) *UNode {
	// 在找father的过程中，沿途所有节点加入当前容器，便于后面扁平化处理
	path := make([]*UNode, 0)
	// 当前节点的父亲不是指向自己，进行循环
	for cur != set.RootFatherMap[cur] {
		path = append(path, cur)
		// 向上移动
		cur = set.RootFatherMap[cur]
	}
	// 循环结束，cur此时是最上的代表节点
	// 把沿途所有节点拍平，都指向当前最上方的代表节点
	for len(path) != 0 {
		for i := len(path) - 1; i >= 0; i-- {
			set.RootFatherMap[path[i]] = cur
		}
	}
	return cur
}

// IsSameSet 判断两个元素是否在同一个并查集中
func (set *UnionSet) IsSameSet(a, b int) bool {
	// 先检查a和b有没有登记
	if _, ok := set.Nodes[a]; !ok {
		return false
	}
	if _, ok := set.Nodes[b]; !ok {
		return false
	}

	// 比较a的最上的代表点和b最上的代表点
	return set.FindFather(set.Nodes[a]) == set.FindFather(set.Nodes[b])
}

// Union 合并两个元素
func (set *UnionSet) Union(a, b int) {
	// 先检查a和b有没有都登记过
	if _, ok := set.Nodes[a]; !ok {
		return
	}
	if _, ok := set.Nodes[b]; !ok {
		return
	}

	// 找到a的最上面的代表点
	aHead := set.FindFather(set.Nodes[a])
	// 找到b的最上面的代表点
	bHead := set.FindFather(set.Nodes[b])
	// 只有两个最上代表点内存地址不相同，需要union
	if aHead != bHead {
		// 由于aHead和bHead都是最上面的代表点，那么在sizeMap里可以拿到大小
		aSetSize := set.SizeMap[aHead]
		bSetSize := set.SizeMap[bHead]
		var big *UNode
		var small *UNode
		// 哪个小，哪个挂在下面
		if aSetSize >= bSetSize {
			big = aHead
			small = bHead
		} else {
			big = bHead
			small = aHead
		}

		// 把小集合直接挂到大集合的最上面的代表节点下面
		set.RootFatherMap[small] = big
		// 大集合的代表节点的size要吸收掉小集合的size
		set.SizeMap[big] = aSetSize + bSetSize
		// 把被吸收掉的小set删除掉
		delete(set.SizeMap, small)
	}
}
