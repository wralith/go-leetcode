package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test not testing anything though, be careful!
func TestCloneGraph(t *testing.T) {

	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node4 := Node{Val: 4}

	node1.Neighbors = []*Node{&node2, &node4}
	node2.Neighbors = []*Node{&node1, &node3}
	node3.Neighbors = []*Node{&node2, &node4}
	node4.Neighbors = []*Node{&node1, &node3}

	clone := cloneGraph(&node1)
	require.NotSame(t, node1, clone)
}
