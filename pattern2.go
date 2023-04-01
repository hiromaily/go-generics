package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

// int Stringer
type newInt int

func (n newInt) String() string {
	return strconv.Itoa(int(n))
}

// float64 Stringer
type newFloat float64

func (n newFloat) String() string {
	//return fmt.Sprintf("%f", n)
	return strconv.FormatFloat(float64(n), 'f', 2, 64)
}

// `type parameters`
func Stringers[T Stringer](val []T) []string {
	var result []string
	for _, v := range val {
		result = append(result, v.String())
	}
	return result
}

func pattern2() {
	fmt.Println("----- Pattern 2 -----")
	fmt.Println(Stringers([]newInt{1, 2, 3, 4}))
	fmt.Println(Stringers([]newFloat{1.5, 2.3, 3.2, 4.567}))
}
