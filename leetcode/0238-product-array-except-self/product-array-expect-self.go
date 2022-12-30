package leetcode

func productExceptSelf(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		res = append(res, 1)
	}

	prefix, postfix := 1, 1
	for i := 0; i < len(nums); i++ {
		res[i] = prefix
		prefix *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}
