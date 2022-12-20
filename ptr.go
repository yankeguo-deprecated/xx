package xx

// Ptr create a pointer to a value
func Ptr[T any](v T) *T {
	return &v
}

// PtrSlice create a new slice with detached pointers to original slice element
func PtrSlice[T any](si []T) []*T {
	so := make([]*T, len(si), len(si))
	for i, _v := range si {
		v := _v
		so[i] = &v
	}
	return so
}

// PtrMap create a new map with detached pointers to original map element
func PtrMap[T comparable, U any](mi map[T]U) map[T]*U {
	mo := make(map[T]*U)
	for k, _v := range mi {
		v := _v
		mo[k] = &v
	}
	return mo
}
