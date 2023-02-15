package leetcode

func subsets(nums []int) (res [][]int) {
	var curr []int
	var dfs func(i int)

	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, curr...))
			return
		}
		curr = append(curr, nums[i])
		dfs(i + 1)
		curr = curr[:len(curr)-1]
		dfs(i + 1)
	}

	dfs(0)
	return res
}
