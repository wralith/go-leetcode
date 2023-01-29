package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_removeNthFromEnd(t *testing.T) {
	list := types.MakeListFromInts(1, 2, 3, 4, 5)
	got := removeNthFromEnd(list, 2)
	want := types.MakeListFromInts(1, 2, 3, 5)
	require.Equal(t, got, want)

	list = types.MakeListFromInts(1)
	got = removeNthFromEnd(list, 1)
	want = types.MakeListFromInts()
	require.Equal(t, got, want)

	list = types.MakeListFromInts(1, 2)
	got = removeNthFromEnd(list, 1)
	want = types.MakeListFromInts(1)
	require.Equal(t, got, want)

	list = types.MakeListFromInts(1, 2, 3, 4, 5)
	got = removeNthFromEnd2(list, 2)
	want = types.MakeListFromInts(1, 2, 3, 5)
	require.Equal(t, got, want)
}

func BenchmarkNthFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := types.MakeListFromInts(1, 2, 3, 4, 5)
		removeNthFromEnd(list, 2)
	}
}

func BenchmarkNthFromEnd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := types.MakeListFromInts(1, 2, 3, 4, 5)
		removeNthFromEnd2(list, 2)
	}
}
