package leetcode

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var left, right = 0, 1
	var counter = 0

	// [1, 1, 2, 3, 3, 3, 4]
	//           ^     ^->^    // Left will stay, right continue
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
			counter = 0
		} else if nums[left] == nums[right] && counter < 1 {
			counter++
			left++
			nums[left] = nums[right]
		}
		right++
	}

	return left + 1

}

// Lovely Answer
/*
	slow := 2
	fast := 2
	for ; fast < len(nums); {
		if nums[slow - 2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
*/
