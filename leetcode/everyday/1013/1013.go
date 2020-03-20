package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _,v := range A {
		sum += v
	}
	if sum % 3 != 0{
		return false
	}
	fmt.Println("sum=",sum)
	temp,i := 0,0
	for ;i<len(A);i++ {
		temp += A[i]
		fmt.Println("1temp=",temp)
		if temp == sum/3{
			i++
			break
		}
	}
	if temp != sum/3 {
		return false
	}

	for ;i<len(A);i++{
		temp+=A[i]
		fmt.Println("2temp=",temp)
		if temp == (sum*2)/3{
			return true
		}
	}
	return false


}
func main() {
	fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
}
