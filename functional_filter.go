package main

func functionalFilter[T any](list []T, filter func(T) bool) []T {
	res := make([]T, 0, len(list))
	for _, el := range list {
		if filter(el) {
			res = append(res, el)
		}
	}
	return res
}