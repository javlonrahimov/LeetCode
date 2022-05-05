package main

type MyStack struct {
	list []int
	top  int
}

func Constructor() MyStack {
	return MyStack{list: make([]int, 10)}
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
	this.top = x
}

func (this *MyStack) Pop() int {
	if !this.Empty() {
		this.list = this.list[:len(this.list)-1]
		pop := this.top
		if len(this.list) != 0 {
			this.top = this.list[len(this.list)-1]
		} else {
			this.top = 0
		}
		return pop
	}
	return 0
}

func (this *MyStack) Top() int {
	return this.top
}

func (this *MyStack) Empty() bool {
	return len(this.list) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
