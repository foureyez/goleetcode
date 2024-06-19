package stack

type MinStack struct {
	elems []Elem
}

type Elem struct {
	val    int
	minVal int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	minVal := val
	if len(this.elems) > 0 {
		prevElem := this.elems[len(this.elems)-1]
		if prevElem.minVal < minVal {
			minVal = prevElem.minVal
		}
	}

	elem := Elem{
		val:    val,
		minVal: minVal,
	}
	this.elems = append(this.elems, elem)
}

func (this *MinStack) Pop() {
	this.elems = this.elems[:len(this.elems)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1].val
}

func (this *MinStack) GetMin() int {
	return this.elems[len(this.elems)-1].minVal
}
