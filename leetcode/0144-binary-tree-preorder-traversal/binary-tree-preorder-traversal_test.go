package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_preorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *types.TreeNode
		want []int
	}{
		{
			name: "Ex",
			root: types.Ints2TreeNode([]int{1, types.NULL, 2, 3}),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		got := preorderTraversal(tt.root)
		require.Equal(t, tt.want, got)
	}
}
