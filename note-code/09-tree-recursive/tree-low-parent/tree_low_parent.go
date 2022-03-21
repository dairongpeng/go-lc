package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// LowestAncestorByMap 给定两个树节点，求这两个节点的最低公共祖先
func (head *Node) LowestAncestorByMap(o1 *Node, o2 *Node) *Node {
	if head == nil {
		return nil
	}
	// key的父节点是value
	parentMap := make(map[*Node]*Node, 0)
	parentMap[head] = nil
	// 递归填充map
	fillParentMap(head, parentMap)
	// 辅助set
	nodeSet := make(map[*Node]string, 0)
	cur := o1
	nodeSet[cur] = ""
	// nodeSet存入的是沿途所有的父节点
	for parent, ok := parentMap[cur]; ok; {
		nodeSet[parent] = ""
	}

	cur = o2
	// o2的某个父节点在parentSet中，就是我们要找的节点
	for _, ok := parentMap[cur]; !ok; {
		// 继续向上
		cur = parentMap[cur]
	}
	return cur
}

func fillParentMap(head *Node, parentMap map[*Node]*Node) {
	if head.Left != nil {
		parentMap[head.Left] = head
		fillParentMap(head.Left, parentMap)
	}
	if head.Right != nil {
		parentMap[head.Right] = head
		fillParentMap(head.Right, parentMap)
	}
}

type Info struct {
	// o1和o2的最初交汇点，如果不是在当前这颗X节点的树上，返回空
	Ans *Node
	// 在当前子树上，是否发现过o1和o2
	FindO1 bool
	FindO2 bool
}

// LowestAncestorByProcess 递归二叉树判断任意两个节点的最低公共祖先
func (head *Node) LowestAncestorByProcess(o1, o2 *Node) *Node {
	return Process(head, o1, o2).Ans
}

func Process(node, o1, o2 *Node) *Info {
	// o1和o2不为空，那么空树上的Info如下
	if node == nil {
		return &Info{
			Ans:    nil,
			FindO1: false,
			FindO2: false,
		}
	}
	leftInfo := Process(node.Left, o1, o2)
	rightInfo := Process(node.Right, o1, o2)

	// 构建node自身需要返回的Info
	// node为头的树上是否发现了o1
	findO1 := node == o1 || leftInfo.FindO1 || rightInfo.FindO1
	// node为头的树上是否发现了o2
	findO2 := node == o2 || leftInfo.FindO2 || rightInfo.FindO2
	// 	O1和O2最初的交汇点在哪？

	// 1) 在左树上已经提前交汇了,最初交汇点保留左树的
	var ans *Node = nil
	if leftInfo.Ans != nil {
		ans = leftInfo.Ans
	}
	// 2) 在右树上已经提前交汇了，最初交汇点保留右树的
	if rightInfo.Ans != nil {
		ans = rightInfo.Ans
	}
	// 3) 没有在左树或者右树上提前交汇
	if ans == nil {
		// 但是找到了o1和o2，那么交汇点就是X自身
		if findO1 && findO2 {
			ans = node
		}
	}
	return &Info{
		Ans:    ans,
		FindO1: findO1,
		FindO2: findO2,
	}
}
