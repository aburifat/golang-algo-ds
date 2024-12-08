package main

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (st *Stack[T]) Push(x T) {
	st.vals = append(st.vals, x)
}

func (st *Stack[T]) Pop() (T, bool) {
	if len(st.vals) == 0 {
		var zero T
		return zero, false
	}
	top := st.vals[len(st.vals)-1]
	st.vals = st.vals[:len(st.vals)-1]
	return top, true
}

func main() {
	var st Stack[int]
	st.Push(10)
	st.Push(20)
	st.Push(30)

	for {
		if v, ok := st.Pop(); ok {
			fmt.Println(v, ok)
		} else {
			break
		}
	}
}
