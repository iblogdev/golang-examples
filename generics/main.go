package main

import "fmt"

type Number interface {
	int | float64
}

func main() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1.0, 2.3))
}

func Sum[T Number](inputs ...T) T {
	var sum T
	for _, input := range inputs {
		sum = sum + input
	}
	return sum
}
