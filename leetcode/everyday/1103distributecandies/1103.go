package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	res := make([]int,num_people)
	var curpos int//当前位置
	for candies > 0{
		purcandies := curpos +1//最多分的糖果
		if candies > purcandies{
			res[curpos%num_people]+= purcandies
			candies -= purcandies
		}else{
			res[curpos%num_people]+= candies
			candies = 0
		}
		curpos++
	}
	return res



}

func main() {
	fmt.Println(distributeCandies(7,4))
}
