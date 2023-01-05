package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dailyTemperatures(t *testing.T) {
	got := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	want := []int{1, 1, 4, 2, 1, 1, 0, 0}
	require.Equal(t, got, want)

	got = dailyTemperatures([]int{30, 40, 50, 60})
	want = []int{1, 1, 1, 0}
	require.Equal(t, got, want)

	got = dailyTemperatures([]int{30, 60, 90})
	want = []int{1, 1, 0}
	require.Equal(t, got, want)
}
