package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "Example 1",
			arg:  "abcabcbb",
			want: 3,
		},
		{
			name: "Example 2",
			arg:  "bbbbb",
			want: 1,
		},
		{
			name: "Example 3",
			arg:  "pwwkew",
			want: 3,
		},
		{
			name: "Example 4",
			arg:  "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.arg)
			require.Equal(t, tt.want, got)
		})
	}
}
