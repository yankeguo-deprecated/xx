package xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepeat(t *testing.T) {
	require.Equal(t, []string{"a", "a"}, Repeat(2, "a"))
	require.Equal(t, []int{5, 5}, Repeat(2, 5))
}

func TestSliceMap(t *testing.T) {
	s := []string{"aaa", "bb", "c"}
	require.Equal(t, []int{3, 2, 1}, SliceMap(s, func(t string) int {
		return len(t)
	}))
}

func TestSliceToSet(t *testing.T) {
	require.Equal(t, map[string]struct{}{
		"hello": {},
		"world": {},
	}, SliceToSet([]string{"world", "hello"}))
}

func TestSliceToMap(t *testing.T) {
	require.Equal(t, map[string]string{
		"hello":   "world",
		"goodbye": "world",
	}, SliceToMap([]string{
		"hello", "world", "goodbye", "world",
	}))
}

func TestSliceFilter(t *testing.T) {
	require.Equal(t, []string{"hell", "dull"}, SliceFilter([]string{"hell", "hello", "dull"}, Not(Eq("hello"))))
}
