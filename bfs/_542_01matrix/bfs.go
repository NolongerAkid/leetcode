package _542_01matrix

var dir = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
func updateMatrixbfs(matrix [][]int) [][]int {
	var queue [][]int
	for i := 0; i < len(matrix); i++{
		for j := 0; j<len(matrix[0]); j++ {
			if matrix[i][j] == 0{
				queue = append(queue, []int{i,j})
			}else{
				matrix[i][j] = -1
			}
		}
	}
	for len(queue)>0{
		x := queue[0][0]
		y := queue[0][1]
		queue = queue[1:]
		for _,d := range dir{
			nx := x + d[0]
			ny := y + d[1]
			if nx >=0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] == -1{
				matrix[nx][ny] = matrix[x][y] + 1
				queue = append(queue, []int{nx,ny})
			}
		}
	}
	return matrix

}
