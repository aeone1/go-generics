package main

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64
}

func max[T Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func maxWithTernary[T Ordered](x, y T) T {
	return ternary(x > y, x, y)
}