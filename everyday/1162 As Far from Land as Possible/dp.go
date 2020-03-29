package main

import "fmt"

func MaxDistanceDp(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				continue
			} else {
				grid[i][j] =minInt(val(grid, i-1, j), val(grid, i, j-1)) + 1

			}

		}
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			} else {
				grid[i][j] = minInt(grid[i][j],minInt(val(grid, i+1, j), val(grid, i, j+1)) + 1)

			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			max = maxInt(max, grid[i][j])

		}
	}
	fmt.Println(grid)
	if max >= 100{
		return -1
	}

	return max-1

}
func val(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return 100
	}
	return grid[x][y]

}
func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}

}
func minInt(a,b int) int {
	if a <= b{
		return a
	}else{
		return b
	}
}
