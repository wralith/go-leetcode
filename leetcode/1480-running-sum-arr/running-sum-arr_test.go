package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_runningSum(t *testing.T) {
	got := runningSum([]int{1, 2, 3, 4})
	want := []int{1, 3, 6, 10}
	require.Equal(t, want, got)

	got = runningSum([]int{1, 1, 1, 1, 1})
	want = []int{1, 2, 3, 4, 5}
	require.Equal(t, want, got)

	got = runningSum([]int{3, 1, 2, 10, 1})
	want = []int{3, 4, 6, 16, 17}
	require.Equal(t, want, got)
}
