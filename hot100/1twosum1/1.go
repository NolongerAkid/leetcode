package main

import "fmt"

func twoSum(nums []int, target int) []int {
	imap := make(map[int]int)
	for i,v := range nums{
		res := imap[v]
		if res == 0{
			fmt.Printf("imap[%d]=0 ",v)
			imap[target-v] = i+1
			fmt.Println(imap)
		}else{
			return []int{imap[v]-1,i}
		}
	}
	fmt.Println(imap)
	return []int{-1,-1}
}
func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
}
