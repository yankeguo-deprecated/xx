package xx

import (
	"github.com/stretchr/testify/require"
	"go/token"
	"testing"
)

func TestCompare(t *testing.T) {
	require.True(t, Compare(1, token.EQL).Do(1))
	require.False(t, Compare(1, token.EQL).Do(2))
}

func TestAnd(t *testing.T) {
	require.True(t, And(Gt(1), Lt(5)).Do(3))
	require.False(t, And(Gt(1), Lt(5)).Do(13))
}

func TestOr(t *testing.T) {
	require.False(t, Or(Lt(1), Gt(5)).Do(3))
	require.True(t, Or(Lt(1), Gt(5)).Do(13))
}

func TestNot(t *testing.T) {
	require.True(t, Not(Lt(1), Gt(5)).Do(3))
	require.False(t, Not(Lt(1), Gt(5)).Do(13))
}

func TestEq(t *testing.T) {
	require.True(t, Eq(1).Do(1))
	require.False(t, Eq(1).Do(2))
}

func TestNeq(t *testing.T) {
	require.False(t, Neq(1).Do(1))
	require.True(t, Neq(1).Do(2))
}

func TestGt(t *testing.T) {
	require.True(t, Gt(1).Do(2))
	require.False(t, Gt(1).Do(0))

	require.True(t, Gt(1.1).Do(2))
	require.False(t, Gt(1.1).Do(0))

	require.True(t, Gt("d").Do("e"))
	require.False(t, Gt("d").Do("c"))
}

func TestLt(t *testing.T) {
	require.False(t, Lt(1).Do(2))
	require.True(t, Lt(1).Do(0))

	require.False(t, Lt(1.1).Do(2))
	require.True(t, Lt(1.1).Do(0))

	require.False(t, Lt("d").Do("e"))
	require.True(t, Lt("d").Do("c"))
}

func TestGte(t *testing.T) {
	require.True(t, Gte(1).Do(1))
	require.True(t, Gte(1).Do(2))
	require.False(t, Gte(1).Do(0))

	require.True(t, Gte(1.1).Do(1.1))
	require.True(t, Gte(1.1).Do(2))
	require.False(t, Gte(1.1).Do(0))

	require.True(t, Gte("d").Do("d"))
	require.True(t, Gte("d").Do("e"))
	require.False(t, Gte("d").Do("c"))
}

func TestLte(t *testing.T) {
	require.False(t, Lte(1).Do(2))
	require.True(t, Lte(1).Do(0))
	require.True(t, Lte(1).Do(1))

	require.False(t, Lte(1.1).Do(2))
	require.True(t, Lte(1.1).Do(0))
	require.True(t, Lte(1.1).Do(1.1))

	require.False(t, Lte("d").Do("e"))
	require.True(t, Lte("d").Do("c"))
	require.True(t, Lte("d").Do("d"))
}
