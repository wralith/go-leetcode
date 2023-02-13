package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProfit(t *testing.T) {
	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	require.Equal(t, 5, got)

	got = maxProfit([]int{7, 6, 4, 3, 1})
	require.Equal(t, 0, got)
}
