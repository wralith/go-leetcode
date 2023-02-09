package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkInclusion(t *testing.T) {
	got := checkInclusion("ab", "eidbaooo")
	require.True(t, got)

	got = checkInclusion("ab", "eidboaoo")
	require.False(t, got)
}
