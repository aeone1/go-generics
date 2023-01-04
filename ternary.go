package main

func ternary[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}