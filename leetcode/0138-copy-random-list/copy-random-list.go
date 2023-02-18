package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copy := map[*Node]*Node{}

	curr := head
	for curr != nil {
		clone := &Node{Val: curr.Val}
		copy[curr] = clone
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		clone := copy[curr]
		clone.Next = copy[curr.Next]
		clone.Random = copy[curr.Random]
		curr = curr.Next
	}

	return copy[head]
}
