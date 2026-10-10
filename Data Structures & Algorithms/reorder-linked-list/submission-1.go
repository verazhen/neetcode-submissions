func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. Calculate the size of the first half
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}
	leftSize := (length + 1) / 2

	// 2. Move curr to the first node of the second half
	curr := head
	var prev *ListNode

	for i := 0; i < leftSize; i++ {
		prev = curr
		curr = curr.Next
	}

	// Split into two lists
	prev.Next = nil

	// 3. Reverse the second half
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