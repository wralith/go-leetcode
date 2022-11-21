package leetcode

import "github.com/wralith/go-leetcode/types"

func sortedArrayToBST(nums []int) *types.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	rootVal := nums[mid]
	root := &types.TreeNode{Val: rootVal}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

/* // -- Also accepted solution in leetcode:
// But tests won't work since it generates different tree, use commented ones

func sortedArrayToBST(nums []int) *types.TreeNode {
	start, end := 0, len(nums)-1
	return buildBST(nums, start, end)

}

func buildBST(nums []int, start, end int) *types.TreeNode {

	if start > end {
		return nil
	}

	mid := (start + end) / 2

	cur := &types.TreeNode{Val: nums[mid]}

	cur.Left = buildBST(nums, start, mid-1)
	cur.Right = buildBST(nums, mid+1, end)

	return cur

}

*/
