package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *types.ListNode
		l2 *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want *types.ListNode
	}{
		{
			name: "Example 1",
			args: args{
				l1: types.MakeListFromInts(2, 4, 3),
				l2: types.MakeListFromInts(5, 6, 4),
			},
			want: types.MakeListFromInts(7, 0, 8),
		},
		{
			name: "Example 2",
			args: args{
				l1: types.MakeListFromInts(0),
				l2: types.MakeListFromInts(0),
			},
			want: types.MakeListFromInts(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			require.Equal(t, tt.want, got)
		})
	}
}
