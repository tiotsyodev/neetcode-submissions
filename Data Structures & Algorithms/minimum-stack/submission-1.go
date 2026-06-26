type MinStack struct {
	ms []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		ms: []int{},
		minStack: []int{}, 
	}
}

func (this *MinStack) Push(val int) {

	this.ms = append(this.ms, val)

	if len(this.minStack) == 0 || this.minStack[len(this.minStack) - 1] >= val {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack) - 1])
	}
}

func (this *MinStack) Pop() {
	if len(this.ms) == 0 {
        return
    }

	this.ms = this.ms[:len(this.ms) - 1]
	this.minStack = this.minStack[:len(this.minStack) - 1]
}

func (this *MinStack) Top() int {
	if len(this.ms) == 0 {
		return 0
	}
	return this.ms[len(this.ms) - 1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
        return 0
    }
    return this.minStack[len(this.minStack)-1]
}
