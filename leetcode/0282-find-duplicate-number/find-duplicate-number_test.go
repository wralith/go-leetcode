package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDuplicate(t *testing.T) {
	got := findDuplicate([]int{1, 3, 4, 2, 2})
	require.Equal(t, 2, got)

	got = findDuplicate([]int{3, 1, 3, 4, 2})
	require.Equal(t, 3, got)
}
