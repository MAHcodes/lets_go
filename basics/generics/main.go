package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func SumOfNums[T Number](nums []T) (n T) {
	for _, num := range nums {
		n += num
	}
	return
}

func main() {
	ints := []int64{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(SumOfNums(ints), SumOfNums(floats))
}
