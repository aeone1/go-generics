package main

func functionalMap[T any](list []T, accumulator func(T) T) []T {
	var res = make([]T, len(list))
	for i, el := range list {
		res[i] = accumulator(el)
	}
	return res
}