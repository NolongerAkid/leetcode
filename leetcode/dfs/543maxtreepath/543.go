package main


 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func diameterOfBinaryTree(root *TreeNode) int {
	res := 1
	depth(root,&res)
	return res-1

}

func depth(node *TreeNode,res *int) int{
	if node == nil {
		return 0
	}
	left := depth(node.Left,res)
	right := depth(node.Right,res)
	maxlength := max(left,right)+1
	*res = max(*res,left+right+1)
	return maxlength


}

func max(a,b int) int{
	if a >= b {
		return a
	}else{
		return b
	}
}
func main() {
	
}
