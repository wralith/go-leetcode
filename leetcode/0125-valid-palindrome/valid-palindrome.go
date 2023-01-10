package leetcode

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isLatinOrNumber(s[left]) {
			left++
		} else if !isLatinOrNumber(s[right]) {
			right--
		} else if toUpperASCII(s[left]) == toUpperASCII(s[right]) {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

func isLatinOrNumber(char byte) bool {
	return ((char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9'))
}

func toUpperASCII(char byte) byte {
	if char >= 'a' && char <= 'z' {
		return char - 32
	}
	return char
}
