package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := types.Ints2TreeNode([]int{6, 2, 8, 0, 4, 7, 9, types.NULL, types.NULL, 3, 5})
	p := types.Ints2TreeNode([]int{2, 0, 4, types.NULL, types.NULL, 3, 5})
	q := types.Ints2TreeNode([]int{8, 7, 9})
	got := lowestCommonAncestor(root, p, q)

	require.Equal(t, 6, got.Val)
}
