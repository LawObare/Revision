package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type values interface {
	constraints.Ordered
}

func add[T values](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(5, 6))
	fmt.Println(add(5.5, 6.6))
	fmt.Println(add(5.93749274, 6.946294))
}
