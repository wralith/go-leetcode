package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_reorderList(t *testing.T) {
	head := types.MakeListFromInts(1, 2, 3, 4)
	want := types.MakeListFromInts(1, 4, 2, 3)
	reorderList(head)
	require.Equal(t, want, head)

	head = types.MakeListFromInts(1, 2, 3, 4, 5)
	want = types.MakeListFromInts(1, 5, 2, 4, 3)
	reorderList(head)
	require.Equal(t, want, head)
}
