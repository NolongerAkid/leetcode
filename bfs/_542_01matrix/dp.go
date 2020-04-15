package _542_01matrix

func updateMatrix(matrix [][]int) [][]int {
	r := len(matrix)
	l := len(matrix[0])
	res := make([][]int,r)
	for i:=0;i<r;i++{
		res[i] = make([]int,l)
		for j := 0; j < l; j++{
			if matrix[i][j] != 0{
				res[i][j] = 10000
				if i != 0 {
					res[i][j] = min(res[i][j],res[i-1][j]+1)
				}
				if j != 0 {
					res[i][j] = min(res[i][j],res[i][j-1]+1)
				}
			}
		}
	}
	for i:= r-1;i>=0;i--{
		for j := l-1;j >= 0; j--{
			if matrix[i][j] != 0{
				if i < r-1 {
					res[i][j] = min(res[i][j],res[i+1][j]+1)
				}
				if j < l-1 {
					res[i][j] = min(res[i][j],res[i][j+1]+1)
				}
			}
		}
	}
	return res
}
func min(a,b int)int {
	if a <= b {
		return a
	}else{
		return b
	}
}