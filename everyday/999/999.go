package main

var dir = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
func numRookCaptures(board [][]byte) int {
	s := findStart(board)
	if s[0] == -1 && s[1] == -1 {
		return 0
	}
	count := 0
	for _,v := range dir{
		x := s[0]
		y := s[1]

		x += v[0]
		y += v[1]
		for x >=0 && y >= 0 && x <8 && y < 8{

			t := board[x][y]
			if t == '.'{
				x += v[0]
				y += v[1]
				continue
			}else if t == 'B'{
				break
			}else if t == 'p'{
				x += v[0]
				y += v[1]
				count += 1
				break
			}

		}
	}
	return count


}

func findStart(board [][]byte) []int{
	for i,m := range board{
		for j,n := range m{
			if string(n) == "R"{
				return []int{i,j}
			}
		}
	}
	return []int{-1,-1}
}

func main() {
	
}
