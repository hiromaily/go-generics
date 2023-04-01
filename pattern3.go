package main

import "fmt"

// any is alias of interface{}
type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	v := make(Stack[T], 0)
	return &v
}

func (s *Stack[T]) Push(x T) {
	(*s) = append((*s), x)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func pattern3() {
	fmt.Println("----- Pattern 3 -----")

	// string stack
	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")
	fmt.Println(strStack.Pop()) // world
	fmt.Println(strStack.Pop()) // hello

	// int stack
	intStack := NewStack[int]()
	intStack.Push(100)
	intStack.Push(200)
	fmt.Println(intStack.Pop()) // 200
	fmt.Println(intStack.Pop()) // 1000
}
