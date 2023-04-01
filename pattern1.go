package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// unions
type Number interface {
	int | int32 | int64 | float32 | float64
}

type List interface {
	[]int
}

// any can not be used for like `<` / `>` operator
func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

//	type Ordered interface {
//		Integer | Float | ~string
//	}
func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func GNumMin[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Loop[T List](x T) {
	for range x {
	}
}
func pattern1() {
	fmt.Println("----- Pattern 1 -----")

	// non-generic func
	smaller := Min(1.5, 4.8)
	fmt.Printf("smaller float value: %f\n", smaller)

	// generic func with int
	intSmaller := GMin[int](2, 3)
	fmt.Printf("smaller integer value: %d\n", intSmaller)

	stringSmaller := GMin[string]("foo", "bar")
	fmt.Printf("smaller string value: %s\n", stringSmaller)

	// generic func with int
	intSmaller2 := GNumMin[int](2, 3)
	fmt.Printf("smaller integer value: %d\n", intSmaller2)

	Loop[[]int]([]int{1, 2, 3})
}
