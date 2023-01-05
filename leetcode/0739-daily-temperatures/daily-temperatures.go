package leetcode

func dailyTemperatures(temperatures []int) (res []int) {
	s := TempStack{}
	for i := 0; i < len(temperatures); i++ {
		res = append(res, 0)
	}

	for i, cur := range temperatures {
		for s.len() > 0 && cur > s.last() {
			_, index := s.pop()
			res[index] = (i - index)
		}
		s.append(cur, i)
	}

	return
}

type TempStack struct {
	Stack [][]int
}

func (s TempStack) len() int {
	return len(s.Stack)
}

func (s TempStack) last() int {
	return s.Stack[s.len()-1][1]
}

func (s *TempStack) append(t, i int) {
	s.Stack = append(s.Stack, []int{i, t})
}

func (s *TempStack) pop() (temp, index int) {
	n := len(s.Stack) - 1
	index = s.Stack[n][0]
	temp = s.Stack[n][1]
	s.Stack = s.Stack[:n]
	return
}
