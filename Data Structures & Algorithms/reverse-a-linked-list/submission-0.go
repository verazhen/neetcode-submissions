/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	//record next
	//redirect next to prev
	//record prev as curr
	//move curr

	var prev *ListNode
	curr := head
	var next *ListNode

	for {
		if curr == nil{
			return prev
		}

		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}
