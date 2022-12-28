package xx

// Repeat create a slice with a value repeated n times
func Repeat[T any](n int, v T) []T {
	out := make([]T, n, n)
	for i := range out {
		out[i] = v
	}
	return out
}

// SliceMap convert a slice to another by mapping each element
func SliceMap[T any, U any](s []T, fn F11[T, U]) []U {
	var o []U
	if s == nil {
		return o
	}
	o = make([]U, len(s), len(s))
	for i := range s {
		o[i] = fn(s[i])
	}
	return o
}

// SliceToSet convert a slice of comparables into a SliceToSet-like map
func SliceToSet[T comparable](s []T) map[T]struct{} {
	o := make(map[T]struct{})
	for _, k := range s {
		o[k] = struct{}{}
	}
	return o
}

// SliceToMap convert a slice to map with key-values
func SliceToMap[T comparable](s []T) map[T]T {
	m := make(map[T]T)
	for i := 0; i < len(s); i += 2 {
		m[s[i]] = s[i+1]
	}
	return m
}

// SliceFilter returns a new slice with element filtered
func SliceFilter[T any](s []T, fn F11[T, bool]) []T {
	var o []T
	for _, v := range s {
		if fn(v) {
			o = append(o, v)
		}
	}
	return o
}

// SliceFlatten flatten a slice of slice
func SliceFlatten[T any](s [][]T) []T {
	var o []T
	for _, items := range s {
		o = append(o, items...)
	}
	return o
}
