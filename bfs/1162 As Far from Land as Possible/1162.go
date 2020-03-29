package main

import "fmt"

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxDistance(grid [][]int) int {
	queue := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	b := false
	cur := []int{0,0}
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		for _, d := range dir {
			nx := cur[0] + d[0]
			ny := cur[1] + d[1]
			if nx<0||nx>=len(grid)||ny<0||ny>=len(grid[0])||grid[nx][ny]!=0{
				continue
			}
			b = true
			grid[nx][ny] = grid[cur[0]][cur[1]]+1
			queue = append(queue,[]int{nx,ny})

		}
	}
	fmt.Println(grid)
	if !b {
		return -1
	}
	return grid[cur[0]][cur[1]]-1

}
func main() {
	fmt.Println(maxDistance([][]int{{1,0,0},{0,0,0},{0,0,0}}))

}
