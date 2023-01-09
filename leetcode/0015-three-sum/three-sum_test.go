package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSum(t *testing.T) {
	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	require.Equal(t, want, got)

	want = [][]int{{0, 0, 0}}
	got = threeSum([]int{0, 0, 0})
	require.Equal(t, want, got)

	want = [][]int{{0, 0, 0}}
	got = threeSum([]int{0, 0, 0, 0})
	require.Equal(t, want, got)
}
