func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	// Continue until both lists are exhausted.
	// Start sum with carry, then add values only when the node exists.
	// This avoids separate l1Val/l2Val variables and keeps nil checks together
	// with pointer movement for better readability.
	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next

		carry = sum / 10
	}

	// If the most significant digit produces a carry,
	// append one extra node.
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}