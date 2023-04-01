package main

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](xs ...T) Set[T] {
	s := make(Set[T])
	for _, xs := range xs {
		s.Add(xs)
	}
	return s
}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func (s Set[T]) Includes(x T) bool {
	_, ok := s[x]
	return ok
}

func pattern4() {
	fmt.Println("----- Pattern 4 -----")

	s := NewSet(1, 2, 3)
	fmt.Println(s.Includes(5)) // false
	s.Add(5)
	fmt.Println(s.Includes(5)) // true
	s.Remove(3)
	fmt.Println(s.Includes(3)) // false
}
