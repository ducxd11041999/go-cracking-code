package palindrome

import model "my-code/go-cracking-code/topic/link-list/base"

func Palindrome(head *model.ListNode) bool {
	clone := ReverseList(head)
	return IsPalindrome(head, clone)
}

// 1 2 3 4 5 6
func ReverseList(head *model.ListNode) *model.ListNode {

	var clone *model.ListNode
	for head != nil {
		node := model.NewList(head.Val)
		node.Next = clone
		clone = node

		head = head.Next
	}

	return clone
}

func IsPalindrome(h1, h2 *model.ListNode) bool {
	if h1 == nil || h2 == nil {
		return false
	}

	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}

		h1 = h1.Next
		h2 = h2.Next
	}

	return true
}

func IsPalindromeUsingStack(head *model.ListNode) bool {
	fast := head
	slow := head
	stack := []int{}

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)

		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if val != slow.Val {
			return false
		}

		slow = slow.Next
	}

	return true
}
