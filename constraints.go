package main

import "golang.org/x/exp/constraints"

func maxWithConstraints[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}