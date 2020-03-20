package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	if S == ""{
		return S
	}
	res := ""
	cur := S[0]
	temp := 1
	for i:=1;i<len(S);i++ {
		if S[i] == cur{
			temp ++
		}else{
			res += string(cur)
			res += strconv.Itoa(temp)
			cur = S[i]
			temp = 1
		}

	}
	res += string(cur)
	res += strconv.Itoa(temp)
	if len(res) >= len(S){
		return S
	}else{
		return res
	}


}
func main() {
	fmt.Println(compressString("aabcccccaa"))
	
}
