package main
var dir = [][]int{{-1,0},{1,0},{0,1},{0,-1}}
func surfaceArea(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	count := 0
	for i,v := range grid{
		for j,k :=range v{
			count += 6*k
			if k >0 {
				count -= 2*(k-1) //自身重合面
			}

			for _,d :=range dir{
				nr := i + d[0]
				nc := j + d[1]
				if nc >=0 && nc < c && nr >= 0 && nr < r{
					count -= min(k,grid[nr][nc])//与旁边重合面
				}
			}
		}
	}
	return count

}

func min(a,b int) int {
	if a <= b{
		return a
	}else{
		return b
	}
}
func main() {
	
}
