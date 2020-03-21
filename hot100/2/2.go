package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
	}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0,nil}
	temp := head
	cur,carry := 0,0


	for l1 != nil && l2 != nil {
		cur = l1.Val + l2.Val + carry
		if cur >= 10 {
			cur = cur % 10
			carry = 1
		}else {
			carry = 0
		}
		temp.Next = & ListNode{cur,nil}
		l1 = l1.Next
		l2 = l2.Next
		temp = temp.Next
	}
	for l1 != nil {
		cur = l1.Val  + carry
		if cur >= 10 {
			cur = cur % 10
			carry = 1
		}else {
			carry = 0
		}
		temp.Next = & ListNode{cur,nil}
		l1 = l1.Next
		temp = temp.Next
	}
	for l2 != nil{
		cur = l2.Val  + carry
		if cur >= 10 {
			cur = cur % 10
			carry = 1
		}else {
			carry = 0
		}
		temp.Next = &ListNode{cur,nil}
		l2 = l2.Next
		temp = temp.Next
	}
	if l1==nil && l2 == nil && carry != 0 {
		temp.Next = &ListNode{1,nil}
	}
	return head.Next



}
func length(head *ListNode) int {
	len := 0
	for head !=nil{
		len ++
		head = head.Next
	}
	return len
}
func main() {
	/*head := &ListNode{0,nil}
	temp := head
	fmt.Printf("%p,%p",head,temp)*/
	l1 := &ListNode{2,nil}
	l1.Next = &ListNode{3,nil}
	//l1.Next.Next = &ListNode{4,nil}
	l2 := &ListNode{5,nil}
	l2.Next = &ListNode{7,nil}
	l2.Next.Next = &ListNode{9,nil}
	l3 := addTwoNumbers(l1,l2)
	for l3 != nil {
		fmt.Print(l3.Val)
		l3 = l3.Next
	}
}
