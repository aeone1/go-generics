package main

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func sum[T Number](list []T) T {
	var res T
	for _, i := range list {
		res += i
	}
	return res
}