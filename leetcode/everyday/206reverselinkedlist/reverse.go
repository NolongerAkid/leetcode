package main

type ListNode struct {
	Val int
	Next *ListNode
}
func nonrecursion(head *ListNode) *ListNode{
	if head == nil {
		return nil
	}
	pre := new(ListNode)
	for head != nil {
		post := head.Next
		head.Next = pre
		pre = head
		head = post
	}
	return pre
}
func recursion(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	var newHead *ListNode = recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
func main() {
	
}
