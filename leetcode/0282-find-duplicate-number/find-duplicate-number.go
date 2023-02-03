package leetcode

// Problem wants constant space complexity, so no hash table :(

// Floyd's tortoise and hare. Go Wiki for more!
// Floyd's cycle-finding algorithm...
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	// Search where slow and fast intersects first!
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// Hare and tortoise
	// We wait them to run and meet at some point
	// When they meet we took tortoise and put it at the beginning
	// Now hare should also move as slow as tortoise
	// We wait them to meet again and this is the place we are looking for!

	// So: Put slow back to the start
	// 	   Fast stays but it also will move slow from now on!
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	// Where they met is the beginning of the cycle
	return slow // or return fast they are at the same position
}

// For more details and visual explanations:
// https://leetcode.com/problems/find-the-duplicate-number/solutions/127594/official-solution/
