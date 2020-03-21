package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))

}
func orangesRotting(grid [][]int) int {
	var rotten [][]int
	var time, all, rottens int
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i, v := range grid {
		for j, k := range v {
			if k != 0 {
				all++
			}
			if k == 2 {
				rotten = append(rotten, []int{i, j})
				rottens++
			}

		}
	}
	if all < rottens {
		return 0
	}

	for len(rotten)>0 {
		//var currotten [][]int
		curlen := len(rotten)
		for i:=0;i<curlen;i++ {
			cur := rotten[0]
			rotten = rotten[1:]
			for _, v := range dir {
				x := cur[0] + v[0]
				y := cur[1] + v[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					rotten = append(rotten, []int{x, y})
					rottens++
				}
			}
		}
		time++
		fmt.Println(rotten)
	}

	if rottens < all {
		fmt.Printf("rottens:%d, all:%d ", rottens, all)
		return -1
	}

	return time - 1

}
