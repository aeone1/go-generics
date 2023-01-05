package main

func functionalMap[T any](list []T, accumulator func(T) T) []T {
	var res = make([]T, 0, len(list))
	for _, el := range list {
		res = append(res, accumulator(el))
	}
	return res
}