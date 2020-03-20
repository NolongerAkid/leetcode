package main
type TreeNode struct {
	     Val int
	     Left *TreeNode
	    Right *TreeNode
	 }
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func isValidBST(root *TreeNode) bool {

	return helper(root,INT_MIN,INT_MAX)
}
func helper(root *TreeNode,lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper{
		return false
	}

	return helper(root.Left,lower,root.Val)&&helper(root.Right,root.Val,upper)


}

