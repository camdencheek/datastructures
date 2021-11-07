package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/camdencheek/datastructures/compare"
)

func CompareInts(a, b int) compare.Result {
	switch {
	case a < b:
		return compare.Less
	case a > b:
		return compare.Greater
	default:
		return compare.Equal
	}
}

func TestBinaryHeap(t *testing.T) {
	bh := New(CompareInts)
	bh.Push(2)
	bh.Push(5)
	bh.Push(3)
	require.Equal(t, 5, *bh.Pop())
	require.Equal(t, 3, *bh.Pop())
	require.Equal(t, 2, *bh.Pop())
	require.Nil(t, bh.Pop())
}
