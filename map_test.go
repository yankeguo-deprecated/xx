package xx

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestMapKeys(t *testing.T) {
	keys := MapKeys(map[string]int{
		"world": 3,
		"hello": 2,
	})
	sort.Strings(keys)
	require.Equal(t, []string{"hello", "world"}, keys)
}

func TestMapVals(t *testing.T) {
	keys := MapVals(map[string]int{
		"world": 3,
		"hello": 2,
	})
	sort.Ints(keys)
	require.Equal(t, []int{2, 3}, keys)
}
