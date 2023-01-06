package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	require.Equal(t, got, want)

	got = twoSum([]int{3, 2, 4}, 6)
	want = []int{1, 2}
	require.Equal(t, got, want)

	got = twoSum([]int{3, 3}, 6)
	want = []int{0, 1}
	require.Equal(t, got, want)
}
