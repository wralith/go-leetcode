package leetcode

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, token := range tokens {
		v, isValidInt := strconv.Atoi(token)
		if isValidInt == nil {
			stack.PushBack(v)
			continue
		}

		n2 := stack.Remove(stack.Back()).(int)
		n1 := stack.Remove(stack.Back()).(int)
		switch token {
		case "+":
			stack.PushBack(n1 + n2)
		case "-":
			stack.PushBack(n1 - n2)
		case "*":
			stack.PushBack(n1 * n2)
		case "/":
			stack.PushBack(n1 / n2)
		}
	}

	return stack.Remove(stack.Back()).(int)
}

func evalRPNSlice(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		v, isValidInt := strconv.Atoi(token)
		if isValidInt == nil {
			stack = append(stack, v)
			continue
		}

		n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		switch token {
		case "+":
			stack = append(stack, n1+n2)
		case "-":
			stack = append(stack, n1-n2)
		case "*":
			stack = append(stack, n1*n2)
		case "/":
			stack = append(stack, n1/n2)
		}
	}

	return stack[0]
}
