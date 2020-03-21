package main
type MyStack struct {
	queue1 []int
	queue2 []int

}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{[]int{},[]int{},}

}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue1 = append(this.queue1,x)

}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {

	var res int
	if len(this.queue1) != 0 {
		res = this.queue1[len(this.queue1)-1]
		this.queue2 = this.queue1[0:len(this.queue1)-1]
		this.queue1 = this.queue1[0:0]
	}else{
		res = this.queue2[len(this.queue2)-1]
		this.queue1 = this.queue2[0:len(this.queue2)-1]
		this.queue2= this.queue2[0:0]

	}
	return res

}


/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.queue1) != 0{
		return this.queue1[len(this.queue1)-1]
	}else{
		return this.queue2[len(this.queue2)-1]
	}

}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1)==0&&len(this.queue2)==0

}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	println("Pop:",obj.Pop())
	println("Top:",obj.Top())
	println("Empty:",obj.Empty())


}
