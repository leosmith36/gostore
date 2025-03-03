package store

func isZero[T comparable](val T) bool {
	var zero T

	return val == zero
}

func getZero[T any](val T) T {
	return *new(T)
}
