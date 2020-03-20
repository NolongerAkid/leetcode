package main

import "fmt"

func lengthOfLIS(nums []int) int {
	longest := 0
	dp := make([]int,len(nums))
	dp[0] = 1
	for i,v := range nums {
		for j := i-1 ; j >= 0 ; j--{
			if v > nums[j]{
				dp[i] = maxInt(dp[i],dp[j]+1)
				longest = maxInt(dp[i],longest)
				//break
			}
		}

	}
	fmt.Println(dp)
	return longest

}
func maxInt(a,b int) int{
	if a >=b {
		return a
	}else {
		return b
	}
}

func main() {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
}
