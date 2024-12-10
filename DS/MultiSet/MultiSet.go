package main

import "fmt"

type Set[T comparable] struct {
	mp   map[T]int
	size int
}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{
		mp:   map[T]int{},
		size: 0,
	}

	for _, item := range items {
		s.mp[item]++
		s.size++
	}

	return s
}

func (s *Set[T]) Insert(x T) {
	s.mp[x]++
	s.size++
}

func (s *Set[T]) Remove(x T) {
	if value, ok := s.mp[x]; ok {
		if value > 1 {
			s.mp[x]--
		} else {
			delete(s.mp, x)
		}
		s.size--
	}
}

func (s Set[T]) Contains(x T) (int, bool) {
	ct, ok := s.mp[x]
	return ct, ok
}

func (s Set[T]) Size() int {
	return s.size
}

func (s Set[T]) CountDistinct() int {
	return len(s.mp)
}

func (s1 Set[T]) Union(s2 Set[T]) Set[T] {
	s := Set[T]{
		mp:   map[T]int{},
		size: 0,
	}

	for key := range s1.mp {
		s.mp[key]++
	}

	for key := range s2.mp {
		s.mp[key]++
	}
	s.size = s1.size + s2.size
	return s
}

func (s1 Set[T]) Intersection(s2 Set[T]) Set[T] {
	s := Set[T]{
		mp:   map[T]int{},
		size: 0,
	}

	if len(s2.mp) < len(s1.mp) {
		s1, s2 = s2, s1
	}

	for key := range s1.mp {
		if _, ok := s2.mp[key]; ok {
			s.mp[key] = min(s1.mp[key], s2.mp[key])
			s.size += s.mp[key]
		}
	}

	return s
}

func (s1 Set[T]) Difference(s2 Set[T]) Set[T] {
	s := Set[T]{
		mp:   map[T]int{},
		size: 0,
	}

	for key := range s1.mp {
		if ct, ok := s2.mp[key]; !ok {
			s.mp[key] = s1.mp[key]
			s.size += s.mp[key]
		} else if ct < s1.mp[key] {
			s.mp[key] = s1.mp[key] - ct
			s.size += s.mp[key]
		}
	}

	return s
}

func (s *Set[T]) Clear() {
	clear(s.mp)
	s.size = 0
}

func (s Set[T]) IsEmpty() bool {
	return s.size == 0
}

func main() {
	s1 := NewSet(1, 2, 2, 3, 4)
	fmt.Println("Set 1:", s1.mp, "Size:", s1.Size(), "Distinct Count:", s1.CountDistinct())

	s1.Insert(2)
	s1.Insert(5)
	fmt.Println("After inserting 2 and 5:", s1.mp, "Size:", s1.Size(), "Distinct Count:", s1.CountDistinct())

	s1.Remove(2)
	fmt.Println("After removing 2:", s1.mp, "Size:", s1.Size(), "Distinct Count:", s1.CountDistinct())
	s1.Remove(5)
	fmt.Println("After removing 5:", s1.mp, "Size:", s1.Size(), "Distinct Count:", s1.CountDistinct())

	count, exists := s1.Contains(2)
	fmt.Println("Contains 2:", exists, "Count:", count)
	count, exists = s1.Contains(5)
	fmt.Println("Contains 5:", exists, "Count:", count)

	s2 := NewSet(3, 4, 4, 5, 6)
	fmt.Println("Set 2:", s2.mp, "Size:", s2.Size(), "Distinct Count:", s2.CountDistinct())

	unionSet := s1.Union(s2)
	fmt.Println("Union of Set 1 and Set 2:", unionSet.mp, "Size:", unionSet.Size(), "Distinct Count:", unionSet.CountDistinct())

	intersectionSet := s1.Intersection(s2)
	fmt.Println("Intersection of Set 1 and Set 2:", intersectionSet.mp, "Size:", intersectionSet.Size(), "Distinct Count:", intersectionSet.CountDistinct())

	differenceSet := s1.Difference(s2)
	fmt.Println("Difference of Set 1 and Set 2:", differenceSet.mp, "Size:", differenceSet.Size(), "Distinct Count:", differenceSet.CountDistinct())

	s1.Clear()
	fmt.Println("After clearing Set 1:", s1.mp, "Size:", s1.Size(), "Is Empty:", s1.IsEmpty())
}
