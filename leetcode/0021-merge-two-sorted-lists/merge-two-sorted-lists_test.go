package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_mergeTwoLists(t *testing.T) {
	v1 := types.MakeListFromInts(1, 2, 4)
	v2 := types.MakeListFromInts(1, 3, 4)
	want := types.MakeListFromInts(1, 1, 2, 3, 4, 4)
	got := mergeTwoLists(v1, v2)
	require.Equal(t, types.MakeIntSliceFromList(want), types.MakeIntSliceFromList(got))

	v1 = types.MakeListFromInts()
	v2 = types.MakeListFromInts(0)
	want = types.MakeListFromInts(0)
	got = mergeTwoLists(v1, v2)
	require.Equal(t, types.MakeIntSliceFromList(want), types.MakeIntSliceFromList(got))

	v1 = types.MakeListFromInts()
	v2 = types.MakeListFromInts()
	want = types.MakeListFromInts()
	got = mergeTwoLists(v1, v2)
	require.Equal(t, types.MakeIntSliceFromList(want), types.MakeIntSliceFromList(got))
}
