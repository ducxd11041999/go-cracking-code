package sumlist

import (
	model "my-code/go-cracking-code/topic/link-list/base"
)

func sum2List(l1 *model.ListNode, l2 *model.ListNode, carry int) *model.ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	rsNode := &model.ListNode{}
	sum := carry
	if l1 != nil {
		sum += l1.Val
	}

	if l2 != nil {
		sum += l2.Val
	}

	rsNode.Val = sum % 10
	if l1 != nil || l2 != nil {
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		addNext := sum2List(l1, l2, sum/10)
		rsNode.Next = addNext
	}

	return rsNode
}
