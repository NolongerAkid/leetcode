package main

import "fmt"

func countCharacters(words []string, chars string) int {
	m := make(map[string]int)
	for _,c := range chars {
		m[string(c)] ++
	}
	count := 0
	for _,v := range words {
		temp := cloneTags(m)//需要深拷贝
		t:=0
		for _,c := range v {

			if temp[string(c)] == 0 {
				break
			}else{
				temp[string(c)]--
				t++
			}
		}
		if t == len(v){
			count += t
			/*fmt.Println("match "+v)
			fmt.Println(m)*/
		}
	}
	return count


}
func cloneTags(tags map[string]int) map[string]int {
	cloneTags := make(map[string]int)
	for k, v := range tags {
		cloneTags[k] = v
	}
	return cloneTags
}


func main() {
	fmt.Println(countCharacters([]string{"cat","bt","hat","tree"},"atach" ))
}
