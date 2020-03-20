package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	max := 0
	mark := 0
	for i,v := range s {
		if m[string(v)] == 0{
			max = maxInt(max,i+1-mark)
			m[string(v)] = i+1
			//fmt.Printf("1i=%d,mark=%d,max=%d\n",i+1,mark,max)
		}else{
			max = maxInt(max,i+1-maxInt(m[string(v)],mark))
			mark = maxInt(mark,m[string(v)])
			m[string(v)] = i +1
			//fmt.Printf("2i=%d,mark=%d,max=%d\n",i+1,mark,max)


		}
	}


	return max


}

func maxInt(a,b int) int {
	if a>=b{
		return a
	}else{
		return b
	}
}
func main() {
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))

}
