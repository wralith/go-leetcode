package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxConsecutiveOnes2(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example 1",
			args: []int{1, 0, 1, 1, 0, 1},
			want: 4,
		},
		{
			name: "Base Case 1",
			args: []int{0},
			want: 1,
		},
		{
			name: "Base Case 2 2",
			args: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes2(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
