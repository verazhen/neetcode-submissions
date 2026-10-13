/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}

	groupStart := head
	groupEnd := head
	prevGroupEnd := dummy

	count := 1

	for groupEnd != nil {
		// Find the end of a complete k-group
		if count < k {
			groupEnd = groupEnd.Next
			count++
			continue
		}

		nextGroupStart := groupEnd.Next

		// Reverse current group
		prev := nextGroupStart
		curr := groupStart

		for curr != nextGroupStart {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// Connect previous group to the new head
		prevGroupEnd.Next = groupEnd

		// Move to the next group
		prevGroupEnd = groupStart
		groupStart = nextGroupStart
		groupEnd = nextGroupStart
		count = 1
	}

	return dummy.Next
}