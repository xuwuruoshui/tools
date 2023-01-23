package utils

func Filter[T any](f func(T) bool, src []T) []T {
	var dst []T
	for _, v := range src {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

func Map[S, T any](src []S, f func(S) T) []T {
	dst := make([]T, len(src))
	for i, v := range src {
		dst[i] = f(v)
	}
	return dst
}
