package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	got := subsets([]int{1, 2, 3})
	want := [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{3},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}
	require.ElementsMatch(t, want, got)

	got = subsets([]int{0})
	want = [][]int{{}, {0}}
	require.ElementsMatch(t, want, got)
}
