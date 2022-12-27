package leetcode

var m map[rune]rune = map[rune]rune{'}': '{', ']': '[', ')': '('}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if _, ok := m[v]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == m[v] {
				stack = pop(stack)
			} else {
				return false
			}
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}

func pop(stack []rune) []rune {
	if len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}

	return stack
}
