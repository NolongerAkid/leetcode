package main
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	degree := make([]int,n)
	link := make([][]int,n)
	for _,v := range edges{
		degree[v[0]]++
		degree[v[1]]++
		link[v[0]] = append(link[v[0]],v[1])
		link[v[1]] = append(link[v[1]],v[0])
	}
	queue := make([]int,0)
	for i,v := range degree{
		if v == 1{
			queue = append(queue,i)
		}
	}
	rst := n
	for rst>2{
		l := len(queue)
		rst -= l

		for i:=0;i<l;i++{
			c := queue[0]
			queue = queue[1:]

			for _,v := range link[c]{
				if degree[v]>0{
					degree[v]--
				}
				if degree[v] == 1{
					queue = append(queue, v)
				}

			}

		}

	}
	return queue

}
func main() {
	
}
