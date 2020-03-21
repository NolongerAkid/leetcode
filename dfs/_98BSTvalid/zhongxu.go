package main

import "fmt"

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []TreeNode{*root}
	pre := ^int(^uint(0) >> 1)
	for len(stack)!= 0{
		for root != nil && root.Left != nil {
			root = root.Left
			stack = append(stack,*root)
		}
		temp := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		if temp.Val <= pre{
			return false
		}
		pre = temp.Val
		root = temp.Right
		if root != nil {
			stack = append(stack,*root)
		}

	}
	return true

}
func main(){
	root := TreeNode{2,nil,nil}
	root.Left = &TreeNode{1,nil,nil}
	root.Right = &TreeNode{3,nil,nil}
	fmt.Println(isValidBST2(&root))
}