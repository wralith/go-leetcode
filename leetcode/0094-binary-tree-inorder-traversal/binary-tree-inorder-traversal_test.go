package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *types.TreeNode
		want []int
	}{
		{
			name: "Example 1",
			root: types.Ints2TreeNode([]int{1, types.NULL, 2, 3}),
			want: []int{1, 3, 2},
		},
		{
			name: "Example 1",
			root: types.Ints2TreeNode([]int{}),
			want: []int(nil),
		},
		{
			name: "Example 3",
			root: types.Ints2TreeNode([]int{1}),
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inorderTraversal(tt.root)
			require.Equal(t, tt.want, got)
		})
	}
}
