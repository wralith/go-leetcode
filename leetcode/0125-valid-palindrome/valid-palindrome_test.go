package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	got := isPalindrome("A man, a plan, a canal: Panama")
	require.True(t, got)

	got = isPalindrome("race a car")
	require.False(t, got)

	got = isPalindrome(" ")
	require.True(t, got)
}
