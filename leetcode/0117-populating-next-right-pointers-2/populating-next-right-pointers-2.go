package leetcode

import "github.com/wralith/go-leetcode/types"

type Node *types.TreeNodeWithNext

func connect(root Node) Node {
	head := root

	for head != nil {
		new := Node(&types.TreeNodeWithNext{}) // TODO: It is ugly
		temp := new

		for head != nil { // iterate after head changed again
			if head.Left != nil { // if there is a left child
				temp.Next = head.Left // temp will point to that
				temp = temp.Next
			}
			if head.Right != nil { // if there is a right child
				temp.Next = head.Right // connect where the temp is (left) to right
				temp = temp.Next
			}
			head = head.Next // jump to the right side of the tree
		}
		head = new.Next
	}

	return root
}
