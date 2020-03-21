package main

import "fmt"

func maxProfit(prices []int) int {
	maxprofix := 0
	min := prices[0]
	for _,v := range prices {
		if v < min {
			min = v
		}else{
			maxprofix = max(maxprofix,v-min)
		}
	}
	return maxprofix
}
func max(a,b int) int {
	if a > b {
		return a
	}else{
		return b
	}
}
func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
