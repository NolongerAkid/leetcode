package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2){
		str1,str2 = str2,str1
	}
	for i :=len(str2);i>=0;i--{
		temp := str2[0:i]
		if len(temp)*strings.Count(str1,temp) == len(str1)&&len(temp)*strings.Count(str2,temp) == len(str2){
			return temp
		}
	}
	return ""

}
func main() {
	fmt.Println(gcdOfStrings("LEET","CODE"))
}
