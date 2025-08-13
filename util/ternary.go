package util

func Ternary[T any](condition bool, first, second T) T {
	if condition {
		return first
	}

	return second
}
