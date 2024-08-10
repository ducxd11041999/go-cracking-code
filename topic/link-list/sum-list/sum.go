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
		addNext := sum2List(l1.Next, l2.Next, sum/10)
		rsNode.Next = addNext
	}

	return rsNode
}
