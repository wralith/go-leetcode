package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func TestIsValidBST(t *testing.T) {
	tree := types.Ints2TreeNode([]int{2, 1, 3})
	require.Equal(t, true, isValidBST(tree))

	tree = types.Ints2TreeNode([]int{5, 1, 4, types.NULL, types.NULL, 3, 6})
	require.Equal(t, false, isValidBST(tree))
}
