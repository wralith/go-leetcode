package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxArea(t *testing.T) {
	got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	want := 49
	require.Equal(t, want, got)

	got = maxArea([]int{1, 1})
	want = 1
	require.Equal(t, want, got)

	got = maxArea([]int{2, 3, 4, 5, 18, 17, 6})
	want = 17
	require.Equal(t, want, got)
}
