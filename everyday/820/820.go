package main

import (
	"fmt"
	"sort"
)
type wordsslice []string
func(ws wordsslice) Len()int{
	return len(ws)
}
func(ws wordsslice) Swap(i,j int){
	ws[i],ws[j] = ws[j],ws[i]
}
func(ws wordsslice)Less(i,j int)bool{
	return len(ws[i])<len(ws[j])
}
type Node struct{
	val string
	children []*Node
}
func(n *Node) Insert(s string){
	child := Node{
		val : s,
	}
	n.children = append(n.children,&child)
}
func Instertword(root *Node,word string) int {
	b := false
	cn := root

	for i := len(word)-1;i>=0;i--{
		l := len(cn.children)
		j :=0
		for ;j<l;j++{
			if string(word[i]) == cn.children[j].val{
				cn = cn.children[j]
				break
			}
		}
		if j == l{//新的节点 新的单词
			b = true
			cn.Insert(string(word[i]))
			cn = cn.children[l]
		}

	}
	if b{
		return len(word)+1
	}else{
		return 0
	}
}
func minimumLengthEncoding(words []string) int {
	var ws wordsslice
	ws = words
	sort.Sort(ws)//按长度排序
	fmt.Println(ws)
	count := 0
	root := new(Node)
	for i := len(ws)-1;i>=0;i--{
		count += Instertword(root,ws[i])
	}
	return count

}

func main() {
	println(minimumLengthEncoding([]string{"atime", "time", "btime"}))
}
