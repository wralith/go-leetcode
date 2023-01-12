package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_carFleet(t *testing.T) {
	want := 3
	got := carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
	require.Equal(t, want, got)

	want = 1
	got = carFleet(10, []int{3}, []int{3})
	require.Equal(t, want, got)

	want = 1
	got = carFleet(100, []int{0, 2, 4}, []int{4, 2, 1})
	require.Equal(t, want, got)
}
