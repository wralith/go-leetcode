package leetcode

type MinStack struct {
	data []int
	min  map[int]int
}

func MinStackConstructor() MinStack {
	return MinStack{
		data: []int{},
		min:  make(map[int]int),
	}
}

func (this *MinStack) Push(val int) {
	n := len(this.data)

	// adding to min map
	if len(this.data) == 0 {
		this.min[n] = val
	} else {
		if this.min[n-1] > val {
			this.min[n] = val
		} else {
			this.min[n] = this.min[n-1]
		}
	}

	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	n := len(this.data)
	if n == 0 {
		return
	}

	delete(this.min, n-1)
	this.data = this.data[:n-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.data)-1]
}
