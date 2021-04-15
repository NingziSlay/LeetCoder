package main

import "fmt"

func main() {
	n1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}

	n2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	n3 := addTwoNumbers(n1, n2)
	fmt.Println(n3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return foo(l1, l2, 0)
}

func foo(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	result := &ListNode{}
	result.Val = l1.Val + l2.Val + carry
	carry = result.Val / 10
	result.Val = result.Val % 10
	result.Next = foo(l1.Next, l2.Next, carry)
	return result
}
