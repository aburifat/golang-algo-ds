package main

import "fmt"

type Queue[T any] struct {
	vals []T
}

func (q *Queue[T]) Push(x T) {
	q.vals = append(q.vals, x)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.vals) == 0 {
		var zero T
		return zero, false
	}
	top := q.vals[0]
	q.vals = q.vals[1:]
	return top, true
}

func main() {
	var q Queue[int]
	q.Push(10)
	q.Push(20)
	q.Push(30)

	for {
		if v, ok := q.Pop(); ok {
			fmt.Println(v, ok)
		} else {
			break
		}
	}
}
