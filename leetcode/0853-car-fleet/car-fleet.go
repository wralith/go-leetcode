package leetcode

import (
	"sort"
)

type car struct {
	position int
	speed    int
}

type SortByPos []car

func (a SortByPos) Len() int           { return len(a) }
func (a SortByPos) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByPos) Less(i, j int) bool { return a[i].position > a[j].position }

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n < 2 {
		return n
	}

	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{position[i], speed[i]}
	}
	sort.Sort(SortByPos(cars))

	stack := []float32{}

	for _, car := range cars {
		stack = append(stack, (float32(target-car.position))/float32(car.speed))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1] // pop
		}
	}

	return len(stack)
}
