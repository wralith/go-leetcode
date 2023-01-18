package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_reverseList(t *testing.T) {
	list := types.MakeListFromInts(1, 2, 3, 4, 5)
	want := types.MakeListFromInts(5, 4, 3, 2, 1)
	require.Equal(t, want, reverseList(list))
	list = types.MakeListFromInts(1, 2, 3, 4, 5)
	require.Equal(t, want, reverseListRecursive(list))
	list = types.MakeListFromInts(1, 2, 3, 4, 5)
	require.Equal(t, want, reverseListRecursiveHelper(list))

	list = types.MakeListFromInts(1, 2)
	want = types.MakeListFromInts(2, 1)
	require.Equal(t, want, reverseList(list))
	list = types.MakeListFromInts(1, 2)
	require.Equal(t, want, reverseListRecursive(list))
	list = types.MakeListFromInts(1, 2)
	require.Equal(t, want, reverseListRecursiveHelper(list))

	list = types.MakeListFromInts()
	want = types.MakeListFromInts()
	require.Equal(t, want, reverseList(list))
	list = types.MakeListFromInts()
	require.Equal(t, want, reverseListRecursive(list))
	list = types.MakeListFromInts()
	require.Equal(t, want, reverseListRecursiveHelper(list))
}
