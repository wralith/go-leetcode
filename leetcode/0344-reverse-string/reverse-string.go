package leetcode

func reverseString(s []byte) {
	// Go snippet for reverse array =)
	// To pointers go through center from both ends while changing bytes
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}
