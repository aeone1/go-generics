package main

func reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, el := range list {
		init = accumulator(init, el)
	}
	return init
}