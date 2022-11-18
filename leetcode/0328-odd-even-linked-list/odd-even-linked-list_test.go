package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_oddEvenList(t *testing.T) {
	tests := []struct {
		name string
		head *types.ListNode
		want *types.ListNode
	}{
		{
			name: "Example 1",
			head: types.MakeListFromInts(1, 2, 3, 4, 5),
			want: types.MakeListFromInts(1, 3, 5, 2, 4),
		},
		{
			name: "Example 2",
			head: types.MakeListFromInts(2, 1, 3, 5, 6, 4, 7),
			want: types.MakeListFromInts(2, 3, 6, 7, 1, 5, 4),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := oddEvenList(test.head)
			require.Equal(t, test.want, got)
		})
	}
}
