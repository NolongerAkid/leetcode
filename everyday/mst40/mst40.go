package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[0:k]

}

func main() {

	fmt.Println(getLeastNumbers([]int{1, 2, 3}, 2))

}
