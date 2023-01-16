package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinEeatingSpeed(t *testing.T) {
	want := 4
	got := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	require.Equal(t, want, got)

	want = 30
	got = minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)
	require.Equal(t, want, got)

	want = 23
	got = minEatingSpeed([]int{30, 11, 23, 4, 20}, 6)
	require.Equal(t, want, got)
}
