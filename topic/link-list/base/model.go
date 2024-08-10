package model

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func (list *ListNode) Append(l *ListNode) *ListNode {
	if list == nil {
		list = l
		return list
	}

	list.Next = l
	return list.Next
}

func (list *ListNode) AppendIntSelf(val int) *ListNode {
	if list == nil {
		list = &ListNode{
			Val: val,
		}
		return list
	}

	itor := list
	for itor.Next != nil {
		itor = itor.Next
	}

	itor.Next = &ListNode{Val: val}
	return list
}

func (list *ListNode) AppendInt(val int) *ListNode {
	if list == nil {
		list = &ListNode{
			Val: val,
		}

		return list
	}

	list.Next = &ListNode{
		Val: val,
	}

	return list.Next
}

func (list *ListNode) AddHead(l *ListNode) *ListNode {
	if list == nil {
		list = l
		return list
	}

	l.Next = list
	return list.Next
}

func (list *ListNode) Print() {
	curr := list
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
