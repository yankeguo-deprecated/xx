package xx

// Eq returns a func that returns true if incoming value is same
func Eq[T comparable](v T) F11[T, bool] {
	return func(v1 T) bool {
		return v1 == v
	}
}

// Neq returns a func that returns true if incoming value is not same
func Neq[T comparable](v T) F11[T, bool] {
	return func(v1 T) bool {
		return v1 != v
	}
}
