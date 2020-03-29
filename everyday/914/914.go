package main

import "sort"

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _,v :=range(deck){
		_,ok := m[v]
		if !ok {
			m[v] = 1
		}else{
			m[v] ++
		}
	}

	var vs []int
	for _,v := range m{
		vs = append(vs,v)
	}
	if len(vs) == 1{
		if vs[0] > 1{
			return true
		}else{
			return false
		}
	}
	sort.Ints(vs)
	for vs[0] != vs[len(vs)-1]{
		sort.Ints(vs)
		for i := len(vs)-1;i>=1;i--{
			if vs[i]%vs[i-1]==0{
				vs[i] = vs[i-1]
			}else{
				vs[i] = vs[i]%vs[i-1]
			}
		}
	}
	if(vs[0]==1){
		return false
	}

	return true


}
func main() {
	
}
