/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Curr := l1
	l2Curr := l2
	dummy := &ListNode{}
	l3Curr := dummy
	carry := 0
	for l1Curr != nil || l2Curr != nil{
		l1Val := 0
		if l1Curr != nil{
			l1Val = l1Curr.Val
		}

		l2Val := 0
		if l2Curr != nil{
			l2Val = l2Curr.Val
		}

		sum := l1Val + l2Val + carry
		digit := sum%10
		carry = sum/10

		sumNode := &ListNode{
			Val: digit,
		}

		l3Curr.Next = sumNode
		l3Curr = l3Curr.Next

		if l1Curr != nil{
			l1Curr = l1Curr.Next
		}

		if l2Curr != nil{
			l2Curr = l2Curr.Next
		}

	}

	if carry > 0{
		l3Curr.Next = &ListNode{
			Val: carry,
		}
	}

	return dummy.Next
}
