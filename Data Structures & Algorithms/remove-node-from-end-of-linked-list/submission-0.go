/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. found the node to removed, stay at prev node
	// 1,(2),3,(4) fast node goes n step ahead,
	// slow one will be right on the prev node while fast node on the last node

	dummy := &ListNode{}
	dummy.Next = head
	slow := dummy
	fast := dummy

	for i:=0;i<n;i++{
		fast = fast.Next
	}

	for fast.Next != nil{
		slow = slow.Next
		fast = fast.Next
	}


	// 2. by removing the node, we need to stay at the previous node first
	// also noted that we need a dummy node incase the removed node is the first node
	slow.Next = slow.Next.Next



	return dummy.Next
}
