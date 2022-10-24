package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPeakElement(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example 1",
			args: []int{1, 2, 3, 1},
			want: 2,
		},
		{
			name: "Example 2",
			args: []int{1, 2, 1, 3, 5, 6, 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
