package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wralith/go-leetcode/types"
)

func Test_isBalanced(t *testing.T) {
	tree := types.Ints2TreeNode([]int{3, 9, 20, types.NULL, types.NULL, 15, 7})
	got := isBalanced(tree)
	assert.True(t, got)

	tree = types.Ints2TreeNode([]int{1, 2, 2, 3, 3, types.NULL, types.NULL, 4, 4})
	got = isBalanced(tree)
	assert.False(t, got)

	tree = types.Ints2TreeNode([]int{1, types.NULL, 2, types.NULL, 3})
	got = isBalanced(tree)
	assert.False(t, got)
}
