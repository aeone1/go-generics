package main


func isContained[T comparable](n T, list []T) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}