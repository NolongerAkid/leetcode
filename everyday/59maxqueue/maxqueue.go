package main

type MaxQueue struct {
	max []int //max[i]代表i之后的最大值
	queue []int
	maxAll int

}


func Constructor() MaxQueue {
	return MaxQueue{}

}


func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		this.maxAll = -1 //题目要求 队列为空时返回-1
	}
	return this.maxAll
}


func (this *MaxQueue) Push_back(value int)  {
	if value > this.maxAll{
		this.maxAll = value
	}
	this.queue = append(this.queue,value)
	this.max = append(this.max,value)
	for l := len(this.max)-2;l>=0;l--{
		if this.max[l] >value{
			break
		}else{
			this.max[l] = value
		}
	}
}


func (this *MaxQueue) Pop_front() int {
	var res int
	if len(this.queue) == 0{
		return -1
	}else{
		res = this.queue[0]
	}

	this.queue = this.queue[1:]
	this.max = this.max[1:]
	if len(this.max) == 0 {
		this.maxAll = -1
	}else{
		this.maxAll = this.max[0]
	}

	return res
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
func main() {
	
}
