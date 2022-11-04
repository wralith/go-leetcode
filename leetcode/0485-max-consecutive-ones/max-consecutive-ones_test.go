package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example 1",
			args: []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			name: "Example 2",
			args: []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxConsecutiveOnes(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
