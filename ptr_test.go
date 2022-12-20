package xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPtr(t *testing.T) {
	v := 2
	nv := Ptr(v)
	v = 3
	require.Equal(t, 2, *nv)
}

func TestPtrSlice(t *testing.T) {
	si := []int{1, 2, 3}
	so := PtrSlice(si)
	*so[0] = 2
	require.Equal(t, 1, si[0])
	require.Equal(t, 2, si[1])
	require.Equal(t, 3, si[2])
	require.Equal(t, 2, *so[0])
	require.Equal(t, 2, *so[1])
	require.Equal(t, 3, *so[2])
}

func TestPtrMap(t *testing.T) {
	mi := map[int]int{1: 1, 2: 2, 3: 3}
	mo := PtrMap(mi)
	*mo[1] = 2
	require.Equal(t, 1, mi[1])
	require.Equal(t, 2, mi[2])
	require.Equal(t, 3, mi[3])
	require.Equal(t, 2, *mo[1])
	require.Equal(t, 2, *mo[2])
	require.Equal(t, 3, *mo[3])
}
