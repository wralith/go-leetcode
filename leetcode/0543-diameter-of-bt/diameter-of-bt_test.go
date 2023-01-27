package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	tree := types.Ints2TreeNode([]int{1, 2, 3, 4, 5})
	got := diameterOfBinaryTree(tree)
	require.Equal(t, 3, got)

	tree = types.Ints2TreeNode([]int{1, 2})
	got = diameterOfBinaryTree(tree)
	require.Equal(t, 1, got)

	tree = types.Ints2TreeNode([]int{2, 3, types.NULL, 1})
	got = diameterOfBinaryTree(tree)
	require.Equal(t, 2, got)
}
