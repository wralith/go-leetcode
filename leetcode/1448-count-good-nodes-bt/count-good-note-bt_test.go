package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_goodNodes(t *testing.T) {
	root := types.Ints2TreeNode([]int{3, 1, 4, 3, types.NULL, 1, 5})
	require.Equal(t, 4, goodNodes(root))

	root = types.Ints2TreeNode([]int{3, 3, types.NULL, 4, 2})
	require.Equal(t, 3, goodNodes(root))

	root = types.Ints2TreeNode([]int{1})
	require.Equal(t, 1, goodNodes(root))

	root = types.Ints2TreeNode([]int{2, types.NULL, 4, 10, 8, types.NULL, types.NULL, 4})
	require.Equal(t, 4, goodNodes(root))
}
