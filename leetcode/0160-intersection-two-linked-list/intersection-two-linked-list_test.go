package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *types.ListNode
		headB *types.ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 2",
			args: args{
				headA: types.MakeListFromInts(),
				headB: types.MakeListFromInts(),
			},
			want: []int(nil),
		},
		{
			name: "Example 3",
			args: args{
				headA: types.MakeListFromInts(2, 6, 4),
				headB: types.MakeListFromInts(1, 5),
			},
			want: []int(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList := getIntersectionNode(tt.args.headA, tt.args.headB)
			got := types.MakeIntSliceFromList(gotList)

			require.Equal(t, tt.want, got)
		})
	}
}
