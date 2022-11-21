package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_sortedArrayToBST(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want *types.TreeNode
	}{
		{
			name: "Example 1",
			nums: []int{-10, -3, 0, 5, 9},
			want: types.Ints2TreeNode([]int{0, -3, 9, -10, types.NULL, 5}),
			// want: types.Ints2TreeNode([]int{0, -10, 5, types.NULL, -3, types.NULL, 9}),
		},
		{
			name: "Example 2",
			nums: []int{1, 3},
			want: types.Ints2TreeNode([]int{3, 1}),
			// want: types.Ints2TreeNode([]int{1, types.NULL, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
