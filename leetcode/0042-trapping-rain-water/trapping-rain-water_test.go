package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trap(t *testing.T) {
	got := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	want := 6
	require.Equal(t, want, got)

	got = trap([]int{4, 2, 0, 3, 2, 5})
	want = 9
	require.Equal(t, want, got)

	got = trap([]int{4, 2, 3})
	want = 1
	require.Equal(t, want, got)
}
