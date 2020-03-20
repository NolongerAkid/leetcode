package main

import (
	"fmt"
)

type ListNode struct {
	    Val int
	    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur,carry := 0,0
	//var temp *ListNode
	head := ListNode{0,nil}
	temp := &head
	fmt.Printf("%p,%p\n",&head,temp)
	for l1 != nil && l2 != nil {
		cur = l1.Val + l2.Val + carry
		if cur > 10 {
			cur = cur % 10
			carry = 1
		}else{
			carry = 0
		}
		temp.Next = &ListNode{cur,nil}
		//temp.Val = cur
		//fmt.Println(temp.Val)
		l1 = l1.Next
		l2 = l2.Next

		temp = temp.Next
		/*if l1 != nil && l2 != nil{
			cur = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next

		}else if l1 == nil {
			cur = l2.Val + carry
			l2 = l2.Next
		}else {//l2 == nil
			cur = l1.Val + carry
			l1 = l1.Next
		}
		if cur > 10 {
			cur = cur % 10
			carry = 1
		}
		res = &ListNode{cur,nil}
		res = res.Next
		*/


	}

	fmt.Println(head.Val)
	fmt.Println(head.Next.Val)
	return head.Next

}
func main() {
	l1 := &ListNode{2,nil}
	l1.Next = &ListNode{3,nil}
	l1.Next.Next = &ListNode{4,nil}
	l2 := &ListNode{5,nil}
	l2.Next = &ListNode{6,nil}
	l2.Next.Next = &ListNode{4,nil}
	l3 := addTwoNumbers(l1,l2)
	for l1 != nil {
		fmt.Print(l1.Val)
		l1 = l1.Next
	}
	fmt.Println()
	for l2 != nil {
		fmt.Print(l2.Val)
		l2 = l2.Next
	}
	fmt.Println()
	for l3 != nil {
		fmt.Print(l3.Val)
		l3 = l3.Next
	}
}
