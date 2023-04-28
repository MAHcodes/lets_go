package main

import "golang.org/x/exp/constraints"

type Number interface{
	constraints.Integer | constraints.Float
}

func Add[T Number](nums ...T) (sum T) {
	for _, i := range nums {
		sum += i
	}
	return
}
