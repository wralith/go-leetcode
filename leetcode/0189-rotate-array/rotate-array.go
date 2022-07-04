package leetcode

// rotate rotates an array to right by given value with O(n) time and space complexity i guess...
func rotate(nums []int, k int) {

	if k > len(nums) || len(nums) == 1 {
		k = k % len(nums)
	}

	startPart := nums[:len(nums)-k]
	endPart := nums[len(nums)-k:]

	newNums := append(endPart, startPart...) // Rotates here

	for i := range nums {
		nums[i] = newNums[i] // Mutates nums
	}

}
