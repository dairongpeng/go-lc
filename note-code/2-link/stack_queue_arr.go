package main

type Que struct {
	// 队列的底层结构
	arr []int
}

func (q *Que) push(v int) {
	q.arr = append(q.arr, v)
}

func (q *Que) poll() (int, bool) {
	if len(q.arr) == 0 {
		return 0, false
	}
	v := q.arr[0]
	q.arr = q.arr[1:]
	return v, true
}

//func main () {
//	q := Que{}
//	q.push(1)
//	q.push(9)
//	q.push(3)
//	if poll, ok := q.poll(); ok {
//		fmt.Println(poll)
//	}
//	if poll, ok := q.poll(); ok {
//		fmt.Println(poll)
//	}
//	if poll, ok := q.poll(); ok {
//		fmt.Println(poll)
//	}
//	if poll, ok := q.poll(); ok {
//		fmt.Println(poll)
//	}
//}
