func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}


	// 1. Move curr to the first node of the second half
	slow := head
	fast := head

	for fast.Next!=nil && fast.Next.Next!=nil  {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split into two lists
	curr := slow.Next
	slow.Next = nil

	// 2. Reverse the second half
	var reversed *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = reversed
		reversed = curr
		curr = next
	}

	// 4. Merge two lists
	curr = head
	for reversed != nil {
		nextCurr := curr.Next
		nextReversed := reversed.Next

		curr.Next = reversed
		reversed.Next = nextCurr

		curr = nextCurr
		reversed = nextReversed
	}
}