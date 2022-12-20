package xx

// MapKeys extract keys of map
func MapKeys[T comparable, U any](m map[T]U) []T {
	var o []T
	for k := range m {
		o = append(o, k)
	}
	return o
}

// MapVals extract values of map
func MapVals[T comparable, U any](m map[T]U) []U {
	var o []U
	for _, v := range m {
		o = append(o, v)
	}
	return o
}
