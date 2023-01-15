package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestConsecutive(t *testing.T) {
	got := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	want := 4
	require.Equal(t, want, got)

	got = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	want = 9
	require.Equal(t, want, got)

	got = longestConsecutive([]int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3})
	want = 5
	require.Equal(t, want, got)
}
