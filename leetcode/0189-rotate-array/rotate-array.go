package leetcode

// rotate rotates an array to right by given value with O(n) time and space complexity i guess...
// func rotate(nums []int, k int) {

// 	if k > len(nums) || len(nums) == 1 {
// 		k = k % len(nums)
// 	}

// 	startPart := nums[:len(nums)-k]
// 	endPart := nums[len(nums)-k:]

// 	newNums := append(endPart, startPart...) // Rotates here

// 	for i := range nums {
// 		nums[i] = newNums[i] // Mutates nums
// 	}

// }

// Two Pointer solution with O(1) space complexity, I guess...
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums)     // Reverses whole array    [1, 2, 3, 4, 5] -> [5, 4, 3, 2, 1]
	reverse(nums[:k]) // Reverses the left side  [5, 4...]       -> [4, 5...]
	reverse(nums[k:]) // Reverses the right side [...3, 2, 1]    -> [...1, 2, 3]

}

func reverse(nums []int) {
	start := 0
	end := len(nums) - 1

	for start < len(nums)/2 {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
