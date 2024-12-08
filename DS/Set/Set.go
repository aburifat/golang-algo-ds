package main

import "fmt"

type Set[T comparable] map[T]bool

func NewSet[T comparable](items ...T) Set[T] {
	set := map[T]bool{}

	for _, item := range items {
		set[item] = true
	}

	return set
}

func (s Set[T]) Insert(x T) {
	s[x] = true
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func (s Set[T]) Contains(x T) bool {
	_, ok := s[x]
	return ok
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s1 Set[T]) Union(s2 Set[T]) Set[T] {
	s := map[T]bool{}

	for key := range s1 {
		s[key] = true
	}

	for key := range s2 {
		s[key] = true
	}

	return s
}

func (s1 Set[T]) Intersection(s2 Set[T]) Set[T] {
	s := map[T]bool{}

	if len(s2) < len(s1) {
		s1, s2 = s2, s1
	}

	for key := range s1 {
		if _, ok := s2[key]; ok {
			s[key] = true
		}
	}

	return s
}

func (s1 Set[T]) Difference(s2 Set[T]) Set[T] {
	s := map[T]bool{}

	for key := range s1 {
		if _, ok := s2[key]; !ok {
			s[key] = true
		}
	}

	return s
}

func (s Set[T]) Clear() {
	clear(s)
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func main() {
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(3, 4, 5, 6)

	fmt.Println("Set1:", set1)
	fmt.Println("Set2:", set2)

	set1.Insert(5)
	fmt.Println("After Insert(5) in Set1:", set1)

	set1.Remove(2)
	fmt.Println("After Remove(2) from Set1:", set1)

	fmt.Println("Set1 Contains(3):", set1.Contains(3))
	fmt.Println("Set1 Contains(6):", set1.Contains(6))

	fmt.Println("Size of Set1:", set1.Size())

	union := set1.Union(set2)
	fmt.Println("Union of Set1 and Set2:", union)

	intersection := set1.Intersection(set2)
	fmt.Println("Intersection of Set1 and Set2:", intersection)

	difference := set1.Difference(set2)
	fmt.Println("Difference of Set1 and Set2:", difference)

	set1.Clear()
	fmt.Println("After Clear Set1:", set1)
	fmt.Println("Is Set1 Empty?:", set1.IsEmpty())
}
