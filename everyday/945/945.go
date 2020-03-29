package main

import "sort"

func minIncrementForUnique(A []int) int {
	count := 0
	sort.Ints(A)
	for i := 1;i<len(A);i++{
		if A[i] <= A[i-1]{
			t := A[i]
			A[i] = A[i-1] + 1
			count += A[i] - t
		}
	}
	return count

}
func main() {
	
}
