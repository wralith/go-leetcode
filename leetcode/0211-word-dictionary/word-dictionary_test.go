package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordDictionary(t *testing.T) {
	dict := Constructor()

	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	require.False(t, dict.Search("pad"))
	require.True(t, dict.Search("bad"))
	require.True(t, dict.Search(".ad"))
	require.True(t, dict.Search("b.."))
}
