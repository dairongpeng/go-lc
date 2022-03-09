// 利用双向链表实现双端队列
package main

// DoubleEndsNode 双端队列节点
type DoubleEndsNode struct {
	val  int
	pre  *DoubleEndsNode
	next *DoubleEndsNode
}

// DoubleEndsList 双端队列接口
type DoubleEndsList interface {
	// AddFromHead 从头部添加节点
	AddFromHead(v int)
	// AddFromBottom 从尾部添加节点
	AddFromBottom(v int)
	// PopFromHead 从头部弹出节点
	PopFromHead() (int, bool)
	// PopFromBottom 从尾部弹出节点
	PopFromBottom() (int, bool)
	// IsEmpty 双端队列是否为空
	IsEmpty() bool
}

type DoubleEndsQueue struct {
	head *DoubleEndsNode
	tail *DoubleEndsNode
}

func (q *DoubleEndsQueue) AddFromHead(v int) {
	cur := &DoubleEndsNode{
		val: v,
	}
	if q.head == nil {
		q.head = cur
		q.tail = cur
	} else {
		cur.next = q.head
		q.head.pre = cur
		q.head = cur
	}
}

func (q *DoubleEndsQueue) AddFromBottom(v int) {
	cur := &DoubleEndsNode{
		val: v,
	}
	if q.head == nil {
		q.head = cur
		q.tail = cur
	} else {
		q.tail.next = cur
		cur.pre = q.tail
		q.tail = cur
	}
}

func (q *DoubleEndsQueue) PopFromHead() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	v := q.head.val
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return v, true
	} else {
		h := q.head
		q.head = q.head.next
		q.head.pre = nil
		h.next = nil
		return v, true
	}
}

func (q *DoubleEndsQueue) PopFromBottom() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	v := q.tail.val
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return v, true
	} else {
		t := q.tail
		q.tail = q.tail.pre
		q.tail.next = nil
		t.pre = nil
		return v, true
	}
}

func (q *DoubleEndsQueue) IsEmpty() bool {
	return q.head == nil
}

// Stack 利用双端队列实现栈
type Stack struct {
	qu *DoubleEndsQueue
}

func (s *Stack) push(v int) {
	s.qu.AddFromHead(v)
}

func (s *Stack) pop() (int, bool) {
	return s.qu.PopFromHead()
}

func (s *Stack) peek() (int, bool) {
	if s.qu.IsEmpty() {
		return 0, false
	}
	return s.qu.head.val, true
}

func (s *Stack) IsEmpty() bool {
	return s.qu.IsEmpty()
}

// Queue 利用双端队列实现队列
type Queue struct {
	qu *DoubleEndsQueue
}

func (q *Queue) push(v int) {
	q.qu.AddFromHead(v)
}

func (q *Queue) poll() (int, bool) {
	return q.qu.PopFromBottom()
}

func (q *Queue) IsEmpty() bool {
	return q.qu.IsEmpty()
}

type MinStack struct {
	data *Stack
	min  *Stack
}

func (s *MinStack) push(v int) {
	// min栈只保存最小的v，当然这里也可以设计成min栈和data栈同步上升的策略
	if s.min.IsEmpty() {
		s.min.push(v)
	} else if c, ok := s.min.peek(); ok {
		if v <= c { // 小于等于都入栈，弹出的时候等于也同步弹出min栈
			s.min.push(v)
		} else {
			s.min.push(c)
		}
	}
	// 数据栈稳步上升
	s.data.push(v)
}

func (s *MinStack) pop() (int, bool) {
	if s.data.IsEmpty() {
		return 0, false
	}
	v, _ := s.data.pop()
	if m, ok := s.min.peek(); ok {
		if m == v {
			s.min.pop()
		}
	}
	return v, true
}

func (s *MinStack) getMin() (int, bool) {
	if s.min.IsEmpty() {
		return 0, false
	}
	return s.min.peek()
}
