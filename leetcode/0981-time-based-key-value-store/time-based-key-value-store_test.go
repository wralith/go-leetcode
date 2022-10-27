package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTimeBasedKeyValueStore(t *testing.T) {
	timeMap := make(TimeMap)
	timeMap.Set("foo", "bar", 1)

	got := timeMap.Get("foo", 1)
	require.Equal(t, "bar", got)

	got = timeMap.Get("foo", 3)
	require.Equal(t, "bar", got)

	timeMap.Set("foo", "bar2", 4)

	got = timeMap.Get("foo", 4)
	require.Equal(t, "bar2", got)
}
