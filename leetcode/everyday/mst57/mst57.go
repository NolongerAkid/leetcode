package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	l,r := 1,2
	sum := 3
	var res [][]int
	for l<r && l < target/2 {
		var temp []int
		for sum <= target {
			r++
			sum+=r
			if sum == target{
				for i:=l ; i<=r ; i++{
					temp = append(temp,i)
				}
				res = append(res,temp)
			}
		}

		for sum > target && l < r {
			sum -= l
			l++
			if sum == target{
				temp = []int{}
				for i:=l ; i<=r ; i++{

					temp = append(temp,i)
				}
				res = append(res,temp)
			}
		}




	}
	return res

}
func main() {
	fmt.Println(findContinuousSequence(15))
}
