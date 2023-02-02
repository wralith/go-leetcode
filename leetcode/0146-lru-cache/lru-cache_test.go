package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	lru := Constructor(2)

	putAndTestCacheState(t, lru, map[int]int{1: 1}, 1, 1)
	putAndTestCacheState(t, lru, map[int]int{1: 1, 2: 2}, 2, 2)

	got := lru.Get(1)
	require.Equal(t, 1, got)

	putAndTestCacheState(t, lru, map[int]int{1: 1, 3: 3}, 3, 3)

	got = lru.Get(2)
	require.Equal(t, -1, got)

	putAndTestCacheState(t, lru, map[int]int{4: 4, 3: 3}, 4, 4)

	got = lru.Get(1)
	require.Equal(t, -1, got)

	got = lru.Get(3)
	require.Equal(t, 3, got)
	got = lru.Get(4)
	require.Equal(t, 4, got)
}

func putAndTestCacheState(t *testing.T, lru LRUCache, want map[int]int, key, value int) {
	lru.Put(key, value)
	got := retrieveMapFromCache(lru.cache)
	require.Equal(t, want, got)
}

func retrieveMapFromCache(cache map[int]*Node) map[int]int {
	res := make(map[int]int, len(cache))
	for key, value := range cache {
		res[key] = value.value
	}
	return res
}
