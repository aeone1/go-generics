package main

import "fmt"

func main() {
	// fmt.Println(max(6, 5))
	// fmt.Println(maxWithConstraints(2, 5))
	// fmt.Println(maxWithConstraints("abc", "abd"))

	// fmt.Println(maxWithTernary(2, 5))

	// list := []string{"a", "b", "c"}
	// fmt.Println(isContained("a", list))
	// fmt.Println(isContained("d", list))

	// listNum := []int{1, 2, 3, 4, 5}
	// fmt.Println(sum(listNum))
	// fmt.Println(sum([]float32{}))

	// list := []int{1, 2, 3, 4, 5}
	// sum := func(x, y int) int { return x + y}
	// mul := func(x, y int) int { return x * y}
	// res := reduce(list, mul, 1)
	// fmt.Println(res)

	list := []int{1, 2, 3, 4, 5}
	mulBy2or3 := func(n int) int {
		if (n % 2 == 0) {
			return n * 2
		}
		return n * 3
	}
	res := functionalMap(list, mulBy2or3)
	fmt.Println(res, len(res), cap(res))

	// list := []int{1, 2, 3, 4, 5}
	// filterEven := func(n int) bool { return n % 2 == 0 }
	// res := functionalFilter(list, filterEven)
	// fmt.Println(res, len(res), cap(res))
}

