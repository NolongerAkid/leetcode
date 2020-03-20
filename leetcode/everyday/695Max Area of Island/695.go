package main



var dir [][]int = [][]int{{-1,0},{1,0},{0,1},{0,-1}}
func maxAreaOfIsland(grid [][]int) int {
	cur,max := 0,0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			bfs(i,j,grid,&cur)
			max = maxInt(cur,max)
			cur = 0
		}
	}
	return max

}
func bfs(x,y int,grid [][]int,cur *int){
	if x>=0 && x < len(grid)&&y>=0&&y<len(grid[0])&&grid[x][y]==1{
		grid[x][y] = 0
		*cur ++
		for _,v := range dir{
			bfs(x+v[0],y+v[1],grid,cur)
		}

	}




}
func maxInt(a,b int) int{
	if a>b {
		return a
	}else{
		return b
	}
}

func main() {
	
}
