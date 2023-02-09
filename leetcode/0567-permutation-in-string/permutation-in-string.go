package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	left, count := 0, newAsciiArr(s1)

	for right := range s2 {
		count[s2[right]-'a']--
		// All zero
		if count == [26]int{} {
			return true
		}

		if right+1 >= len(s1) {
			count[s2[left]-'a']++
			left++
		}
	}

	return false
}

func newAsciiArr(s string) (arr [26]int) {
	for i := range s {
		arr[s[i]-'a']++
	}
	return
}
