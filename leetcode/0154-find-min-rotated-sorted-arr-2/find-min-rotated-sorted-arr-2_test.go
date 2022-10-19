package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{7, 7, 7, 1, 6, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin2(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
