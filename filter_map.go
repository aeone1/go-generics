package main

func functionalFilter[T any](list []T, filter func(T) bool) []T {
	var res []T
	for _, el := range list {
		if filter(el) {
			res = append(res, el)
		}
	}
	return res
}